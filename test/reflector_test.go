package test

import (
	"fmt"
	"testing"

	protoval "github.com/vedadiyan/protoval/pkg"
)

func TestReflect(t *testing.T) {
	vc := make(map[string]func() (bool, error))
	test := Test{}
	test.FieldA = "okk"
	test.TestItem = append(test.TestItem, &TestItem{TestField: "okkkk"})
	err := protoval.Reflect("validate", &test, vc)
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
