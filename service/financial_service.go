package service

import (
	"context"
	"financial/model"
	"financial/repository"
)

// FinancialService handles business logic related to financial records.
type FinancialService struct {
	repo *repository.FinancialRepository
}

// NewFinancialService initializes a new FinancialService.
func NewFinancialService(repo *repository.FinancialRepository) *FinancialService {
	return &FinancialService{repo: repo}
}

// GetFinancialById retrieves a financial record by its ID.
func (s *FinancialService) GetFinancialById(ctx context.Context, id int) (*model.Financial, error) {
	// Pass the context to the repository
	return s.repo.GetFinancialById(ctx, id)
}
