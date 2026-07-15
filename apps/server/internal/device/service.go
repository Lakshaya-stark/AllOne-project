package device

import (
	"context"
	"time"

	"allone/server/internal/models"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Register(
	ctx context.Context,
	userID uuid.UUID,
	req RegisterRequest,
) error {

	device := &models.Device{
		ID:         uuid.New(),
		UserID:     userID,
		Name:       req.Name,
		Platform:   req.Platform,
		DeviceType: req.DeviceType,
		PublicKey:  req.PublicKey,
		CreatedAt:  time.Now(),
	}

	return s.repo.Register(ctx, device)
}

func (s *Service) ListDevices(
	ctx context.Context,
	userID uuid.UUID,
) ([]models.Device, error) {

	return s.repo.ListByUser(ctx, userID)
}