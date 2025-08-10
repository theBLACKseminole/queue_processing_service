# 🚀 Docker Quick Start Guide

## Быстрый запуск за 3 шага

### 1. Запуск всех сервисов
```bash
docker-compose up -d
```

### 2. Ожидание готовности (10 секунд)
```bash
# Windows
Start-Sleep -Seconds 10

# Linux/Mac
sleep 10
```

### 3. Тестирование API
```bash
# Windows
.\quick_test.ps1

# Linux/Mac
./quick_test.sh
```

## 📋 Что запускается

- **PostgreSQL** - база данных на порту 5432
- **Redis** - очередь задач на порту 6379  
- **Go приложение** - API сервер на порту 8080

## 🔍 Проверка статуса

```bash
# Статус всех сервисов
docker-compose ps

# Логи приложения
docker-compose logs -f app

# Логи базы данных
docker-compose logs -f postgres

# Логи Redis
docker-compose logs -f redis
```

## 🛑 Остановка

```bash
# Остановить все сервисы
docker-compose down

# Остановить и удалить данные
docker-compose down -v
```

## 🚨 Если что-то не работает

1. **Проверьте Docker:**
   ```bash
   docker --version
   docker-compose --version
   ```

2. **Проверьте порты:**
   ```bash
   # Windows
   netstat -an | findstr :8080
   
   # Linux/Mac
   netstat -an | grep :8080
   ```

3. **Перезапустите сервисы:**
   ```bash
   docker-compose restart
   ```

4. **Проверьте логи:**
   ```bash
   docker-compose logs
   ```

## 📱 API Endpoints

После запуска API доступен по адресу: `http://localhost:8080`

- `POST /api/v1/tasks` - создать задачу
- `GET /api/v1/tasks` - получить все задачи  
- `GET /api/v1/tasks/{id}` - получить задачу по ID
- `GET /api/v1/queue/length` - получить длину очереди

## 🎯 Пример использования

```bash
# Создать задачу
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Моя задача", "description": "Описание"}'

# Получить все задачи
curl http://localhost:8080/api/v1/tasks
```

## 📚 Подробная документация

См. [README.md](README.md) для полного описания проекта.
