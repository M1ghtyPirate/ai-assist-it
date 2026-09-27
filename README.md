# ai-assist-it

Краткое описание
----------------
Простой демонстрационный HTTP-сервис на Go для управления пользователями и их задачами (todos). Предоставляет REST API для CRUD-операций над пользователями и задачами и автоматически сгенерированную документацию Swagger.

Установка
---------
Требования:
- Go 1.20+ (или актуальная версия Go)
- Git (опционально)

Установка инструмента для генерации Swagger (swag):
- Рекомендуемый способ (в локальный модульный кэш):
  ```powershell
  go install github.com/swaggo/swag/cmd/swag@latest
  ```

- Альтернатива (без установки):
  ```powershell
  go run github.com/swaggo/swag/cmd/swag@latest init -g main.go -o docs
  ```

Генерация документации (при наличии swag):
- ```powershell
  swag init -g main.go -o docs
  ```
или
- ```powershell
  go run github.com/swaggo/swag/cmd/swag@latest init -g main.go -o docs
  ```

В проекте уже присутствует файл tasks.vs.json с готовыми задачами для Visual Studio. В нём определены следующие задачи (Task Runner):

- Launch — запускает сервер командой: go run main.go
- Generate Swagger Docs — генерирует документацию командой: swag init -g main.go -o docs
- Install Swagger — устанавливает инструмент swag командой: go install github.com/swaggo/swag/cmd/swag@latest

Вы можете выполнять эти задачи через меню Tasks в Visual Studio или Task Runner.
После запуска задачи Launch сервер можно остановить комбинацией клавиш Ctrl+Break в окне выполнения задачи.


Список зависимостей
-------------------
Основные зависимости проекта (указаны с версиями из go.mod):

- github.com/go-chi/chi/v5 v5.0.10 — легковесный роутер HTTP для Go
- github.com/swaggo/http-swagger/v2 v2.0.2 — Swagger UI handler для Go
- github.com/swaggo/swag v1.16.2 — инструмент генерации Swagger документации (dev dependency)

В go.mod также присутствуют ряд indirect зависимостей, используемых внутри этих пакетов (go run/go build подтянет их автоматически).


Запуск
------
Запустить сервер в рабочей директории проекта:
```powershell
go run .
```
Сервер по умолчанию слушает http://localhost:8080
Swagger UI доступен по адресу: http://localhost:8080/swagger/

Структура проекта
-----------------
- main.go                    — точка входа, маршрутизация и подключение Swagger UI
- internal/handlers/         — HTTP-хендлеры (users, todos)
- internal/models/           — определения моделей и запросов
- docs/                      — сгенерированные swagger.json / swagger.yaml и docs.go

API
---
Base path: /api/v1

Пользователи (users):
- GET  /api/v1/users         — список пользователей
- GET  /api/v1/users/{id}    — получить пользователя по ID
- POST /api/v1/users         — создать пользователя
- PUT  /api/v1/users/{id}    — обновить пользователя
- DELETE /api/v1/users/{id}  — удалить пользователя

Задачи (todos):
- GET  /api/v1/todos         — список задач
- GET  /api/v1/todos/{id}    — получить задачу по ID
- POST /api/v1/todos         — создать задачу
- PUT  /api/v1/todos/{id}    — обновить задачу
- DELETE /api/v1/todos/{id}  — удалить задачу

Форматы ответов
---------------
Все успешные ответы обёрнуты в структуру:
```json
{
  "success": true,
  "data": <payload>
}
```
Ошибки возвращают:
```json
{
  "error": "error_code",
  "message": "Описание ошибки"
}
```

Примеры запросов (curl) и примерные ответы
-----------------------------------------
1) Создать пользователя
```bash
curl -s -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","name":"Иван Иванов"}'
```

Пример успешного ответа (201):
```json
{
  "success": true,
  "data": {
	"id": 1,
	"email": "user@example.com",
	"name": "Иван Иванов",
	"created_at": "2026-01-01T12:00:00Z"
  }
}
```

2) Получить список пользователей
```bash
curl -s http://localhost:8080/api/v1/users
```

Пример успешного ответа (200):
```json
{
  "success": true,
  "data": [
	{"id":1,"email":"user@example.com","name":"Иван Иванов","created_at":"2026-01-01T12:00:00Z"}
  ]
}
```

3) Создать задачу
```bash
curl -s -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"user_id":1,"title":"Купить молоко","description":"2 литра","done":false}'
```

Пример успешного ответа (201):
```json
{
  "success": true,
  "data": {
	"id": 1,
	"user_id": 1,
	"title": "Купить молоко",
	"description": "2 литра",
	"done": false,
	"created_at": "2026-01-01T12:00:00Z",
	"updated_at": "2026-01-01T12:00:00Z"
  }
}
```

4) Получить задачу по ID
```bash
curl -s http://localhost:8080/api/v1/todos/1
```

Пример успешного ответа (200):
```json
{
  "success": true,
  "data": { /* объект задачи как выше */ }
}
```

5) Ошибочный запрос (пример)
```bash
curl -s http://localhost:8080/api/v1/users/invalid
```

Пример ответа (400):
```json
{
  "error": "bad_request",
  "message": "Неверный ID"
}
```

Замечания
---------
- При изменении моделей или аннотаций нужно регенерировать docs командой swag init.
- Файлы в docs/ перезаписываются при каждой генерации; если вносите ручные правки, сохраняйте их отдельно.


Лицензия и контакты
--------------------
Задание по оформлению API-проекта и README при помощи ИИ-ассистента.

Ассистент: GitHub Copilot
