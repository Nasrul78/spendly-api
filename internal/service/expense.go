package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/nasrul78/spendly-api/internal/domain"
	"github.com/nasrul78/spendly-api/internal/repository"
)

type ExpenseService struct {
	expenseRepo *repository.ExpenseRepository
}

func NewExpenseService(expenseRepo *repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{expenseRepo: expenseRepo}
}

func (s *ExpenseService) Create(ctx context.Context, userID string, req *domain.CreateExpenseRequest) (*domain.Expense, error) {
	_, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, domain.ErrInvalidDate
	}

	expense, err := s.expenseRepo.Create(ctx, userID, req.CategoryID, req.Amount, req.Note, req.Date)
	if err != nil {
		return nil, err
	}

	result := &domain.Expense{
		ID:        expense.ID.Bytes,
		Amount:    expense.Amount,
		Date:      expense.Date.Time,
		CreatedAt: expense.CreatedAt.Time,
		UpdatedAt: expense.UpdatedAt.Time,
	}

	if expense.CategoryID.Valid {
		cid := uuid.UUID(expense.CategoryID.Bytes)
		result.CategoryID = &cid
	}

	if expense.Note.Valid {
		n := expense.Note.String
		result.Note = &n
	}

	return result, nil
}

func (s *ExpenseService) GetAllByUserID(ctx context.Context, userID string, filter domain.ExpenseFilter) (*domain.PaginatedExpenseResponse, error) {
	if filter.Page == 0 {
		filter.Page = 1
	}

	if filter.Limit == 0 {
		filter.Limit = 10
	}

	rows, err := s.expenseRepo.GetAllByUserID(ctx, userID, filter)
	if err != nil {
		return nil, err
	}

	count, err := s.expenseRepo.CountAllByUserID(ctx, userID, filter)
	if err != nil {
		return nil, err
	}

	expenses := make([]domain.Expense, len(rows))
	for i, row := range rows {
		expense := domain.Expense{
			ID:        row.ID.Bytes,
			Amount:    row.Amount,
			Date:      row.Date.Time,
			CreatedAt: row.CreatedAt.Time,
			UpdatedAt: row.UpdatedAt.Time,
		}

		if row.CategoryID.Valid {
			cid := uuid.UUID(row.CategoryID.Bytes)
			expense.CategoryID = &cid
		}

		if row.CategoryName.Valid {
			cn := row.CategoryName.String
			expense.CategoryName = &cn
		}

		if row.Note.Valid {
			n := row.Note.String
			expense.Note = &n
		}

		expenses[i] = expense
	}

	result := &domain.PaginatedExpenseResponse{
		Data:  expenses,
		Total: count,
		Page:  filter.Page,
		Limit: filter.Limit,
	}

	return result, nil
}

func (s *ExpenseService) GetByID(ctx context.Context, userID, expenseID string) (*domain.Expense, error) {
	expense, err := s.expenseRepo.GetByID(ctx, userID, expenseID)
	if err != nil {
		return nil, err
	}

	result := &domain.Expense{
		ID:        expense.ID.Bytes,
		Amount:    expense.Amount,
		Date:      expense.Date.Time,
		CreatedAt: expense.CreatedAt.Time,
		UpdatedAt: expense.UpdatedAt.Time,
	}

	if expense.CategoryID.Valid {
		cid := uuid.UUID(expense.CategoryID.Bytes)
		result.CategoryID = &cid
	}

	if expense.CategoryName.Valid {
		cn := expense.CategoryName.String
		result.CategoryName = &cn
	}

	if expense.Note.Valid {
		n := expense.Note.String
		result.Note = &n
	}

	return result, nil
}

func (s *ExpenseService) Update(ctx context.Context, userID, expenseID string, req *domain.UpdateExpenseRequest) (*domain.Expense, error) {
	expense, err := s.expenseRepo.Update(ctx, userID, expenseID, req.CategoryID, req.Amount, req.Note, req.Date)
	if err != nil {
		return nil, err
	}

	result := &domain.Expense{
		ID:        expense.ID.Bytes,
		Amount:    expense.Amount,
		Date:      expense.Date.Time,
		CreatedAt: expense.CreatedAt.Time,
		UpdatedAt: expense.UpdatedAt.Time,
	}

	if expense.CategoryID.Valid {
		cid := uuid.UUID(expense.CategoryID.Bytes)
		result.CategoryID = &cid
	}

	if expense.Note.Valid {
		n := expense.Note.String
		result.Note = &n
	}

	return result, nil
}

func (s *ExpenseService) Delete(ctx context.Context, userID, expenseID string) error {
	if err := s.expenseRepo.Delete(ctx, userID, expenseID); err != nil {
		return err
	}
	return nil
}

func (s *ExpenseService) GetSummaryByUserID(ctx context.Context, userID string, filter domain.ExpenseSummaryFilter) (*domain.ExpenseSummaryResponse, error) {
	rows, err := s.expenseRepo.GetSummaryByUserID(ctx, userID, filter)
	if err != nil {
		return nil, err
	}

	var grandTotal int64
	breakdown := make([]domain.ExpenseSummaryItem, len(rows))
	for i, row := range rows {
		var categoryName string
		if row.CategoryName.Valid {
			categoryName = row.CategoryName.String
		} else {
			categoryName = "Uncategorized"
		}

		breakdown[i] = domain.ExpenseSummaryItem{
			Category: categoryName,
			Total:    row.TotalAmount,
		}
		grandTotal += row.TotalAmount
	}

	result := &domain.ExpenseSummaryResponse{
		Total:     grandTotal,
		Breakdown: breakdown,
	}

	if filter.From != nil {
		result.From = *filter.From
	}

	if filter.To != nil {
		result.To = *filter.To
	}

	return result, nil
}
