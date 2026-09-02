package application

import (
	"context"

	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/domain"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(
	repo domain.UserRepository,
) *UserService {

	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUser(
	ctx context.Context,
	id int64,
) (*domain.User, error) {

	return s.repo.GetByID(
		ctx,
		id,
	)
}
