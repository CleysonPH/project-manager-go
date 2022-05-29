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

// FindAll implements ProjectService
func (s *projectService) FindAll() ([]*domain.Project, error) {
	projects, err := s.projectRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return projects, nil
}
