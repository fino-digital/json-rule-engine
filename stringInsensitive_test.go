package tests_test

import (
	"testing"

	tests "github.com/fino-digital/json-rule-engine"
)

func TestInsensitiveStringTestWrongType(t *testing.T) {
	test, err := tests.NewCaseInsensitiveStringTest("test")

	err = test.Run(123)
	if err == nil {
		t.Error("should have failed when number was passed as input")
	}
}

func TestInsensitiveStringTestWrongInput(t *testing.T) {
	test, err := tests.NewCaseInsensitiveStringTest("test")

	err = test.Run("pew")
	if err == nil {
		t.Error("should have failed when non matching input was passed")
	}
}

func TestInsensitiveStringTestCorrectInput(t *testing.T) {
	test, err := tests.NewCaseInsensitiveStringTest("Test")

	err = test.Run("test")
	if err != nil {
		t.Error("test should not have failed when given correct input")
	}
}
