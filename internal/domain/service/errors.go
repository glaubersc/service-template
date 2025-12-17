package service

type DomainError struct {
	Code    string
	Message string
	Details map[string]any
}

func (e DomainError) Error() string {
	return e.Message
}

func New(code, message string, details map[string]any) DomainError {
	return DomainError{
		Code:    code,
		Message: message,
		Details: details,
	}
}
