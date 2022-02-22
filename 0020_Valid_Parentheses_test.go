package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, v.want, output, "Input %v", v.input)
	}
}
