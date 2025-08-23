package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"callisto/internal/request"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/", s.HelloWorldHandler)
	router.Get("/users", s.getAllUsers)

	router.Get("/health", s.healthHandler)

	return router
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	request.JSON(w, http.StatusOK, map[string]string{"message": "Hello World!"})
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	request.JSON(w, http.StatusOK, jsonResp)
}

func (s *Server) getAllUsers(w http.ResponseWriter, r *http.Request) {
	request.JSON(w, http.StatusOK, map[string]string{"message": "Hello User!"})
}
