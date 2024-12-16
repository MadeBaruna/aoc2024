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
var moves []rune
var cur Vec

func main() {
	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(strings.TrimSpace(content), "\n")
	grid = make([][]rune, 0)
	moves = make([]rune, 0)

	gridPart := true
	for i := 0; i < len(rows); i++ {
		if rows[i] == "" {
			gridPart = false
		}

		if gridPart {
			cols := []rune(rows[i])
			for j, c := range cols {
				if c == '@' {
					cur = Vec{j, i}
				}
			}
			grid = append(grid, cols)
			continue
		}

		moves = append(moves, []rune(rows[i])...)
	}

	printGrid()
	fmt.Println(string(moves))
	fmt.Println(cur)

	for _, m := range moves {
		dir := getDirection(m)
		move(cur, dir)
		// printGrid()
	}
	calculateBoxes()
}

func move(pos, dir Vec) bool {
	fmt.Println(pos, dir)
	next := Vec{pos.x + dir.x, pos.y + dir.y}
	if grid[next.y][next.x] == '#' {
		return false
	}
	if grid[next.y][next.x] == '.' {
		shift(pos, next)
		return true
	}
	if move(next, dir) {
		shift(pos, next)
		return true
	}
	return false
}

func shift(a, b Vec) {
	if grid[a.y][a.x] == '@' {
		cur.x = b.x
		cur.y = b.y
	}
	grid[b.y][b.x] = grid[a.y][a.x]
	grid[a.y][a.x] = '.'
}

func getDirection(r rune) Vec {
	switch r {
	case '^':
		return Vec{0, -1}
	case 'v':
		return Vec{0, 1}
	case '<':
		return Vec{-1, 0}
	case '>':
		return Vec{1, 0}
	}
	return Vec{0, 0}
}

func calculateBoxes() {
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'O' {
				total += 100*i + j
			}
		}
	}
	fmt.Println(total)
}

func printGrid() {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
}
