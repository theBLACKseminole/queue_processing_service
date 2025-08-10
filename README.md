# Queue Processing Service

Простой Go сервер с чистой архитектурой для управления очередью задач. Позволяет добавлять задачи в очередь, обрабатывать их (с имитацией 5-секундного процесса) и уведомлять об удалении из очереди.

## 🏗️ Архитектура

Проект построен по принципам Clean Architecture:

- **`domain`** - бизнес-логика и модели данных
- **`repository`** - слой доступа к данным (PostgreSQL + Redis)
- **`service`** - бизнес-логика приложения
- **`handler`** - HTTP обработчики (Gin)
- **`worker`** - фоновый обработчик очереди
- **`config`** - конфигурация приложения
- **`database`** - инициализация подключений

## 🚀 Быстрый запуск с Docker

### Предварительные требования

- Docker
- Docker Compose

### 1. Клонирование репозитория

```bash
git clone <repository-url>
cd queue_processing_service
```

### 2. Запуск всех сервисов

```bash
docker-compose up -d
```

Эта команда запустит:
- **PostgreSQL** на порту 5432
- **Redis** на порту 6379
- **Go приложение** на порту 8080

### 3. Проверка статуса

```bash
docker-compose ps
```

### 4. Просмотр логов

```bash
# Все сервисы
docker-compose logs -f

# Только Go приложение
docker-compose logs -f app

# Только база данных
docker-compose logs -f postgres

# Только Redis
docker-compose logs -f redis
```

### 5. Остановка сервисов

```bash
docker-compose down
```

### 6. Полная очистка (включая данные)

```bash
docker-compose down -v
```



## 📡 API Endpoints

После запуска сервер будет доступен по адресу: `http://localhost:8080`

### Создание задачи
```bash
POST /api/v1/tasks
Content-Type: application/json

{
  "title": "Название задачи",
  "description": "Описание задачи"
}
```

### Получение всех задач
```bash
GET /api/v1/tasks
```

### Получение задачи по ID
```bash
GET /api/v1/tasks/{id}
```

### Получение длины очереди
```bash
GET /api/v1/queue/length
```

## 📁 Структура проекта

```
queue_processing_service/
├── cmd/
│   └── main.go                 # Точка входа приложения
├── internal/
│   ├── config/
│   │   └── config.go          # Конфигурация
│   ├── database/
│   │   └── database.go        # Инициализация БД
│   ├── domain/
│   │   └── task.go            # Модель задачи
│   ├── handler/
│   │   └── task_handler.go    # HTTP обработчики
│   ├── repository/
│   │   ├── postgres_repository.go  # PostgreSQL репозиторий
│   │   └── redis_repository.go     # Redis репозиторий
│   ├── service/
│   │   └── task_service.go    # Бизнес-логика
│   └── worker/
│       └── queue_worker.go    # Обработчик очереди
├── docker-compose.yml          # Docker Compose конфигурация
├── Dockerfile                  # Docker образ приложения
├── go.mod                      # Go модули
├── go.sum                      # Go зависимости
├── .gitignore                  # Git исключения
├── env.example                 # Пример переменных окружения
└── README.md                   # Документация
```

## 🔄 Жизненный цикл задачи

1. **Создание**: задача добавляется в PostgreSQL и помещается в Redis очередь
2. **Ожидание**: задача находится в статусе `pending`
3. **Обработка**: воркер забирает задачу из очереди, статус меняется на `processing`
4. **Имитация**: выполняется 5-секундная имитация обработки
5. **Завершение**: статус меняется на `completed`, задача удаляется из очереди

