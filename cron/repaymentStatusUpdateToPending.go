package cron

import (
	"billing-engine/services"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

// This cron is used to update the status of repayment for this week from LOCKED to PENDING state
func UpdateRepaymentStatusFromLockedToPending(services *services.Services) {
	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", func() { //runs every day once or @midnight
		ctx := gin.Context{}
		count, dbErr := services.DB.UpdateRepaymentsFromLockedToPending(&ctx)
		if dbErr != nil {
			panic(fmt.Sprint("Error while updating repayment status from locked to pending in db : ", dbErr))
		} else {
			fmt.Println("Successfully updated repayment status from locked to pending :", count)
		}
	})
	if err != nil {
		panic(fmt.Sprint("cron UpdateRepaymentStatusFromLockedToPending error : ", err))
	}
	c.Start()
}
