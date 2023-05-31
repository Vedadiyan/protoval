package protoval

type ValidationError struct {
	fieldName string
	message   string
}

func (protovalError ValidationError) Error() string {
	return protovalError.message
}

func NewError(fieldName string, message string) *ValidationError {
	validationError := ValidationError{
		fieldName: fieldName,
		message:   message,
	}
	return &validationError
}
