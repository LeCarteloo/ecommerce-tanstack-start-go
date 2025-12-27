package mocks

import (
	"context"

	repo "github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/mock"
)

type MockQuerier struct {
	mock.Mock
}

func (m *MockQuerier) GetUserByID(ctx context.Context, id pgtype.UUID) (repo.GetUserByIDRow, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(repo.GetUserByIDRow), args.Error(1)
}

var _ repo.Querier = (*MockQuerier)(nil)
