package service

import (
	"context"
	"log"
	"queue_processing_service/internal/domain"
	"queue_processing_service/internal/repository"
	"time"
)

type TaskService struct {
	postgresRepo *repository.PostgresRepository
	redisRepo    *repository.RedisRepository
}

func NewTaskService(postgresRepo *repository.PostgresRepository, redisRepo *repository.RedisRepository) *TaskService {
	return &TaskService{
		postgresRepo: postgresRepo,
		redisRepo:    redisRepo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, title, description string) (*domain.Task, error) {
	task := &domain.Task{
		Title:       title,
		Description: description,
		Status:      domain.StatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.postgresRepo.Create(task); err != nil {
		return nil, err
	}

	if err := s.redisRepo.PushToQueue(ctx, task); err != nil {
		return nil, err
	}

	log.Printf("Task created and added to queue: %s", title)
	return task, nil
}

func (s *TaskService) ProcessTask(ctx context.Context) error {
	task, err := s.redisRepo.PopFromQueue(ctx)
	if err != nil {
		return err
	}

	task.Status = domain.StatusProcessing
	task.UpdatedAt = time.Now()
	if err := s.postgresRepo.Update(task); err != nil {
		return err
	}

	log.Printf("Processing task: %s", task.Title)

	// Имитация долгой обработки
	time.Sleep(5 * time.Second)

	task.Status = domain.StatusCompleted
	task.UpdatedAt = time.Now()
	completedAt := time.Now()
	task.CompletedAt = &completedAt

	if err := s.postgresRepo.Update(task); err != nil {
		return err
	}

	log.Printf("Task completed: %s", task.Title)
	return nil
}

func (s *TaskService) GetAllTasks() ([]domain.Task, error) {
	return s.postgresRepo.GetAll()
}

func (s *TaskService) GetTaskByID(id uint) (*domain.Task, error) {
	return s.postgresRepo.GetByID(id)
}

func (s *TaskService) GetQueueLength(ctx context.Context) (int64, error) {
	return s.redisRepo.GetQueueLength(ctx)
}
