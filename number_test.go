package tests_test

import (
	"testing"

	tests "github.com/fino-digital/json-rule-engine"
)

func TestNumberTestWrongType(t *testing.T) {
	test := tests.NewNumberTest(123.)

	err := test.Run("test")
	if err == nil {
		t.Error("should have failed when string was passed as input")
	}
}

func TestNumberTestWrongInput(t *testing.T) {
	test := tests.NewNumberTest(123.)

	err := test.Run(321.)
	if err == nil {
		t.Error("should have failed when non matching input was passed")
	}
}

func TestNumberTestCorrectInput(t *testing.T) {
	test := tests.NewNumberTest(123.)

	err := test.Run(123.)
	if err != nil {
		t.Error("test should not have failed when given correct input")
	}
}

func TestGreaterThanTestWrongType(t *testing.T) {
	test := tests.NewGreaterThanTest(123.)

	err := test.Run("test")
	if err == nil {
		t.Error("should have failed when string was passed as input")
	}
}

func TestGreaterThanTestWrongInput(t *testing.T) {
	test := tests.NewGreaterThanTest(123.)

	err := test.Run(12.)
	if err == nil {
		t.Error("should have failed when non matching input was passed")
	}
}

func TestGreaterThanTestCorrectInput(t *testing.T) {
	test := tests.NewGreaterThanTest(123.)

	err := test.Run(1234.)
	if err != nil {
		t.Error("test should not have failed when given correct input")
	}
}

func TestLessThanTestWrongType(t *testing.T) {
	test := tests.NewLessThanTest(123.)

	err := test.Run("test")
	if err == nil {
		t.Error("should have failed when string was passed as input")
	}
}

func TestLessThanTestWrongInput(t *testing.T) {
	test := tests.NewLessThanTest(123.)

	err := test.Run(1234.)
	if err == nil {
		t.Error("should have failed when non matching input was passed")
	}
}

func TestLessThanTestCorrectInput(t *testing.T) {
	test := tests.NewLessThanTest(123.)

	err := test.Run(12.)
	if err != nil {
		t.Error("test should not have failed when given correct input")
	}
}

func TestGreaterThanEqualTestWrongType(t *testing.T) {
	test := tests.NewGreaterThanEqualTest(123.)

	err := test.Run("test")
	if err == nil {
		t.Error("should have failed when string was passed as input")
	}
}

func TestGreaterThanEqualTestWrongInput(t *testing.T) {
	test := tests.NewGreaterThanEqualTest(123.)

	err := test.Run(12.)
	if err == nil {
		t.Error("should have failed when non matching input was passed")
	}
}

func TestGreaterThanEqualTestCorrectInput(t *testing.T) {
	test := tests.NewGreaterThanEqualTest(123.)

	err := test.Run(123.)
	if err != nil {
		t.Error("test should not have failed when given correct input")
	}

	err = test.Run(1234.)
	if err != nil {
		t.Error("test should not have failed when given correct input")
	}
}

func TestLessThanEqualTestWrongType(t *testing.T) {
	test := tests.NewLessThanEqualTest(123.)

	err := test.Run("test")
	if err == nil {
		t.Error("should have failed when string was passed as input")
	}
}

func TestLessThanEqualTestWrongInput(t *testing.T) {
	test := tests.NewLessThanEqualTest(123.)

	err := test.Run(1234.)
	if err == nil {
		t.Error("should have failed when non matching input was passed")
	}
}

func TestLessThanEqualTestCorrectInput(t *testing.T) {
	test := tests.NewLessThanEqualTest(123.)

	err := test.Run(123.)
	if err != nil {
		t.Error("test should not have failed when given correct input")
	}

	err = test.Run(12.)
	if err != nil {
		t.Error("test should not have failed when given correct input")
	}
}
