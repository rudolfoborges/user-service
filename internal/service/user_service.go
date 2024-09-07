package service

import (
	"context"
	"errors"

	"github.com/amandakeren/user-service/internal/entity"
	"github.com/amandakeren/user-service/internal/gateway/repository"
)

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Create(ctx context.Context, input CreateUserInput) error {
	exists, err := s.userRepository.ExistsByEmail(ctx, input.Email)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already exists")
	}

	user := entity.NewUser(input.Name, input.Email, input.Password)

	err = s.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
