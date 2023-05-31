package protoval

import (
	"bytes"
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Reflect(message proto.Message, vc map[string]func() bool) {
	reflector := message.ProtoReflect()
	fields := reflector.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		switch field.Kind() {
		case protoreflect.MessageKind:
			{
				Reflect(reflector.Get(field).Message().Interface(), vc)
			}
		default:
			{
				options := field.Options().(*descriptorpb.FieldOptions)
				proto.RangeExtensions(options, func(et protoreflect.ExtensionType, i interface{}) bool {
					fullName := et.TypeDescriptor().FullName()
					if fullName == "validate" {
						rule, ok := i.(string)
						if !ok {
							return true
						}
						value := reflector.Get(field).Interface()
						vc[field.TextName()] = func() bool {
							_ = rule
							_ = value
							return true
						}
						return false
					}
					return true
				})
			}
		}
	}
}

func Trim(str string) string {
	return strings.TrimSuffix(strings.TrimPrefix(str, " "), " ")
}

func GetRule(str string, rules map[string]any) error {
	if len(str) == 0 {
		return nil
	}
	split := strings.Split(str, ":")
	if len(split) == 1 {
		rules[str] = true
		return nil
	}
	if len(split) == 2 {
		rules[split[0]] = split[1]
		return nil
	}
	return fmt.Errorf("invalid rule")
}

func ExprParser(expr string) {
	var buffer bytes.Buffer
	rules := make(map[string]any)

	for i := 0; i < len(expr); i++ {
		c := expr[i]
		switch c {
		case '/':
			{
				if buffer.Len() > 0 {
					panic("invalid validation pattern")
				}
				hold := false
			LOOP:
				for x := i + 1; x < len(expr); x++ {
					c := expr[x]
					if hold {
						hold = false
						goto WRITE
					}
					switch c {
					case '\\':
						{
							hold = true
						}
					case '/':
						{
							i = x
							break LOOP
						}
					}
				WRITE:
					buffer.WriteByte(c)
				}
				str := buffer.String()
				buffer.Reset()
				rules["regex"] = str
			}
		case '|':
			{
				str := Trim(buffer.String())
				buffer.Reset()
				err := GetRule(str, rules)
				if err != nil {
					panic(err)
				}
			}
		default:
			{
				buffer.WriteByte(c)
			}
		}
	}
	str := Trim(buffer.String())
	buffer.Reset()
	err := GetRule(str, rules)
	if err != nil {
		panic(err)
	}
	c := 10
	_ = c
}
