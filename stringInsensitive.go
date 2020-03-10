package tests

import (
	"fmt"
	"reflect"
	"strings"
)

type stringInsensitiveTest struct {
	expected string
}

func (st *stringInsensitiveTest) Run(input interface{}) error {
	castedString, ok := castString(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't a string", input, reflect.TypeOf(input))
	}

	if !strings.EqualFold(castedString, st.expected) {
		return fmt.Errorf("%q is not case insensitive equal to %q", castedString, st.expected)
	}

	return nil

}

func (st *stringInsensitiveTest) IsMetaTest() bool {
	return false
}

func NewCaseInsensitiveStringTest(expected interface{}) (Test, error) {
	if castedString, ok := castString(expected); ok {
		return &stringInsensitiveTest{
			expected: castedString,
		}, nil
	}

	return nil, fmt.Errorf("couldn't create test for value %#v", expected)
}
