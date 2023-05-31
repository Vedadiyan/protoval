package protoval

import (
	"google.golang.org/protobuf/proto"
)

type ValidationContext struct {
	validationTag string
	message       proto.Message
	errors        []error
}

func New(validationTag string, message proto.Message) *ValidationContext {
	validationContext := ValidationContext{
		validationTag: validationTag,
		message:       message,
		errors:        make([]error, 0),
	}
	return &validationContext
}

func (vc *ValidationContext) Validate() error {
	validators := make(map[string]func() error)
	err := Reflect(vc.validationTag, vc.message, validators, "", -1)
	if err != nil {
		return err
	}
	for _, validator := range validators {
		err := validator()
		if err != nil {
			if _, ok := err.(ValidationError); !ok {
				return err
			}
			vc.errors = append(vc.errors, err)
		}
	}
	return nil
}

func (vc ValidationContext) IsValid() bool {
	return len(vc.errors) == 0
}

func (vc ValidationContext) Errors() []error {
	return vc.errors
}
