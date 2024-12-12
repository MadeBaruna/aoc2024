package main

import (
	"fmt"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

type Vector struct {
	x, y int
}

func main() {
	content, err := lib.ReadFile(true)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")

	grid := make([][]rune, len(lines))
	nodes := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
		nodes[i] = make([]rune, len(line))
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != '.' {
				find(Vector{x, y}, grid, nodes)
			}
		}
	}

	printGrid(grid)
	fmt.Println()
	printGrid(nodes)
	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if nodes[y][x] != 0 {
				total++
			}
		}
	}
	fmt.Println(total)
}

func find(pos Vector, grid, nodes [][]rune) {
	c := grid[pos.y][pos.x]

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if x == pos.x && y == pos.y {
				continue
			}

			if grid[y][x] == c {
				dist := distance(pos, Vector{x, y})
				dest := destination(pos, dist)
				if dest.x < 0 || dest.x >= len(grid[y]) || dest.y < 0 || dest.y >= len(grid) {
					continue
				}
				nodes[dest.y][dest.x] = '#'
			}
		}
	}
}

func distance(a, b Vector) Vector {
	return Vector{a.x - b.x, a.y - b.y}
}

func destination(a, b Vector) Vector {
	return Vector{a.x + b.x, a.y + b.y}
}

func printGrid(grid [][]rune) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(grid[y][x]))
			}
		}
		fmt.Println()
	}
}
