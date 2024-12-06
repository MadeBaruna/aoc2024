package main

import (
	"fmt"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

type Dir struct {
	x, y int
}

var total int = 0

func main() {
	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")
	grid := make([][]string, len(lines))

	for i := 0; i < len(grid); i++ {
		grid[i] = strings.Split(lines[i], "")
	}
	fmt.Println(grid)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "A" {
				check(grid, i, j)
			}
		}
	}

	fmt.Println(total)
}

func check(grid [][]string, i, j int) {
	// X _ X
	// _ A _
	// X _ X

	// check diagonal 1
	if !(walk(grid, Dir{-1, -1}, "M", i, j) && walk(grid, Dir{1, 1}, "S", i, j)) &&
		!(walk(grid, Dir{-1, -1}, "S", i, j) && walk(grid, Dir{1, 1}, "M", i, j)) {
		return
	}
	// check diagonal 2
	if !(walk(grid, Dir{1, -1}, "M", i, j) && walk(grid, Dir{-1, 1}, "S", i, j)) &&
		!(walk(grid, Dir{1, -1}, "S", i, j) && walk(grid, Dir{-1, 1}, "M", i, j)) {
		return
	}

	total++
}

func walk(grid [][]string, dir Dir, char string, i, j int) bool {
	x := j + dir.x
	y := i + dir.y
	// fmt.Println(dir, i, j, x, y, char, grid[y][x])

	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[i]) {
		return false
	}

	return grid[y][x] == char
}
