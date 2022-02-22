package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input []string
		want string
	}{
		{
			[]string{"", ""},
			"",
		},
		{
			[]string{"", "a"},
			"",
		},
		{
			[]string{"a", "ab"},
			"a",
		},
		{
			[]string{"ab", "abd", "abc"},
			"ab",
		},
	}

	for _, v := range(tests) {
		output := longestCommonPrefix(v.input)
		assert.Equal(t, v.want, output, "Input %v", v.input)
	}
}
