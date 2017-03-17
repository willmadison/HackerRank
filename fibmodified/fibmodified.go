package main

import (
	"fmt"
	"math/big"
)

var resultsByTerm map[int]*big.Int = map[int]*big.Int{}

func main() {
	var firstTerm int
	var secondTerm int
	var desiredTerm int

	fmt.Scanf("%v %v %v", &firstTerm, &secondTerm, &desiredTerm)

	a := big.NewInt(int64(firstTerm))
	b := big.NewInt(int64(secondTerm))

	currentTerm := 2

	for currentTerm < desiredTerm {
		c := big.NewInt(0)

		c.Mul(b, b)
		c.Add(c, a)

		a, b = b, c

		currentTerm++
	}

	fmt.Println(b)
}
