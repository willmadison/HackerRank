package candles_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/HackerRank/candles"
)

// https://www.hackerrank.com/challenges/birthday-cake-candles/problem

func TestBirthdayCakeCandles(t *testing.T) {
	cases := []struct {
		name        string
		given       []int
		expected    int
		shouldError bool
	}{
		{
			"list has various numbers",
			[]int{1, 2, 3, 4},
			1,
			false,
		},
		{
			"list numbers are all the same",
			[]int{4, 4, 4, 4},
			4,
			false,
		},
		{
			"list has no numbers",
			[]int{},
			0,
			true,
		},
	}

	// Is what I'm doing making sense?
	// Basically planning to do a table driven test for each of the scenarios you talked about above...
	// Would the other case[s] be put in the struct above?
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := candles.BirthdayCakeCandles(tc.given)

			if err != nil && !tc.shouldError {
				t.Errorf("should not have errored in this scenario")
			}

			if err == nil && tc.shouldError {
				t.Errorf("should have errored in this scenario")
			}

			assert.Equal(t, tc.expected, actual)

		})
	}
}
