package repository

import (
	"time"

	"github.com/cleysonph/project-manager/domain"
)

func NewInMemoryProjectRepository() ProjectRepository {
	return &inMemoryProjectRepository{}
}

var inMemoryProjects = []*domain.Project{
	{
		ID:          1,
		Title:       "Project 1",
		Description: "Project 1 description",
		StartAt:     time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		FinishAt:    time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
		ConcludedAt: time.Time{},
		Status:      domain.ProjectStatusCreated,
	},
}

type inMemoryProjectRepository struct{}

// FindAll implements ProjectRepository
func (*inMemoryProjectRepository) FindAll() ([]*domain.Project, error) {
	return inMemoryProjects, nil
}
