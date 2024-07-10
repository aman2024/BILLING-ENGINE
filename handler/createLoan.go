package handler

import (
	"billing-engine/domain"
	"billing-engine/services"
	"errors"
	"net/http"

	"billing-engine/business"

	"github.com/gin-gonic/gin"
)

func CreateLoanHandler(c *gin.Context, services *services.Services) {
	req, err := validateCreateLoanReq(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":          false,
			"error":           "INVALID_REQ",
			"error_statement": err.Error(),
		})
		return
	}
	res, err := business.CreateLoan(c, services, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":          false,
			"error":           "CREATE_LOAN_ERROR",
			"error_statement": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   res,
	})
}

func validateCreateLoanReq(c *gin.Context) (*domain.CreateLoanReq, error) {
	req := domain.CreateLoanReq{}
	err := c.BindJSON(&req)
	if err != nil {
		return nil, err
	}
	if userId, ok := c.Get("userId"); ok {
		req.UserId = userId.(string)
	} else {
		return nil, errors.New("invalid userId")
	}

	if req.Amount <= 0 {
		return nil, errors.New("amount should be greater than 0")
	}

	return &req, nil
}
