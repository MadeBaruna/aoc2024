package main

import (
	"fmt"
	"image"
	"image/color"
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

	for i := 0; i < 100000; i++ {
		for j := 0; j < len(robots); j++ {
			move(&robots[j])
		}
		printGrid(i + 1)
	}
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

func printGrid(index int) {
	grid := make([][]int, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]int, w)
	}

	img := image.NewRGBA(image.Rect(0, 0, w, h))
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	for _, robot := range robots {
		// fmt.Println(robot)
		grid[robot.p.y][robot.p.x]++
	}

	l := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			// browsing the images on file browser, saw the images at 11k+ index
			// but wrong answer, there is box around the tree, so here
			// trying to fine straight line
			if grid[i][j] == 0 {
				l = 0
			} else {
				l++
			}
			if l > 10 {
				fmt.Println("Found at", index)
			}

			img.Set(j, i, black)
			if grid[i][j] != 0 {
				img.Set(j, i, white)
			}
		}
	}

	// filename := fmt.Sprintf("images/%d.png", index)
	// f, err := os.Create(filename)
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// png.Encode(f, img)
	// fmt.Println("Image saved to", filename)
}
