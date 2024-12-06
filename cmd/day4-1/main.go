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
			if grid[i][j] == "X" {
				walk(grid, Dir{-1, -1}, "X", i, j)
				walk(grid, Dir{0, -1}, "X", i, j)
				walk(grid, Dir{1, -1}, "X", i, j)
				walk(grid, Dir{1, 0}, "X", i, j)
				walk(grid, Dir{1, 1}, "X", i, j)
				walk(grid, Dir{0, 1}, "X", i, j)
				walk(grid, Dir{-1, 1}, "X", i, j)
				walk(grid, Dir{-1, 0}, "X", i, j)
			}
		}
	}

	fmt.Println(total)
}

func walk(grid [][]string, dir Dir, cur string, i, j int) {
	x := j + dir.x
	y := i + dir.y
	// fmt.Println(dir, i, j, x, y, cur)

	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[i]) {
		return
	}

	switch cur {
	case "X":
		if grid[y][x] == "M" {
			walk(grid, dir, "M", y, x)
		}
		break
	case "M":
		if grid[y][x] == "A" {
			walk(grid, dir, "A", y, x)
		}
		break
	case "A":
		if grid[y][x] == "S" {
			total += 1
		}
		break
	}
}
