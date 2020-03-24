package tests

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// $all
type allTest struct {
	test objectTest
}

func NewAllTest(input interface{}) (Test, error) {
	castedObject, ok := castObject(input)
	if !ok {
		return nil, fmt.Errorf("value for $all is not a an object (%#v %v)", input, reflect.TypeOf(input))
	}

	allObjectTest, err := NewObjectTest(castedObject)
	if err != nil {
		return nil, fmt.Errorf("failed to create all test: %v", err)
	}

	objectTest, ok := allObjectTest.(*objectTest)
	if !ok {
		return nil, errors.New("failed to get objecttest")
	}

	return &allTest{
		test: *objectTest,
	}, nil
}

func (at *allTest) Run(input interface{}) error {
	castedArray, ok := castArray(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't an array", input, reflect.TypeOf(input))
	}

	if len(castedArray) == 0 {
		return fmt.Errorf("input (%#v %v) is empty", input, reflect.TypeOf(input))
	}

	var errs []error

	errs = at.test.runMetaTests(castedArray)

	for idx, element := range castedArray {
		fieldErrors, err := at.test.runFieldTests(element)
		if err != nil {
			return err
		}

		if len(fieldErrors) == 0 {
			continue
		}

		errs = append(errs, &fieldError{
			field: strconv.Itoa(idx),
			fieldError: &objectTestError{
				errors: fieldErrors,
			}})
	}

	if len(errs) > 0 {
		return &objectTestError{
			errors: errs,
		}
	}

	return nil
}

func (at *allTest) IsMetaTest() bool {
	return true
}
