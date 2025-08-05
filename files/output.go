package files

import (
	"github.com/alem-platform/ap"
)

// Функция для проверки победы
func CheckWin(grid [][]rune, revealed [][]bool) bool {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '*' && !revealed[i][j] {
				return false
			}
		}
	}
	return true
}

// Функция для печати статистики
func PrintStats(h, w, bombs, moves int) {
	// Печать статистики игрока
	PrintMessage("Ваша статистика:\n")

	// Печать размера карты
	PrintMessage("- Размер карты: ")
	PrintNumber(h)
	ap.PutRune('x')
	PrintNumber(w)
	ap.PutRune('\n')

	// Печать количества бомб
	PrintMessage("- Количество бомб: ")
	PrintNumber(bombs)
	ap.PutRune('\n')

	// Печать количества ходов
	PrintMessage("- Количество ходов: ")
	PrintNumber(moves)
	ap.PutRune('\n')
}

