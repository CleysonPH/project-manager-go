package domain

import "time"

func NewProject(
	id uint64,
	title string,
	description string,
	startAt time.Time,
	finishAt time.Time,
	concludedAt time.Time,
	status ProjectStatus,
) (*Project, error) {
	p := &Project{
		ID:          id,
		Title:       title,
		Description: description,
		StartAt:     startAt,
		FinishAt:    finishAt,
		ConcludedAt: concludedAt,
		Status:      status,
	}
	if err := validate(p); err != nil {
		return nil, err
	}
	return p, nil
}

type Project struct {
	ID          uint64
	Title       string
	Description string
	StartAt     time.Time
	FinishAt    time.Time
	ConcludedAt time.Time
	Status      ProjectStatus
}

func validate(p *Project) error {
	if p.StartAt.After(p.ConcludedAt) {
		return NewDomainError("StartAt cannot be after ConcludedAt")
	}
	if p.Title == "" {
		return NewDomainError("Title cannot be empty")
	}
	if p.Status == ProjectStatusConcluded && p.ConcludedAt.IsZero() {
		return NewDomainError("ConcludedAt cannot be empty")
	}
	if p.Status == ProjectStatusCreated && !p.StartAt.IsZero() {
		return NewDomainError("StartAt cannot be set")
	}
	if p.Status != ProjectStatusCreated && p.Status != ProjectStatusConcluded {
		return NewDomainError("Status must be Created or Concluded")
	}
	return nil
}

type ProjectStatus string

const (
	ProjectStatusCreated   = "CREATED"
	ProjectStatusConcluded = "CONCLUDED"
)
