package tests

import (
	"billing-engine/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type DbClientMock struct {
	mock.Mock
}

// ReadLoanInfo mocks the ReadLoanInfo method
func (m *DbClientMock) ReadLoanInfo(c *gin.Context, userId string, loanId int) ([]domain.ViewLoanEntity, error) {
	args := m.Called(c, userId, loanId)
	return args.Get(0).([]domain.ViewLoanEntity), args.Error(1)
}

// Insert mocks the Insert method
func (m *DbClientMock) Insert(c *gin.Context, tableName string, columns []string, values []interface{}) (int64, error) {
	args := m.Called(c, tableName, columns, values)
	return args.Get(0).(int64), args.Error(1)
}

// UpdateStatus mocks the UpdateStatus method
func (m *DbClientMock) UpdateLoanStatusToApproved(c *gin.Context, req *domain.ApproveLoanReq) (int64, error) {
	args := m.Called(c, req)
	return args.Get(0).(int64), args.Error(1)
}
func (m *DbClientMock) UpdateLoanToIsDelinquentOrTotalTerms(c *gin.Context, loanId int, isDelinquent bool, totalTerms int) (int64, error) {
	args := m.Called(c, loanId, isDelinquent, totalTerms)
	return args.Get(0).(int64), args.Error(1)
}
func (m *DbClientMock) UpdateLoanStatusAndAmountPaid(c *gin.Context, status string, amountPaid float64, loanId int) (int64, error) {
	args := m.Called(c, status, amountPaid, loanId)
	return args.Get(0).(int64), args.Error(1)
}

func (m *DbClientMock) UpdateRepaymentStatus(c *gin.Context, id int, loanId int, termNo int, status string) (int64, error) {
	args := m.Called(c, id, loanId, termNo, status)
	return args.Get(0).(int64), args.Error(1)
}

func (m *DbClientMock) UpdateRepaymentsFromLockedToPending(c *gin.Context) (int64, error) {
	args := m.Called(c)
	return args.Get(0).(int64), args.Error(1)
}

func (m *DbClientMock) ReadPendingRepayments(c *gin.Context) ([]domain.ViewRepaymentEntity, error) {
	args := m.Called(c)
	return args.Get(0).([]domain.ViewRepaymentEntity), args.Error(1)
}
func (m *DbClientMock) ReadRepaymentsInfo(c *gin.Context, loanId int, termNo int) ([]domain.ViewRepaymentEntity, error) {
	args := m.Called(c, loanId, termNo)
	return args.Get(0).([]domain.ViewRepaymentEntity), args.Error(1)
}
