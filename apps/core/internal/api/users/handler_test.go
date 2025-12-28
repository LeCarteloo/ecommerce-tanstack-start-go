package users

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	repo "github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/adapters/postgresql/sqlc"
	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/api/users/mocks"
	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/apperrors"
	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/respond"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandlerGetUserByID(t *testing.T) {
	userID := pgtype.UUID{
		Bytes: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		Valid: true,
	}
	missingUserID := pgtype.UUID{
		Bytes: [16]byte{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		Valid: true,
	}

	setup := func() (*chi.Mux, *mocks.MockUserService) {
		mockService := new(mocks.MockUserService)
		userHandler := NewHandler(mockService)

		r := chi.NewRouter()
		r.Get("/users/{userId}", userHandler.GetUserByID)

		return r, mockService
	}

	t.Run("returns 200 and user JSON when user exists", func(t *testing.T) {
		expectedUser := repo.GetUserByIDRow{
			ID:        userID,
			Username:  "user",
			Email:     "",
			Role:      "user",
			CreatedAt: pgtype.Timestamptz{},
		}

		r, mockService := setup()

		mockService.On("GetUserByID", mock.Anything, userID).Return(expectedUser, nil)

		req := httptest.NewRequest(http.MethodGet, "/users/"+userID.String(), nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

		var response repo.GetUserByIDRow
		err := json.NewDecoder(rec.Body).Decode(&response)
		require.NoError(t, err)
		assert.EqualValues(t, expectedUser, response)

		mockService.AssertExpectations(t)
	})

	t.Run("returns 400 and error message for invalid userId format", func(t *testing.T) {
		r, mockService := setup()

		req := httptest.NewRequest(http.MethodGet, "/users/invalid-uuid", nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		var response respond.ErrorResponse
		err := json.NewDecoder(rec.Body).Decode(&response)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, response.Message, apperrors.ErrInvalidIdFormat.Error())

		mockService.AssertNotCalled(t, "GetUserByID", mock.Anything, mock.Anything)
	})

	t.Run("returns 404 and error message if user does not exist", func(t *testing.T) {
		r, mockService := setup()

		mockService.On("GetUserByID", mock.Anything, missingUserID).Return(repo.GetUserByIDRow{}, apperrors.ErrUserNotFound)

		req := httptest.NewRequest(http.MethodGet, "/users/"+missingUserID.String(), nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		var response respond.ErrorResponse
		err := json.NewDecoder(rec.Body).Decode(&response)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Contains(t, response.Message, apperrors.ErrUserNotFound.Error())
	})

	t.Run("returns 500 and error message if there is unexpected error", func(t *testing.T) {
		r, mockService := setup()

		mockService.On("GetUserByID", mock.Anything, userID).Return(repo.GetUserByIDRow{}, errors.New("database error"))

		req := httptest.NewRequest(http.MethodGet, "/users/"+userID.String(), nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		var response respond.ErrorResponse
		err := json.NewDecoder(rec.Body).Decode(&response)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, response.Message, apperrors.ErrUnexpected.Error())
	})
}
