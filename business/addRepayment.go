package business

import (
	"billing-engine/domain"
	"billing-engine/services"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func AddRepayment(c *gin.Context, services *services.Services, req *domain.AddRepaymentReq) (*domain.AddRepaymentRes, error) {
	var res domain.AddRepaymentRes

	loanEntity, err := services.DB.ReadLoanInfo(c, req.UserId, req.LoanId)
	if err != nil {
		return nil, err
	}
	if len(loanEntity) != 1 {
		return nil, errors.New("no loan found")
	}
	if loanEntity[0].Status != domain.STATUS_APPROVED {
		if loanEntity[0].Status == domain.STATUS_PENDING {
			return nil, errors.New("loan is not APPROVED yet")
		} else if loanEntity[0].Status == domain.STATUS_PAID {
			return nil, errors.New("loan already PAID ")
		} else {
			return nil, errors.New(fmt.Sprint("loan not in APPROVED state, current state: ", loanEntity[0].Status))
		}
	}
	repaymentEntity, err := services.DB.ReadRepaymentsInfo(c, req.LoanId, req.TermNo)
	if err != nil {
		return nil, err
	}
	if len(repaymentEntity) != 1 {
		return nil, errors.New("no term entry found")
	}
	if repaymentEntity[0].Status != domain.STATUS_PENDING {
		if loanEntity[0].Status == domain.STATUS_LOCKED {
			return nil, errors.New("term is not ENABLED/ACTIVE yet")
		} else if loanEntity[0].Status == domain.STATUS_PAID {
			return nil, errors.New("term is already PAID ")
		} else {
			return nil, errors.New(fmt.Sprint("term not in ACTIVE state, current state: ", repaymentEntity[0].Status))
		}
	}
	if repaymentEntity[0].RepaymentAmount != float64(req.TermAmount) {
		return nil, errors.New(fmt.Sprint("termAmount is not correct. It should be :", repaymentEntity[0].RepaymentAmount))
	}
	dayDiff, err := DaysDifferenceBwTodayAndRepaymentDate(repaymentEntity[0].RepaymentDate)
	if err != nil {
		return nil, err
	}
	if dayDiff >= 7 || dayDiff < 0 {
		return nil, errors.New("this term is not ACTIVE yet")
	}
	count, err := services.DB.UpdateRepaymentStatus(c, repaymentEntity[0].Id, repaymentEntity[0].LoanId, repaymentEntity[0].TermNo, domain.STATUS_PAID)
	if err != nil || count != 1 {
		return nil, errors.New("error while updating repaymentStatus to PAID ")
	}
	status := domain.STATUS_APPROVED
	if (float64(loanEntity[0].RepaymentAmount) - loanEntity[0].AmountPaid) == req.TermAmount {
		status = domain.STATUS_PAID
	}
	amountPaid := loanEntity[0].AmountPaid + req.TermAmount
	count, err = services.DB.UpdateLoanStatusAndAmountPaid(c, status, amountPaid, req.LoanId)
	if err != nil || count != 1 {
		return nil, errors.New("error while updating loan amount and status")
	}
	res.LoanId = req.LoanId
	res.Balance = float64(loanEntity[0].RepaymentAmount) - amountPaid

	return &res, nil

}

func DaysDifferenceBwTodayAndRepaymentDate(repaymentDate string) (int, error) {
	parsedRepaymentDate, err := time.Parse("2006-01-02 15:04:05", repaymentDate)
	if err != nil {
		return 0, err
	}
	diff := time.Since(parsedRepaymentDate).Hours() / 24
	fmt.Println(diff)
	return int(diff), nil

}
