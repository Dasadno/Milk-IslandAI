// Package world provides the EventBus for inter-agent and world communication.
//
// EventBus — pub-sub система для развязанной коммуникации между агентами,
// оркестратором и внешними системами (API). Агенты не вызывают друг друга
// напрямую — они публикуют и подписываются на события.

package world

import (
	"sync"
	"time"
)

// -----------------------------------------------------------------------------
// EventBus — шина событий с топиковой маршрутизацией
// -----------------------------------------------------------------------------
// Паттерн publish-subscribe. Каждое событие имеет Topic, определяющий,
// каким подписчикам оно будет доставлено. AffectedAgents дополнительно
// фильтрует доставку — событие получают только указанные агенты.
//
// Буферизованная очередь (queue) обеспечивает неблокирующий Publish().
// Background горутина processEvents() вычитывает из очереди и раздаёт подписчикам.

type EventBus struct {
	// subscribers — маппинг топик → список подписчиков.
	// Каждый подписчик имеет канал доставки и опциональный фильтр.
	Subscribers map[EventTopic][]Subscriber

	// eventLog — история всех событий для API-эндпоинтов и отладки.
	EventLog []WorldEvent

	// queue — буферизованный канал для неблокирующего Publish().
	// processEvents() вычитывает и маршрутизирует.
	Queue chan WorldEvent

	// mu — RWMutex для потокобезопасного доступа к subscribers и eventLog.
	mu sync.RWMutex
}

// Subscriber — подписчик на события определённого топика.
type Subscriber struct {
	// ID — идентификатор подписчика (обычно agent ID или "dashboard").
	ID string

	// Channel — канал доставки событий. EventBus пишет, подписчик читает.
	Channel chan WorldEvent

	// Filter — опциональная функция фильтрации.
	// Если != nil, событие доставляется только если Filter возвращает true.
	// Используется для тонкой настройки — например, агент подписан на TopicGlobal,
	// но хочет только события типа "discovery".
	Filter func(WorldEvent) bool
}

// EventTopic — именованный канал маршрутизации событий.
type EventTopic string

const (
	// TopicGlobal — глобальные мировые события (катастрофы, открытия, праздники).
	// Все агенты обычно подписаны.
	TopicGlobal EventTopic = "global"

	// TopicInteraction — события взаимодействий между агентами.
	// Подписчики: участники взаимодействия + дашборд.
	TopicInteraction EventTopic = "interaction"

	// TopicMoodChange — события смены настроения агента.
	// Подписчики: дашборд (для обновления UI) + близкие агенты.
	TopicMoodChange EventTopic = "mood_change"

	// TopicGoalUpdate — события изменения целей агента.
	// Подписчики: дашборд.
	TopicGoalUpdate EventTopic = "goal_update"

	// TopicMemory — события формирования/консолидации воспоминаний.
	// Подписчики: дашборд + аналитика.
	TopicMemory EventTopic = "memory"

	// TopicRelationship — события изменения отношений.
	// Подписчики: дашборд (обновление графа) + участники.
	TopicRelationship EventTopic = "relationship"

	// TopicSystem — системные события (пауза, возобновление, сброс).
	// Подписчики: все агенты + дашборд.
	TopicSystem EventTopic = "system"
)

// -----------------------------------------------------------------------------
// WorldEvent — событие в мире симуляции
// -----------------------------------------------------------------------------
// Единый формат для всех событий: мировых, агентных, системных.
// Сохраняется в eventLog и таблице events в БД.

type WorldEvent struct {
	// ID — уникальный идентификатор события (UUID).
	ID string

	// Topic — топик маршрутизации (определяет, кто получит событие).
	Topic EventTopic

	// Type — конкретный тип события внутри топика.
	// TopicGlobal: "disaster", "celebration", "discovery", "weather_change"
	// TopicInteraction: "conversation", "debate", "conflict"
	// TopicSystem: "pause", "resume", "agent_joined", "agent_left"
	Type string

	// Source — кто/что сгенерировало событие.
	// Agent ID для агентных событий, "system" для системных, "api" для инъекций.
	Source string

	// AffectedAgents — список ID агентов, которым доставить событие.
	// Пустой список = broadcast всем подписчикам топика.
	AffectedAgents []string

	// Payload — произвольные данные события.
	// Структура зависит от Type: для "conversation" — лог диалога,
	// для "disaster" — описание и масштаб, для "mood_change" — old/new mood.
	Payload map[string]any

	// Timestamp — когда событие произошло (simulation time).
	Timestamp time.Time

	// Tick — номер тика симуляции, в котором событие произошло.
	Tick int64

	// Priority — приоритет доставки.
	// 0 = обычный, 1+ = повышенный (доставляется первым).
	// Системные события (pause/resume) имеют высокий приоритет.
	Priority int
}
