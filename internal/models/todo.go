package models

import "time"

// @Description Представляет задачу, принадлежащую пользователю.
type Todo struct {
	// Уникальный идентификатор задачи.
	ID          int       `json:"id" example:"1"`
	// Идентификатор пользователя-владельца задачи.
	UserID      int       `json:"user_id" example:"1"`
	// Заголовок задачи.
	Title       string    `json:"title" example:"Купить молоко"`
	// Подробное описание задачи.
	Description string    `json:"description" example:"Купить 2 литра молока в магазине"`
	// Флаг завершённости задачи.
	Done        bool      `json:"done" example:"false"`
	// Время создания задачи.
	CreatedAt   time.Time `json:"created_at" example:"2026-01-01T12:00:00Z"`
	// Время последнего обновления задачи.
	UpdatedAt   time.Time `json:"updated_at" example:"2026-01-02T15:30:00Z"`
}

// @Description Структура запроса для создания новой задачи.
type CreateTodoRequest struct {
	// Идентификатор пользователя, которому принадлежит новая задача.
	UserID      int    `json:"user_id" validate:"required" example:"1"`
	// Заголовок новой задачи.
	Title       string `json:"title" validate:"required,max=200" example:"Купить молоко"`
	// Описание новой задачи (опционально).
	Description string `json:"description" validate:"max=1000" example:"Купить 2 литра молока"`
	// Начальное состояние завершённости (по умолчанию false).
	Done        bool   `json:"done" example:"false"`
}

// @Description Структура запроса для частичного обновления задачи.
type UpdateTodoRequest struct {
	// Новый заголовок (опционально).
	Title       *string `json:"title" validate:"omitempty,max=200" example:"Купить хлеб"`
	// Новое описание (опционально).
	Description *string `json:"description" validate:"omitempty,max=1000" example:"Купить 1 багет"`
	// Новое состояние завершённости (опционально).
	Done        *bool   `json:"done" example:"true"`
}
