package queens_test

import (
	"testing"

	"github.com/willmadison/HackerRank/queens"
)

func TestAttackableLocations(t *testing.T) {
	q := queens.Location{
		Row:    4,
		Column: 4,
	}
	dimensions := queens.Dimension{
		Width:  4,
		Length: 4,
	}
	obstacles := []queens.Location{}

	n := queens.AttackableLocations(q, dimensions, obstacles)

	if n != 9 {
		t.Error("want:", 9, "got:", n)
	}
}
