package models

import "time"

// @Description Представляет пользователя приложения.
type User struct {
	// Уникальный идентификатор пользователя.
	ID        int       `json:"id"`
	// Электронная почта пользователя.
	Email     string    `json:"email" example:"user@example.com"`
	// Отображаемое имя пользователя.
	Name      string    `json:"name" example:"Иван Иванов"`
	// Время создания записи пользователя.
	CreatedAt time.Time `json:"created_at" example:"2026-01-01T12:00:00Z"`
}

// @Description Структура запроса для создания нового пользователя.
type CreateUserRequest struct {
	// Электронная почта нового пользователя.
	Email string `json:"email" validate:"required,email" example:"user@example.com"`
	// Отображаемое имя нового пользователя.
	Name  string `json:"name" validate:"required,max=100" example:"Иван Иванов"`
}

// @Description Структура запроса для частичного обновления пользователя.
type UpdateUserRequest struct {
	// Новая электронная почта (опционально).
	Email *string `json:"email" validate:"omitempty,email" example:"new@example.com"`
	// Новое отображаемое имя (опционально).
	Name  *string `json:"name" validate:"omitempty,max=100" example:"Пётр Петров"`
}

// @Description Стандартная структура ошибки, возвращаемая обработчиками.
type ErrorResponse struct {
	// Краткий машинно-читаемый код или имя ошибки.
	Error   string `json:"error" example:"bad_request"`
	// Подробное человекочитаемое описание ошибки.
	Message string `json:"message" example:"Невалидный запрос"`
}

// @Description Универсальная обёртка успешного ответа.
type SuccessResponse struct {
	// Флаг успешности операции.
	Success bool        `json:"success" example:"true"`
	// Полезная нагрузка ответа (опционально).
	Data    interface{} `json:"data,omitempty"`
}
