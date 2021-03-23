package tests

import (
	"fmt"
	"reflect"
	"strings"
)

// $startsWithIgnoreCase
type startsWithInsensitiveTest struct {
	expectedStart string
}

func NewStartsWithInsensitiveTest(input interface{}) (Test, error) {
	castedString, ok := castString(input)
	if !ok {
		return nil, fmt.Errorf("value for $startsWithIgnoreCase is not a string (%#v %v)", input, reflect.TypeOf(input))
	}

	lowerCaseString := strings.ToLower(castedString)
	return &startsWithInsensitiveTest{
		expectedStart: lowerCaseString,
	}, nil
}

func (swt *startsWithInsensitiveTest) Run(input interface{}) error {
	castedString, ok := castString(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't a string", input, reflect.TypeOf(input))
	}

	if !strings.HasPrefix(strings.ToLower(castedString), swt.expectedStart) {
		return fmt.Errorf("%q does not start with %q", castedString, swt.expectedStart)
	}

	return nil

}

func (swt *startsWithInsensitiveTest) IsMetaTest() bool {
	return true
}
