package main

import (
	"fmt"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

type Vec struct {
	x, y int
}

var grid [][]rune
var marked [][]bool

func main() {
	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(strings.TrimSpace(content), "\n")

	grid = make([][]rune, len(rows))
	marked = make([][]bool, len(rows))

	for i, n := range rows {
		grid[i] = []rune(n)
		marked[i] = make([]bool, len(n))
	}

	printGrid(grid)

	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if !marked[y][x] {
				total += start(Vec{x, y})
			}
		}
	}
	fmt.Println(total)
}

func start(start Vec) int {
	char := grid[start.y][start.x]

	total, fence := group(Vec{start.x, start.y}, char)
	// fmt.Println(total, fence)
	return total * fence
}

func group(cell Vec, char rune) (int, int) {
	// fmt.Println(cell, string(char))

	if cell.x < 0 || cell.x >= len(grid[0]) || cell.y < 0 || cell.y >= len(grid) {
		return 0, 1
	}

	if marked[cell.y][cell.x] {
		if grid[cell.y][cell.x] != char {
			return 0, 1
		} else {
			return 0, 0
		}
	}

	if grid[cell.y][cell.x] != char {
		return 0, 1
	}

	marked[cell.y][cell.x] = true

	_total := 1
	_fence := 0
	t, f := group(Vec{cell.x, cell.y - 1}, char)
	_total += t
	_fence += f
	t, f = group(Vec{cell.x + 1, cell.y}, char)
	_total += t
	_fence += f
	t, f = group(Vec{cell.x, cell.y + 1}, char)
	_total += t
	_fence += f
	t, f = group(Vec{cell.x - 1, cell.y}, char)
	_total += t
	_fence += f
	return _total, _fence
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
