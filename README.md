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

## 🔧 Конфигурация

### Переменные окружения

Создайте файл `.env` на основе `env.example`:

```bash
# Server Configuration
SERVER_PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=queue_service
DB_SSLMODE=disable

# Redis Configuration
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
```

### Docker Compose конфигурация

Основные настройки в `docker-compose.yml`:

- **PostgreSQL**: база данных `queue_service`, пользователь `postgres`, пароль `password`
- **Redis**: без пароля, база данных 0
- **Go приложение**: порт 8080

## 🛠️ Разработка

### Локальный запуск (без Docker)

1. Установите Go 1.21+
2. Установите PostgreSQL и Redis
3. Создайте базу данных `queue_service`
4. Запустите:

```bash
go mod tidy
go run cmd/main.go
```

### Тестирование

```bash
# Тест API (PowerShell)
.\test_api.ps1

# Тест API (Bash)
./test_api.sh
```

## 📊 Мониторинг

### Проверка здоровья сервисов

```bash
# PostgreSQL
docker exec -it queue_processing_service-postgres-1 psql -U postgres -d queue_service -c "SELECT version();"

# Redis
docker exec -it queue_processing_service-redis-1 redis-cli ping
```

### Просмотр данных

```bash
# Подключение к PostgreSQL
docker exec -it queue_processing_service-postgres-1 psql -U postgres -d queue_service

# Подключение к Redis
docker exec -it queue_processing_service-redis-1 redis-cli
```

## 🚨 Устранение неполадок

### Проблемы с подключением к базе данных

1. Проверьте, что PostgreSQL запущен:
   ```bash
   docker-compose ps postgres
   ```

2. Проверьте логи:
   ```bash
   docker-compose logs postgres
   ```

3. Перезапустите сервис:
   ```bash
   docker-compose restart postgres
   ```

### Проблемы с Redis

1. Проверьте статус:
   ```bash
   docker-compose ps redis
   ```

2. Проверьте логи:
   ```bash
   docker-compose logs redis
   ```

### Проблемы с приложением

1. Проверьте логи приложения:
   ```bash
   docker-compose logs app
   ```

2. Перезапустите:
   ```bash
   docker-compose restart app
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

## 📝 Лицензия

MIT License
