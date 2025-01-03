package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chiragthapa777/students-api/internal/config"
	student "github.com/chiragthapa777/students-api/internal/http/handlers"
	"github.com/chiragthapa777/students-api/internal/storage/sqlite"
)

func main() {
	// load config

	cfg := config.MustLoad()

	// database setup

	_, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal("Error while initializing db", err)
	}

	slog.Info("storage initialized")

	// setup router

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	// setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	log.Printf("Server started : %s \n", cfg.HTTPServer.Addr)

	// until the program receive interrupt signal from the os, for graceful shutdown

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatalf("Failed to start server: %s", err.Error())
		}
	}()

	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown server gracefully", slog.String("error : %s", err.Error()))
	}

}

// error with db package gcc required
