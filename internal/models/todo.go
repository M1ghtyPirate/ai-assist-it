package models

import "time"

// Представляет задачу, принадлежащую пользователю.
type Todo struct {
	// Уникальный идентификатор задачи.
	ID          int       `json:"id"`
	// Идентификатор пользователя-владельца задачи.
	UserID      int       `json:"user_id"`
	// Заголовок задачи.
	Title       string    `json:"title"`
	// Подробное описание задачи.
	Description string    `json:"description"`
	// Флаг завершённости задачи.
	Done        bool      `json:"done"`
	// Время создания задачи.
	CreatedAt   time.Time `json:"created_at"`
	// Время последнего обновления задачи.
	UpdatedAt   time.Time `json:"updated_at"`
}

// Структура запроса для создания новой задачи.
type CreateTodoRequest struct {
	// Идентификатор пользователя, которому принадлежит новая задача.
	UserID      int    `json:"user_id" validate:"required"`
	// Заголовок новой задачи.
	Title       string `json:"title" validate:"required,max=200"`
	// Описание новой задачи (опционально).
	Description string `json:"description" validate:"max=1000"`
	// Начальное состояние завершённости (по умолчанию false).
	Done        bool   `json:"done"`
}

// Структура запроса для частичного обновления задачи.
type UpdateTodoRequest struct {
	// Новый заголовок (опционально).
	Title       *string `json:"title" validate:"omitempty,max=200"`
	// Новое описание (опционально).
	Description *string `json:"description" validate:"omitempty,max=1000"`
	// Новое состояние завершённости (опционально).
	Done        *bool   `json:"done"`
}
