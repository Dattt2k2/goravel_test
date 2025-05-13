package services

import (
	"goravel/app/http/helpers"
	"goravel/app/models"
	"goravel/app/repositories"

	"errors"
)

type AuthService interface {
	Register(email, password string) (*models.User, string, string, error)
	SignIn(email, password string) (*models.User, string, string, error)
	SignOut() error
	RefreshToken(refreshToken string) (*models.User, string, string, error)
	FindUserById(id uint64) (*models.User, error)
}

type AuthServiceImpl struct {
	authRepo repositories.AuthRepository
}

func NewAuthService(authRepo repositories.AuthRepository) AuthService {
	return &AuthServiceImpl{
		authRepo: authRepo,
	}
}

func (s *AuthServiceImpl) Register(email, passowrd string) (*models.User, string, string, error) {
	user, err := s.authRepo.Register(email, passowrd)
	if err != nil {
		return nil, "", "", err
	}
	accessToken, refreshToken, err := helpers.GenerateTokens(*user)
	if err != nil {
		return nil, "", "", errors.New("failed to generate tokens")

	}
	return user, accessToken, refreshToken, nil
}

func (s *AuthServiceImpl) SignIn(email, password string) (*models.User, string, string, error) {
	user, err := s.authRepo.SignIn(email, password)
	if err != nil {
		return nil, "", "", err
	}
	accessToken, refreshToken, err := helpers.GenerateTokens(*user)
	if err != nil {
		return nil, "", "", errors.New("failed to generate tokens")
	}
	return user, accessToken, refreshToken, nil
}

// RefreshToken service refreshes the tokens using a refresh token
func (s *AuthServiceImpl) RefreshToken(refreshToken string) (*models.User, string, string, error) {
	// Repository trả về (accessToken, newRefreshToken, user, err)
	accessToken, newRefreshToken, user, err := s.authRepo.RefreshToken(refreshToken)
	if err != nil {
		return nil, "", "", err
	}

	// Controller mong đợi (user, accessToken, newRefreshToken, err)
	return user, accessToken, newRefreshToken, nil
}

func (s *AuthServiceImpl) FindUserById(id uint64) (*models.User, error) {
	return s.authRepo.FindUserById(id)
}

// SignOut xử lý đăng xuất người dùng
func (s *AuthServiceImpl) SignOut() error {
	// Gọi đến repository để thực hiện các thao tác cần thiết (nếu có)
	return s.authRepo.SignOut()
}
