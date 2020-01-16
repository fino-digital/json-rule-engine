package tests_test

import (
	"testing"

	tests "github.com/fino-digital/json-rule-engine"
)

func TestContainsTestWrongExpectedType(t *testing.T) {
	_, err := tests.NewContainsTest(map[string]interface{}{"hey": 1.})

	if err == nil {
		t.Error("test should have not accepted map as expected value")
	}
}

func TestContainsTestWrongType(t *testing.T) {
	test, err := tests.NewContainsTest([]interface{}{1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run(123)
	if err == nil {
		t.Error("should have failed when number was passed as input")
	}
}

func TestContainsTestWrongInputType(t *testing.T) {
	test, err := tests.NewContainsTest([]interface{}{1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run(1.1)
	if err == nil {
		t.Error("should have failed when non array was passed as input")
	}
}

func TestContainsTestWrongInput(t *testing.T) {
	test, err := tests.NewContainsTest([]interface{}{1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run([]interface{}{2., "asd"})
	if err == nil {
		t.Error("should have failed when non matching input was passed", err)
	}
}

func TestContainsTestCorrectInput(t *testing.T) {
	test, err := tests.NewContainsTest([]interface{}{1., "asd"})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run([]interface{}{
		"pew", "asd", 1.,
	})
	if err != nil {
		t.Error("should not have failed when correct input was passed: ", err)
	}
}

func TestContainsSomeTestWrongExpectedType(t *testing.T) {
	_, err := tests.NewContainsSomeTest(map[string]interface{}{"hey": 1.})

	if err == nil {
		t.Error("test should have not accepted map as expected value")
	}
}

func TestContainsSomeTestWrongType(t *testing.T) {
	test, err := tests.NewContainsSomeTest([]interface{}{1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run(123)
	if err == nil {
		t.Error("should have failed when number was passed as input")
	}
}

func TestContainsSomeTestWrongInputType(t *testing.T) {
	test, err := tests.NewContainsSomeTest([]interface{}{1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run(1.1)
	if err == nil {
		t.Error("should have failed when non array was passed as input")
	}
}

func TestContainsSomeTestWrongInput(t *testing.T) {
	test, err := tests.NewContainsSomeTest([]interface{}{1., "pew"})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run([]interface{}{2., "asd"})
	if err == nil {
		t.Error("should have failed when non matching input was passed", err)
	}
}

func TestContainsSomeTestCorrectInput(t *testing.T) {
	test, err := tests.NewContainsSomeTest([]interface{}{1., "asd"})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run([]interface{}{
		"pew", 1.,
	})
	if err != nil {
		t.Error("should not have failed when correct input was passed: ", err)
	}
}

func TestContainsNoneTestWrongExpectedType(t *testing.T) {
	_, err := tests.NewContainsNoneTest(map[string]interface{}{"hey": 1.})

	if err == nil {
		t.Error("test should have not accepted map as expected value")
	}
}

func TestContainsNoneTestWrongType(t *testing.T) {
	test, err := tests.NewContainsNoneTest([]interface{}{1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run(123)
	if err == nil {
		t.Error("should have failed when number was passed as input")
	}
}

func TestContainsNoneTestWrongInputType(t *testing.T) {
	test, err := tests.NewContainsNoneTest([]interface{}{1.})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run(1.1)
	if err == nil {
		t.Error("should have failed when non array was passed as input")
	}
}

func TestContainsNoneTestWrongInput(t *testing.T) {
	test, err := tests.NewContainsNoneTest([]interface{}{1., "pew"})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run([]interface{}{"pew", "asd"})
	if err == nil {
		t.Error("should have failed when non matching input was passed", err)
	}
}

func TestContainsNoneTestCorrectInput(t *testing.T) {
	test, err := tests.NewContainsNoneTest([]interface{}{1., "asd"})
	if err != nil {
		t.Error("failed to create test", err)
		return
	}

	err = test.Run([]interface{}{
		"pew", 2.,
	})
	if err != nil {
		t.Error("should not have failed when correct input was passed: ", err)
	}
}
