package tests

import (
	"fmt"
	"math"
	"reflect"
)

// $length
type lengthTest struct {
	expectedLen int
}

func NewLengthTest(input interface{}) (Test, error) {
	castedNumber, ok := castNumber(input)
	if !ok {
		return nil, fmt.Errorf("value for $length is not a number (%#v %v)", input, reflect.TypeOf(input))
	}

	// we need to check whether an integer was supplied
	if math.Floor(castedNumber) != castedNumber {
		return nil, fmt.Errorf("provided length for $length is not an integer (%f)", castedNumber)
	}

	return &lengthTest{
		expectedLen: int(castedNumber),
	}, nil
}

func (lt *lengthTest) Run(input interface{}) error {
	castedArray, ok := castArray(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't an array", input, reflect.TypeOf(input))
	}

	if len(castedArray) != lt.expectedLen {
		return fmt.Errorf("input array length (%d) is not equal expected length (%d)", len(castedArray), lt.expectedLen)
	}

	return nil

}

func (lt *lengthTest) IsMetaTest() bool {
	return true
}
