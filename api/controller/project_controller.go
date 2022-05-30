package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cleysonph/project-manager/api/presenter"
	"github.com/cleysonph/project-manager/domain"
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

func (c *ProjectController) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var project *presenter.Project
		if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		p, err := c.projectService.Create(&domain.Project{
			Title:       project.Title,
			Description: project.Description,
			StartAt:     project.StartAt,
			FinishAt:    project.FinishAt,
			ConcludedAt: project.ConcludedAt,
			Status:      domain.ProjectStatusCreated,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
		toJson := &presenter.Project{
			ID:          p.ID,
			Title:       p.Title,
			Description: p.Description,
			StartAt:     p.StartAt,
			FinishAt:    p.FinishAt,
			ConcludedAt: p.ConcludedAt,
			Status:      string(p.Status),
		}
		if err := json.NewEncoder(w).Encode(toJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	})
}
