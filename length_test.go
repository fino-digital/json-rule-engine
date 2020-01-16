package tests_test

import (
	"testing"

	"gitlab.com/fino/data-service-validator/tester/pkg/tests"
)

func TestLengthTestWrongExpectedValue(t *testing.T) {
	_, err := tests.NewStartsWithTest(123)

	if err == nil {
		t.Error("test should have only accepted strings as input")
	}
}

func TestLengthTestWrongType(t *testing.T) {
	test, err := tests.NewStartsWithTest("test")
	if err != nil {
		t.Error("failed to create test", err)
	}

	err = test.Run(123)
	if err == nil {
		t.Error("should have failed when number was passed as input")
	}
}

func TestLengthTestWrongInput(t *testing.T) {
	test, err := tests.NewStartsWithTest("test")
	if err != nil {
		t.Error("failed to create test", err)
	}

	err = test.Run("pew")
	if err == nil {
		t.Error("should have failed when non matching input was passed")
	}
}

func TestLengthTestCorrectInput(t *testing.T) {
	test, err := tests.NewStartsWithTest("test")
	if err != nil {
		t.Error("failed to create test", err)
	}

	err = test.Run("test2")
	if err != nil {
		t.Error("test should not have failed when given correct input")
	}
}
