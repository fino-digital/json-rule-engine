package tests_test

import (
	"testing"

	tests "github.com/fino-digital/json-rule-engine"
)

func TestStartsWithTestWrongExpectedValue(t *testing.T) {
	_, err := tests.NewStartsWithTest(123)

	if err == nil {
		t.Error("test should have only accepted strings as input")
	}
}

func TestStartsWithTestWrongType(t *testing.T) {
	test, err := tests.NewStartsWithTest("test")
	if err != nil {
		t.Error("failed to create test", err)
	}

	err = test.Run(123)
	if err == nil {
		t.Error("should have failed when number was passed as input")
	}
}

func TestStartsWithTestWrongInput(t *testing.T) {
	test, err := tests.NewStartsWithTest("test")
	if err != nil {
		t.Error("failed to create test", err)
	}

	err = test.Run("pew")
	if err == nil {
		t.Error("should have failed when non matching input was passed")
	}
}

func TestStartsWithTestCorrectInput(t *testing.T) {
	test, err := tests.NewStartsWithTest("test")
	if err != nil {
		t.Error("failed to create test", err)
	}

	err = test.Run("test2")
	if err != nil {
		t.Error("test should not have failed when given correct input")
	}
}
