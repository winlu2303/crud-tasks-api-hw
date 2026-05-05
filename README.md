# REST API-service (CRUD tasks)
*Простое REST API для управления задачами на Go с in-memory хранилищем*

## Особенности
1. **Высокая производительность** - In-memory хранение с потокобезопасностью
2. **Полный CRUD** - Создание, чтение, обновление и удаление задач
3. **Потокобезопасность** - Использование `sync.RWMutex` для конкурентного доступа
4. **Логирование** - Автоматическое логирование всех запросов
5. **Чистая архитектура** - Разделение на слои (handlers, storage, models)
6. **Zero dependencies** - Только стандартная библиотека Go

## Требования
- Go 1.21 или выше
- Git

## Структура проекта 
```
crud-tasks-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handlers/
│   │   └── tasks.go
│   ├── models/
│   │   └── task.go
│   ├── storage/
│   │   ├── storage.go
│   │   └── memory.go
│   └── middleware/
│       └── logger.go
├── go.mod
├── go.sum
├── Dockerfile
├── openapi.yaml
├── .gitignore
├── README.md
└── manual-tests.md
```

## Запуск и установка

### Клонирование репозитория
```bash
git clone https://github.com/winlu2303/crud-tasks-api.git
cd crud-tasks-api
```
### Установка зависимостей
```bash
go mod download
```
### Запуск сервера
```bash
go run cmd/server/main.go
```

## Проверка работы
```
# Health check
curl http://localhost:8080/health

# Ожидаемый ответ:
{"status":"ok"}
```
## API Эндпоинты
|   Метод   |	Эндпоинт    |	Описание              |	  Код успеха  |
|:---------:|:-------------:|:-----------------------:|:-------------:|
|   GET     |	/tasks      |	Получить все задачи   |	    200 OK    |
|:---------:|:-------------:|:-----------------------:|:-------------:|
|   POST    |	/tasks	    |   Создать новую задачу  |    201 Created|
|:---------:|:-------------:|:-----------------------:|:-------------:|
|   GET     |	/tasks/{id}	|   Получить задачу по ID |	    200 OK    |
|:---------:|:-------------:|:-----------------------:|:-------------:|
|   PUT     |	/tasks/{id}	|   Обновить задачу	      |     200 OK    |
|:---------:|:-------------:|:-----------------------:|:-------------:|
|   DELETE  |	/tasks/{id}	|   Удалить задачу        |	204 No Content|
|:---------:|:-------------:|:-----------------------:|:-------------:|
|   GET     |	/health	    |Проверка здоровья сервера| 	200 OK    |
|:---------:|:-------------:|:-----------------------:|:-------------:|

## Модель данных
```json
{
  "id": 1,
  "title": "Купить продукты",
  "done": false,
  "created_at": "2024-01-15T10:30:00+03:00"
}
```

## Примеры использования
### GET /tasks - Получить все задачи
```bash
curl -X GET http://localhost:8080/tasks
```

### POST /tasks - Создать задачу
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Сходить в Озон, и забрать товары"}'
```

### GET /tasks/{id} - Получить задачу
```bash
curl -X GET http://localhost:8080/tasks/1
```

### PUT /tasks/{id} - Обновить задачу
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Купить кофий и сметану", "done": true}'
```

### DELETE /tasks/{id} - Удалить задачу
```bash
curl -X DELETE http://localhost:8080/tasks/1
```
---
## Тестирование 
### Полный тестовый сценарий
```bash
#!/bin/bash

echo " Тестирование REST API-service (CRUD tasks)"
echo "==========================================="

# 1. Health check
echo -e "\n1 Health check:"
curl -s http://localhost:8080/health | jq

# 2. Создание задачи
echo -e "\n2 Создание задачи:"
curl -s -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Тестовая задача"}' | jq

# 3. Получение всех задач
echo -e "\n3 Все задачи:"
curl -s http://localhost:8080/tasks | jq

# 4. Получение задачи по ID
echo -e "\n4 Задача с ID=1:"
curl -s http://localhost:8080/tasks/1 | jq

# 5. Обновление задачи
echo -e "\n5 Обновление задачи:"
curl -s -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Обновленная задача", "done": true}' | jq

# 6. Удаление задачи
echo -e "\n6 Удаление задачи:"
curl -s -X DELETE http://localhost:8080/tasks/1

# 7. Проверка после удаления
echo -e "\n7 Все задачи после удаления:"
curl -s http://localhost:8080/tasks | jq

echo -e "\n\(\checkmark \) Тестирование завершено!"
```
---
## Docker
```bash
# Сборка образа
docker build -t tasks-api .

# Запуск контейнера
docker run -p 8080:8080 tasks-api

# Проверить работу
curl http://localhost:8080/health
```