package application

import (
	"context"

	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/domain"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/password"
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


func (s *UserService) CreateUser(
	ctx context.Context,
	user *domain.User,
	rawPassword string,
) error {


	hash, err := password.Hash(
		rawPassword,
	)


	if err != nil {
		return err
	}


	user.PasswordHash = hash


	return s.repo.Create(
		ctx,
		user,
	)
}