package service

import (
	"context"
	"errors"
	"task_tracker_api/internal/model"

	"task_tracker_api/internal/repository"
	"task_tracker_api/internal/util"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService struct {
	users *repository.UserRepository
	jwt   *util.JWTManager
}

func NewAuthService(users *repository.UserRepository, jwt *util.JWTManager) *AuthService {
	return &AuthService{users: users, jwt: jwt}
}

func (s *AuthService) Register(ctx context.Context, email, password string) error {
	hash, err := util.HashPassword(password)
	if err != nil {
		return err
	}

	return s.users.Create(ctx, &model.User{
		Email:        email,
		PasswordHash: hash,
		Role:         "user",
	})
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.users.GetByEmail(ctx, email)
	if err != nil || user == nil { // проверяем и ошибку, и nil
		return "", ErrInvalidCredentials
	}

	if err := util.CheckPassword(user.PasswordHash, password); err != nil {
		return "", ErrInvalidCredentials
	}

	return s.jwt.Generate(int64(user.ID), user.Role)
}
