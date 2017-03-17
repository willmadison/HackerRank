package main

import (
	"fmt"
	"strings"
)

const (
	pad   = " "
	pound = "#"
)

func main() {
	var height int

	fmt.Scanf("%v", &height)

	for i := 1; i <= height; i++ {
		padding := height - i
		fmt.Println(getStairRow(i, padding))
	}
}

func getStairRow(numPounds, padding int) string {
	stairRow := []string{}

	for i := 0; i < padding; i++ {
		stairRow = append(stairRow, pad)
	}

	for i := 0; i < numPounds; i++ {
		stairRow = append(stairRow, pound)
	}

	return strings.Join(stairRow, "")
}
