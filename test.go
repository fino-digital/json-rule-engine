package tests

import "fmt"

type Test interface {
	Run(input interface{}) error
	IsMetaTest() bool
}

// NewTest generates a new test for a given input array or object
func NewTest(in interface{}) (Test, error) {
	castedArray, ok := castArray(in)
	if ok {
		return NewArrayTest(castedArray)
	}

	castedObject, ok := castObject(in)
	if ok {
		return NewObjectTest(castedObject)
	}

	return nil, fmt.Errorf("input was not an object or array: %v", in)
}
