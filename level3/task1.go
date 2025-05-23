package main

import (
	"fmt"
)

func CountLengthAndBytes(first, second string) string {
	sum := first + second
	return fmt.Sprintf("Объединённая строка: %s. Количество байт: %d. Количество символов: %d.", sum, len(sum), len([]rune(sum)))
}
