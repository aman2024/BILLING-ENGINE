package db

import (
	"billing-engine/domain"

	"github.com/gin-gonic/gin"
)

type SQLDbQuery interface {
	ReadLoanInfo(c *gin.Context, userId string, loanId int) ([]domain.ViewLoanEntity, error)
	Insert(c *gin.Context, tableName string, columns []string, values []interface{}) (int64, error)
	UpdateLoanStatusToApproved(c *gin.Context, req *domain.ApproveLoanReq) (int64, error)
	UpdateLoanToIsDelinquentOrTotalTerms(c *gin.Context, loanId int, isDelinquent bool, totalTerms int) (int64, error)
	UpdateLoanStatusAndAmountPaid(c *gin.Context, status string, amountPaid float64, loanId int) (int64, error)
	UpdateRepaymentStatus(c *gin.Context, id int, loanId int, termNo int, status string) (int64, error)
	UpdateRepaymentsFromLockedToPending(c *gin.Context) (int64, error)
	ReadPendingRepayments(c *gin.Context) ([]domain.ViewRepaymentEntity, error)
	ReadRepaymentsInfo(c *gin.Context, loanId int, termNo int) ([]domain.ViewRepaymentEntity, error)
}
