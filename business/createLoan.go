package business

import (
	"billing-engine/domain"
	"billing-engine/services"

	"github.com/gin-gonic/gin"
)

func CreateLoan(c *gin.Context, services *services.Services, req *domain.CreateLoanReq) (*domain.CreateLoanRes, error) {
	var res domain.CreateLoanRes
	repaymentAmount := int(float64(req.Amount) * 1.10)
	result, err := services.DB.Insert(c, "billing_info", []string{"user_id", "amount", "initial_term", "rate", "repayment_amount", "status", "total_terms"}, []interface{}{req.UserId, req.Amount, domain.TERM, domain.RATE, repaymentAmount, domain.STATUS_PENDING, domain.TERM})
	if err != nil {
		return nil, err
	} else {
		res.LoanId = result
	}

	return &res, err

}
