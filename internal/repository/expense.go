package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/nasrul78/spendly-api/internal/db"
	"github.com/nasrul78/spendly-api/internal/domain"
)

type ExpenseRepository struct {
	queries *db.Queries
}

func NewExpenseRepository(pool *pgxpool.Pool) *ExpenseRepository {
	return &ExpenseRepository{queries: db.New(pool)}
}

func (r *ExpenseRepository) Create(ctx context.Context, userID string, categoryID *string, amount int64, note *string, date string) (*db.Expense, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}

	cid, err := parseOptionalUUID(categoryID)
	if err != nil {
		return nil, err
	}

	n, err := parseOptionalText(note)
	if err != nil {
		return nil, err
	}

	d, err := parseDate(date)
	if err != nil {
		return nil, err
	}

	expense, err := r.queries.CreateExpense(ctx, db.CreateExpenseParams{
		UserID:     uid,
		CategoryID: cid,
		Amount:     amount,
		Note:       n,
		Date:       d,
	})
	if err != nil {
		return nil, MapError(err)
	}
	return &expense, nil
}

func (r *ExpenseRepository) GetAllByUserID(ctx context.Context, userID string, filter domain.ExpenseFilter) ([]db.GetExpensesByUserIDRow, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}

	from, err := parseOptionalDate(filter.From)
	if err != nil {
		return nil, err
	}

	to, err := parseOptionalDate(filter.To)
	if err != nil {
		return nil, err
	}

	cid, err := parseOptionalUUID(filter.CategoryID)
	if err != nil {
		return nil, err
	}

	offset := (filter.Page - 1) * filter.Limit

	expenses, err := r.queries.GetExpensesByUserID(ctx, db.GetExpensesByUserIDParams{
		UserID:  uid,
		Column2: from,
		Column3: to,
		Column4: cid,
		Limit:   int32(filter.Limit),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, MapError(err)
	}
	return expenses, nil
}

func (r *ExpenseRepository) GetByID(ctx context.Context, userID, expenseID string) (*db.GetExpenseByIDRow, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}

	eid, err := parseUUID(expenseID)
	if err != nil {
		return nil, err
	}

	expense, err := r.queries.GetExpenseByID(ctx, db.GetExpenseByIDParams{
		ID:     eid,
		UserID: uid,
	})
	if err != nil {
		return nil, MapError(err)
	}
	return &expense, nil
}

func (r *ExpenseRepository) CountAllByUserID(ctx context.Context, userID string, filter domain.ExpenseFilter) (int64, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return 0, err
	}

	from, err := parseOptionalDate(filter.From)
	if err != nil {
		return 0, err
	}

	to, err := parseOptionalDate(filter.To)
	if err != nil {
		return 0, err
	}

	cid, err := parseOptionalUUID(filter.CategoryID)
	if err != nil {
		return 0, err
	}

	count, err := r.queries.CountExpensesByUserID(ctx, db.CountExpensesByUserIDParams{
		UserID:  uid,
		Column2: from,
		Column3: to,
		Column4: cid,
	})
	if err != nil {
		return 0, MapError(err)
	}
	return count, nil
}

func (r *ExpenseRepository) Update(ctx context.Context, userID, expenseID string, categoryID *string, amount int64, note *string, date string) (*db.Expense, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}

	eid, err := parseUUID(expenseID)
	if err != nil {
		return nil, err
	}

	cid, err := parseOptionalUUID(categoryID)
	if err != nil {
		return nil, err
	}

	n, err := parseOptionalText(note)
	if err != nil {
		return nil, err
	}

	d, err := parseDate(date)
	if err != nil {
		return nil, err
	}

	expense, err := r.queries.UpdateExpense(ctx, db.UpdateExpenseParams{
		ID:         eid,
		UserID:     uid,
		CategoryID: cid,
		Amount:     amount,
		Note:       n,
		Date:       d,
	})
	if err != nil {
		return nil, MapError(err)
	}
	return &expense, nil
}

func (r *ExpenseRepository) Delete(ctx context.Context, userID, expenseID string) error {
	uid, err := parseUUID(userID)
	if err != nil {
		return err
	}

	eid, err := parseUUID(expenseID)
	if err != nil {
		return err
	}

	tag, err := r.queries.DeleteExpense(ctx, db.DeleteExpenseParams{
		ID:     eid,
		UserID: uid,
	})
	if err != nil {
		return MapError(err)
	}

	if tag == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *ExpenseRepository) GetSummaryByUserID(ctx context.Context, userID string, filter domain.ExpenseSummaryFilter) ([]db.GetExpenseSummaryByUserIDRow, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}

	from, err := parseOptionalDate(filter.From)
	if err != nil {
		return nil, err
	}

	to, err := parseOptionalDate(filter.To)
	if err != nil {
		return nil, err
	}

	summary, err := r.queries.GetExpenseSummaryByUserID(ctx, db.GetExpenseSummaryByUserIDParams{
		UserID:  uid,
		Column2: from,
		Column3: to,
	})
	if err != nil {
		return nil, MapError(err)
	}
	return summary, nil
}
