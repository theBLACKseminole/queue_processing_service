package worker

import (
	"context"
	"log"
	"queue_processing_service/internal/service"
)

type QueueWorker struct {
	taskService *service.TaskService
	stopChan    chan struct{}
}

func NewQueueWorker(taskService *service.TaskService) *QueueWorker {
	return &QueueWorker{
		taskService: taskService,
		stopChan:    make(chan struct{}),
	}
}

func (w *QueueWorker) Start(ctx context.Context) {
	log.Println("Queue worker started")

	for {
		select {
		case <-ctx.Done():
			log.Println("Queue worker stopped by context")
			return
		case <-w.stopChan:
			log.Println("Queue worker stopped")
			return
		default:
			if err := w.taskService.ProcessTask(ctx); err != nil {
				log.Printf("Error processing task: %v", err)
			}
		}
	}
}

func (w *QueueWorker) Stop() {
	close(w.stopChan)
}
