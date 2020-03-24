package tests

import (
	"fmt"
	"reflect"
	"strconv"
)

func createSomeTestsForInput(input interface{}) ([]Test, error) {
	inputArray, ok := castArray(input)
	if !ok {
		return nil, fmt.Errorf("input wasn't an array: %v", input)
	}

	var tests []Test
	for index, value := range inputArray {
		someTest, err := NewSomeTest(value)
		if err != nil {
			return nil, fmt.Errorf("failed to create test for index %d (%v): %v", index, value, err)
		}

		tests = append(tests, someTest)
	}

	return tests, nil
}

type containsTest struct {
	someTests []Test
}

func (ct *containsTest) Run(input interface{}) error {
	inputArray, ok := castArray(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't an array", input, reflect.TypeOf(input))
	}

	if len(inputArray) == 0 {
		return fmt.Errorf("input (%#v %v) is empty", input, reflect.TypeOf(input))
	}

	var errs []error
	for _, test := range ct.someTests {
		err := test.Run(input)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return &objectTestError{
			errors: errs,
		}
	}

	return nil
}

func (ct *containsTest) IsMetaTest() bool {
	return false
}

func NewContainsTest(input interface{}) (Test, error) {
	tests, err := createSomeTestsForInput(input)
	if err != nil {
		return nil, err
	}

	return &containsTest{
		someTests: tests,
	}, nil
}

type containsSomeTest struct {
	someTests []Test
}

func (ct *containsSomeTest) Run(input interface{}) error {
	inputArray, ok := castArray(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't an array", input, reflect.TypeOf(input))
	}

	if len(inputArray) == 0 {
		return fmt.Errorf("input (%#v %v) is empty", input, reflect.TypeOf(input))
	}

	var errs []error
	for _, test := range ct.someTests {
		err := test.Run(input)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) == len(ct.someTests) {
		return &objectTestError{
			errors: errs,
		}
	}

	return nil
}

func (ct *containsSomeTest) IsMetaTest() bool {
	return false
}

func NewContainsSomeTest(input interface{}) (Test, error) {
	tests, err := createSomeTestsForInput(input)
	if err != nil {
		return nil, err
	}

	return &containsSomeTest{
		someTests: tests,
	}, nil
}

type containsNoneTest struct {
	someTests []Test
}

func (ct *containsNoneTest) Run(input interface{}) error {
	inputArray, ok := castArray(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't an array", input, reflect.TypeOf(input))
	}

	if len(inputArray) == 0 {
		return nil
	}

	var errs []error
	for idx, test := range ct.someTests {
		err := test.Run(input)
		if err == nil {
			errs = append(errs, &fieldError{
				field:      strconv.Itoa(idx),
				fieldError: fmt.Errorf("should not have been contained in array"),
			})
		}
	}

	if len(errs) > 0 {
		return &objectTestError{
			errors: errs,
		}
	}

	return nil
}

func (ct *containsNoneTest) IsMetaTest() bool {
	return false
}

func NewContainsNoneTest(input interface{}) (Test, error) {
	tests, err := createSomeTestsForInput(input)
	if err != nil {
		return nil, err
	}

	return &containsNoneTest{
		someTests: tests,
	}, nil
}
