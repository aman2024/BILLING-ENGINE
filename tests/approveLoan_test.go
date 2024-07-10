package tests

import (
	"billing-engine/business"
	"billing-engine/domain"
	"billing-engine/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestApproveLoan(t *testing.T) {

	t.Run("TestApproveLoan", func(t *testing.T) {

		ReadLoanInfoResp := []domain.ViewLoanEntity{{LoanId: 1, Amount: 1000, Status: domain.STATUS_PENDING}}

		dbClient := new(DbClientMock)
		dbClient.On("ReadLoanInfo", mock.Anything, mock.Anything, mock.Anything).Return(ReadLoanInfoResp, nil)
		dbClient.On("UpdateLoanStatusToApproved", mock.Anything, mock.Anything).Return(int64(1), nil)
		dbClient.On("Insert", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(int64(1), nil)

		service := services.Services{
			DB: dbClient,
		}
		req := domain.ApproveLoanReq{
			LoanId: 1,
		}
		res, err := business.ApproveLoan(&gin.Context{}, &service, &req)

		expectedRes := domain.ApproveLoanRes{LoanId: 1, Status: "APPROVED"}
		assert.NoError(t, err)
		assert.Equal(t, &expectedRes, res)
	})

	t.Run("TestApproveLoan", func(t *testing.T) {

		ReadLoanInfoResp := []domain.ViewLoanEntity{{LoanId: 1, Amount: 1000, Status: domain.STATUS_PAID}}

		dbClient := new(DbClientMock)
		dbClient.On("ReadLoanInfo", mock.Anything, mock.Anything, mock.Anything).Return(ReadLoanInfoResp, nil)

		service := services.Services{
			DB: dbClient,
		}
		req := domain.ApproveLoanReq{
			LoanId: 1,
		}
		_, err := business.ApproveLoan(&gin.Context{}, &service, &req)

		assert.Error(t, err)
	})

}
