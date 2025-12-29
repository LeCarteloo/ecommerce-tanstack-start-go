package users

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/api/users/dto"
	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/apperrors"
	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/respond"
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
		respond.WriteError(w, http.StatusBadRequest, apperrors.ErrInvalidIdFormat.Error())
		return
	}

	user, err := h.service.GetUserByID(r.Context(), userId)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserNotFound) {
			respond.WriteError(w, http.StatusNotFound, err.Error())
			return
		}

		slog.Error("get user by id: unexpected failure",
			"user_id", userId.String(),
			"error", err,
		)
		respond.WriteError(w, http.StatusInternalServerError, apperrors.ErrUnexpected.Error())
		return
	}

	respond.Write(w, http.StatusOK, user)
}

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var decodedUserParams dto.CreateUserParams
	if err := respond.Read(r, &decodedUserParams); err != nil {
		respond.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.service.RegisterUser(r.Context(), decodedUserParams)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserAlreadyExists) {
			respond.WriteError(w, http.StatusConflict, err.Error())
			return
		}

		slog.Error("register user: unexpected failure",
			"email", decodedUserParams.Email,
			"error", err,
		)
		respond.WriteError(w, http.StatusInternalServerError, apperrors.ErrUnexpected.Error())
		return
	}

	respond.Write(w, http.StatusCreated, user)
}
