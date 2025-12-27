package users

import (
	"context"
	"testing"

	repo "github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/adapters/postgresql/sqlc"
	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/api/users/mocks"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.MockQuerier)
	userID := pgtype.UUID{
		Bytes: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
	}

	t.Run("should return user when user exists", func(t *testing.T) {
		userService := NewUserService(mockRepo)
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
}
