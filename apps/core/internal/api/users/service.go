package users

import (
	"context"
	"errors"

	repo "github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/adapters/postgresql/sqlc"
	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/api/users/dto"
	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/apperrors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

func NewService(r repo.Querier) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetUserByID(ctx context.Context, userId pgtype.UUID) (repo.GetUserByIDRow, error) {
	user, err := s.repo.GetUserByID(ctx, userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return repo.GetUserByIDRow{}, apperrors.ErrUserNotFound
		}

		return repo.GetUserByIDRow{}, err
	}

	return user, nil
}

func (s *Service) RegisterUser(ctx context.Context, params dto.CreateUserParams) (repo.RegisterUserRow, error) {
	user, err := s.repo.RegisterUser(ctx, repo.RegisterUserParams{
		Username: params.Username,
		Email:    params.Email,
		Password: params.Password,
	})
	if err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return repo.RegisterUserRow{}, apperrors.ErrUserAlreadyExists
			}
		}

		return repo.RegisterUserRow{}, err
	}

	return user, nil
}

type UserService interface {
	GetUserByID(ctx context.Context, userId pgtype.UUID) (repo.GetUserByIDRow, error)
	RegisterUser(ctx context.Context, params dto.CreateUserParams) (repo.RegisterUserRow, error)
}

type Service struct {
	repo repo.Querier
}
