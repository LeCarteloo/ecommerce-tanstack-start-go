package mocks

import (
	"context"

	repo "github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUserByID(ctx context.Context, userId pgtype.UUID) (repo.GetUserByIDRow, error) {
	args := m.Called(ctx, userId)

	return args.Get(0).(repo.GetUserByIDRow), args.Error(1)
}
