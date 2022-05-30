package repository

import "github.com/cleysonph/project-manager/domain"

type ProjectRepository interface {
	FindAll() ([]*domain.Project, error)
	Create(project *domain.Project) (*domain.Project, error)
}
