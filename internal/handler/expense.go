package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/nasrul78/spendly-api/internal/domain"
	"github.com/nasrul78/spendly-api/internal/middleware"
	"github.com/nasrul78/spendly-api/internal/service"
)

type ExpenseHandler struct {
	expenseService *service.ExpenseService
	validate       *validator.Validate
}

func NewExpenseHandler(expenseService *service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		expenseService: expenseService,
		validate:       validator.New(),
	}
}

func (h *ExpenseHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)

	var req domain.CreateExpenseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validate.Struct(req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	expense, err := h.expenseService.Create(r.Context(), userID, &req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, expense)
}

func (h *ExpenseHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)
	expenseID := chi.URLParam(r, "id")

	expense, err := h.expenseService.GetByID(r.Context(), userID, expenseID)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, expense)
}

func (h *ExpenseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	limit, _ := strconv.Atoi(q.Get("limit"))

	filter := domain.ExpenseFilter{
		Page:  page,
		Limit: limit,
	}

	if v := q.Get("from"); v != "" {
		filter.From = &v
	}
	if v := q.Get("to"); v != "" {
		filter.To = &v
	}
	if v := q.Get("category_id"); v != "" {
		filter.CategoryID = &v
	}

	expenses, err := h.expenseService.GetAllByUserID(r.Context(), userID, filter)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, expenses)
}

func (h *ExpenseHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)
	expenseID := chi.URLParam(r, "id")

	var req domain.UpdateExpenseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validate.Struct(req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	expense, err := h.expenseService.Update(r.Context(), userID, expenseID, &req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, expense)
}

func (h *ExpenseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)
	expenseID := chi.URLParam(r, "id")

	if err := h.expenseService.Delete(r.Context(), userID, expenseID); err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ExpenseHandler) GetSummaryByUserID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)
	q := r.URL.Query()

	var from, to *string
	if v := q.Get("from"); v != "" {
		from = &v
	}
	if v := q.Get("to"); v != "" {
		to = &v
	}

	filter := domain.ExpenseSummaryFilter{
		From: from,
		To:   to,
	}

	summary, err := h.expenseService.GetSummaryByUserID(r.Context(), userID, filter)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, summary)
}
