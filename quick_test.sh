#!/bin/bash

# Quick API Test for Queue Processing Service
# Run this after starting the services with: docker-compose up -d

echo "🚀 Quick API Test"
echo "================="

# Wait for services to be ready
echo "⏳ Waiting for services to be ready..."
sleep 10

# Test 1: Create a task
echo -e "\n📝 Creating a task..."
if response=$(curl -s -X POST "http://localhost:8080/api/v1/tasks" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Test Task",
    "description": "Quick test task"
  }' 2>/dev/null); then
    echo "✅ Task created successfully!"
    echo "   Response: $response"
else
    echo "❌ Failed to create task"
fi

# Test 2: Get all tasks
echo -e "\n📋 Getting all tasks..."
if tasks=$(curl -s -X GET "http://localhost:8080/api/v1/tasks" 2>/dev/null); then
    echo "✅ Found tasks:"
    echo "$tasks" | jq -r '.[] | "   - \(.title) (Status: \(.status))"' 2>/dev/null || echo "$tasks"
else
    echo "❌ Failed to get tasks"
fi

# Test 3: Get queue length
echo -e "\n📊 Getting queue length..."
if queue_length=$(curl -s -X GET "http://localhost:8080/api/v1/queue/length" 2>/dev/null); then
    echo "✅ Queue length: $queue_length"
else
    echo "❌ Failed to get queue length"
fi

echo -e "\n🎉 Quick test completed!"
echo "Check the logs with: docker-compose logs -f"
