package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"jinji/internal/database"
	"jinji/internal/env"
	"jinji/internal/routes"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultIdleTimeout    = time.Minute
	defaultReadTimeout    = 5 * time.Second
	defaultWriteTimeout   = 10 * time.Second
	defaultShutdownPeriod = 30 * time.Second
)

func NewServer(db *database.Queries) *http.Server {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", env.GetInt("PORT", 5000)),
		Handler:      routes.Routes(db),
		IdleTimeout:  defaultIdleTimeout,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	return server
}

func NewDB(dsn string) *database.Queries {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Hour)

	return database.New(db)
}

func GracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownPeriod)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}
