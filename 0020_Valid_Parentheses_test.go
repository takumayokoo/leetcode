package main

import (
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
		AssertEqual(t, v.input, v.want, output)
	}
}
