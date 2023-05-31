package protoval

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Reflect(validationTag string, message proto.Message, vc map[string]func() (bool, error)) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	reflector := message.ProtoReflect()
	fields := reflector.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		options := field.Options().(*descriptorpb.FieldOptions)
		proto.RangeExtensions(options, func(et protoreflect.ExtensionType, i interface{}) bool {
			fullName := et.TypeDescriptor().FullName()
			if fullName == protoreflect.FullName(validationTag) {
				rule, ok := i.(string)
				if !ok {
					return true
				}
				value := reflector.Get(field).Interface()
				rules, err := ExprParser(rule)
				if err != nil {
					panic(err)
				}
				vc[field.TextName()] = func() (bool, error) {
					output := true
					for name, rule := range rules {
						validator, ok := _rules[name]
						if !ok {
							return false, fmt.Errorf("%s: rule not found", name)
						}
						result, err := validator(value, rule)
						if err != nil {
							return false, err
						}
						output = output && result
					}
					return output, nil
				}
				return false
			}
			return true
		})
		if field.Kind() == protoreflect.MessageKind {
			if field.IsList() {
				list := reflector.Get(field).List()
				for i := 0; i < list.Len(); i++ {
					err := Reflect(validationTag, list.Get(i).Message().Interface(), vc)
					if err != nil {
						return err
					}
				}
				continue
			}
			err := Reflect(validationTag, reflector.Get(field).Message().Interface(), vc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
