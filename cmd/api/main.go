package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"task_tracker_api/internal/config"
	"task_tracker_api/internal/repository"
	"task_tracker_api/internal/server"
	"task_tracker_api/internal/service"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	err := godotenv.Load("/home/flora/GolandProjects/task_tracker_api/.env")
	if err != nil {
		log.Println(".env file not found, using system environment")
	} else {
		fmt.Println(".env file found")
	}

	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: log/slog

	logger := setupLogger(cfg.Env)
	logger.Info("started task_tracker_api", slog.String("env", cfg.Env))
	logger.Debug("debug messages are enabled")

	// Todo: init storage: postgresql
	dsn := os.Getenv("DATABASE_URL")
	logger.Info(dsn)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	taskRepo := repository.NewTaskRepository(db)

	taskService := service.NewTaskService(taskRepo)

	srv := server.New(taskService, logger)

	srv.Start(os.Getenv("APP_PORT"))

}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return logger
}
