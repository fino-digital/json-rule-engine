package tests_test

import (
	"testing"

	"gitlab.com/fino/data-service-validator/tester/pkg/tests"
)

func TestNoneTestUnsupportedField(t *testing.T) {
	_, err := tests.NewNoneTest([]interface{}{1})

	if err == nil {
		t.Error("test should have not accpted array as expected value")
	}
}

func TestNoneTestWrongType(t *testing.T) {
	test, err := tests.NewNoneTest(map[string]interface{}{"test": 1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run(123)
	if err == nil {
		t.Error("should have failed when number was passed as input")
	}
}

func TestNoneTestWrongInputType(t *testing.T) {
	test, err := tests.NewNoneTest(map[string]interface{}{"test": 1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run(1.1)
	if err == nil {
		t.Error("should have failed when non matching input was passed")
	}
}

func TestNoneTestCorrectInput(t *testing.T) {
	test, err := tests.NewNoneTest(map[string]interface{}{"test": 1., "woop": "asd"})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run([]interface{}{map[string]interface{}{"test": "heh"}})
	if err != nil {
		t.Error("should have succeeded when non matching input was passed", err)
	}
}

func TestNoneTestWrongInput(t *testing.T) {
	test, err := tests.NewNoneTest(map[string]interface{}{"test": 1., "woop": "asd"})
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
			"woop": "asdaaa",
		},
	})

	if err == nil {
		t.Error("should not have when matching input was passed: ", err)
	}
}
