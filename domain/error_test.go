package domain

import "testing"

func TestNewDomainError(t *testing.T) {
	e := NewDomainError("Test error")
	if e.Error() != "Test error" {
		t.Errorf("Expected error to be Test error, got %s", e.Error())
	}
}
