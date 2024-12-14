package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

type Vec struct {
	x, y int
}

type Robot struct {
	p Vec
	v Vec
}

var robots []Robot

var w, h int = 101, 103

func main() {
	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(strings.TrimSpace(content), "\n")
	robots = make([]Robot, len(rows))

	re := regexp.MustCompile(`(-?\d+)`)
	for i, n := range rows {
		matches := re.FindAllString(n, -1)
		_x, _ := strconv.Atoi(matches[0])
		_y, _ := strconv.Atoi(matches[1])
		_vx, _ := strconv.Atoi(matches[2])
		_vy, _ := strconv.Atoi(matches[3])

		robots[i] = Robot{Vec{_x, _y}, Vec{_vx, _vy}}
	}

	printGrid()
	for i := 0; i < 100; i++ {
		for j := 0; j < len(robots); j++ {
			move(&robots[j])
			// printGrid()
		}
	}
	printGrid()
}

func move(robot *Robot) {
	robot.p.x += robot.v.x
	robot.p.y += robot.v.y

	if robot.p.x < 0 {
		tmp := robot.p.x
		robot.p.x = w
		robot.p.x += tmp
	}
	if robot.p.y < 0 {
		tmp := robot.p.y
		robot.p.y = h
		robot.p.y += tmp
	}
	if robot.p.x >= w {
		robot.p.x %= w
	}
	if robot.p.y >= h {
		robot.p.y %= h
	}
}

func printGrid() {
	grid := make([][]int, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]int, w)
	}

	for _, robot := range robots {
		// fmt.Println(robot)
		grid[robot.p.y][robot.p.x]++
	}

	quad := make([]int, 4)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == 0 {
				fmt.Print(".")
			} else {
				if i < h/2 && j < w/2 {
					quad[0] += grid[i][j]
				} else if i < h/2 && j > w/2 {
					quad[1] += grid[i][j]
				} else if i > h/2 && j < w/2 {
					quad[2] += grid[i][j]
				} else if i > h/2 && j > w/2 {
					quad[3] += grid[i][j]
				}

				fmt.Print(grid[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()

	total := 1
	for _, q := range quad {
		total *= q
	}
	fmt.Println(quad, total)
}
