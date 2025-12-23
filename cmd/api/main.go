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
	"task_tracker_api/internal/util"

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
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file not found, using system environment")
	} else {
		fmt.Println(".env file found")
	}

	cfg := config.MustLoad()

	logger := setupLogger(cfg.Env)
	logger.Debug("logger initialized")

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is empty")
	}

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close db: %v", err)
		}
	}()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// repositories (sqlx!)
	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	// jwt manager
	jwtManager := util.NewJWTManager(
		cfg.Auth.JWTSecret,
		cfg.Auth.AccessTokenTTL,
	)

	// services
	authService := service.NewAuthService(userRepo, jwtManager)
	taskService := service.NewTaskService(taskRepo)

	// server
	srv := server.New(
		taskService,
		authService,
		logger,
	)

	srv.Start(os.Getenv("APP_PORT"))
}

func setupLogger(env string) *slog.Logger {
	switch env {
	case envLocal:
		return slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		return slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		return slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		return slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}
}
