package tests

import (
	"errors"
	"fmt"
	"strings"
)

type fieldError struct {
	field      string
	fieldError error
}

func (fe *fieldError) Error() string {
	return fmt.Sprintf("test for field %q failed: %v", fe.field, fe.fieldError.Error())
}

type objectTestError struct {
	errors []error
}

func traceTestError(errorMessage string, indentation int, err error) string {
	var fe *fieldError
	if errors.As(err, &fe) {
		return traceTestError(errorMessage+strings.Repeat("\t", indentation)+fe.field+".", indentation, fe.fieldError)
	}
	var oe *objectTestError
	if errors.As(err, &oe) {
		return errorMessage + "\n" + oe.IndentedError(indentation+1)
	}

	return strings.TrimSuffix(errorMessage, ".") + ": " + fmt.Sprintf("%v\n", err)
}

func (otr *objectTestError) IndentedError(indents int) string {
	errorMessage := ""
	for _, err := range otr.errors {
		errorMessage = traceTestError(errorMessage, indents, err)
	}

	return errorMessage
}

func (otr *objectTestError) Error() string {
	errorMessage := "\n"
	for _, err := range otr.errors {
		errorMessage = traceTestError(errorMessage, 0, err)
	}

	return errorMessage
}

type objectTest struct {
	metaTests  map[string]Test
	fieldTests map[string]Test
}

func (ot *objectTest) runMetaTests(input interface{}) []error {
	var errs []error
	for key, test := range ot.metaTests {
		err := test.Run(input)

		if err != nil {
			errs = append(errs, &fieldError{key, err})
		}
	}

	return errs
}

func (ot *objectTest) runFieldTests(input interface{}) ([]error, error) {
	var errs []error
	if len(ot.fieldTests) > 0 {
		castedMap, ok := castObject(input)
		if !ok {
			return nil, fmt.Errorf("input (%#v) wasn't an object", input)
		}

		for key, test := range ot.fieldTests {
			value := castedMap[key]
			if value == nil {
				errs = append(errs, &fieldError{key, errors.New("field does not exist")})
				continue
			}

			err := test.Run(value)

			if err != nil {
				errs = append(errs, &fieldError{key, err})
			}
		}
	}

	return errs, nil
}

func (ot *objectTest) Run(input interface{}) error {
	errs := ot.runMetaTests(input)

	fieldTestErrors, err := ot.runFieldTests(input)
	if err != nil {
		return err
	}

	errs = append(errs, fieldTestErrors...)
	if err != nil {
		return err
	}

	if len(errs) > 0 {
		return &objectTestError{
			errors: errs,
		}
	}

	return nil
}

func (ot *objectTest) IsMetaTest() bool {
	return false
}

func fieldTestForValue(input interface{}) (Test, error) {
	if castedString, ok := castString(input); ok {
		return NewStringTest(castedString), nil
	}

	if castedNumber, ok := castNumber(input); ok {
		return NewNumberTest(castedNumber), nil
	}

	if castedObject, ok := castObject(input); ok {
		test, err := NewObjectTest(castedObject)
		if err != nil {
			return nil, err
		}

		return test, nil
	}

	if castedArray, ok := castArray(input); ok {
		test, err := NewArrayTest(castedArray)
		if err != nil {
			return nil, err
		}

		return test, nil
	}

	return nil, fmt.Errorf("couldn't create test for value %#v", input)
}

func NewObjectTest(input map[string]interface{}) (Test, error) {
	tests := map[string]Test{}
	metaTests := map[string]Test{}
	for k, v := range input {
		if strings.HasPrefix(k, "$") {
			test, err := NewSpecialTest(k, v)
			if err != nil {
				return nil, err
			}

			metaTests[k] = test
			continue
		}

		test, err := fieldTestForValue(v)
		if err != nil {
			return nil, fmt.Errorf("couldn't create test for field %q with value %#v: %v", k, v, err)
		}

		tests[k] = test
	}

	return &objectTest{
		metaTests:  metaTests,
		fieldTests: tests,
	}, nil
}
