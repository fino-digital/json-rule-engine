package tests

import (
	"fmt"
	"reflect"
)

type stringTest struct {
	expected string
}

func (st *stringTest) Run(input interface{}) error {
	castedString, ok := castString(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't a string", input, reflect.TypeOf(input))
	}

	if castedString != st.expected {
		return fmt.Errorf("%q is not equal to %q", castedString, st.expected)
	}

	return nil

}

func (st *stringTest) IsMetaTest() bool {
	return false
}

func NewStringTest(expected string) Test {
	return &stringTest{
		expected: expected,
	}
}
