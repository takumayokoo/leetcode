package main

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{
			[]string{"hello", "ll"},
			2,
		},
		{
			[]string{"aaaaa", "bba"},
			-1,
		},
		{
			[]string{"", ""},
			0,
		},
		{
			[]string{"", "a"},
			-1,
		},
		{
			[]string{"a", ""},
			0,
		},
		{
			[]string{"a", "a"},
			0,
		},
	}

	for _, v := range tests {
		output := strStr(v.input[0], v.input[1])
		AssertEqual(t, v.want, output, fmt.Sprintf("\nInput %v\nExpected %v\nOutput %v\n", v.input, v.want, output))
	}
}
