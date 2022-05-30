package domain

func NewDomainError(message string) error {
	return &DomainError{Message: message}
}

type DomainError struct {
	Message string
}

// Error implements error
func (err *DomainError) Error() string {
	return err.Message
}
