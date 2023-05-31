package test

import (
	"fmt"
	"testing"

	protoval "github.com/vedadiyan/protoval/pkg"
)

func TestReflect(t *testing.T) {
	test := Test{}
	test.FieldA = "okk"
	test.TestItem = append(test.TestItem, &TestItem{TestField: "okkkkkkkkkkkk"})
	vc, err := protoval.Validate("validate", &test)
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
