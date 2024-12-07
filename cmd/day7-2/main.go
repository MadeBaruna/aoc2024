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
	comb := make([]int, length)

	for {
		res := item.nums[0]
		for i, c := range comb {
			if c == 0 {
				res += item.nums[i+1]
			} else if c == 1 {
				res *= item.nums[i+1]
			} else {
				_res, err := strconv.Atoi(strconv.Itoa(res) + strconv.Itoa(item.nums[i+1]))
				if err != nil {
					panic(err)
				}
				res = _res
			}
		}
		if res == item.target {
			return item.target
		}

		for i := length - 1; i >= 0; i-- {
			comb[i]++
			if comb[i] < 3 {
				break
			}
			comb[i] = 0
			if i == 0 {
				return 0
			}
		}
	}
}
