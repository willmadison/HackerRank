package between

func Count(a, b []int) int {
	factors := commonFactors(b)
	between := []int{}

	for _, f := range factors {
		isBetween := true

		for _, n := range a {
			if f%n != 0 {
				isBetween = false
				break
			}
		}

		if isBetween {
			between = append(between, f)
		}
	}

	return len(between)
}

func commonFactors(numbers []int) []int {
	factors := []int{1}
	max := findMax(numbers)

	i := 2

	for i <= max {
		common := true

		for _, n := range numbers {
			if n%i != 0 {
				common = false
				break
			}
		}

		if common {
			factors = append(factors, i)
		}

		i++
	}

	return factors
}

func findMax(numbers []int) int {
	max := numbers[0]

	for _, n := range numbers {
		if n > max {
			max = n
		}
	}

	return max
}
