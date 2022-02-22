package main

import (
	"fmt"
	"testing"
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
		AssertEqual(t, v.want, output, fmt.Sprintf("\nInput %v\nExpected %v\nOutput %v\n", v.input, v.want, output))
	}
}
