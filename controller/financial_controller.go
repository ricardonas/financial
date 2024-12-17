package controller

import (
	"financial/model"
	"financial/service"
	"github.com/gin-gonic/gin"
)

// FinancialController handles HTTP requests related to financial records.
type FinancialController struct {
	service *service.FinancialService
}

// NewFinancialController initializes a new FinancialController.
func NewFinancialController(service *service.FinancialService) *FinancialController {
	return &FinancialController{service: service}
}

// GetFinancialById is the method that retrieves financial details by ID.
func (c *FinancialController) GetFinancialById(ctx *gin.Context, id int) (*model.Financial, error) {
	// Call the service to get the financial details and return the result.
	financial, err := c.service.GetFinancialById(ctx, id)
	if err != nil {
		return nil, err
	}
	return financial, nil
}
