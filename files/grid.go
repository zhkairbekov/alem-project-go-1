package files

import (
	"fmt"
	"math/rand"

	"github.com/alem-platform/ap"
)

// Функция для ввода карты вручную
func GetGrid(h, w int) [][]rune {
	grid := make([][]rune, h)
	for i := 0; i < h; i++ {
		var line string
		fmt.Scanf("%s", &line)
		if len(line) != w {
			return nil
		}
		grid[i] = []rune(line)
	}
	return grid
}

// Функция для генерации случайной карты
func GenerateGrid(h, w int) [][]rune {
	grid := make([][]rune, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]rune, w)
		for j := 0; j < w; j++ {
			if rand.Intn(5) == 0 {
				grid[i][j] = '*'
			} else {
				grid[i][j] = '.'
			}
		}
	}
	return grid
}

// Функция для раскрытия клетки
func RevealCells(grid [][]rune, revealed [][]bool, x, y int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || revealed[x][y] {
		return
	}
	revealed[x][y] = true
	if CountAdjacentBombs(grid, x, y) == 0 {
		directions := [][2]int{
			{-1, -1},
			{-1, 0},
			{-1, 1},
			{0, -1},
			{0, 1},
			{1, -1},
			{1, 0},
			{1, 1},
		}
		for _, d := range directions {
			RevealCells(grid, revealed, x+d[0], y+d[1])
		}
	}
}

// Print the grid with colorful numbers
func PrintGrid(grid [][]rune, h, w, cellH, cellW int, revealed [][]bool) {
	// 1. Вывод номеров столбцов (по центру ячеек)
	for c := 1; c <= w; c++ {
		num := c
		WriteNumber(num, 8)
	}
	ap.PutRune('\n')

	// 2. Верхняя граница (крыша)
	PrintMessage("    ") // Отступ перед линией
	for t := 0; t < (cellW*w + w - 1); t++ {
		PrintMessage("_")
	}
	ap.PutRune('\n')

	// 3. Вывод ячеек (строки)
	for i := 0; i < h; i++ {
		PrintMessage("   ") // Отступ перед первой колонкой
		// Верхняя линия ячеек
		for j := 0; j < w; j++ {
			ap.PutRune('|')
			if revealed == nil || !revealed[i][j] {
				PrintMessage("XXXXXXX") // Закрытая ячейка
			} else {
				PrintMessage("       ") // Открытая ячейка
			}
		}
		ap.PutRune('|')
		ap.PutRune('\n')

		// Средняя строка (число, бомба или пустая)
		num := i + 1
		PrintNumber(num)
		if num > 9 {
			PrintMessage(" ")
		}
		if num < 10 {
			PrintMessage("  ")
		}

		for j := 0; j < w; j++ {
			ap.PutRune('|')
			if revealed == nil || !revealed[i][j] {
				PrintMessage("XXXXXXX") // Закрытая ячейка
			} else if grid[i][j] == '*' {
				PrintMessage("   *   ") // Бомба по центру
			} else {
				bombCount := CountAdjacentBombs(grid, i, j)
				if bombCount == 0 {
					PrintMessage("       ") // Пустая ячейка
				} else {
					// Выводим число по центру
					PrintMessage("   ")
					ColorizeNumber(bombCount)
					PrintMessage("   ")
				}
			}
		}
		ap.PutRune('|')
		ap.PutRune('\n')

		// Нижняя граница ячеек с учетом условий
		PrintMessage("   ") // Отступ перед первой колонкой
		for j := 0; j < w; j++ {
			ap.PutRune('|')
			if revealed == nil || !revealed[i][j] {
				PrintMessage("XXXXXXX") // Закрытая ячейка
			} else {
				PrintMessage("_______") // Открытая ячейка
			}
		}
		ap.PutRune('|')
		ap.PutRune('\n')
	}
}
