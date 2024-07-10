package routes

import (
	"billing-engine/handler"
	"billing-engine/middleware"

	"billing-engine/services"

	"github.com/gin-gonic/gin"
)

func Init(services *services.Services) *gin.Engine {
	router := gin.New()
	router.NoRoute(func(c *gin.Context) {})

	api := router.Group("/api/v1")

	api.POST("/loan/create", middleware.UserAuth(), func(ctx *gin.Context) { handler.CreateLoanHandler(ctx, services) })
	api.GET("/loan/view", middleware.UserAuth(), func(ctx *gin.Context) { handler.ViewLoanHandler(ctx, services) })
	api.PUT("/loan/add/repayment", middleware.UserAuth(), func(ctx *gin.Context) { handler.AddRepaymentHandler(ctx, services) })

	adminApi := router.Group("api/v1/admin")
	adminApi.PUT("/loan/approve", middleware.AdminAuth(), func(ctx *gin.Context) { handler.ApproveLoanHandler(ctx, services) })

	return router

}
