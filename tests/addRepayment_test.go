package tests

import (
	"billing-engine/business"
	"billing-engine/domain"
	"billing-engine/services"
	"errors"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddRepayment(t *testing.T) {

	t.Run("TestAddRepayment", func(t *testing.T) {

		ReadLoanInfoResp := []domain.ViewLoanEntity{{LoanId: 1, Amount: 1000, Status: domain.STATUS_APPROVED, RepaymentAmount: 1100}}
		ReadRepaymentsInfoResp := []domain.ViewRepaymentEntity{{Id: 1, LoanId: 1, RepaymentAmount: 22, Status: domain.STATUS_PENDING, RepaymentDate: time.Now().Format("2006-01-02 15:04:05")}}
		dbClient := new(DbClientMock)
		dbClient.On("ReadLoanInfo", mock.Anything, mock.Anything, mock.Anything).Return(ReadLoanInfoResp, nil)
		dbClient.On("ReadRepaymentsInfo", mock.Anything, mock.Anything, mock.Anything).Return(ReadRepaymentsInfoResp, nil)

		dbClient.On("UpdateRepaymentStatus", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(int64(1), nil)
		dbClient.On("UpdateLoanStatusAndAmountPaid", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(int64(1), nil)

		service := services.Services{
			DB: dbClient,
		}
		req := domain.AddRepaymentReq{
			LoanId:     1,
			UserId:     "1",
			TermNo:     1,
			TermAmount: 22,
		}
		res, err := business.AddRepayment(&gin.Context{}, &service, &req)

		expectedRes := domain.AddRepaymentRes{LoanId: 1, Balance: 1078}
		assert.NoError(t, err)
		assert.Equal(t, &expectedRes, res)
	})

	t.Run("TestAddRepayment", func(t *testing.T) {

		ReadLoanInfoResp := []domain.ViewLoanEntity{{LoanId: 1, Amount: 500, Status: domain.STATUS_APPROVED}}

		dbClient := new(DbClientMock)
		dbClient.On("ReadLoanInfo", mock.Anything, mock.Anything, mock.Anything).Return(ReadLoanInfoResp, nil)
		dbClient.On("ReadRepaymentsInfo", mock.Anything, mock.Anything, mock.Anything).Return([]domain.ViewRepaymentEntity{}, errors.New("no term entry found"))

		service := services.Services{
			DB: dbClient,
		}
		req := domain.AddRepaymentReq{
			LoanId:     1,
			UserId:     "1",
			TermNo:     2,
			TermAmount: 100,
		}
		_, err := business.AddRepayment(&gin.Context{}, &service, &req)

		assert.Error(t, err)
	})

}
