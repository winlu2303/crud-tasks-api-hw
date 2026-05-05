# Manual Tests

## Подготовка к тестированию

```bash
# Запуск сервера
go run cmd/server/main.go

# Сервер будет доступен на http://localhost:8080
```

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

## Результаты тестирования

|       Тест            |	Ожидаемый код   |	Фактический код	    |   Статус  |
|   Health check	    |       200         |   	200             |   PASS    |
|   Создание задачи	    |       201         |   	201             |   PASS    |
|  Получение всех задач |       200         |   	200             |   PASS    |
|   Получение по ID	    |       200         |   	200             |   PASS    |
|   Обновление задачи   |       200         |   	200             |   PASS    |
|   Удаление задачи	    |       204         |   	204             |   PASS    |
|   Без title	        |       400         |   	400             |   PASS    |
|   Неверный JSON	    |       400         |   	400             |   PASS    |
|   Задача не найдена	|       404         |   	404             |   PASS    |
|   Неверный I	        |       400         |   	400             |   PASS    |
| Неподдерживаемый метод|	    405         |   	405             |   PASS    |
