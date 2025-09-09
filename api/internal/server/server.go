package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "modernc.org/sqlite"

	"callisto/internal/database"
)

type Application struct {
	port int
	db   *database.Queries
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	db, err := sql.Open("sqlite", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	app := &Application{
		port: port,
		db:   database.New(db),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.port),
		Handler:      app.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
