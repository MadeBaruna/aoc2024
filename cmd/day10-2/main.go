package main

import (
	"fmt"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

func main() {
	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")

	grid := make([][]int, len(lines))

	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, c := range line {
			n := int(c - '0')
			grid[i][j] = n
		}
	}

	printGrid(grid)

	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 0 {
				total += walk(grid, x, y-1, 0)
				total += walk(grid, x, y+1, 0)
				total += walk(grid, x-1, y, 0)
				total += walk(grid, x+1, y, 0)
			}
		}
	}
	fmt.Println(total)
}

func walk(grid [][]int, x, y, cur int) int {
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return 0
	}

	if grid[y][x] != cur+1 {
		return 0
	}
	if grid[y][x] == 9 {
		return 1
	}

	cur = grid[y][x]
	res := walk(grid, x, y-1, cur)
	res += walk(grid, x, y+1, cur)
	res += walk(grid, x-1, y, cur)
	res += walk(grid, x+1, y, cur)
	return res
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}
