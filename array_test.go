package tests_test

import (
	"testing"

	"gitlab.com/fino/data-service-validator/tester/pkg/tests"
)

func TestArrayTestUnsupportedField(t *testing.T) {
	_, err := tests.NewArrayTest([]interface{}{1})

	if err == nil {
		t.Error("test should have not accpted integer as field")
	}
}

func TestArrayTestWrongType(t *testing.T) {
	test, err := tests.NewArrayTest([]interface{}{1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run(123)
	if err == nil {
		t.Error("should have failed when number was passed as input")
	}
}

func TestArrayTestWrongInput(t *testing.T) {
	test, err := tests.NewArrayTest([]interface{}{1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run([]interface{}{1.1})
	if err == nil {
		t.Error("should have failed when non matching input was passed")
	}
}

func TestArrayTestCorrectInput(t *testing.T) {
	test, err := tests.NewArrayTest([]interface{}{1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run([]interface{}{1.})
	if err != nil {
		t.Error("test should not have failed when given correct input")
	}
}
