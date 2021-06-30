package diffsquares

import "math"

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

func SquareOfSum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += int(math.Pow(float64(i), 2.0))
	}
	return sum
}
