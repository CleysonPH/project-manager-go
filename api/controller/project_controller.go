package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cleysonph/project-manager/api/presenter"
	"github.com/cleysonph/project-manager/usecase"
)

func NewProjectController(projectService usecase.ProjectService) ProjectController {
	return ProjectController{
		projectService: projectService,
	}
}

type ProjectController struct {
	projectService usecase.ProjectService
}

func (c *ProjectController) FindAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		projects, err := c.projectService.FindAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		var toJson []*presenter.Project
		for _, p := range projects {
			toJson = append(toJson, &presenter.Project{
				ID:          p.ID,
				Title:       p.Title,
				Description: p.Description,
				StartAt:     p.StartAt,
				FinishAt:    p.FinishAt,
				ConcludedAt: p.ConcludedAt,
				Status:      string(p.Status),
			})
		}
		if err := json.NewEncoder(w).Encode(toJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	})
}
