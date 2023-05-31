package test

import (
	"testing"

	protoval "github.com/vedadiyan/protoval/pkg"
)

func TestReflect(t *testing.T) {
	protoval.Reflect(&Test{FieldA: "ok"}, make(map[string]func() bool))
}

func TestPattern(t *testing.T) {
	protoval.ExprParser("required|max_len:100|min_len:100|/[Aa-Zz]\\/\\\\%$/|UPPERCASE")
}
