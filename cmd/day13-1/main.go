package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

type Button struct {
	x, y int
}

func main() {
	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(strings.TrimSpace(content), "\n")

	re := regexp.MustCompile(`(\d+)`)

	a := Button{0, 0}
	b := Button{0, 0}
	counter := 0
	total := 0
	for _, n := range rows {
		if n == "" {
			continue
		}
		counter++

		matches := re.FindAllString(n, -1)
		_x, _ := strconv.Atoi(matches[0])
		_y, _ := strconv.Atoi(matches[1])
		if counter == 1 {
			a = Button{_x, _y}
		} else if counter == 2 {
			b = Button{_x, _y}
		} else if counter == 3 {
			counter = 0
			total += calc(a, b, _x, _y)
		}
	}
	fmt.Println(total)
}

func calc(a, b Button, x, y int) int {
	fmt.Println(a, b, x, y)

	min := int(^uint(0) >> 1)
	found := false
	for _a := 0; _a < 100; _a++ {
		for _b := 0; _b < 100; _b++ {
			_x := a.x*_a + b.x*_b
			_y := a.y*_a + b.y*_b
			if x != _x || y != _y {
				continue
			}

			cur := _a*3 + _b
			if cur < min {
				found = true
				min = cur
			}
		}
	}

	if found {
		return min
	}
	return 0
}
