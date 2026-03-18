package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/nasrul78/spendly-api/internal/db"
	"github.com/nasrul78/spendly-api/internal/domain"
)

type CategoryRepository struct {
	queries *db.Queries
}

func NewCategoryRepository(pool *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{queries: db.New(pool)}
}

func (r *CategoryRepository) Create(ctx context.Context, userID, name string) (*db.Category, error) {
	uid := pgtype.UUID{}
	uid.Scan(userID)

	category, err := r.queries.CreateCategory(ctx, db.CreateCategoryParams{
		UserID: uid,
		Name:   name,
	})
	if err != nil {
		return nil, MapError(err)
	}
	return &category, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, userID, categoryID string) (*db.Category, error) {
	uid := pgtype.UUID{}
	uid.Scan(userID)

	cid := pgtype.UUID{}
	cid.Scan(categoryID)

	category, err := r.queries.GetCategoryByID(ctx, db.GetCategoryByIDParams{
		UserID: uid,
		ID:     cid,
	})
	if err != nil {
		return nil, MapError(err)
	}
	return &category, nil
}

func (r *CategoryRepository) GetAllByUserID(ctx context.Context, userID string) ([]db.Category, error) {
	uid := pgtype.UUID{}
	uid.Scan(userID)

	categories, err := r.queries.GetCategoriesByUserID(ctx, uid)
	if err != nil {
		return nil, MapError(err)
	}
	return categories, nil
}

func (r *CategoryRepository) Update(ctx context.Context, userID, categoryID, name string) (*db.Category, error) {
	uid := pgtype.UUID{}
	uid.Scan(userID)

	cid := pgtype.UUID{}
	cid.Scan(categoryID)

	category, err := r.queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		Name:   name,
		ID:     cid,
		UserID: uid,
	})
	if err != nil {
		return nil, MapError(err)
	}
	return &category, nil
}

func (r *CategoryRepository) Delete(ctx context.Context, userID, categoryID string) error {
	uid := pgtype.UUID{}
	uid.Scan(userID)

	cid := pgtype.UUID{}
	cid.Scan(categoryID)

	tag, err := r.queries.DeleteCategory(ctx, db.DeleteCategoryParams{
		ID:     cid,
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
