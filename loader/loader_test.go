package loader_test

import (
	"testing"

	"github.com/fino-digital/json-rule-engine/loader"
)

func TestLoadFromFile(t *testing.T) {
	test, err := loader.LoadFromFile("./test.qjson")
	if err != nil {
		t.Error(err)
		return
	}

	err = test.Run(map[string]interface{}{
		"pew": map[string]interface{}{
			"lol": "asd",
		},
	})

	if err != nil {
		t.Error(err)
		return
	}
}
