package grains

import (
	"errors"
	"math"
)

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("error")
	}

	return uint64(math.Pow(2, float64(n-1))), nil
}

func Total() uint64 {
	var total uint64 = 0

	for i := 1; i <= 64; i++ {
		result, _ := Square(i)
		total += result
	}
	return total
}
