package main

import (
	"fmt"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

type Vec struct {
	x, y int
}

type Item struct {
	c rune
	p Vec
}

var grid [][]*Item
var moves []rune
var cur Vec

func main() {
	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(strings.TrimSpace(content), "\n")
	grid = make([][]*Item, 0)
	moves = make([]rune, 0)

	gridPart := true
	for i := 0; i < len(rows); i++ {
		if rows[i] == "" {
			gridPart = false
		}

		if gridPart {
			cols := make([]*Item, 0)
			for _, c := range rows[i] {
				x := len(cols)
				y := i

				if c == '@' {
					cur = Vec{len(cols), i}
					cols = append(cols, []*Item{{'@', Vec{x, y}}, {'.', Vec{x + 1, y}}}...)
				} else if c == '#' {
					cols = append(cols, []*Item{{'#', Vec{x, y}}, {'#', Vec{x + 1, y}}}...)
				} else if c == 'O' {
					box := &Item{'X', Vec{x, y}}
					cols = append(cols, []*Item{box, box}...)
				} else {
					cols = append(cols, []*Item{{'.', Vec{x, y}}, {'.', Vec{x + 1, y}}}...)
				}
			}
			grid = append(grid, cols)
			continue
		}

		moves = append(moves, []rune(rows[i])...)
	}

	printGrid('@')
	fmt.Println(string(moves))
	fmt.Println(cur)

	for _, m := range moves {
		// fmt.Print("\033[H\033[2J")
		// printGrid(m)
		dir := getDirection(m)
		fmt.Println(string(m))
		res := grid[cur.y][cur.x].check(dir)
		if res {
			grid[cur.y][cur.x].move(dir)
		}
		// fmt.Println()
		// printGrid(m)
		// fmt.Println()
		// time.Sleep(200 * time.Millisecond)
	}
	printGrid('@')
	calculateBoxes()
}

func (item *Item) check(dir Vec) bool {
	target := Vec{item.p.x + dir.x, item.p.y + dir.y}
	if target.x < 0 || target.x >= len(grid[0]) || target.y < 0 || target.y >= len(grid) {
		return false
	}

	if item.c == '.' {
		return true
	}
	if item.c == '#' {
		return false
	}

	tC := grid[target.y][target.x]
	fmt.Println("? ", item, " -> ", tC)
	if item == tC {
		target.x += dir.x
		target.y += dir.y
		tC = grid[target.y][target.x]
	}

	if item.c != 'X' || dir.x != 0 {
		if tC.c == 'X' {
			return tC.check(dir)
		}

		if tC.c == '.' {
			return true
		}

		if tC.c == '#' {
			return false
		}
	}

	// fmt.Println(string(item.c), dir, string(tC.c))

	if dir.y != 0 {
		tL := grid[item.p.y+dir.y][item.p.x]
		tR := grid[item.p.y+dir.y][item.p.x+1]
		_tL := tL.check(dir)
		_tR := tR.check(dir)
		return _tL && _tR
	}

	return true
}

func (item *Item) move(dir Vec) {
	target := Vec{item.p.x + dir.x, item.p.y + dir.y}
	tC := grid[target.y][target.x]
	fmt.Println("- ", item, " -> ", target)

	if item.c == 'X' {
		if dir.y != 0 {
			tL := grid[item.p.y+dir.y][item.p.x]
			tR := grid[item.p.y+dir.y][item.p.x+1]
			tL.move(dir)
			if tL != tR {
				tR.move(dir)
			}
		} else {
			if item == tC {
				target.x += dir.x
				target.y += dir.y
				tC = grid[target.y][target.x]
			}
			tC.move(dir)
		}

		grid[item.p.y][item.p.x] = &Item{'.', item.p}
		grid[item.p.y][item.p.x+1] = &Item{'.', item.p}
		item.p.x += dir.x
		item.p.y += dir.y
		grid[item.p.y][item.p.x] = item
		grid[item.p.y][item.p.x+1] = item
	} else if item.c == '@' {
		if tC.c == 'X' {
			tC.move(dir)
		}
		grid[item.p.y][item.p.x] = &Item{'.', item.p}
		item.p.x += dir.x
		item.p.y += dir.y
		grid[item.p.y][item.p.x] = item
		cur.x = item.p.x
		cur.y = item.p.y
	}
}

func (item *Item) String() string {
	return fmt.Sprintf(`%c:%d,%d`, item.c, item.p.x, item.p.y)
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
			if grid[i][j].p.x == j && grid[i][j].c == 'X' {
				total += 100*i + j
			}
		}
	}
	fmt.Println(total)
}

func printGrid(m rune) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].c == 'X' {
				if grid[i][j].p.x == j {
					// fmt.Print("X")
					fmt.Print("[")
				} else {
					// fmt.Print(grid[i][j].p.x)
					fmt.Print("]")
				}
			} else if grid[i][j].c == '@' {
				fmt.Print(string(m))
			} else {
				fmt.Print(string(grid[i][j].c))
			}
		}
		fmt.Println()
	}
}
