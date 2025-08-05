package files

import (
	"github.com/alem-platform/ap"
)

// Функция для раскрашивания числа в зависимости от его значения
func ColorizeNumber(bombCount int) {
	var colorCode string
	switch bombCount {
	case 1:
		colorCode = "\033[34m" // Синий
	case 2:
		colorCode = "\033[32m" // Зеленый
	case 3:
		colorCode = "\033[31m" // Красный
	case 4:
		colorCode = "\033[35m" // Пурпурный
	case 5:
		colorCode = "\033[33m" // Желтый
	case 6:
		colorCode = "\033[36m" // Голубой
	default:
		colorCode = "\033[0m" // Без цвета
	}

	// Печатаем код цвета
	for _, ch := range colorCode {
		ap.PutRune(ch)
	}

	// Печатаем число
	ap.PutRune(rune('0' + bombCount))

	// Сброс цвета
	for _, ch := range "\033[0m" {
		ap.PutRune(ch)
	}
}
