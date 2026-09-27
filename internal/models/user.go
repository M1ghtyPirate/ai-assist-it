package models

import "time"

// Представляет пользователя приложения.
type User struct {
	// Уникальный идентификатор пользователя.
	ID        int       `json:"id"`
	// Электронная почта пользователя.
	Email     string    `json:"email"`
	// Отображаемое имя пользователя.
	Name      string    `json:"name"`
	// Время создания записи пользователя.
	CreatedAt time.Time `json:"created_at"`
}

// Структура запроса для создания нового пользователя.
type CreateUserRequest struct {
	// Электронная почта нового пользователя.
	Email string `json:"email" validate:"required,email"`
	// Отображаемое имя нового пользователя.
	Name  string `json:"name" validate:"required,max=100"`
}

// Структура запроса для частичного обновления пользователя.
type UpdateUserRequest struct {
	// Новая электронная почта (опционально).
	Email *string `json:"email" validate:"omitempty,email"`
	// Новое отображаемое имя (опционально).
	Name  *string `json:"name" validate:"omitempty,max=100"`
}

// Стандартная структура ошибки, возвращаемая обработчиками.
type ErrorResponse struct {
	// Краткий машинно-читаемый код или имя ошибки.
	Error   string `json:"error"`
	// Подробное человекочитаемое описание ошибки.
	Message string `json:"message"`
}

// Универсальная обёртка успешного ответа.
type SuccessResponse struct {
	// Флаг успешности операции.
	Success bool        `json:"success"`
	// Полезная нагрузка ответа (опционально).
	Data    interface{} `json:"data,omitempty"`
}
