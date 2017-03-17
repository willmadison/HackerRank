package main

import (
	"fmt"
	"math"
)

func main() {
	var numTestCases int

	fmt.Scanf("%v", &numTestCases)

	for i := 0; i < numTestCases; i++ {
		var lowerBound float64
		var upperBound float64

		fmt.Scanf("%v %v", &lowerBound, &upperBound)

		numPerfectSquares := int(math.Floor(math.Sqrt(upperBound)) - math.Ceil(math.Sqrt(lowerBound)) + 1)

		fmt.Println(numPerfectSquares)
	}
}
