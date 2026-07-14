package device

import (
	"context"

	"allone/server/internal/models"
)

type Repository interface {
	Register(ctx context.Context, device *models.Device) error
	GetByID(ctx context.Context, id string) (*models.Device, error)
	ListByUser(ctx context.Context, userID string) ([]models.Device, error)
}