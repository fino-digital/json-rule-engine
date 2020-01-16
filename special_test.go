package tests_test

import (
	"errors"
	"testing"

	tests "github.com/fino-digital/json-rule-engine"
)

func TestSpecialTestUnknownKeyword(t *testing.T) {
	_, err := tests.NewSpecialTest("test", 1.)

	if !errors.Is(err, tests.UnknownKeywordError) {
		t.Error("test should have only accepted strings as input")
	}
}

func TestSpecialTestCorrectInput(t *testing.T) {
	_, err := tests.NewSpecialTest("$length", 1.)
	if err != nil {
		t.Error("failed to create test", err)
	}
}
