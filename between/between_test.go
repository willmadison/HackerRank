package between_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/HackerRank/between"
)

// https://www.hackerrank.com/challenges/between-two-sets/problem

func TestBetween(t *testing.T) {
	cases := []struct {
		name     string
		a, b     []int
		expected int
	}{
		{
			"First example case",
			[]int{2, 6},
			[]int{24, 36},
			2,
		},
		{
			"Second example case",
			[]int{2, 4},
			[]int{16, 32, 96},
			3,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := between.Count(tc.a, tc.b)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
