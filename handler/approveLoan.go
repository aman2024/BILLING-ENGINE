package handler

import (
	"billing-engine/business"
	"billing-engine/domain"
	"billing-engine/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApproveLoanHandler(c *gin.Context, services *services.Services) {
	req, err := validateApproveLoanReq(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":          false,
			"error":           "INVALID_REQ",
			"error_statement": err.Error(),
		})
		return
	}
	res, err := business.ApproveLoan(c, services, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":          false,
			"error":           "APPROVE_LOAN_ERROR",
			"error_statement": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   res,
	})
}

func validateApproveLoanReq(c *gin.Context) (*domain.ApproveLoanReq, error) {
	req := domain.ApproveLoanReq{}
	err := c.BindJSON(&req)
	if err != nil {
		return nil, err
	}
	if adminId, ok := c.Get("adminId"); ok {
		req.AdminId = adminId.(string)
	} else {
		return nil, errors.New("invalid adminId")
	}

	if req.LoanId <= 0 {
		return nil, errors.New("invalid loanId")
	}

	return &req, nil
}
