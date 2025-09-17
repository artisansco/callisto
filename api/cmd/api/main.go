package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"jinji/internal/database"
	"jinji/internal/env"
	"jinji/internal/server"
)

type config struct {
	smtp struct {
		host     string
		port     int
		username string
		password string
		from     string
	}
}

type application struct {
	config config
	port   int
	db     *database.Queries
	wg     sync.WaitGroup
}

func main() {
	var cfg config

	port := env.GetInt("PORT", 5000)
	dsn := env.GetString("DATABASE_URL", "jinji.db?_foreign_keys=on")
	cfg.smtp.host = env.GetString("SMTP_HOST", "example.smtp.host")
	cfg.smtp.port = env.GetInt("SMTP_PORT", 25)
	cfg.smtp.username = env.GetString("SMTP_USERNAME", "example_username")
	cfg.smtp.password = env.GetString("SMTP_PASSWORD", "password")
	cfg.smtp.from = env.GetString("SMTP_FROM", "Example Name <no_reply@example.org>")

	// mailer, err := smtp.NewMailer(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.from)
	// if err != nil {
	// 	return err
	// }

	db := server.NewDB(dsn)
	app := &application{
		config: cfg,
		port:   port,
		db:     db,
	}

	app.serve()
}

func (app *application) serve() error {
	serv := server.NewServer(app.db)

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go server.GracefulShutdown(serv, done)

	log.Printf("Running on port %s", serv.Addr)
	err := serv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")

	return nil
}
