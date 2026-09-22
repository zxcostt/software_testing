package factorial

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name          string
		input         int //входнные данные
		expected      int //ожидаемый результат
		expectedError bool
	}{
		{name: "Факториал 5", input: 5, expected: 120, expectedError: false},
		{name: "Факториал 0", input: 0, expected: 1, expectedError: false},
		{name: "Отрицательное число", input: -3, expected: 0, expectedError: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := factorial(tc.input)

			//Проверяем ожидание ошибки
			if tc.expectedError {
				if err == nil {
					t.Errorf("для %d ожидали ошибку, но получили nil", tc.input)
				}
				return
			}

			//Проверяем неожидаемую ошибку
			if err != nil {
				t.Errorf("для %d получили неожиданную ошибку: %v", tc.input, err)
			}

			//
			if result != tc.expected {
				t.Errorf("для %d ожидали %d, но получили %d", tc.input, tc.expected, result)
			}
		})
	}
}
