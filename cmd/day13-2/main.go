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
			total += calc(a, b, _x+10000000000000, _y+10000000000000)
		}
	}
	fmt.Println(total)
}

// 94a + 22b = 8400
// 34a + 67b = 5400
func calc(a, b Button, x, y int) int {
	fmt.Println(a, b, x, y)

	a1, b1, t1 := a.x, b.x, x
	a2, b2, t2 := a.y, b.y, y
	// fmt.Println(a1, b1, t1)
	// fmt.Println(a2, b2, t2)

	a1 *= a.y
	b1 *= a.y
	t1 *= a.y

	a2 *= a.x
	b2 *= a.x
	t2 *= a.x

	// fmt.Println(a1, b1, t1)
	// fmt.Println(a2, b2, t2)

	_b := (t1 - t2) / (b1 - b2)
	_a := (t1 - b1*_b) / a1

	if a.x*_a+b.x*_b != x || a.y*_a+b.y*_b != y {
		fmt.Println("not found")
		return 0
	}

	fmt.Println(_a, _b)
	return _a*3 + _b
}
