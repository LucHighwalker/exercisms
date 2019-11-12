package diffsquares

import (
	"math"
)

func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return int(math.Pow(float64(sum), 2.0))
}

func SumOfSquares(n int) int {
	nf := float64(n)

	sum := 0.0
	for i := 1.0; i <= nf; i++ {
		sum += math.Pow(i, 2.0)
	}

	return int(sum)
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
