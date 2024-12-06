package main

import (
	"fmt"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

var grid = make([][]rune, 0)
var cur = []int{0, 0}

func main() {
	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")
	grid = make([][]rune, len(lines))

	for i, line := range lines {
		col := []rune(line)
		grid[i] = col

		for j, r := range col {
			if r == '^' {
				cur[0] = j
				cur[1] = i
			}
		}
	}

	fmt.Println(cur)
	walk()

	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'X' {
				total++
			}
		}
	}

	fmt.Println(total)
}

func walk() {
	dir := []int{0, -1}
	out := false

	grid[cur[1]][cur[0]] = 'X'

	for !out {
		cur[0] += dir[0]
		cur[1] += dir[1]
		if cur[0] < 0 || cur[0] >= len(grid[0]) || cur[1] < 0 || cur[1] >= len(grid) {
			out = true
			break
		}
		if grid[cur[1]][cur[0]] == '#' {
			cur[0] -= dir[0]
			cur[1] -= dir[1]
			dir = turn(dir)
			cur[0] += dir[0]
			cur[1] += dir[1]
		}
		grid[cur[1]][cur[0]] = 'X'
	}
}

func turn(dir []int) []int {
	if dir[0] == 0 && dir[1] == -1 {
		return []int{1, 0}
	} else if dir[0] == 1 && dir[1] == 0 {
		return []int{0, 1}
	} else if dir[0] == 0 && dir[1] == 1 {
		return []int{-1, 0}
	} else {
		return []int{0, -1}
	}
}
