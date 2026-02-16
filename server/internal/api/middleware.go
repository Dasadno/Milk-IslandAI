// Package api provides middleware and error types for the HTTP server.
//
// Middleware-функции обрабатывают cross-cutting concerns:
// логирование, CORS, rate limiting, recovery от паник, трассировка запросов.

package api

// -----------------------------------------------------------------------------
// APIError — стандартизированная ошибка API
// -----------------------------------------------------------------------------
// Все ошибки API возвращаются в этом формате. ErrorHandler middleware
// перехватывает внутренние ошибки и оборачивает их в APIError,
// скрывая детали реализации в production-режиме.

type APIError struct {
	// Code — машино-читаемый код ошибки ("AGENT_NOT_FOUND", "BAD_REQUEST").
	// Фронтенд использует для i18n и специфичной обработки.
	Code string `json:"code"`

	// Message — человеко-читаемое описание ошибки.
	Message string `json:"message"`

	// Details — опциональные детали (валидационные ошибки, debug info).
	// В production — nil, в development — стек-трейс и контекст.
	Details any `json:"details,omitempty"`
}

// Error реализует интерфейс error для удобства использования в Go-коде.
func (e *APIError) Error() string {
	return e.Message
}

// Предопределённые коды ошибок.
const (
	ErrCodeBadRequest     = "BAD_REQUEST"      // 400 — невалидный JSON, отсутствуют поля
	ErrCodeNotFound       = "NOT_FOUND"        // 404 — агент/событие/связь не найдены
	ErrCodeConflict       = "CONFLICT"         // 409 — дублирование (например, связь уже есть)
	ErrCodeRateLimited    = "RATE_LIMITED"      // 429 — превышен лимит запросов
	ErrCodeInternalError  = "INTERNAL_ERROR"    // 500 — внутренняя ошибка сервера
)
