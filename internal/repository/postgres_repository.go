package repository

import (
	"queue_processing_service/internal/domain"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(task *domain.Task) error {
	return r.db.Create(task).Error
}

func (r *PostgresRepository) GetByID(id uint) (*domain.Task, error) {
	var task domain.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *PostgresRepository) Update(task *domain.Task) error {
	return r.db.Save(task).Error
}

func (r *PostgresRepository) GetAll() ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *PostgresRepository) GetPending() ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.db.Where("status = ?", domain.StatusPending).Find(&tasks).Error
	return tasks, err
}
