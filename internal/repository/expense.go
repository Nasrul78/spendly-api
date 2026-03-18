package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
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
	uid := pgtype.UUID{}
	uid.Scan(userID)

	cid := pgtype.UUID{}
	if categoryID != nil {
		cid.Scan(*categoryID)
	}

	n := pgtype.Text{}
	if note != nil {
		n.Scan(*note)
	}

	d := pgtype.Date{}
	d.Scan(date)

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

func (r *ExpenseRepository) GetByID(ctx context.Context, userID, expenseID string) (*db.GetExpenseByIDRow, error) {
	uid := pgtype.UUID{}
	uid.Scan(userID)

	eid := pgtype.UUID{}
	eid.Scan(expenseID)

	expense, err := r.queries.GetExpenseByID(ctx, db.GetExpenseByIDParams{
		ID:     eid,
		UserID: uid,
	})
	if err != nil {
		return nil, MapError(err)
	}
	return &expense, nil
}

func (r *ExpenseRepository) GetAllByUserID(ctx context.Context, userID string, filter domain.ExpenseFilter) ([]db.GetExpensesByUserIDRow, error) {
	uid := pgtype.UUID{}
	uid.Scan(userID)

	from := pgtype.Date{}
	if filter.From != nil {
		from.Scan(*filter.From)
	}

	to := pgtype.Date{}
	if filter.To != nil {
		to.Scan(*filter.To)
	}

	cid := pgtype.UUID{}
	if filter.CategoryID != nil {
		cid.Scan(*filter.CategoryID)
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

func (r *ExpenseRepository) CountAllByUserID(ctx context.Context, userID string, filter domain.ExpenseFilter) (int64, error) {
	uid := pgtype.UUID{}
	uid.Scan(userID)

	from := pgtype.Date{}
	if filter.From != nil {
		from.Scan(*filter.From)
	}

	to := pgtype.Date{}
	if filter.To != nil {
		to.Scan(*filter.To)
	}

	cid := pgtype.UUID{}
	if filter.CategoryID != nil {
		cid.Scan(*filter.CategoryID)
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
	uid := pgtype.UUID{}
	uid.Scan(userID)

	eid := pgtype.UUID{}
	eid.Scan(expenseID)

	cid := pgtype.UUID{}
	if categoryID != nil {
		cid.Scan(*categoryID)
	}

	n := pgtype.Text{}
	if note != nil {
		n.Scan(*note)
	}

	d := pgtype.Date{}
	d.Scan(date)

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
	uid := pgtype.UUID{}
	uid.Scan(userID)

	eid := pgtype.UUID{}
	eid.Scan(expenseID)

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
	uid := pgtype.UUID{}
	uid.Scan(userID)

	from := pgtype.Date{}
	if filter.From != nil {
		from.Scan(*filter.From)
	}

	to := pgtype.Date{}
	if filter.To != nil {
		to.Scan(*filter.To)
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
