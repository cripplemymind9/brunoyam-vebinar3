package ma

import (
	"github.com/cripplemymind9/brunoyam-vebinar3/homework/internal/repository"
	"github.com/cripplemymind9/brunoyam-vebinar3/homework/internal/server"
	"github.com/go-chi/chi"
	"net/http"
	"log"
)

func main() {
	repo, err := repository.New()
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(":8080", repo)

	router := chi.NewRouter()

	router.Route("/", func(router chi.Router) {
		router.Get("/tasks", server.GetAllTasksHandler)
		router.Post("/task", server.SaveTaskHandler)
		router.Put("/task", server.UpdateTaskHandler)
		router.Get("/task", server.GetTaskHandler)
		router.Delete("/task", server.DeleteTaskHandler)
	})

	HTTPServer := &http.Server{}
	HTTPServer.Handler = router
	HTTPServer.Addr = ":8080"
	HTTPServer.ListenAndServe()
}