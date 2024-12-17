package repository

import (
	"context"
	"financial/model"
	"github.com/jackc/pgx/v5"
)

// FinancialRepository represents a repository that interacts with the financial table in the database.
type FinancialRepository struct {
	db *pgx.Conn
}

// NewFinancialRepository initializes a new FinancialRepository with the database connection.
func NewFinancialRepository(ctx context.Context, connString string) (*FinancialRepository, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	return &FinancialRepository{db: conn}, nil
}

// GetFinancialById retrieves a financial record by its ID.
func (repo *FinancialRepository) GetFinancialById(ctx context.Context, id int) (*model.Financial, error) {
	var financial model.Financial

	// Query the database for the financial record
	err := repo.db.QueryRow(ctx, "SELECT id, name, value, due_date, paid_at FROM financial WHERE id=$1", id).Scan(
		&financial.ID,
		&financial.Name,
		&financial.Value,
		&financial.DueDate,
		&financial.PaidAt,
	)
	if err != nil {
		return nil, err
	}

	return &financial, nil
}
