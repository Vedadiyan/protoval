package protoval

import (
	"fmt"
	"strconv"
)

var (
	_rules map[string]func(value any, rule string) (bool, error)
)

func init() {
	_rules = make(map[string]func(value any, rule string) (bool, error))
	_rules["required"] = Required
	_rules["max_len"] = MaxLen
	_rules["min_len"] = MinLen
}

func Register(name string, fn func(value any, rule string) (bool, error)) {
	_rules[name] = fn
}

func Required(value any, rule string) (bool, error) {
	return value != nil, nil
}

func MinLen(value any, rule string) (bool, error) {
	str, ok := value.(string)
	if !ok {
		return false, fmt.Errorf("expected string but recieved %T", value)
	}
	i, err := strconv.ParseInt(rule, 10, 32)
	if err != nil {
		return false, err
	}
	return len(str) >= int(i), nil
}

func MaxLen(value any, rule string) (bool, error) {
	str, ok := value.(string)
	if !ok {
		return false, fmt.Errorf("expected string but recieved %T", value)
	}
	i, err := strconv.ParseInt(rule, 10, 32)
	if err != nil {
		return false, err
	}
	return len(str) <= int(i), nil
}
