package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, v.want, output, "Input %v", v.input)
	}
}
