package tests

// castNumber tries to cast to a float64
func castNumber(in interface{}) (float64, bool) {
	casted, ok := in.(float64)
	return casted, ok
}

// castString tries to cast to a string
func castString(in interface{}) (string, bool) {
	casted, ok := in.(string)
	return casted, ok
}

// castObject tries to cast to a generic object
func castObject(in interface{}) (map[string]interface{}, bool) {
	casted, ok := in.(map[string]interface{})
	return casted, ok
}

// castArray tries to cast to an array
func castArray(in interface{}) ([]interface{}, bool) {
	casted, ok := in.([]interface{})
	return casted, ok
}
