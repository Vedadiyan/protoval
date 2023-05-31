package test

import (
	"fmt"
	"testing"

	protoval "github.com/vedadiyan/protoval/pkg"
)

func TestReflect(t *testing.T) {
	vc := make(map[string]func() (bool, error))
	err := protoval.Reflect(&Test{FieldA: "ooooooooooooook"}, vc)
	if err != nil {
		t.FailNow()
	}
	for key, value := range vc {
		result, err := value()
		if err != nil {
			t.FailNow()
		}
		fmt.Println(key, result)
	}
}

func TestPattern(t *testing.T) {
	protoval.ExprParser("required|  max_len : 100 |min_len:100|/[Aa-Zz]\\/\\\\%$/|UPPERCASE")
}
