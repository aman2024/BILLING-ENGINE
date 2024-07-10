package db

import (
	"billing-engine/domain"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type DbClient struct {
	Client *sql.DB
}

func (dbClient *DbClient) ReadLoanInfo(c *gin.Context, userId string, loanId int) ([]domain.ViewLoanEntity, error) {
	var res []domain.ViewLoanEntity
	conn, err := dbClient.Client.Conn(c)
	if err != nil {
		return res, err
	}
	defer conn.Close()

	qbuilder := squirrel.Select("id", "amount", "initial_term", "rate", "repayment_amount", "status", "amount_paid", "is_delinquent", "total_terms", "created_at", "updated_at", "approved_at").
		From("billing_info")

	if userId != "" {
		qbuilder = qbuilder.Where("user_id = ?", userId)
	}
	if loanId != 0 {
		qbuilder = qbuilder.Where("id = ?", loanId)
	}

	query, qargs, err := qbuilder.ToSql()
	if err != nil {
		return res, err
	}
	rows, err := conn.QueryContext(c, query, qargs...)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		entity := domain.ViewLoanEntity{}
		err := rows.Scan(&entity.LoanId, &entity.Amount, &entity.InitialTerm, &entity.Rate, &entity.RepaymentAmount, &entity.Status, &entity.AmountPaid, &entity.IsDelinquent, &entity.TotalTerms, &entity.CreatedAt, &entity.UpdatedAt, &entity.ApprovedAt)
		if err != nil {
			return res, err
		}
		res = append(res, entity)
	}
	return res, nil

}

func (dbClient *DbClient) Insert(c *gin.Context, tableName string, columns []string, values []interface{}) (int64, error) {
	conn, err := dbClient.Client.Conn(c)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	qbuilder := squirrel.Insert(tableName).
		Columns(columns...).
		Values(values...)

	query, qargs, err := qbuilder.ToSql()

	if err != nil {
		return 0, err
	}
	res, err := conn.ExecContext(c, query, qargs...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (dbClient *DbClient) UpdateLoanStatusToApproved(c *gin.Context, req *domain.ApproveLoanReq) (int64, error) {
	conn, err := dbClient.Client.Conn(c)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	qbuilder := squirrel.Update("billing_info").
		Set("status", domain.STATUS_APPROVED).
		Set("admin_approver_id", req.AdminId).
		Set("approved_at", time.Now()).
		Where(squirrel.Eq{"id": req.LoanId})

	query, qargs, err := qbuilder.ToSql()

	if err != nil {
		return 0, err
	}
	res, err := conn.ExecContext(c, query, qargs...)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (dbClient *DbClient) UpdateLoanStatusAndAmountPaid(c *gin.Context, status string, amountPaid float64, loanId int) (int64, error) {
	conn, err := dbClient.Client.Conn(c)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	qbuilder := squirrel.Update("billing_info").
		Set("status", status).
		Set("amount_paid", amountPaid).
		Where(squirrel.Eq{"id": loanId})

	query, qargs, err := qbuilder.ToSql()

	if err != nil {
		return 0, err
	}
	res, err := conn.ExecContext(c, query, qargs...)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (dbClient *DbClient) UpdateLoanToIsDelinquentOrTotalTerms(c *gin.Context, loanId int, isDelinquent bool, totalTerms int) (int64, error) {
	conn, err := dbClient.Client.Conn(c)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	qbuilder := squirrel.Update("billing_info").
		Where(squirrel.Eq{"id": loanId})

	if isDelinquent {
		qbuilder = qbuilder.Set("is_delinquent", isDelinquent)
	}
	if totalTerms != 0 {
		qbuilder = qbuilder.Set("total_terms", totalTerms)

	}

	query, qargs, err := qbuilder.ToSql()

	if err != nil {
		return 0, err
	}
	res, err := conn.ExecContext(c, query, qargs...)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (dbClient *DbClient) UpdateRepaymentsFromLockedToPending(c *gin.Context) (int64, error) {
	conn, err := dbClient.Client.Conn(c)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	today := time.Now().Format("2006-01-02")
	qbuilder := squirrel.Update("repayment_info").
		Set("status", domain.STATUS_PENDING).
		Where(squirrel.Eq{"DATE(repayment_date)": today}).
		Where(squirrel.Eq{"status": domain.STATUS_LOCKED})

	query, qargs, err := qbuilder.ToSql()

	if err != nil {
		return 0, err
	}
	res, err := conn.ExecContext(c, query, qargs...)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (dbClient *DbClient) ReadPendingRepayments(c *gin.Context) ([]domain.ViewRepaymentEntity, error) {
	var res []domain.ViewRepaymentEntity
	conn, err := dbClient.Client.Conn(c)
	if err != nil {
		return res, err
	}
	defer conn.Close()

	today := time.Now().Format("2006-01-02")

	qbuilder := squirrel.Select("id", "loan_id", "repayment_amount", "term_no", "repayment_date", "status").
		From("repayment_info").
		Where(squirrel.Eq{"status": domain.STATUS_PENDING}).
		Where("DATE(repayment_date) = DATE_SUB(?, INTERVAL 7 DAY)", today)

	query, qargs, err := qbuilder.ToSql()
	if err != nil {
		return res, err
	}
	rows, err := conn.QueryContext(c, query, qargs...)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		entity := domain.ViewRepaymentEntity{}
		err := rows.Scan(&entity.Id, &entity.LoanId, &entity.RepaymentAmount, &entity.TermNo, &entity.RepaymentDate, &entity.Status)
		if err != nil {
			return res, err
		}
		res = append(res, entity)
	}
	return res, nil
}

func (dbClient *DbClient) ReadRepaymentsInfo(c *gin.Context, loanId int, termNo int) ([]domain.ViewRepaymentEntity, error) {
	var res []domain.ViewRepaymentEntity
	conn, err := dbClient.Client.Conn(c)
	if err != nil {
		return res, err
	}
	defer conn.Close()

	qbuilder := squirrel.Select("id", "loan_id", "repayment_amount", "term_no", "repayment_date", "status").
		From("repayment_info")

	if loanId != 0 {
		qbuilder = qbuilder.Where("loan_id = ?", loanId)
	}
	if termNo != 0 {
		qbuilder = qbuilder.Where("term_no = ?", termNo)
	}

	query, qargs, err := qbuilder.ToSql()
	if err != nil {
		return res, err
	}
	rows, err := conn.QueryContext(c, query, qargs...)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		entity := domain.ViewRepaymentEntity{}
		err := rows.Scan(&entity.Id, &entity.LoanId, &entity.RepaymentAmount, &entity.TermNo, &entity.RepaymentDate, &entity.Status)
		if err != nil {
			return res, err
		}
		res = append(res, entity)
	}
	return res, nil
}

func (dbClient *DbClient) UpdateRepaymentStatus(c *gin.Context, id int, loanId int, termNo int, status string) (int64, error) {
	conn, err := dbClient.Client.Conn(c)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	qbuilder := squirrel.Update("repayment_info").
		Set("status", status).
		Where(squirrel.Eq{"id": id}).
		Where(squirrel.Eq{"loan_id": loanId}).
		Where(squirrel.Eq{"term_no": termNo})

	query, qargs, err := qbuilder.ToSql()

	if err != nil {
		return 0, err
	}
	res, err := conn.ExecContext(c, query, qargs...)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
