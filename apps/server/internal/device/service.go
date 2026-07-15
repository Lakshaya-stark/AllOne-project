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
) (*models.Device, error) {

	device := &models.Device{
		ID:         uuid.New(),
		UserID:     userID,
		Name:       req.Name,
		Platform:   req.Platform,
		DeviceType: req.DeviceType,
		PublicKey:  req.PublicKey,
		CreatedAt:  time.Now(),
	}

	if err := s.repo.Register(ctx, device); err != nil {
		return nil, err
	}

	return device, nil
}

func (s *Service) ListDevices(
	ctx context.Context,
	userID uuid.UUID,
) ([]models.Device, error) {

	return s.repo.ListByUser(ctx, userID)
}