package main

import (
	"fmt"
	"strconv"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

type Item struct {
	target int
	nums   []int
}

func main() {
	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")

	items := make([]Item, len(lines))

	for i, line := range lines {
		l := strings.Split(line, ":")
		_target := strings.TrimSpace(l[0])
		target, err := strconv.Atoi(_target)
		if err != nil {
			panic(err)
		}

		_nums := strings.Split(strings.TrimSpace(l[1]), " ")
		nums := make([]int, len(_nums))
		for i, num := range _nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			nums[i] = n
		}

		items[i] = Item{target: target, nums: nums}
	}

	total := 0
	for _, item := range items {
		total += check(item)
	}
	fmt.Println(total)
}

func check(item Item) int {
	length := len(item.nums) - 1
	for i := 0; i < (1 << length); i++ {
		res := item.nums[0]
		for j := 0; j < length; j++ {
			if i&(1<<j) != 0 {
				res *= item.nums[j+1]
			} else {
				res += item.nums[j+1]
			}
		}
		if res == item.target {
			return item.target
		}
	}
	return 0
}

// ++++
// *+++
// +*++
// ++*+
// +++*
// **++
// *+*+
// *++*
// +**+
// +*+*
// ++**
// ***+
// **+*
// *+**
// +***
// ****
