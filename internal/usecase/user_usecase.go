// internal/usecase/user_usecase.go
package usecase

import (
	"errors"
	"time"

	"happy_backend/internal/entities"
	"happy_backend/internal/repository"
	"happy_backend/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo             repository.UserRepository
	jwtSecret        string
	jwtRefreshSecret string
	jwtExpiry        time.Duration
	jwtRefreshExpiry time.Duration
}

func NewUserUsecase(
	r repository.UserRepository,
	jwtSecret, jwtRefreshSecret string,
	jwtExpiry, jwtRefreshExpiry time.Duration,
) *UserUsecase {
	return &UserUsecase{
		repo:             r,
		jwtSecret:        jwtSecret,
		jwtRefreshSecret: jwtRefreshSecret,
		jwtExpiry:        jwtExpiry,
		jwtRefreshExpiry: jwtRefreshExpiry,
	}
}
func (u *UserUsecase) GetByID(id string) (*entities.User, error) {
	return u.repo.GetByID(id)
}
func (u *UserUsecase) Secret() string {
	return u.jwtSecret
}
func (u *UserUsecase) Register(user *entities.User) (*entities.User, string, string, error) {
	existing, _ := u.repo.GetByEmail(user.Email)
	if existing != nil {
		return nil, "", "", errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", "", errors.New("failed to hash password")
	}
	user.Password = string(hashed)

	if err := u.repo.Create(user); err != nil {
		return nil, "", "", err
	}

	// Generate tokens right after successful signup
	accessToken, err := jwt.GenerateToken(u.jwtSecret, u.jwtExpiry, user.ID)
	if err != nil {
		return nil, "", "", errors.New("failed to create access token")
	}
	refreshToken, err := jwt.GenerateToken(u.jwtRefreshSecret, u.jwtRefreshExpiry, user.ID)
	if err != nil {
		return nil, "", "", errors.New("failed to create refresh token")
	}

	return user, accessToken, refreshToken, nil
}

func (u *UserUsecase) Login(email, password string) (*entities.User, string, string, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil || user == nil {
		return nil, "", "", errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", "", errors.New("invalid credentials")
	}

	accessToken, err := jwt.GenerateToken(u.jwtSecret, u.jwtExpiry, user.ID)
	if err != nil {
		return nil, "", "", errors.New("failed to create access token")
	}
	refreshToken, err := jwt.GenerateToken(u.jwtRefreshSecret, u.jwtRefreshExpiry, user.ID)
	if err != nil {
		return nil, "", "", errors.New("failed to create refresh token")
	}

	return user, accessToken, refreshToken, nil
}

func (u *UserUsecase) RefreshAccessToken(refreshToken string) (string, error) {
	claims, err := jwt.ValidateToken(refreshToken, u.jwtRefreshSecret)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}
	userID, _ := claims["user_id"].(string)
	if userID == "" {
		return "", errors.New("invalid refresh token payload")
	}
	newAccessToken, err := jwt.GenerateToken(u.jwtSecret, u.jwtExpiry, userID)
	if err != nil {
		return "", errors.New("failed to generate access token")
	}
	return newAccessToken, nil
}
