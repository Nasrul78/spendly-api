package repository

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

func parseUUID(s string) (pgtype.UUID, error) {
	var uid pgtype.UUID
	if err := uid.Scan(s); err != nil {
		return pgtype.UUID{}, fmt.Errorf("invalid uuid %q: %w", s, err)
	}
	return uid, nil
}

func parseOptionalUUID(s *string) (pgtype.UUID, error) {
	var uid pgtype.UUID
	if s == nil {
		return uid, nil
	}
	if err := uid.Scan(s); err != nil {
		return pgtype.UUID{}, fmt.Errorf("invalid uuid %q: %w", *s, err)
	}
	return uid, nil
}

func parseDate(s string) (pgtype.Date, error) {
	var d pgtype.Date
	if err := d.Scan(s); err != nil {
		return pgtype.Date{}, fmt.Errorf("invalid date %q: %w", s, err)
	}
	return d, nil
}

func parseOptionalDate(s *string) (pgtype.Date, error) {
	var d pgtype.Date
	if s == nil {
		return d, nil
	}
	if err := d.Scan(s); err != nil {
		return pgtype.Date{}, fmt.Errorf("invalid date %q: %w", *s, err)
	}
	return d, nil
}

func parseOptionalText(s *string) (pgtype.Text, error) {
	var t pgtype.Text
	if s == nil {
		return t, nil
	}
	if err := t.Scan(s); err != nil {
		return pgtype.Text{}, fmt.Errorf("invalid text %q: %w", *s, err)
	}
	return t, nil
}
