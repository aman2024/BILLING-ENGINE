package handler

import (
	"billing-engine/business"
	"billing-engine/domain"
	"billing-engine/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRepaymentHandler(c *gin.Context, services *services.Services) {
	req, err := validateAddRepaymentReq(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":          false,
			"error":           "INVALID_REQ",
			"error_statement": err.Error(),
		})
		return
	}
	res, err := business.AddRepayment(c, services, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":          false,
			"error":           "ADD_REPAYMENT_ERROR",
			"error_statement": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   res,
	})
}

func validateAddRepaymentReq(c *gin.Context) (*domain.AddRepaymentReq, error) {
	req := domain.AddRepaymentReq{}
	err := c.BindJSON(&req)
	if err != nil {
		return nil, err
	}
	if userId, ok := c.Get("userId"); ok {
		req.UserId = userId.(string)
	} else {
		return nil, errors.New("invalid userId")
	}

	if req.LoanId <= 0 {
		return nil, errors.New("invalid loanId")
	}

	if req.TermAmount <= 0 {
		return nil, errors.New("termAmount should be greater than 0")
	}
	if req.TermNo <= 0 {
		return nil, errors.New("termNo should be greater than 0")
	}

	return &req, nil
}
