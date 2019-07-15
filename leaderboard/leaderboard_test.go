package leaderboard_test

import (
	"testing"

	"github.com/willmadison/hackerrank/leaderboard"

	"github.com/stretchr/testify/assert"
)

func TestLeaderboardClimb(t *testing.T) {
	cases := []struct {
		given struct {
			scoreboard  []int
			aliceScores []int
		}
		expected []int
	}{
		{
			given: struct {
				scoreboard  []int
				aliceScores []int
			}{
				scoreboard:  []int{100, 100, 50, 40, 40, 20, 10},
				aliceScores: []int{5, 25, 50, 120},
			},
			expected: []int{6, 4, 2, 1},
		},
	}

	for _, c := range cases {
		actual := leaderboard.Climb(c.given.scoreboard, c.given.aliceScores)
		assert.Equal(t, c.expected, actual)
	}

}
