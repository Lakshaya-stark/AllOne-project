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
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
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

func (s *Service) Login(ctx context.Context, req LoginRequest) error {

	user, err := s.repo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)
}