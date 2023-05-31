package protoval

import "fmt"

type ValidationError string

func (protovalError ValidationError) Error() string {
	return string(protovalError)
}

func ErrorF(format string, a ...any) ValidationError {
	return ValidationError(fmt.Sprintf(format, a...))
}
