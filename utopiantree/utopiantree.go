package main

import "fmt"

type Season int

const (
	Spring Season = iota
	Summer
)

func main() {
	var numTestCases int

	fmt.Scanf("%v", &numTestCases)

	var numCycles int

	for i := 0; i < numTestCases; i++ {
		fmt.Scanf("%v", &numCycles)
		fmt.Println(determineHeightAfter(numCycles))
	}
}

func determineHeightAfter(numCycles int) int {
	height := 1

	season := Spring

	for i := 0; i < numCycles; i++ {
		switch season {
		case Spring:
			height *= 2
			season = Summer
		case Summer:
			height++
			season = Spring
		}
	}

	return height
}
