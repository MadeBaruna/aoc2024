package main

import (
	"fmt"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

type Pos struct {
	x, y int
}

func main() {
	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")
	grid := make([][]rune, len(lines))
	start := Pos{0, 0}

	for i, line := range lines {
		col := []rune(line)
		grid[i] = col

		for j, r := range col {
			if r == '^' {
				start = Pos{j, i}
			}
		}
	}

	fmt.Println(start)
	try(start, grid)
}

func try(start Pos, grid [][]rune) {

	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' || grid[i][j] == '^' {
				continue
			}
			grid[i][j] = '#'
			total += walk(start, grid)
			grid[i][j] = '.'
		}
	}

	fmt.Println(total)
}

func walk(start Pos, grid [][]rune) int {
	visit := make([][]Pos, len(grid))
	for i := 0; i < len(grid); i++ {
		visit[i] = make([]Pos, len(grid[i]))
	}

	dir := Pos{0, -1}
	cur := start

	for {
		cur.x += dir.x
		cur.y += dir.y
		if cur.x < 0 || cur.x >= len(grid[0]) || cur.y < 0 || cur.y >= len(grid) {
			break
		}
		if grid[cur.y][cur.x] == '#' {
			cur.x -= dir.x
			cur.y -= dir.y

			if visit[cur.y][cur.x].x == dir.x && visit[cur.y][cur.x].y == dir.y {
				return 1
			}

			visit[cur.y][cur.x].x = dir.x
			visit[cur.y][cur.x].y = dir.y

			turn(&dir)
			cur.x += dir.x
			cur.y += dir.y

			// corner case need to turn again
			// .#..
			// .^#.
			// ....
			if grid[cur.y][cur.x] == '#' {
				cur.x -= dir.x
				cur.y -= dir.y
				turn(&dir)
				cur.x += dir.x
				cur.y += dir.y
			}
		}
	}

	return 0
}

func turn(dir *Pos) {
	if dir.x == 0 && dir.y == -1 {
		dir.x = 1
		dir.y = 0
	} else if dir.x == 1 && dir.y == 0 {
		dir.x = 0
		dir.y = 1
	} else if dir.x == 0 && dir.y == 1 {
		dir.x = -1
		dir.y = 0
	} else {
		dir.x = 0
		dir.y = -1
	}
}

// func printGrid(grid [][]rune) {
// 	for i := 0; i < len(grid); i++ {
// 		fmt.Println(string(grid[i]))
// 	}
// }

// func printVisit(visit [][]Pos) {
// 	for i := 0; i < len(visit); i++ {
// 		for j := 0; j < len(visit[i]); j++ {
// 			fmt.Print(visit[i][j])
// 		}
// 		fmt.Println()
// 	}
// }
