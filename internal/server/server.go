package server

import (
	"log/slog"
	"net/http"
	"task_tracker_api/internal/handler"
	"task_tracker_api/internal/service"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	taskService service.TaskService
	authService *service.AuthService
	logger      *slog.Logger
}

func New(taskService service.TaskService, authService *service.AuthService, logger *slog.Logger) *Server {
	return &Server{
		taskService: taskService,
		authService: authService,
		logger:      logger,
	}
}

func (s *Server) Start(port string) {
	s.logger.Info("http server started", slog.String("port", port))

	r := chi.NewRouter()

	taskHandler := handler.NewTaskHandler(s.taskService)
	authHandler := handler.NewAuthHandler(s.authService)

	r.Post("/auth/register", authHandler.Register)
	r.Post("/auth/login", authHandler.Login)

	r.Post("/tasks", taskHandler.Create)
	r.Get("/tasks/{id}", taskHandler.Get)
	r.Get("/tasks", taskHandler.GetAll)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		s.logger.Error("http server failed", slog.Any("error", err))
		panic(err)
	}

}
