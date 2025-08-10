package repository

import (
	"context"
	"encoding/json"
	"queue_processing_service/internal/domain"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{client: client}
}

func (r *RedisRepository) PushToQueue(ctx context.Context, task *domain.Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		return err
	}
	return r.client.LPush(ctx, "task_queue", data).Err()
}

func (r *RedisRepository) PopFromQueue(ctx context.Context) (*domain.Task, error) {
	result, err := r.client.BRPop(ctx, 0, "task_queue").Result()
	if err != nil {
		return nil, err
	}

	var task domain.Task
	err = json.Unmarshal([]byte(result[1]), &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *RedisRepository) GetQueueLength(ctx context.Context) (int64, error) {
	return r.client.LLen(ctx, "task_queue").Result()
}
