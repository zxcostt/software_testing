package factorial

import (
	"errors"
)

func factorial(n int) (int, error) {

	if n < 0 {
		return 0, errors.New("Факториал не определён для отрицательных чисел")
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result, nil
}
