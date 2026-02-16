// Package world provides the TimeManager for simulation clock control.
//
// TimeManager управляет понятием времени в симуляции. Время симуляции
// отделено от wall-clock — его можно ускорить, замедлить, поставить на паузу.

package world

import (
	"sync"
	"time"
)

// -----------------------------------------------------------------------------
// TimeManager — менеджер часов симуляции
// -----------------------------------------------------------------------------
// Генерирует тики с заданным интервалом. Оркестратор слушает TickChannel()
// и запускает обработку каждого тика. Поддерживает:
//   - Ускорение/замедление (SetSpeed)
//   - Пауза/возобновление (Pause/Resume)
//   - Пошаговое выполнение (Step) — один тик за раз
//   - Планирование колбэков (ScheduleAt/ScheduleEvery)

type TimeManager struct {
	// currentTick — монотонно возрастающий счётчик тиков.
	CurrentTick int64

	// tickDuration — базовая длительность одного тика (например, 1 секунда).
	// Реальный интервал = tickDuration / speedMultiplier.
	TickDuration time.Duration

	// speedMultiplier — множитель скорости.
	// 2.0 = тики в 2 раза чаще, 0.5 = в 2 раза реже.
	// Ограничен диапазоном 0.1–10.0.
	SpeedMultiplier float64

	// isPaused — приостановлена ли генерация тиков.
	IsPaused bool

	// startTime — когда симуляция была запущена (wall-clock).
	StartTime time.Time

	// simTime — текущее время симуляции (может отличаться от wall-clock).
	SimTime time.Time

	// ticker — системный тикер, генерирующий события с интервалом tickDuration/speed.
	Ticker *time.Ticker

	// scheduledOnce — колбэки, запланированные на конкретный тик.
	// Ключ = номер тика, значение = список функций для вызова.
	ScheduledOnce map[int64][]func()

	// scheduledRecurring — повторяющиеся колбэки.
	// Вызываются каждые N тиков (для рефлексии, статистики, случайных событий).
	ScheduledRecurring []RecurringTask

	// mu — мьютекс для потокобезопасного доступа.
	mu sync.Mutex
}

// TickInfo — информация о текущем тике, передаётся через TickChannel.
// Оркестратор использует эти данные для формирования WorldContext.
type TickInfo struct {
	// Tick — номер текущего тика.
	Tick int64

	// SimTime — время симуляции в момент тика.
	SimTime time.Time

	// Delta — сколько симуляционного времени прошло с прошлого тика.
	Delta time.Duration

	// WallTime — реальное время (wall-clock) в момент тика.
	WallTime time.Time
}

// RecurringTask — повторяющаяся задача, выполняемая каждые Interval тиков.
type RecurringTask struct {
	// Interval — через сколько тиков повторять (например, 100 = каждые 100 тиков).
	Interval int64

	// LastRun — тик последнего выполнения.
	LastRun int64

	// Callback — функция для вызова.
	Callback func()

	// Label — описание задачи для логирования ("memory_consolidation", "mood_decay").
	Label string
}
