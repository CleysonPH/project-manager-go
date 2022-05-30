package main

import (
	"log"
	"net/http"
	"time"

	"github.com/cleysonph/project-manager/api/controller"
	"github.com/cleysonph/project-manager/api/router"
	"github.com/cleysonph/project-manager/repository"
	"github.com/cleysonph/project-manager/usecase"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	projectRepo := repository.NewInMemoryProjectRepository()
	projectService := usecase.NewProjectService(projectRepo)
	projectController := controller.NewProjectController(projectService)

	router.InitProjectRoutes(r, projectController)

	server := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
