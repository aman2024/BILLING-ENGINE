package business

import (
	"billing-engine/domain"
	"billing-engine/services"

	"github.com/gin-gonic/gin"
)

func ViewLoan(c *gin.Context, services *services.Services, req *domain.ViewLoanReq) (*domain.ViewDetailedLoanRes, error) {
	var res domain.ViewDetailedLoanRes
	result, err := services.DB.ReadLoanInfo(c, req.UserId, req.LoanId)
	if err != nil {
		return nil, err
	}

	for _, entity := range result {
		e := domain.ViewDetailedLoanEntity{
			LoanId:          entity.LoanId,
			Amount:          entity.Amount,
			InitialTerm:     entity.InitialTerm,
			Rate:            entity.Rate,
			RepaymentAmount: entity.RepaymentAmount,
			Status:          entity.Status,
			AmountPaid:      entity.AmountPaid,
			IsDelinquent:    entity.IsDelinquent,
			TotalTerms:      entity.TotalTerms,
			CreatedAt:       entity.CreatedAt,
			UpdatedAt:       entity.UpdatedAt,
			ApprovedAt:      entity.ApprovedAt,
		}
		e.OutstandingBalance = float64(e.RepaymentAmount) - e.AmountPaid
		e.SkippedTerms = entity.TotalTerms - entity.InitialTerm
		if entity.Status != domain.STATUS_PENDING {
			repaymentEntity, err := services.DB.ReadRepaymentsInfo(c, int(e.LoanId), 0)
			if err != nil {
				return nil, err
			}
			e.RepaymentsInfo = repaymentEntity
		}
		res.LoanInfo = append(res.LoanInfo, e)
	}

	return &res, err

}
