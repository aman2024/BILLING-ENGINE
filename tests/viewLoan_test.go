package tests

import (
	"billing-engine/business"
	"billing-engine/domain"
	"billing-engine/services"
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestViewLoan(t *testing.T) {
	t.Run("TestViewLoan", func(t *testing.T) {

		ReadLoanInfoResp := []domain.ViewLoanEntity{{LoanId: 1, Amount: 1000, Status: domain.STATUS_PENDING}}

		dbClient := new(DbClientMock)
		dbClient.On("ReadLoanInfo", mock.Anything, mock.Anything, mock.Anything).Return(ReadLoanInfoResp, nil)

		service := services.Services{
			DB: dbClient,
		}
		req := domain.ViewLoanReq{
			UserId: "1",
		}

		resp := domain.ViewDetailedLoanRes{LoanInfo: []domain.ViewDetailedLoanEntity{{LoanId: 1, Amount: 1000, Status: domain.STATUS_PENDING}}}
		res, err := business.ViewLoan(&gin.Context{}, &service, &req)

		assert.NoError(t, err)
		assert.Equal(t, &resp, res)
	})

	t.Run("TestViewLoan", func(t *testing.T) {

		ReadLoanInfoResp := []domain.ViewLoanEntity{{}}
		dbClient := new(DbClientMock)
		dbClient.On("ReadLoanInfo", mock.Anything, mock.Anything, mock.Anything).Return(ReadLoanInfoResp, errors.New("no loan found"))

		service := services.Services{
			DB: dbClient,
		}
		req := domain.ViewLoanReq{
			UserId: "1",
		}
		_, err := business.ViewLoan(&gin.Context{}, &service, &req)

		assert.Error(t, err)
	})

}
