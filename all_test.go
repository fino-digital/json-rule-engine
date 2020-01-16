package tests_test

import (
	"testing"

	tests "github.com/fino-digital/json-rule-engine"
)

func TestAllTestUnsupportedField(t *testing.T) {
	_, err := tests.NewAllTest([]interface{}{1})

	if err == nil {
		t.Error("test should have not accpted array as expected value")
	}
}

func TestAllTestWrongType(t *testing.T) {
	test, err := tests.NewAllTest(map[string]interface{}{"test": 1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run(123)
	if err == nil {
		t.Error("should have failed when number was passed as input")
	}
}

func TestAllTestWrongInputType(t *testing.T) {
	test, err := tests.NewAllTest(map[string]interface{}{"test": 1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run(1.1)
	if err == nil {
		t.Error("should have failed when non matching input was passed")
	}
}

func TestAllTestWrongInput(t *testing.T) {
	test, err := tests.NewAllTest(map[string]interface{}{"test": 1., "woop": "asd"})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run([]interface{}{map[string]interface{}{"test": "heh"}})
	if err == nil {
		t.Error("should have failed when non matching input was passed")
	}
}

func TestAllTestCorrectInput(t *testing.T) {
	test, err := tests.NewAllTest(map[string]interface{}{"test": 1., "woop": "asd"})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run([]interface{}{
		map[string]interface{}{
			"test": 1.,
			"woop": "asd",
		},
		map[string]interface{}{
			"test": 1.,
			"woop": "asd",
		},
	})
	if err != nil {
		t.Error("should not have failed when correct input was passed: ", err)
	}
}
