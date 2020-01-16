package tests

import (
	"fmt"
	"reflect"
	"strings"
)

// $startsWith
type startsWithTest struct {
	expectedStart string
}

func NewStartsWithTest(input interface{}) (Test, error) {
	castedString, ok := castString(input)
	if !ok {
		return nil, fmt.Errorf("value for $startsWith is not a string (%#v %v)", input, reflect.TypeOf(input))
	}
	return &startsWithTest{
		expectedStart: castedString,
	}, nil
}

func (swt *startsWithTest) Run(input interface{}) error {
	castedString, ok := castString(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't a string", input, reflect.TypeOf(input))
	}

	if !strings.HasPrefix(castedString, swt.expectedStart) {
		return fmt.Errorf("%q does not start with %q", castedString, swt.expectedStart)
	}

	return nil

}

func (swt *startsWithTest) IsMetaTest() bool {
	return true
}
