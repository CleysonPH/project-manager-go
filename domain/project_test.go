package domain

import (
	"testing"
	"time"
)

func TestNewProject(t *testing.T) {
	p, err := NewProject(
		1,
		"Project title",
		"Project description",
		time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
		time.Time{},
		ProjectStatusCreated,
	)
	if err != nil {
		t.Error(err)
	}
	if p.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", p.ID)
	}
	if p.Title != "Project title" {
		t.Errorf("Expected Title to be title, got %s", p.Title)
	}
	if p.Description != "Project description" {
		t.Errorf("Expected Description to be description, got %s", p.Description)
	}
	if p.Status != ProjectStatusCreated {
		t.Errorf("Expected Status to be Created, got %s", p.Status)
	}
}

func TestNewProjectWithStartAtAfterBeforeAt(t *testing.T) {
	_, err := NewProject(
		1,
		"Project title",
		"Project description",
		time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Time{},
		ProjectStatusCreated,
	)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestNewProjectWithEmptyTitle(t *testing.T) {
	_, err := NewProject(
		1,
		"",
		"Project description",
		time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
		time.Time{},
		ProjectStatusCreated,
	)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestNewProjectWithConcludedStatusAndEmptyConcludedAt(t *testing.T) {
	_, err := NewProject(
		1,
		"Project title",
		"Project description",
		time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
		time.Time{},
		ProjectStatusConcluded,
	)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestNewProjectWithCreatedStatusAndNotEmptyConcludedAt(t *testing.T) {
	_, err := NewProject(
		1,
		"Project title",
		"Project description",
		time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2022, time.January, 3, 0, 0, 0, 0, time.UTC),
		ProjectStatusCreated,
	)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestNewProjectWithConcludedAtBeforeStartAt(t *testing.T) {
	_, err := NewProject(
		1,
		"Project title",
		"Project description",
		time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2022, time.January, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		ProjectStatusConcluded,
	)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
