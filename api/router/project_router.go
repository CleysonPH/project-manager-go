package router

import (
	"github.com/cleysonph/project-manager/api/controller"
	"github.com/gorilla/mux"
)

func InitProjectRoutes(r *mux.Router, controller controller.ProjectController) {
	r.Handle("/api/v1/projects", controller.FindAll()).Methods("GET")
}
