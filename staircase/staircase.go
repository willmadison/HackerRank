package main

import (
	"bytes"
	"fmt"
)

const (
	pad   = ' '
	pound = '#'
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
	var buf bytes.Buffer

	for i := 0; i < padding; i++ {
		buf.WriteRune(pad)
	}

	for i := 0; i < numPounds; i++ {
		buf.WriteRune(pound)
	}

	return buf.String()
}
