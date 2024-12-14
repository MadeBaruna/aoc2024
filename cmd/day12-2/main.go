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
var width, height int

func main() {
	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(strings.TrimSpace(content), "\n")

	grid = make([][]rune, len(rows)+2)
	marked = make([][]bool, len(rows)+2)

	grid[0] = []rune(strings.Repeat(".", len(rows[0])+2))
	grid[len(rows)+1] = []rune(strings.Repeat(".", len(rows[0])+2))
	marked[0] = make([]bool, len(rows[0])+2)
	marked[len(rows)+1] = make([]bool, len(rows[0])+2)

	for i := 0; i < len(rows); i++ {
		grid[i+1] = []rune("." + rows[i] + ".")
		marked[i+1] = make([]bool, len(rows[i])+2)
	}

	height = len(grid) - 1
	width = len(grid[0]) - 1

	printGrid(grid)

	total := 0
	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
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
		return 0, 0
	}

	if marked[cell.y][cell.x] {
		if grid[cell.y][cell.x] != char {
			return 0, 0
		} else {
			return 0, 0
		}
	}

	if grid[cell.y][cell.x] != char {
		return 0, 0
	}

	marked[cell.y][cell.x] = true

	_total := 1
	_fence := 0

	corner := isCorner(cell, char)
	if corner > 0 {
		// fmt.Println("corner", cell, corner)
		_fence += corner
	}

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

func isCorner(cell Vec, c rune) int {
	count := 0

	if grid[cell.y][cell.x+1] != c && grid[cell.y-1][cell.x] != c {
		// ..
		// X.
		count++
	}
	if grid[cell.y][cell.x-1] != c && grid[cell.y-1][cell.x] != c {
		// ..
		// .X
		count++
	}
	if grid[cell.y][cell.x+1] != c && grid[cell.y+1][cell.x] != c {
		// X.
		// ..
		count++
	}
	if grid[cell.y][cell.x-1] != c && grid[cell.y+1][cell.x] != c {
		// .X
		// ..
		count++
	}

	// Xx
	// x.
	if grid[cell.y][cell.x+1] == c && grid[cell.y+1][cell.x] == c && grid[cell.y+1][cell.x+1] != c {
		count++
	}
	// xX
	// .x
	if grid[cell.y][cell.x-1] == c && grid[cell.y+1][cell.x] == c && grid[cell.y+1][cell.x-1] != c {
		count++
	}
	// x.
	// Xx
	if grid[cell.y-1][cell.x] == c && grid[cell.y][cell.x+1] == c && grid[cell.y-1][cell.x+1] != c {
		count++
	}
	// .x
	// xX
	if grid[cell.y-1][cell.x] == c && grid[cell.y][cell.x-1] == c && grid[cell.y-1][cell.x-1] != c {
		count++
	}

	return count
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
