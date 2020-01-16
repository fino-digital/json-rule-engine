package tests_test

import (
	"testing"

	tests "github.com/fino-digital/json-rule-engine"
)

func TestNewTestWrongInput(t *testing.T) {
	_, err := tests.NewTest(1.)
	if err == nil {
		t.Error("should have failed when non object was passed")
	}
}

func TestNewTestCorrectInput(t *testing.T) {
	_, err := tests.NewTest([]interface{}{1.})
	if err != nil {
		t.Errorf("test creation should have succeeded: %v", err)
	}
}

func TestBuilder(t *testing.T) {
	test, err := tests.NewObjectTest(map[string]interface{}{
		"a": map[string]interface{}{
			"aa": "a",
			"bb": "b",
			"cc": map[string]interface{}{
				"aaa": "c",
			},
		},
		"b": map[string]interface{}{
			"$startsWith": "a",
		},
		"c": map[string]interface{}{
			"$all": map[string]interface{}{"aa": "a"},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	err = test.Run(map[string]interface{}{
		"a": map[string]interface{}{
			"aa": "a",
			"bb": "b",
			"cc": map[string]interface{}{
				"aaa": "c",
			},
		},
		"b": "aa",
		"c": []interface{}{
			map[string]interface{}{
				"aa": "a",
			},
			map[string]interface{}{
				"aa": "a",
			},
		},
	})

	if err != nil {
		t.Error(err)
	}
}
