package business

import (
	"billing-engine/domain"
	"billing-engine/services"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func ApproveLoan(c *gin.Context, services *services.Services, req *domain.ApproveLoanReq) (*domain.ApproveLoanRes, error) {
	var res domain.ApproveLoanRes

	loanEntity, err := services.DB.ReadLoanInfo(c, "", req.LoanId)
	if err != nil {
		return nil, err
	} else if len(loanEntity) != 1 {
		return nil, errors.New("no loan found")
	}

	if loanEntity[0].Status != "PENDING" {
		return nil, errors.New("loan not in PENDING state")
	}

	result, err := services.DB.UpdateLoanStatusToApproved(c, req)
	if err != nil {
		return nil, err
	} else if result != 1 {
		return nil, errors.New("no row affected")
	} else {
		err := CreateRepaymentEntities(c, services, loanEntity[0])
		if err != nil {
			return nil, errors.New(fmt.Sprint("Failed in creating repayment entities", err))
		}
		res.LoanId = int64(req.LoanId)
		res.Status = domain.STATUS_APPROVED
	}

	return &res, nil

}

func CreateRepaymentEntities(c *gin.Context, services *services.Services, loanEntity domain.ViewLoanEntity) error {
	eachRepaymentAmount := float64(loanEntity.RepaymentAmount) / float64(50)

	for i := 1; i <= 50; i++ {
		var status string
		if i == 1 {
			status = domain.STATUS_PENDING
		} else {
			status = domain.STATUS_LOCKED
		}
		date := calculateRepaymentDate(i)
		_, err := services.DB.Insert(c, "repayment_info", []string{"loan_id", "repayment_amount", "term_no", "repayment_date", "status"}, []interface{}{loanEntity.LoanId, eachRepaymentAmount, i, date, status})
		if err != nil {
			return err
		}
	}
	return nil

}

func calculateRepaymentDate(termNo int) time.Time {
	now := time.Now()
	daysToAdd := (termNo - 1) * 7
	return now.AddDate(0, 0, daysToAdd)
}
