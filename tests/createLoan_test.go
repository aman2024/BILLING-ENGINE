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

func TestCreateLoan(t *testing.T) {
	t.Run("TestCreateLoan", func(t *testing.T) {

		dbClient := new(DbClientMock)
		dbClient.On("Insert", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(int64(1), nil)

		service := services.Services{
			DB: dbClient,
		}
		req := domain.CreateLoanReq{
			UserId: "1",
			Amount: 500,
		}
		res, err := business.CreateLoan(&gin.Context{}, &service, &req)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), res.LoanId)
	})

	t.Run("TestCreateLoan", func(t *testing.T) {

		dbClient := new(DbClientMock)
		dbClient.On("Insert", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(int64(0), errors.New("already exists"))

		service := services.Services{
			DB: dbClient,
		}
		req := domain.CreateLoanReq{
			UserId: "1",
			Amount: 500,
		}
		_, err := business.CreateLoan(&gin.Context{}, &service, &req)
		assert.Error(t, err)
	})

}
