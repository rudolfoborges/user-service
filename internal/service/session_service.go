package service

import (
	"context"
	"errors"

	"github.com/amandakeren/user-service/internal/entity"
	"github.com/amandakeren/user-service/internal/gateway/repository"
	"github.com/google/uuid"
)

type LoginOutput struct {
	UserID       uuid.UUID
	UserName     string
	SessionToken string
}

type SessionService struct {
	userRepository    repository.UserRepository
	sessionRepository repository.SessionRepository
}

func NewSessionService(
	userRepository repository.UserRepository,
	sessionRepository repository.SessionRepository,
) *SessionService {
	return &SessionService{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func (s *SessionService) Login(ctx context.Context, email, password string) (*LoginOutput, error) {
	userExists, err := s.userRepository.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !userExists {
		return nil, errors.New("não existe um usuário com esse email")
	}

	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user.ComparePassword(password) != nil {
		return nil, errors.New("password inválido")
	}

	session := entity.NewSession(user.ID)
	err = s.sessionRepository.Create(ctx, session)
	if err != nil {
		return nil, err
	}

	sessionToken, err := session.GenerateToken()
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		UserID:       user.ID,
		UserName:     user.Name,
		SessionToken: sessionToken,
	}, nil
}
