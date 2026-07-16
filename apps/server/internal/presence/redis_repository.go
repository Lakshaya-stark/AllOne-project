package presence

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(
	client *redis.Client,
) *RedisRepository {

	return &RedisRepository{
		client: client,
	}
}

func presenceKey(
	deviceID uuid.UUID,
) string {

	return "presence:device:" + deviceID.String()
}

func (r *RedisRepository) SetOnline(
	ctx context.Context,
	deviceID uuid.UUID,
) error {

	return r.client.Set(
		ctx,
		presenceKey(deviceID),
		"online",
		2*time.Minute,
	).Err()
}

func (r *RedisRepository) SetOffline(
	ctx context.Context,
	deviceID uuid.UUID,
) error {

	return r.client.Del(
		ctx,
		presenceKey(deviceID),
	).Err()
}

func (r *RedisRepository) IsOnline(
	ctx context.Context,
	deviceID uuid.UUID,
) (bool, error) {

	result, err := r.client.Exists(
		ctx,
		presenceKey(deviceID),
	).Result()

	return result == 1, err
}