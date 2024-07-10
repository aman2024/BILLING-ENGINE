package cron

import (
	"billing-engine/domain"
	"billing-engine/services"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

// This cron is used to update the status of repayment for this week from LOCKED to PENDING state
func UpdateRepaymentStatusFromPendingToSkipped(services *services.Services) {
	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", func() { //runs every day once or @midnight
		ctx := gin.Context{}
		res, dbErr := services.DB.ReadPendingRepayments(&ctx)
		if dbErr != nil {
			panic(fmt.Sprint("Error in UpdateRepaymentStatusFromPendingToSkipped while reading PendingRepayments : ", dbErr))
		} else {
			for _, repaymentEntity := range res {
				loanEntity, err := services.DB.ReadLoanInfo(&ctx, "", repaymentEntity.LoanId)
				if err != nil || len(loanEntity) != 1 {
					panic(fmt.Sprint("Error in UpdateRepaymentStatusFromPendingToSkipped while reading loanEntity  : ", err))
				}

				count, err := services.DB.UpdateRepaymentStatus(&ctx, repaymentEntity.Id, repaymentEntity.LoanId, repaymentEntity.TermNo, domain.STATUS_SKIPPED)
				if err != nil || count != 1 {
					panic(fmt.Sprint("Error in UpdateRepaymentStatusFromPendingToSkipped while updating RepaymentStatus : ", err))
				}
				date := calculateRepaymentDate(loanEntity[0].TotalTerms+1, loanEntity[0].ApprovedAt.String)
				_, err = services.DB.Insert(&ctx, "repayment_info", []string{"loan_id", "repayment_amount", "term_no", "repayment_date", "status"}, []interface{}{repaymentEntity.LoanId, repaymentEntity.RepaymentAmount, loanEntity[0].TotalTerms + 1, date, domain.STATUS_LOCKED})
				if err != nil {
					panic(fmt.Sprint("Error in UpdateRepaymentStatusFromPendingToSkipped while creating new RepaymentEntity  : ", err))
				}

				isDeliquent := loanEntity[0].IsDelinquent
				if !isDeliquent && repaymentEntity.TermNo != 1 {
					lastRepaymentEntity, err := services.DB.ReadRepaymentsInfo(&ctx, repaymentEntity.LoanId, repaymentEntity.TermNo-1)
					if err != nil || len(lastRepaymentEntity) != 1 {
						panic(fmt.Sprint("Error in UpdateRepaymentStatusFromPendingToSkipped while reading lastRepaymentEntity: ", err))
					}
					if lastRepaymentEntity[0].Status == domain.STATUS_SKIPPED {
						isDeliquent = true
					}
				}

				count, err = services.DB.UpdateLoanToIsDelinquentOrTotalTerms(&ctx, repaymentEntity.LoanId, isDeliquent, loanEntity[0].TotalTerms+1)
				if err != nil || count != 1 {
					panic(fmt.Sprint("Error in UpdateRepaymentStatusFromPendingToSkipped while updating loan totalTerms : ", err))
				}

			}
		}
	})
	if err != nil {
		panic(fmt.Sprint("cron UpdateRepaymentStatusFromPendingToSkipped error : ", err))
	}
	c.Start()
}

func calculateRepaymentDate(termNo int, fromDate string) time.Time {
	now, error := time.Parse("2006-01-02 15:04:05", fromDate)
	if error != nil {
		panic(fmt.Sprint("Error in UpdateRepaymentStatusFromPendingToSkipped while caluclating  RepaymentDate : ", error))
	}
	daysToAdd := (termNo - 1) * 7
	return now.AddDate(0, 0, daysToAdd)
}
