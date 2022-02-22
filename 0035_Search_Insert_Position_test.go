package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		input []int
		target int
		want  int
	}{
		{
			[]int{1,3,5,6},
			5,
			2,
		},
		{
			[]int{1,3,5,6},
			2,
			1,
		},
		{
			[]int{1,3,5,6},
			7,
			4,
		},
	}

	for _, v := range tests {
		result := searchInsert(v.input, v.target)
		assert.Equal(t, v.want, result, "Input %v\nTarget %v", v.input, v.target)
	}
}
