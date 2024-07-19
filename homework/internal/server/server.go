package server

import (
	"github.com/cripplemymind9/brunoyam-vebinar3/homework/internal/domain/models"
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
)

type Repository interface {
	GetAllTasks() ([]models.Task, error)
	InsertTask(models.Task) error
	UpdateTask(models.Task, int) error
	DeleteTask(int) error
	GetTask(int) (models.Task, error)
}

type Server struct {
	addr 	string
	db 		Repository
}

func NewServer (addr string, db Repository) *Server {
	return &Server{
		addr: addr,
		db: 	db,
	}
}

func (s *Server) GetAllTasksHandler(res http.ResponseWriter, req *http.Request) {
	tasks, err := s.db.GetAllTasks()
	if err != nil {
		http.Error(res, "get tasks error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(res)
	if err := enc.Encode(tasks); err != nil {
		http.Error(res, "internal error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) SaveTaskHandler(res http.ResponseWriter, req *http.Request) {
	var task models.Task
	body := req.Body
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&task); err != nil {
		http.Error(res, "invalid argument", http.StatusBadRequest)
		return
	}

	if err := s.db.InsertTask(task); err != nil {
		http.Error(res, "internal server", http.StatusBadRequest)
		return
	}

	res.Header().Set("Conntent-type", "text/plain")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(fmt.Sprintf("Task %v was saved", task.ID)))
}

func (s *Server) UpdateTaskHandler (res http.ResponseWriter, req *http.Request) {
	param := req.URL.Query().Get("ID")
	ID, err := strconv.Atoi(param)
	if err != nil {
		http.Error(res, "invalid argument", http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := json.NewDecoder(req.Body).Decode(&task); err != nil {
		http.Error(res, "invalid argument", http.StatusBadRequest)
		return
	}

	if err := s.db.UpdateTask(task, ID); err != nil {
		http.Error(res, "internal server", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Conntent-type", "text/plain")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(fmt.Sprintf("Task %v was updated!", task.ID)))
}

func (s *Server) DeleteTaskHandler (res http.ResponseWriter, req *http.Request) {
	param := req.URL.Query().Get("ID")
	ID, err := strconv.Atoi(param)
	if err != nil {
		http.Error(res, "invalid argument", http.StatusBadRequest)
		return
	}

	if err := s.db.DeleteTask(ID); err != nil {
		http.Error(res, "internal server", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-type", "text/plain")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(fmt.Sprintf("Task %v was deleted!", ID + 1)))
}

func (s *Server) GetTaskHandler(res http.ResponseWriter, req *http.Request) {
	param := req.URL.Query().Get("ID")
	ID, err := strconv.Atoi(param)
	if err != nil {
		http.Error(res, "invalid argument", http.StatusBadRequest)
		return
	}

	task, err := s.db.GetTask(ID)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(task)
}