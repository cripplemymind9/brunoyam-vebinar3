package server

import (
	"github.com/cripplemymind9/brunoyam-vebinar3/classwork/internal/domain/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (s *Server) Run() error {
	router := gin.Default()

	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.GET("/", s.GetTasksHadler)
		taskRoutes.POST("/", s.InsertTaskHandler)
		taskRoutes.PUT("/:id", s.UpdateTaskHandler)
		taskRoutes.DELETE("/:id", s.DeleteTaskHandler)
		taskRoutes.GET("/:id", s.GetTaskHandler)
	}

	return router.Run(s.addr)
}

func (s *Server) GetTasksHadler(ctx *gin.Context) {
	tasks, err := s.db.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (s *Server) InsertTaskHandler(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindBodyWithJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	
	if task.Status == "" {
		task.Status = "New"
	}

	if err := s.db.InsertTask(task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.String(http.StatusOK, "Task was saved")
}

func (s *Server) UpdateTaskHandler(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	var task models.Task
	if err := ctx.ShouldBindBodyWithJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	if err := s.db.UpdateTask(task, id); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	ctx.String(http.StatusOK, "Task was updated")
}

func (s *Server) DeleteTaskHandler(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	if err := s.db.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	ctx.String(http.StatusOK, "Task was deleted")
}

func (s *Server) GetTaskHandler (ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	task, err := s.db.GetTask(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
	}

	ctx.JSON(http.StatusOK, task)
}