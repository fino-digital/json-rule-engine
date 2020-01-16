package tests

import (
	"fmt"
	"reflect"
	"strconv"
)

// $none
type noneTest struct {
	test Test
}

func NewNoneTest(input interface{}) (Test, error) {
	castedObject, ok := castObject(input)
	if !ok {
		return nil, fmt.Errorf("value for $none is not a an object (%#v %v)", input, reflect.TypeOf(input))
	}

	noneObjectTest, err := NewObjectTest(castedObject)
	if err != nil {
		return nil, fmt.Errorf("failed to create none test: %v", err)
	}

	return &noneTest{
		test: noneObjectTest,
	}, nil
}

func (at *noneTest) Run(input interface{}) error {
	castedArray, ok := castArray(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't an array", input, reflect.TypeOf(input))
	}

	failed := false
	var errs []error

	for idx, element := range castedArray {
		err := at.test.Run(element)
		if err != nil {
			continue
		}

		failed = true

		errs = append(errs, &fieldError{
			field:      strconv.Itoa(idx),
			fieldError: fmt.Errorf("field was equal to expected value"),
		})
	}

	if failed {
		return &objectTestError{
			errors: errs,
		}
	}

	return nil
}

func (at *noneTest) IsMetaTest() bool {
	return true
}
