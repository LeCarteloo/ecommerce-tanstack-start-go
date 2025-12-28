package users

import (
	"net/http"

	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/apperrors"
	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/json"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type handler struct {
	service UserService
}

func NewHandler(s UserService) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userIdParam := chi.URLParam(r, "userId")

	var userId pgtype.UUID
	if err := userId.Scan(userIdParam); err != nil {
		http.Error(w, apperrors.ErrInvalidIdFormat.Error(), http.StatusBadRequest)
		return
	}

	user, _ := h.service.GetUserByID(r.Context(), userId)

	json.Write(w, http.StatusOK, user)
}
