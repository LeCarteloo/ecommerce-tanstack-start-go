package users

import (
	"context"
	"errors"
	"testing"

	repo "github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/adapters/postgresql/sqlc"
	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/api/users/mocks"
	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/apperrors"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	ctx := context.Background()
	userID := pgtype.UUID{
		Bytes: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
	}

	t.Run("should return user when user exists", func(t *testing.T) {
		mockRepo := new(mocks.MockQuerier)
		userService := NewService(mockRepo)
		expectedUser := repo.GetUserByIDRow{
			ID:        userID,
			Username:  "user",
			Email:     "user@mail.com",
			Role:      "user",
			CreatedAt: pgtype.Timestamptz{},
		}

		mockRepo.On("GetUserByID", ctx, userID).Return(expectedUser, nil)

		user, err := userService.GetUserByID(ctx, userID)
		assert.NoError(t, err)

		if assert.NotNil(t, user) {
			assert.EqualValues(t, expectedUser, user)
		}
	})

	t.Run("should return not found error if not found", func(t *testing.T) {
		mockRepo := new(mocks.MockQuerier)
		userService := NewService(mockRepo)

		mockRepo.On("GetUserByID", ctx, userID).Return(repo.GetUserByIDRow{}, errors.New("not found"))

		user, err := userService.GetUserByID(ctx, userID)
		assert.Error(t, err, apperrors.ErrUserNotFound)
		assert.EqualValues(t, repo.GetUserByIDRow{}, user)
	})
}
