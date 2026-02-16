package middleware

import (
	"net/http"

	"milk/server/internal/api"
)

func TODO(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("endpoint in work"))
}

func NewMux(h *api.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	// AGENTS
	mux.HandleFunc("GET /agents", h.ListAgents)
	mux.HandleFunc("GET /agents/{id}/memory", TODO)
	mux.HandleFunc("GET /agents/{id}/thoughts", TODO)
	mux.HandleFunc("GET /agents/{id}", h.GetAgent)
	mux.HandleFunc("POST /agents/{id}/inject", TODO)

	// RELATIONSHIPS
	mux.HandleFunc("GET /relationships", TODO)
	mux.HandleFunc("GET /relationships/{agentId}", TODO)
	mux.HandleFunc("POST /relationships", TODO)

	// EVENTS
	mux.HandleFunc("GET /events", TODO)
	mux.HandleFunc("POST /events", TODO)
	mux.HandleFunc("GET /events/stream", TODO)

	// WORLD
	mux.HandleFunc("GET /world/status", h.GetWorldStatus)
	mux.HandleFunc("POST /world/control", TODO)
	mux.HandleFunc("GET /world/statistics", TODO)

	// CONTROL PANEL
	mux.HandleFunc("POST /control/spawn", h.SpawnAgent)
	mux.HandleFunc("DELETE /control/agents/{id}", h.DeactivateAgentHandler)
	mux.HandleFunc("POST /control/reset", TODO)

	return mux
}
