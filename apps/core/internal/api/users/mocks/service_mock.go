package mocks

import (
	"context"

	repo "github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/adapters/postgresql/sqlc"
	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/api/users/dto"
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

func (m *MockUserService) RegisterUser(ctx context.Context, params dto.CreateUserParams) (repo.RegisterUserRow, error) {
	args := m.Called(ctx, params)

	return args.Get(0).(repo.RegisterUserRow), args.Error(1)
}
