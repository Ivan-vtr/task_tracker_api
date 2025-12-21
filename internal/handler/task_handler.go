package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"task_tracker_api/internal/model"
	"task_tracker_api/internal/service"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(s service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := h.service.Create(r.Context(), &task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "encoding error", http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	fmt.Println(id)

	task, err := h.service.Get(r.Context(), id)
	fmt.Println(task)
	if err != nil {
		http.Error(w, "not found", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "encoding error", http.StatusInternalServerError)
		return
	}
}
