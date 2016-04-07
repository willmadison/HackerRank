package main

import (
	"fmt"
	"time"
)

const inputFormat = "03:04:05PM"
const outputFormat = "15:04:05"

func main() {
	var timestamp string

	fmt.Scanf("%v", &timestamp)

	givenTime, err := time.Parse(inputFormat, timestamp)

	if err != nil {
		fmt.Println("Encountered an error:", err)
		return
	}

	fmt.Println(givenTime.Format(outputFormat))
}
