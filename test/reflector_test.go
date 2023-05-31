package test

import (
	"fmt"
	"testing"

	protoval "github.com/vedadiyan/protoval/pkg"
)

func TestReflect(t *testing.T) {
	test := Test{}
	test.FieldA = "okk"
	test.TestItem = append(test.TestItem, &TestItem{TestField: "ok"})
	vc := protoval.New("validate", &test)
	err := vc.Validate()
	if err != nil {
		t.FailNow()
	}
	fmt.Println(vc.IsValid())
	for _, err := range vc.Errors() {
		fmt.Println(err.FieldName())
	}
}

func TestPattern(t *testing.T) {
	protoval.ExprParser("required|  max_len : 100 |min_len:100|/[Aa-Zz]\\/\\\\%$/|UPPERCASE")
}
