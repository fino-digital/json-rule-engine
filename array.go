package tests

import (
	"errors"
	"fmt"
	"strconv"
)

type arrayTest struct {
	tests []Test
}

func (ot *arrayTest) Run(input interface{}) error {
	castedArray, ok := castArray(input)
	if !ok {
		return fmt.Errorf("input (%#v) wasn't an array", input)
	}

	var errs []error
	if len(castedArray) != len(ot.tests) {
		return fmt.Errorf("input array length (%d) is not equal to test array length (%d)", len(castedArray), len(ot.tests))
	}

	for index, test := range ot.tests {
		value := castedArray[index]
		if value == nil {
			errs = append(errs, &fieldError{field: strconv.Itoa(index), fieldError: errors.New("index does not exist")})
			continue
		}

		err := test.Run(value)

		if err != nil {
			errs = append(errs, &fieldError{field: strconv.Itoa(index), fieldError: err})
		}
	}

	if len(errs) > 0 {
		return &objectTestError{
			errors: errs,
		}
	}

	return nil
}

func (at *arrayTest) IsMetaTest() bool {
	return false
}

func NewArrayTest(input []interface{}) (Test, error) {
	tests := make([]Test, len(input))

	for index, v := range input {
		if castedString, ok := castString(v); ok {
			tests[index] = NewStringTest(castedString)
			continue
		}

		if castedInt, ok := castNumber(v); ok {
			tests[index] = NewNumberTest(castedInt)
			continue
		}

		if castedObject, ok := castObject(v); ok {
			test, err := NewObjectTest(castedObject)
			if err != nil {
				return nil, err
			}

			tests[index] = test
			continue
		}

		if castedArray, ok := castArray(v); ok {
			test, err := NewArrayTest(castedArray)
			if err != nil {
				return nil, err
			}

			tests[index] = test
			continue
		}

		return nil, fmt.Errorf("couldn't create test for index %d with value %#v", index, v)
	}

	return &arrayTest{
		tests: tests,
	}, nil
}
