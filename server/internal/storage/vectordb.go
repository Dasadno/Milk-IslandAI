// Package storage provides the Vector Database layer for episodic memory.
//
// Обёртка над Chromem-go (или кастомной реализацией) для similarity-based
// поиска воспоминаний. Эпизодические воспоминания хранятся как векторные
// эмбеддинги — агенты вспоминают релевантный опыт, а не просто последний.

package storage

import "time"

// -----------------------------------------------------------------------------
// VectorStore — хранилище векторных эмбеддингов
// -----------------------------------------------------------------------------
// Каждый агент получает отдельную коллекцию (изолированное пространство памяти).
// При Recall() запрос превращается в эмбеддинг через EmbeddingProvider,
// затем ищутся ближайшие соседи (cosine similarity) среди воспоминаний агента.

type VectorStore struct {
	// Embedder — провайдер генерации эмбеддингов.
	// Может быть GigaChat/Ollama (quality) или LocalEmbedder (speed).
	Embedder EmbeddingProvider

	// StoragePath — путь к файлу/директории персистентного хранения Chromem-go.
	StoragePath string
}

// EmbeddingProvider — интерфейс генерации векторных представлений текста.
// Две реализации:
//   1. GigaChat/Ollama — через LLM API (выше качество, нужен сервер)
//   2. LocalEmbedder — TF-IDF/bag-of-words (быстрее, без зависимостей)
type EmbeddingProvider interface {
	// Embed — превращает текст в вектор фиксированной размерности.
	Embed(text string) ([]float32, error)

	// EmbedBatch — batch-версия для массовых операций (начальная загрузка).
	EmbedBatch(texts []string) ([][]float32, error)
}

// -----------------------------------------------------------------------------
// VectorMemory — воспоминание в формате vector store
// -----------------------------------------------------------------------------
// Отличается от MemoryRecord тем, что Content уже имеет эмбеддинг
// и metadata в формате string→string (ограничение Chromem-go).

type VectorMemory struct {
	// ID — уникальный идентификатор (совпадает с MemoryRecord.ID).
	ID string

	// Content — текст воспоминания, по которому строится эмбеддинг.
	Content string

	// EmotionalTag — эмоция в момент формирования (используется как фильтр).
	EmotionalTag string

	// Importance — значимость 0.0–1.0 (используется для порогового фильтра).
	Importance float64

	// Timestamp — когда воспоминание было сформировано.
	Timestamp time.Time

	// RelatedAgents — ID агентов-участников.
	RelatedAgents []string

	// Metadata — дополнительные данные в формате string→string.
	// Chromem-go принимает только string-значения в metadata.
	Metadata map[string]string
}

// VectorSearchResult — результат similarity search.
type VectorSearchResult struct {
	// Memory — найденное воспоминание.
	Memory VectorMemory

	// Similarity — cosine similarity score от 0.0 (не похоже) до 1.0 (идентично).
	// Используется как один из факторов ранжирования при Recall.
	Similarity float32
}

// MemoryFilter — фильтры для SearchWithFilter().
// Все поля опциональны — nil/zero = не фильтровать.
type MemoryFilter struct {
	// After — только воспоминания после этого времени.
	After *time.Time

	// Before — только воспоминания до этого времени.
	Before *time.Time

	// MinImportance — минимальный порог важности.
	MinImportance float64

	// EmotionalTag — фильтр по эмоции ("joy", "fear").
	EmotionalTag string

	// RelatedAgent — должен содержать этого агента в RelatedAgents.
	RelatedAgent string
}
