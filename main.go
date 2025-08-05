package main

import (
	files "alem-project-go-1/files"
	"fmt"
)

func main() {
	var h, w int
	cellH, cellW := 3, 7

	// Выбор режима игры
	var mode int
	files.PrintMessage("Выберите режим игры:\n1. Ввести карту вручную\n2. Сгенерировать случайную карту\nВведите ваш выбор: ")
	fmt.Scanf("%d", &mode)

	if mode > 2 || mode < 1 {
		files.ShowError("Ошибка: неправильный ввод.")
		return
	}

	// Ввод размеров карты
	files.PrintMessage("Введите размер карты (h w): ")
	fmt.Scanf("%d %d", &h, &w)

	// Проверка валидности размера карты
	if h < 3 || w < 3 {
		files.ShowError("Ошибка: неправильный ввод.")
		return
	}

	// Генерация или ввод карты в зависимости от выбранного режима
	var grid [][]rune
	if mode == 1 {
		grid = files.GetGrid(h, w)
	} else if mode == 2 {
		grid = files.GenerateGrid(h, w)
	} else {
		files.ShowError("Ошибка: неправильный ввод.")
		return
	}

	// Проверка наличия бомб на карте
	if files.CountBombs(grid) < 2 {
		files.ShowError("Ошибка: неправильный ввод.")
		return
	}

	// Печать начальной карты
	files.PrintGrid(grid, h, w, cellH, cellW, nil)

	moves := 0
	var x, y int
	revealed := make([][]bool, h)
	for i := range revealed {
		revealed[i] = make([]bool, w)
	}

	// Основной игровой цикл
	for {
		files.PrintMessage("Введите координаты: ")
		fmt.Scanf("%d %d", &x, &y)

		// Проверка корректности ввода
		if x < 1 || x > h || y < 1 || y > w {
			files.PrintMessage("Неправильный ввод.\n")
			continue
		}
		x, y = x-1, y-1
		moves++

		// Перемещение бомбы в случае первого хода, если на выбранной клетке бомба
		if moves == 1 && grid[x][y] == '*' {
			files.MoveBomb(grid, x, y)
		}

		// Завершение игры, если игрок выбрал клетку с бомбой
		if grid[x][y] == '*' {
			files.RevealAllBombs(grid, revealed)
			files.PrintGrid(grid, h, w, cellH, cellW, revealed)
			files.PrintMessage("Вы проиграли!\n")
			files.PrintStats(h, w, files.CountBombs(grid), moves)
			return
		}

		// Открытие клетки
		files.RevealCells(grid, revealed, x, y)

		// Проверка на победу
		if files.CheckWin(grid, revealed) {
			files.PrintGrid(grid, h, w, cellH, cellW, revealed)
			files.PrintMessage("Вы выиграли!\n")
			files.PrintStats(h, w, files.CountBombs(grid), moves)
			return
		}

		files.PrintGrid(grid, h, w, cellH, cellW, revealed)
	}
}
