package main

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			"()",
			true,
		},
		{
			"()[]{}",
			true,
		},
		{
			"(]",
			false,
		},
		{
			"{[]}",
			true,
		},
		{
			"[",
			false,
		},
		{
			"]",
			false,
		},
	}

	for _, v := range tests {
		output := isValid(v.input)
		AssertEqual(t, v.want, output, fmt.Sprintf("\nInput %v\nExpected %v\nOutput %v\n", v.input, v.want, output))
	}
}
