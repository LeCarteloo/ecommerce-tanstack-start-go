package users

import (
	"context"

	repo "github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func NewUserService(r repo.Querier) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetUserByID(ctx context.Context, userID pgtype.UUID) (repo.GetUserByIDRow, error) {
	return repo.GetUserByIDRow{}, nil
}

type UserService interface {
	GetUserByID(ctx context.Context, userID pgtype.UUID) (repo.GetUserByIDRow, error)
}

type Service struct {
	repo repo.Querier
}
