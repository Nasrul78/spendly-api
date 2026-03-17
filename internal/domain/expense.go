package domain

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID           uuid.UUID  `json:"id"`
	UserID       uuid.UUID  `json:"user_id"`
	CategoryID   *uuid.UUID `json:"category_id,omitempty"`
	CategoryName *string    `json:"category,omitempty"`
	Amount       int64      `json:"amount"`
	Note         *string    `json:"note,omitempty"`
	Date         time.Time  `json:"date"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type CreateExpenseRequest struct {
	CategoryID *string `json:"category_id"`
	Amount     int64   `json:"amount" validate:"required,gt=0"`
	Note       *string `json:"note"`
	Date       string  `json:"date" validate:"required"`
}

type UpdateExpenseRequest struct {
	CategoryID *string `json:"category_id"`
	Amount     int64   `json:"amount" validate:"required,gt=0"`
	Note       *string `json:"note"`
	Date       string  `json:"date" validate:"required"`
}

type ExpenseFilter struct {
	From       *string
	To         *string
	CategoryID *string
	Page       int
	Limit      int
}

type ExpenseSummaryFilter struct {
	From *string
	To   *string
}

type ExpenseSummaryItem struct {
	Category string `json:"category"`
	Total    int64  `json:"total"`
}

type ExpenseSummaryResponse struct {
	From      string               `json:"from,omitempty"`
	To        string               `json:"to,omitempty"`
	Total     int64                `json:"total"`
	Breakdown []ExpenseSummaryItem `json:"breakdown"`
}

type PaginatedExpenseResponse struct {
	Data  []Expense `json:"data"`
	Total int64     `json:"total"`
	Page  int       `json:"page"`
	Limit int       `json:"limit"`
}
