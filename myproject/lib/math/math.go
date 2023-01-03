package math

// Sum calculates the sum of a list of numbers
func Sum(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
