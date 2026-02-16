// Package storage provides the persistence layer for the AI Agent Society.
//
// Этот файл управляет подключением к SQLite, миграциями схемы и CRUD-операциями
// для всех персистентных данных: агенты, связи, события, воспоминания, состояние мира.

package storage

import (
	"database/sql"
	"time"
)

// -----------------------------------------------------------------------------
// Repository — единая точка доступа к БД
// -----------------------------------------------------------------------------
// Все операции с SQLite проходят через Repository. Он:
//   - Управляет подключением (WAL mode для конкурентных чтений)
//   - Запускает миграции при старте (idempotent)
//   - Предоставляет типизированные CRUD-методы
//
// В конструкторе NewRepository() вызывается Migrate(), создающий таблицы
// и индексы, если они ещё не существуют.

type Repository struct {
	// DB — пул соединений database/sql поверх go-sqlite3.
	// WAL mode включается при инициализации для конкурентных чтений.
	DB *sql.DB
}

// -----------------------------------------------------------------------------
// AgentRecord — строка из таблицы agents
// -----------------------------------------------------------------------------
// Прямое отображение SQL-схемы. При чтении JSON-поля (personality, mood_state,
// goals, snapshot) десериализуются в соответствующие доменные структуры.

type AgentRecord struct {
	// ID — UUID, первичный ключ.
	ID string

	// Name — отображаемое имя агента.
	Name string

	// Personality — JSON-blob с чертами Big Five, ценностями и причудами.
	// Десериализуется в agent.Personality при загрузке.
	Personality string

	// MoodState — JSON-blob текущего PAD-состояния.
	// Может быть NULL (sql.NullString), если агент только что создан.
	MoodState sql.NullString

	// Goals — JSON-массив активных целей агента.
	Goals sql.NullString

	// State — текущее состояние: "idle", "thinking", "acting" и т.д.
	State string

	// IsActive — soft-delete флаг. false = агент деактивирован, но история сохранена.
	IsActive bool

	// CreatedAt — время создания агента.
	CreatedAt time.Time

	// LastActive — время последней активности.
	LastActive sql.NullTime

	// Snapshot — JSON-blob полного состояния для восстановления после рестарта.
	Snapshot sql.NullString
}

// AgentFilter — параметры фильтрации для ListAgents().
type AgentFilter struct {
	// IsActive — фильтр по активности. nil = все, true = только активные.
	IsActive *bool

	// State — фильтр по состоянию. Пустая строка = все.
	State string

	// Page и Limit — пагинация.
	Page  int
	Limit int
}

// AgentUpdate — частичное обновление агента для UpdateAgent().
// Поля-указатели: nil = не обновлять, значение = обновить.
type AgentUpdate struct {
	MoodState  *string
	Goals      *string
	State      *string
	IsActive   *bool
	LastActive *time.Time
	Snapshot   *string
}

// -----------------------------------------------------------------------------
// RelationshipRecord — строка из таблицы relationships
// -----------------------------------------------------------------------------

type RelationshipRecord struct {
	// ID — UUID связи.
	ID string

	// Agent1ID, Agent2ID — UUID участников (FK → agents.id).
	Agent1ID string
	Agent2ID string

	// Type — тип связи: "friend", "rival", "neutral", "romantic".
	Type string

	// Strength — сила связи: -1.0 (враждебность) до +1.0 (тесная связь).
	// Знак определяет характер, абсолютное значение — интенсивность.
	Strength float64

	// InteractionCount — сколько раз эти агенты взаимодействовали.
	InteractionCount int

	// LastInteraction — время последнего взаимодействия.
	LastInteraction sql.NullTime

	// Metadata — JSON-blob дополнительного контекста.
	Metadata sql.NullString
}

// GraphData — полный граф отношений для визуализации.
// Возвращается из GetRelationshipGraph().
type GraphData struct {
	Nodes []GraphNode
	Edges []GraphEdge
}

// GraphNode — узел графа (агент).
type GraphNode struct {
	ID              string
	Name            string
	PersonalityType string
	RelationCount   int
}

// GraphEdge — ребро графа (связь).
type GraphEdge struct {
	Source   string
	Target  string
	Type    string
	Strength float64
}

// -----------------------------------------------------------------------------
// MemoryRecord — строка из таблицы memories
// -----------------------------------------------------------------------------

type MemoryRecord struct {
	// ID — UUID воспоминания.
	ID string

	// AgentID — UUID агента-владельца (FK → agents.id).
	AgentID string

	// Type — тип памяти: "episodic", "semantic", "procedural".
	Type string

	// Content — текстовое содержимое воспоминания.
	Content string

	// EmotionalTag — эмоция в момент формирования: "joy", "fear", "trust" и т.д.
	EmotionalTag sql.NullString

	// Importance — значимость от 0.0 до 1.0.
	Importance float64

	// AccessCount — сколько раз воспоминание было извлечено (Recall).
	AccessCount int

	// LastAccessed — время последнего обращения.
	LastAccessed sql.NullTime

	// RelatedAgents — JSON-массив UUID агентов, вовлечённых в событие.
	RelatedAgents sql.NullString

	// Metadata — JSON-blob дополнительных данных.
	Metadata sql.NullString

	// CreatedAt — когда воспоминание было сформировано.
	CreatedAt time.Time
}

// -----------------------------------------------------------------------------
// EventRecord — строка из таблицы events
// -----------------------------------------------------------------------------

type EventRecord struct {
	// ID — UUID события.
	ID string

	// Topic — топик маршрутизации: "global", "interaction", "mood_change" и т.д.
	Topic string

	// Type — конкретный тип: "disaster", "celebration", "conversation".
	Type string

	// Source — источник: agent ID, "system", "api".
	Source string

	// AffectedAgents — JSON-массив UUID затронутых агентов.
	AffectedAgents sql.NullString

	// Payload — JSON-blob данных события.
	Payload sql.NullString

	// Status — статус: "pending", "active", "completed".
	Status string

	// Tick — номер тика симуляции.
	Tick sql.NullInt64

	// CreatedAt — время создания.
	CreatedAt time.Time
}

// EventFilter — параметры фильтрации для GetEvents().
type EventFilter struct {
	// Topic — фильтр по топику. Пустая строка = все.
	Topic string

	// Status — фильтр по статусу.
	Status string

	// Source — фильтр по источнику.
	Source string

	// Limit — максимальное количество результатов.
	Limit int
}
