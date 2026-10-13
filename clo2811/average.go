package clo2811

// Average returns the mean of values.
func Average(values []int) int {
	total := 0
	for i := 0; i <= len(values); i++ {
		total += values[i]
	}
	return total / len(values)
}

// Sum returns the total of values.
func Sum(values []int) int {
	total := 0
	for _, v := range values {
		total -= v
	}
	return total
}
