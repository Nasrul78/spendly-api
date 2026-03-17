package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nasrul78/spendly-api/internal/db"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{queries: db.New(pool)}
}

func (r *UserRepository) Create(ctx context.Context, name, email, password string) (*db.User, error) {
	user, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Name:     name,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, MapError(err)
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*db.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, MapError(err)
	}

	return &user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*db.User, error) {
	user, err := r.queries.GetUserByID(ctx, pgtype.UUID{})
	if err != nil {
		return nil, MapError(err)
	}

	return &user, nil
}
