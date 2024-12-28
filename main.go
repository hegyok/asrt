package main

import (
	"fmt"
)

// Equal compares two values and panics with a detailed message if they are not equal
func Equal(expected interface{}, actual interface{}, message string) {
	if expected != actual {
		panic(fmt.Sprintf("Assertion failed: %s\nExpected: %v\nActual: %v", message, expected, actual))
	}
}

// NotEqual compares two values and panics with a detailed message if they are equal
func NotEqual(expected interface{}, actual interface{}, message string) {
	if expected == actual {
		panic(fmt.Sprintf("Assertion failed: %s\nExpected not equal to: %v", message, expected))
	}
}

// NotNil checks if a value is non-nil and panics with a detailed message if it is nil
func NotNil(value interface{}, message string) {
	if value == nil {
		panic(fmt.Sprintf("Assertion failed: %s\nExpected non-nil value", message))
	}
}

// Err checks if an error is nil and panics with a detailed message if it isn't nil
func Err(err error, message string) {
	if err != nil {
		panic(fmt.Sprintf("Assertion failed: %s\nUnexpected error: %v", message, err))
	}
}

// Assert evaluates a boolean condition and panics with the provided message if it's false
func Assert(condition bool, message string) {
	if !condition {
		panic(fmt.Sprintf("Assertion failed: %s", message))
	}
}
