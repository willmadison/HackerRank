package main

import "fmt"

func main() {
	var numTestCases int

	fmt.Scanf("%v", &numTestCases)

	for i := 0; i < numTestCases; i++ {
		var numStudents int
		var threshold int

		fmt.Scanf("%v %v", &numStudents, &threshold)

		var value int

		var numStudentsOnTime int

		for j := 0; j < numStudents; j++ {
			fmt.Scanf("%v", &value)

			if value <= 0 {
				numStudentsOnTime++
			}
		}

		decision := "YES"

		if numStudentsOnTime >= threshold {
			decision = "NO"
		}

		fmt.Println(decision)
	}
}
