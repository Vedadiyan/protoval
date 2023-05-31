package protoval

import "google.golang.org/protobuf/proto"

func Validate(validationTag string, message proto.Message) (map[string]func() (bool, error), error) {
	vc := make(map[string]func() (bool, error))
	err := Reflect(validationTag, message, vc, "", -1)
	if err != nil {
		return nil, err
	}
	return vc, nil
}
