package auth

import (
	"context"
	"time"

	"allone/server/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo Repository
	jwt  *JWTService
}

func NewService(repo Repository, jwt *JWTService) *Service {
	return &Service{
		repo: repo,
		jwt:  jwt,
	}
}
func (s *Service) Register(ctx context.Context, req RegisterRequest) error {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := &models.User{
		ID:           uuid.New(),
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return s.repo.CreateUser(ctx, user)
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {

	user, err := s.repo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, err
	}

	token, err := s.jwt.GenerateAccessToken(user.ID)

	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken: token,
	}, nil
}