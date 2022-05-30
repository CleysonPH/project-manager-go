package usecase

import (
	"github.com/cleysonph/project-manager/domain"
	"github.com/cleysonph/project-manager/repository"
)

func NewProjectService(projectRepository repository.ProjectRepository) ProjectService {
	return &projectService{
		projectRepository: projectRepository,
	}
}

type projectService struct {
	projectRepository repository.ProjectRepository
}

// Create implements ProjectService
func (s *projectService) Create(project *domain.Project) (*domain.Project, error) {
	p, err := s.projectRepository.Create(project)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// FindAll implements ProjectService
func (s *projectService) FindAll() ([]*domain.Project, error) {
	projects, err := s.projectRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return projects, nil
}
