package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcosvfn/url-shortener/internal/infrastructure/http"
	"github.com/marcosvfn/url-shortener/internal/infrastructure/redis"
	"github.com/marcosvfn/url-shortener/internal/usecases"
)

func main() {
	redisRepo, err := redis.NewRedisRepository("redis:6379")
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	urlService := usecases.NewURLService(redisRepo)
	handler := http.NewURLHandler(urlService)

	router := mux.NewRouter()
	router.HandleFunc("/shorten", handler.ShortenURL).Methods("POST")
	router.HandleFunc("/{shortCode}", handler.RedirectURL).Methods("GET")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
