package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"queue_processing_service/internal/config"
	"queue_processing_service/internal/database"
	"queue_processing_service/internal/handler"
	"queue_processing_service/internal/repository"
	"queue_processing_service/internal/service"
	"queue_processing_service/internal/worker"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	// Инициализация баз данных
	db := database.InitPostgres(&cfg.Database)
	redisClient := database.InitRedis(&cfg.Redis)

	// Инициализация репозиториев
	postgresRepo := repository.NewPostgresRepository(db)
	redisRepo := repository.NewRedisRepository(redisClient)

	// Инициализация сервисов
	taskService := service.NewTaskService(postgresRepo, redisRepo)

	// Инициализация хендлеров
	taskHandler := handler.NewTaskHandler(taskService)

	// Инициализация воркера
	queueWorker := worker.NewQueueWorker(taskService)

	// Настройка роутера
	router := gin.Default()

	// API роуты
	api := router.Group("/api/v1")
	{
		api.POST("/tasks", taskHandler.CreateTask)
		api.GET("/tasks", taskHandler.GetAllTasks)
		api.GET("/tasks/:id", taskHandler.GetTaskByID)
		api.GET("/queue/length", taskHandler.GetQueueLength)
	}

	// Запуск воркера в горутине
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go queueWorker.Start(ctx)

	// Настройка HTTP сервера
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("Server started on port %s", cfg.Server.Port)

	// Ожидание сигнала для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Остановка воркера
	queueWorker.Stop()

	// Graceful shutdown сервера
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}
