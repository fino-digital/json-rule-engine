package loader

import (
	"fmt"
	"io/ioutil"

	"github.com/fino-digital/qjson"
	"gitlab.com/fino/data-service-validator/tester/pkg/tests"
)

func LoadFromFile(path string) (tests.Test, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read request file: %v", err)
	}

	parsedData, err := qjson.Unmarshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal qjson: %v", err)
	}

	test, err := tests.NewTest(parsedData)
	if err != nil {
		return nil, fmt.Errorf("parsed file could not be converted to test: %v", err)
	}

	return test, nil
}
