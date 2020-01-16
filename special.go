package tests

import (
	"errors"
	"fmt"
)

var UnknownKeywordError = errors.New("unknown keyword")

func NewSpecialTest(keyword string, input interface{}) (Test, error) {
	switch keyword {
	case "$length":
		return NewLengthTest(input)
	case "$all":
		return NewAllTest(input)
	case "$some":
		return NewSomeTest(input)
	case "$none":
		return NewNoneTest(input)
	case "$contains":
		return NewContainsTest(input)
	case "$containsSome":
		return NewContainsSomeTest(input)
	case "$containsNone":
		return NewContainsNoneTest(input)
	case "$startsWith":
		return NewStartsWithTest(input)
	}

	if number, ok := castNumber(input); ok {
		switch keyword {
		case "$lt":
			return NewLessThanTest(number), nil
		case "$lte":
			return NewLessThanEqualTest(number), nil
		case "$gt":
			return NewGreaterThanTest(number), nil
		case "$gte":
			return NewGreaterThanEqualTest(number), nil
		}
	}

	return nil, fmt.Errorf("%w (%q)", UnknownKeywordError, keyword)
}
