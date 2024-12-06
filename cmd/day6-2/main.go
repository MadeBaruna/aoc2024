package main

import (
	"fmt"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

type Pos struct {
	x, y int
}

var grid = make([][]rune, 0)
var visit = make([][]Pos, 0)
var cur = Pos{0, 0}

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
				cur = Pos{i, j}
			}
		}
	}

	fmt.Println(cur)
	try()
}

func try() {
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' {
				continue
			}
			grid[i][j] = '#'
			total += walk()
			grid[i][j] = '.'
		}
	}
}

func walk() int {
	dir := Pos{0, -1}
	out := false

	// grid[cur.y][cur.x] = 'X'

	for !out {
		cur.x += dir[0]
		cur.y += dir[1]
		if cur[0] < 0 || cur[0] >= len(grid[0]) || cur[1] < 0 || cur[1] >= len(grid) {
			out = true
			break
		}
		if grid[cur[1]][cur[0]] == '#' {
			cur[0] -= dir[0]
			cur[1] -= dir[1]

			if visit[cur[1]][cur[0]][0] == dir[0] && visit[cur[1]][cur[0]][1] == dir[1] {
				return 1
			}

			visit[cur[0]][cur[1]][0] = dir[0]
			visit[cur[0]][cur[1]][1] = dir[1]

			turn(&dir)
			cur[0] += dir[0]
			cur[1] += dir[1]
		}
	}

	return 0
}

func turn(dir *Pos) {
	if dir.x == 0 && dir.y == -1 {
		dir = &Pos{1, 0}
	} else if dir.x == 1 && dir.y == 0 {
		dir = &Pos{0, 1}
	} else if dir.x == 0 && dir.y == 1 {
		dir = &Pos{-1, 0}
	} else {
		dir = &Pos{0, -1}
	}
}
