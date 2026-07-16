package presence

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	redis    *RedisRepository
	postgres *PostgresRepository
}

func NewService(
	redis *RedisRepository,
	postgres *PostgresRepository,
) *Service {

	return &Service{
		redis:    redis,
		postgres: postgres,
	}
}

func (s *Service) SetOnline(
	ctx context.Context,
	deviceID uuid.UUID,
) error {

	return s.redis.SetOnline(
		ctx,
		deviceID,
	)
}

func (s *Service) SetOffline(
	ctx context.Context,
	deviceID uuid.UUID,
) error {

	if err := s.redis.SetOffline(
		ctx,
		deviceID,
	); err != nil {
		return err
	}

	return s.postgres.UpdateLastSeen(
		ctx,
		deviceID,
	)
}

func (s *Service) IsOnline(
	ctx context.Context,
	deviceID uuid.UUID,
) (bool, error) {

	return s.redis.IsOnline(
		ctx,
		deviceID,
	)
}