package files

import (
	"github.com/alem-platform/ap"
)

func WriteNumber(n, width int) {
	var digits [10]rune
	count := 0
	if n == 0 {
		digits[count] = '0'
		count++
	} else {
		for n > 0 {
			digits[count] = rune('0' + (n % 10))
			n /= 10
			count++
		}
	}

	for i := count; i < width; i++ {
		ap.PutRune(' ')
	}

	for i := count - 1; i >= 0; i-- {
		ap.PutRune(digits[i])
	}
}

func PrintNumber(num int) {
	// Если число равно 0, сразу печатаем '0'
	if num == 0 {
		ap.PutRune('0')
		return
	}

	// Для чисел больше 0, сохраняем цифры в срез
	digits := []rune{}
	for num > 0 {
		digits = append([]rune{rune('0' + num%10)}, digits...)
		num /= 10
	}

	// Печатаем цифры с помощью ap.PutRune
	for _, digit := range digits {
		ap.PutRune(digit)
	}
}

// Функция для вывода сообщения
func PrintMessage(msg string) {
	for _, char := range msg {
		ap.PutRune(char)
	}
}

// Функция для вывода ошибки
func ShowError(str string) {
	for _, char := range str {
		ap.PutRune(char)
	}
	ap.PutRune('\n')
}
