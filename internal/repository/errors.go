package repository

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nasrul78/spendly-api/internal/domain"
)

func MapError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ErrNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23503":
			return domain.ErrNotFound // 23503 = foreign_key_violation
		case "23505":
			return domain.ErrConflict // 23505 = unique_violation
		}
	}

	return domain.ErrInternal
}
