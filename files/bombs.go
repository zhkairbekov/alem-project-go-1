package files

import (
	"math/rand"
)

// Функция для подсчета количества бомб на карте
func CountBombs(grid [][]rune) int {
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == '*' {
				count++
			}
		}
	}
	return count
}

// Функция для подсчета бомб вокруг клетки
func CountAdjacentBombs(grid [][]rune, x, y int) int {
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
	count := 0
	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && ny >= 0 && nx < len(grid) && ny < len(grid[0]) && grid[nx][ny] == '*' {
			count++
		}
	}
	return count
}

// Функция для перемещения бомбы
func MoveBomb(grid [][]rune, x, y int) {
	for {
		nx, ny := rand.Intn(len(grid)), rand.Intn(len(grid[0]))
		if grid[nx][ny] != '*' {
			grid[nx][ny] = '*'
			grid[x][y] = '.'
			break
		}
	}
}

// Функция для открытия всех бомб
func RevealAllBombs(grid [][]rune, revealed [][]bool) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '*' {
				revealed[i][j] = true
			}
		}
	}
}
