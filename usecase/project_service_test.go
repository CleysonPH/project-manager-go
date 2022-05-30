package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/cleysonph/project-manager/domain"
	"github.com/cleysonph/project-manager/repository"
)

func newFixtureProject() *domain.Project {
	return &domain.Project{
		ID:          1,
		Title:       "Project 1",
		Description: "Project 1 description",
		StartAt:     time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		FinishAt:    time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
		ConcludedAt: time.Time{},
		Status:      domain.ProjectStatusCreated,
	}
}

func TestProjectService_FindAll(t *testing.T) {
	projectRepository := repository.NewInMemoryProjectRepository()
	projectService := NewProjectService(projectRepository)

	projects, err := projectService.FindAll()
	expectedProjects := []*domain.Project{
		newFixtureProject(),
	}

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(projects) != len(expectedProjects) {
		t.Errorf("Expected %d projects, got %d", len(expectedProjects), len(projects))
	}

	for i, project := range projects {
		if project.ID != expectedProjects[i].ID {
			t.Errorf("Expected project id %d, got %d", expectedProjects[i].ID, project.ID)
		}

		if project.Title != expectedProjects[i].Title {
			t.Errorf("Expected project title %s, got %s", expectedProjects[i].Title, project.Title)
		}

		if project.Description != expectedProjects[i].Description {
			t.Errorf("Expected project description %s, got %s", expectedProjects[i].Description, project.Description)
		}

		if project.StartAt != expectedProjects[i].StartAt {
			t.Errorf("Expected project start at %s, got %s", expectedProjects[i].StartAt, project.StartAt)
		}

		if project.FinishAt != expectedProjects[i].FinishAt {
			t.Errorf("Expected project finish at %s, got %s", expectedProjects[i].FinishAt, project.FinishAt)
		}

		if project.ConcludedAt != expectedProjects[i].ConcludedAt {
			t.Errorf("Expected project concluded at %s, got %s", expectedProjects[i].ConcludedAt, project.ConcludedAt)
		}

		if project.Status != expectedProjects[i].Status {
			t.Errorf("Expected project status %s, got %s", expectedProjects[i].Status, project.Status)
		}
	}
}

type projectRepositoryMock struct{}

func (*projectRepositoryMock) FindAll() ([]*domain.Project, error) {
	return nil, errors.New("error")
}

func (*projectRepositoryMock) Create(project *domain.Project) (*domain.Project, error) {
	return project, nil
}

func TestProjectService_PassRepositoryErr(t *testing.T) {
	projectService := NewProjectService(&projectRepositoryMock{})
	_, err := projectService.FindAll()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
