package main

import (
	"fmt"
	"strings"
)

func main() {
	var numTestCases int
	var numDigits int

	fmt.Scanf("%v", &numTestCases)

	for i := 0; i < numTestCases; i++ {
		fmt.Scanf("%v", &numDigits)
		fmt.Println(getDecentNumberByDigits(numDigits))
	}
}

func getDecentNumberByDigits(numDigits int) string {
	decentNumber := "-1"

	numFives := numDigits
	numThrees := 0

	for numFives%3 != 0 {
		numFives -= 5
	}

	if numFives >= 0 {
		numThrees = numDigits - numFives
		decentNumber = makeDecent(numFives, numThrees)
	}

	return decentNumber
}

func makeDecent(numFives, numThrees int) string {
	digits := []string{}

	for i := 0; i < numFives; i++ {
		digits = append(digits, "5")
	}

	for i := 0; i < numThrees; i++ {
		digits = append(digits, "3")
	}

	return strings.Join(digits, "")
}
