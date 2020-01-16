package tests_test

import (
	"testing"

	"gitlab.com/fino/data-service-validator/tester/pkg/tests"
)

func TestStringTestWrongType(t *testing.T) {
	test := tests.NewStringTest("test")

	err := test.Run(123)
	if err == nil {
		t.Error("should have failed when number was passed as input")
	}
}

func TestStringTestWrongInput(t *testing.T) {
	test := tests.NewStringTest("test")

	err := test.Run("pew")
	if err == nil {
		t.Error("should have failed when non matching input was passed")
	}
}

func TestStringTestCorrectInput(t *testing.T) {
	test := tests.NewStringTest("test")

	err := test.Run("test")
	if err != nil {
		t.Error("test should not have failed when given correct input")
	}
}
