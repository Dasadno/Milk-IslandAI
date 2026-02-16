// Package world provides the simulation orchestrator and world management.
//
// The Orchestrator is the central coordinator of the AI Agent Society simulation.
// It manages the world tick loop, coordinates agent actions, resolves conflicts,
// and maintains global world state. Think of it as the "game master".

package world

import (
	"context"
	"sync"
	"time"
)

// -----------------------------------------------------------------------------
// Orchestrator — центральный координатор симуляции
// -----------------------------------------------------------------------------
// Управляет главным циклом симуляции. Каждый тик:
//   1. TimeManager.Advance() — продвигает часы
//   2. Собирает WorldContext для каждого агента
//   3. Вызывает agent.Tick() для каждого активного агента (параллельно)
//   4. resolveActions() — разрешает конфликты и сопоставляет взаимодействия
//   5. EventBus.processEvents() — доставляет события подписчикам
//   6. Периодические задачи (консолидация памяти, затухание эмоций)
//   7. Сохраняет снапшоты состояния с заданным интервалом

type Orchestrator struct {
	// eventBus — pub-sub система для коммуникации между агентами и системой.
	EventBus *EventBus

	// timeManager — управление часами симуляции (тики, скорость, пауза).
	TimeManager *TimeManager

	// state — текущее состояние мира (тик, скорость, пауза, кол-во агентов).
	State WorldState

	// mu — RWMutex для потокобезопасного доступа к agents и state.
	// Read-lock для GetAgent/ListAgents, write-lock для Register/Remove/Tick.
	mu sync.RWMutex

	// ctx и cancel — контекст для graceful shutdown.
	// Stop() вызывает cancel(), завершая все горутины.
	ctx    context.Context
	cancel context.CancelFunc
}

// WorldState — глобальное состояние симуляции.
// Сериализуется в таблицу world_state (key-value) при сохранении.
type WorldState struct {
	// IsPaused — на паузе ли симуляция. При запуске = true.
	IsPaused bool

	// CurrentTick — текущий тик (монотонно возрастающий счётчик).
	CurrentTick int64

	// Speed — множитель скорости симуляции.
	// 1.0 = реальное время, 2.0 = удвоенная скорость, 0.5 = замедление.
	Speed float64

	// ActiveAgents — количество активных (is_active=true) агентов.
	ActiveAgents int

	// StartedAt — когда симуляция была запущена (для расчёта uptime).
	StartedAt time.Time
}

// -----------------------------------------------------------------------------
// OrchestratorDeps — зависимости для создания Orchestrator
// -----------------------------------------------------------------------------
// Передаются в NewOrchestrator(). Dependency injection для тестируемости.

type OrchestratorDeps struct {
	// EventBus — система событий (должна быть уже инициализирована).
	EventBus *EventBus

	// TimeManager — менеджер времени (должен быть уже инициализирован).
	TimeManager *TimeManager
}

// -----------------------------------------------------------------------------
// ActionResult — результат обработки действия агента оркестратором
// -----------------------------------------------------------------------------
// После resolveActions() каждое действие получает результат:
// успех, отказ (цель занята), или модификация (конфликт разрешён компромиссом).

type ActionResult struct {
	// AgentID — ID агента, чьё действие было обработано.
	AgentID string

	// Success — выполнено ли действие.
	Success bool

	// Reason — причина отказа или модификации ("target agent is busy", "conflict resolved").
	Reason string

	// InteractionResult — результат взаимодействия (если действие было interact).
	InteractionResult *InteractionResult
}

// InteractionResult — результат взаимодействия двух агентов.
type InteractionResult struct {
	// Agent1ID, Agent2ID — участники.
	Agent1ID string
	Agent2ID string

	// Dialogue — лог диалога (пары реплик).
	Dialogue []DialogueTurn

	// RelationshipDelta — как изменились отношения (-1.0 .. +1.0).
	// Положительный → укрепление, отрицательный → ослабление.
	RelationshipDelta float64

	// EmotionalImpact — как взаимодействие повлияло на эмоции каждого агента.
	EmotionalImpact map[string]string // agentID → emotion label
}

// DialogueTurn — одна реплика в диалоге между агентами.
type DialogueTurn struct {
	// SpeakerID — кто говорит.
	SpeakerID string

	// Content — текст реплики.
	Content string

	// Emotion — эмоция говорящего в момент реплики.
	Emotion string

	// Timestamp — время реплики в симуляции.
	Timestamp time.Time
}
