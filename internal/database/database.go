package database

import (
	"context"
	"fmt"
	"log"
	"queue_processing_service/internal/config"
	"queue_processing_service/internal/domain"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(cfg *config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&domain.Task{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("PostgreSQL connected successfully")
	return db
}

func InitRedis(cfg *config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("Redis connected successfully")
	return client
}
