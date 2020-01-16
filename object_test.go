package tests_test

import (
	"testing"

	"gitlab.com/fino/data-service-validator/tester/pkg/tests"
)

func TestObjectTestUnsupportedField(t *testing.T) {
	_, err := tests.NewObjectTest(map[string]interface{}{"test": 1})

	if err == nil {
		t.Error("test should have not accpted integer as field")
	}
}

func TestObjectTestWrongType(t *testing.T) {
	test, err := tests.NewObjectTest(map[string]interface{}{"test": 1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run(123)
	if err == nil {
		t.Error("should have failed when number was passed as input")
	}
}

func TestObjectTestWrongInput(t *testing.T) {
	test, err := tests.NewObjectTest(map[string]interface{}{"test": 1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run(map[string]interface{}{"test": 1.1})
	if err == nil {
		t.Error("should have failed when non matching input was passed")
	}
}

func TestObjectTestCorrectInput(t *testing.T) {
	test, err := tests.NewObjectTest(map[string]interface{}{"test": 1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run(map[string]interface{}{"test": 1.})
	if err != nil {
		t.Error("test should not have failed when given correct input")
	}
}
