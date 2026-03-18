package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/nasrul78/spendly-api/internal/domain"
	"github.com/nasrul78/spendly-api/internal/middleware"
	"github.com/nasrul78/spendly-api/internal/service"
)

type CategoryHandler struct {
	categoryService *service.CategoryService
	validate        *validator.Validate
}

func NewCategoryHandler(categoryService *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
		validate:        validator.New(),
	}
}

// Create godoc
// @Summary      Create a category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body domain.CreateCategoryRequest true "Create category request"
// @Success      201 {object} domain.Category
// @Failure      400 {object} handler.errorResponse
// @Failure      401 {object} handler.errorResponse
// @Failure      409 {object} handler.errorResponse
// @Router       /categories [post]
func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	UserID := r.Context().Value(middleware.UserIDKey).(string)

	var req domain.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validate.Struct(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.categoryService.Create(r.Context(), UserID, &req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, category)
}

// GetAll godoc
// @Summary      Get all categories
// @Tags         categories
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array}  domain.Category
// @Failure      401 {object} handler.errorResponse
// @Router       /categories [get]
func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	UserID := r.Context().Value(middleware.UserIDKey).(string)

	categories, err := h.categoryService.GetAll(r.Context(), UserID)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, categories)
}

// GetByID godoc
// @Summary      Get a category by ID
// @Tags         categories
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Category ID"
// @Success      200 {object} domain.Category
// @Failure      401 {object} handler.errorResponse
// @Failure      404 {object} handler.errorResponse
// @Router       /categories/{id} [get]
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	UserID := r.Context().Value(middleware.UserIDKey).(string)
	categoryID := chi.URLParam(r, "id")

	category, err := h.categoryService.GetByID(r.Context(), UserID, categoryID)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, category)
}

// Update godoc
// @Summary      Update a category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path     string true "Category ID"
// @Param        request body domain.UpdateCategoryRequest true "Update category request"
// @Success      200 {object} domain.Category
// @Failure      400 {object} handler.errorResponse
// @Failure      401 {object} handler.errorResponse
// @Failure      404 {object} handler.errorResponse
// @Router       /categories/{id} [put]
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	UserID := r.Context().Value(middleware.UserIDKey).(string)
	categoryID := chi.URLParam(r, "id")

	var req domain.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validate.Struct(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.categoryService.Update(r.Context(), UserID, categoryID, &req)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, category)
}

// Delete godoc
// @Summary      Delete a category
// @Tags         categories
// @Security     BearerAuth
// @Param        id path string true "Category ID"
// @Success      204
// @Failure      401 {object} handler.errorResponse
// @Failure      404 {object} handler.errorResponse
// @Router       /categories/{id} [delete]
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	UserID := r.Context().Value(middleware.UserIDKey).(string)
	categoryID := chi.URLParam(r, "id")

	if err := h.categoryService.Delete(r.Context(), UserID, categoryID); err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
