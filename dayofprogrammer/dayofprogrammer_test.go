package dayofprogrammer_test

import (
	"testing"

	"github.com/willmadison/hackerrank/dayofprogrammer"

	"github.com/stretchr/testify/assert"
)

func TestDayOfProgrammer(t *testing.T) {
	cases := []struct {
		name  string
		given struct {
			day, year int
		}
		expected string
	}{
		{
			"256th day of 2017",
			struct {
				day, year int
			}{256, 2017},
			"13.09.2017",
		},
		{
			"256th day of 2016",
			struct {
				day, year int
			}{256, 2016},
			"12.09.2016",
		},
		{
			"256th day of 1800",
			struct {
				day, year int
			}{256, 1800},
			"12.09.1800",
		},
		{
			"32nd day of 1918",
			struct {
				day, year int
			}{32, 1918},
			"14.02.1918",
		},
		{
			"33rd day of 1918",
			struct {
				day, year int
			}{33, 1918},
			"15.02.1918",
		},
		{
			"46th day of 1918",
			struct {
				day, year int
			}{46, 1918},
			"28.02.1918",
		},
		{
			"47th day of 1918",
			struct {
				day, year int
			}{47, 1918},
			"01.03.1918",
		},
		{
			"50th day of 1918",
			struct {
				day, year int
			}{50, 1918},
			"04.03.1918",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := dayofprogrammer.Get(c.given.day, c.given.year)
			assert.Equal(t, c.expected, actual)
		})
	}

}
