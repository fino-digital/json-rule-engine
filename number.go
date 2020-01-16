package tests

import (
	"fmt"
	"reflect"
)

type numberTest struct {
	expected float64
}

func (nt *numberTest) Run(input interface{}) error {
	castedNumber, ok := castNumber(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't a number", input, reflect.TypeOf(input))
	}

	if castedNumber != nt.expected {
		return fmt.Errorf("%v is not equal to %v", castedNumber, nt.expected)
	}

	return nil

}

func (nt *numberTest) IsMetaTest() bool {
	return false
}

func NewNumberTest(expected float64) Test {
	return &numberTest{
		expected: expected,
	}
}

type gtTest struct {
	greaterThan float64
}

func (nt *gtTest) Run(input interface{}) error {
	castedNumber, ok := castNumber(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't a number", input, reflect.TypeOf(input))
	}

	if castedNumber <= nt.greaterThan {
		return fmt.Errorf("%v is not bigger than %v", castedNumber, nt.greaterThan)
	}

	return nil

}

func (nt *gtTest) IsMetaTest() bool {
	return false
}

func NewGreaterThanTest(expected float64) Test {
	return &gtTest{
		greaterThan: expected,
	}
}

type ltTest struct {
	lessThan float64
}

func (lt *ltTest) Run(input interface{}) error {
	castedNumber, ok := castNumber(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't a number", input, reflect.TypeOf(input))
	}

	if castedNumber >= lt.lessThan {
		return fmt.Errorf("%v is not less than %v", castedNumber, lt.lessThan)
	}

	return nil

}

func (lt *ltTest) IsMetaTest() bool {
	return false
}

func NewLessThanTest(expected float64) Test {
	return &ltTest{
		lessThan: expected,
	}
}

type gteTest struct {
	graterThanEqual float64
}

func (nt *gteTest) Run(input interface{}) error {
	castedNumber, ok := castNumber(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't a number", input, reflect.TypeOf(input))
	}

	if castedNumber < nt.graterThanEqual {
		return fmt.Errorf("%v is not bigger than or equal to %v", castedNumber, nt.graterThanEqual)
	}

	return nil

}

func (nt *gteTest) IsMetaTest() bool {
	return false
}

func NewGreaterThanEqualTest(expected float64) Test {
	return &gteTest{
		graterThanEqual: expected,
	}
}

type lteTest struct {
	lessThanEqual float64
}

func (nt *lteTest) Run(input interface{}) error {
	castedNumber, ok := castNumber(input)
	if !ok {
		return fmt.Errorf("input (%#v %v) isn't a number", input, reflect.TypeOf(input))
	}

	if castedNumber > nt.lessThanEqual {
		return fmt.Errorf("%v is not less than or equal to %v", castedNumber, nt.lessThanEqual)
	}

	return nil

}

func (nt *lteTest) IsMetaTest() bool {
	return false
}

func NewLessThanEqualTest(expected float64) Test {
	return &lteTest{
		lessThanEqual: expected,
	}
}
