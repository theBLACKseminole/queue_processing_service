# Quick API Test for Queue Processing Service
# Run this after starting the services with: docker-compose up -d

Write-Host "🚀 Quick API Test" -ForegroundColor Green
Write-Host "=================" -ForegroundColor Green

# Wait for services to be ready
Write-Host "⏳ Waiting for services to be ready..." -ForegroundColor Yellow
Start-Sleep -Seconds 10

# Test 1: Create a task
Write-Host "`n📝 Creating a task..." -ForegroundColor Cyan
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/tasks" -Method POST -ContentType "application/json" -Body '{
        "title": "Test Task",
        "description": "Quick test task"
    }'
    Write-Host "✅ Task created successfully!" -ForegroundColor Green
    Write-Host "   ID: $($response.id)" -ForegroundColor White
    Write-Host "   Title: $($response.title)" -ForegroundColor White
    Write-Host "   Status: $($response.status)" -ForegroundColor White
} catch {
    Write-Host "❌ Failed to create task: $($_.Exception.Message)" -ForegroundColor Red
}

# Test 2: Get all tasks
Write-Host "`n📋 Getting all tasks..." -ForegroundColor Cyan
try {
    $tasks = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/tasks" -Method GET
    Write-Host "✅ Found $($tasks.Count) tasks" -ForegroundColor Green
    foreach ($task in $tasks) {
        Write-Host "   - $($task.title) (Status: $($task.status))" -ForegroundColor White
    }
} catch {
    Write-Host "❌ Failed to get tasks: $($_.Exception.Message)" -ForegroundColor Red
}

# Test 3: Get queue length
Write-Host "`n📊 Getting queue length..." -ForegroundColor Cyan
try {
    $queueLength = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/queue/length" -Method GET
    Write-Host "✅ Queue length: $($queueLength.length)" -ForegroundColor Green
} catch {
    Write-Host "❌ Failed to get queue length: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n🎉 Quick test completed!" -ForegroundColor Green
Write-Host "Check the logs with: docker-compose logs -f" -ForegroundColor Yellow
