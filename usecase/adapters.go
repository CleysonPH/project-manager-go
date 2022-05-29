package usecase

import "github.com/cleysonph/project-manager/domain"

type ProjectService interface {
	FindAll() ([]*domain.Project, error)
}
