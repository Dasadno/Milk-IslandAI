// Package gigachat provides a client for the GigaChat LLM API via Ollama.
//
// Клиент инкапсулирует всю коммуникацию с языковой моделью GigaChat.
// Два режима работы:
//   1. Ollama — локальный inference (рекомендуется для хакатона)
//   2. Direct — облачный GigaChat API (нужен API key)

package gigachat

import (
	"net/http"
	"time"
)

// -----------------------------------------------------------------------------
// Client — клиент для GigaChat/Ollama
// -----------------------------------------------------------------------------
// Предоставляет высокоуровневые методы для когнитивных задач агентов:
//   Think()     — размышление (system prompt + контекст + память → мысли)
//   Converse()  — генерация реплики диалога
//   Reflect()   — мета-когнитивная рефлексия
//   Summarize() — суммаризация воспоминаний для консолидации
//   Embed()     — генерация эмбеддингов для vector search

type Client struct {
	// BaseURL — endpoint сервера.
	// Ollama: "http://localhost:11434"
	// GigaChat: "https://gigachat.devices.sberbank.ru/api/v1"
	BaseURL string

	// Model — имя модели для запросов.
	// Ollama: имя модели в локальном реестре (например, "llama3", "mistral")
	// GigaChat: "GigaChat", "GigaChat-Pro", "GigaChat-Max"
	Model string

	// APIKey — ключ авторизации (только для Direct mode).
	// Для Ollama не используется.
	APIKey string

	// HTTPClient — HTTP-клиент с настроенными таймаутами.
	HTTPClient *http.Client

	// Mode — режим работы: Ollama или Direct.
	Mode ClientMode

	// Config — параметры по умолчанию для запросов.
	Config ClientConfig
}

// ClientMode — режим подключения к LLM.
type ClientMode string

const (
	// ModeOllama — локальный Ollama сервер. Бесплатно, без ограничений,
	// но требует запущенный Ollama с загруженной моделью.
	ModeOllama ClientMode = "ollama"

	// ModeDirect — прямой доступ к GigaChat API Сбера.
	// Требует API key и OAuth-авторизацию.
	ModeDirect ClientMode = "direct"
)

// ClientConfig — конфигурация LLM-клиента.
type ClientConfig struct {
	// DefaultTemperature — температура по умолчанию для генерации (0.0–2.0).
	// Низкая (0.1–0.3) = детерминированные, предсказуемые ответы.
	// Высокая (0.8–1.5) = креативные, разнообразные ответы.
	// Для агентов: базово 0.7, модифицируется BrainConfig.CreativityFactor.
	DefaultTemperature float64

	// MaxTokens — максимальное количество токенов в ответе.
	// Ограничивает длину генерации. Для Think() обычно 512–1024.
	MaxTokens int

	// Timeout — таймаут HTTP-запроса к LLM.
	// Рекомендуется 30s для обычных запросов, 60s для длинных рефлексий.
	Timeout time.Duration

	// RetryAttempts — количество повторов при transient failures (timeout, 5xx).
	RetryAttempts int

	// RetryDelay — задержка между повторами.
	RetryDelay time.Duration
}

// ClientOption — функциональная опция для настройки клиента.
type ClientOption func(*Client)

// -----------------------------------------------------------------------------
// CompletionRequest — запрос к LLM на генерацию текста
// -----------------------------------------------------------------------------
// Основной формат запроса. Brain формирует этот объект перед каждым Think().

type CompletionRequest struct {
	// SystemPrompt — «личность» агента.
	// Содержит описание характера, ценностей, причуд и формат ответа.
	// Остаётся стабильным между запросами одного агента.
	SystemPrompt string

	// Messages — история сообщений (контекст разговора).
	// Для Think(): один message с контекстом ситуации.
	// Для Converse(): история диалога (чередование user/assistant).
	Messages []Message

	// Temperature — переопределение DefaultTemperature для этого запроса.
	// nil = использовать дефолт из Config.
	Temperature *float64

	// MaxTokens — переопределение для этого запроса.
	MaxTokens *int

	// Stream — включить стриминг ответа (для real-time мыслей на дашборде).
	Stream bool
}

// Message — одно сообщение в контексте разговора (формат OpenAI/Ollama).
type Message struct {
	// Role — роль отправителя: "system", "user", "assistant".
	// system = системный промпт, user = входной контекст, assistant = ответ LLM.
	Role string `json:"role"`

	// Content — текст сообщения.
	Content string `json:"content"`
}

// CompletionResponse — ответ LLM на запрос генерации.
type CompletionResponse struct {
	// Content — сгенерированный текст.
	Content string

	// TokensUsed — общее количество потреблённых токенов (prompt + completion).
	// Используется для мониторинга расхода ресурсов.
	TokensUsed int

	// Model — какая модель ответила (полезно при использовании fallback-моделей).
	Model string

	// Duration — сколько времени занял запрос.
	Duration time.Duration
}

// -----------------------------------------------------------------------------
// StreamChunk — фрагмент стримингового ответа
// -----------------------------------------------------------------------------
// Читается из канала, возвращённого CompleteStream().
// Brain.StreamThoughts() пробрасывает чанки в SSE-эндпоинт.

type StreamChunk struct {
	// Content — фрагмент текста (может быть одно слово или часть предложения).
	Content string

	// Done — финальный чанк (после него канал закрывается).
	Done bool

	// Error — ошибка стрима. Не-nil = стрим прерван.
	Error error
}
