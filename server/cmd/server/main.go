package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
	"milk/server/cmd/server/middleware"
	"milk/server/data"
	"milk/server/internal/api"
	"milk/server/internal/storage"
)

func main() {
	data.DbConnection()

	repo := storage.NewRepository(data.Db)
	handler := api.NewHandler(repo)

	mux := middleware.NewMux(handler)

	h := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "OPTIONS", "DELETE"},
		AllowedHeaders: []string{"Content-type"},
		Debug:          true,
	}).Handler(mux)

	fmt.Println("server starting on :8080")
	if err := http.ListenAndServe(":8080", h); err != nil {
		fmt.Println("failed to connect server: ", err)
	}
}
