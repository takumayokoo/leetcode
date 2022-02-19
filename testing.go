package main

import (
	"testing"
)

func AssertEqual(t *testing.T, input interface{}, expected interface{}, output interface{}) {
	if output != expected {
		t.Errorf("\nInput %v\nOutput %v\nExpected %v", input, output, expected)
	}
}