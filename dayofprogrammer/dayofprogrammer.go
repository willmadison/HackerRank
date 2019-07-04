package dayofprogrammer

import (
	"fmt"
	"time"
)

func Get(day, year int) string {
	calendar := getCalendarByYear(year)
	return calendar.get(day)
}

func getCalendarByYear(year int) calendar {
	switch {
	case 1700 <= year && year <= 1917:
		return newJulianCalendar(year)
	default:
		return newGregorianCalendar(year)
	}
}

func newJulianCalendar(year int) calendar {
	daysByMonth := map[time.Month]int{
		time.January:   31,
		time.February:  28,
		time.March:     31,
		time.April:     30,
		time.May:       31,
		time.June:      30,
		time.July:      31,
		time.August:    31,
		time.September: 30,
		time.October:   31,
		time.November:  30,
		time.December:  31,
	}

	if isJulianLeapYear(year) {
		daysByMonth[time.February]++
	}

	return calendar{year, daysByMonth}
}

func newGregorianCalendar(year int) calendar {
	daysByMonth := map[time.Month]int{
		time.January:   31,
		time.February:  28,
		time.March:     31,
		time.April:     30,
		time.May:       31,
		time.June:      30,
		time.July:      31,
		time.August:    31,
		time.September: 30,
		time.October:   31,
		time.November:  30,
		time.December:  31,
	}

	if isGregorianLeapYear(year) {
		daysByMonth[time.February]++
	}

	return calendar{year, daysByMonth}
}

func isJulianLeapYear(year int) bool {
	return year%4 == 0
}

func isGregorianLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

type calendar struct {
	year        int
	daysByMonth map[time.Month]int
}

func (c calendar) get(day int) string {
	var daysTranspired int
	months := []time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}

	var month time.Month
	var dayInMonth int

	for _, m := range months {
		daysInMonth := c.daysByMonth[m]

		if c.year == 1918 && m == time.February {
			daysInMonth -= 13
		}

		if daysTranspired+daysInMonth < day {
			daysTranspired += daysInMonth
		} else {
			month = m
			dayInMonth = day - daysTranspired

			if c.year == 1918 && m == time.February {
				dayInMonth += 13
			}
			break
		}
	}

	return fmt.Sprintf("%02d.%02d.%d", dayInMonth, month, c.year)
}
