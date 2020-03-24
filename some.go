package tests

import (
	"fmt"
	"reflect"
	"strconv"
)

// $some
type someTest struct {
	test Test
}

func NewSomeTest(input interface{}) (Test, error) {
	fieldTest, err := fieldTestForValue(input)
	if err != nil {
		return nil, err
	}

	return &someTest{
		test: fieldTest,
	}, nil
}

func (at *someTest) Run(input interface{}) error {
	castedArray, ok := castArray(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't an array", input, reflect.TypeOf(input))
	}

	if len(castedArray) == 0 {
		return fmt.Errorf("input (%#v %v) is empty", input, reflect.TypeOf(input))
	}

	succeeded := false
	var errs []error

	for idx, element := range castedArray {
		err := at.test.Run(element)
		if err == nil {
			succeeded = true
			continue
		}

		errs = append(errs, &fieldError{
			field:      strconv.Itoa(idx),
			fieldError: err,
		})
	}

	if !succeeded {
		return &objectTestError{
			errors: errs,
		}
	}

	return nil
}

func (at *someTest) IsMetaTest() bool {
	return true
}
