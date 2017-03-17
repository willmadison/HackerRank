package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numTestCases int
	var value int

	fmt.Scanf("%v", &numTestCases)

	for i := 0; i < numTestCases; i++ {
		fmt.Scanf("%v", &value)
		fmt.Println(findDivisibleDigits(value))
	}
}

func findDivisibleDigits(number int) int {
	divisibleDigits := 0
	asString := strconv.Itoa(number)

	digits := strings.Split(asString, "")

	for _, digit := range digits {
		n, _ := strconv.Atoi(digit)

		if n > 0 && number%n == 0 {
			divisibleDigits++
		}
	}

	return divisibleDigits
}
