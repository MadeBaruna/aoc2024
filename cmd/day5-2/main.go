package main

import (
	"fmt"
	"strconv"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

func main() {
	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	fmt.Println(content)

	lines := strings.Split(strings.TrimSpace(content), "\n")

	rules := make(map[int][]int)
	updates := make([][]int, 0)

	ruleSection := true
	for _, line := range lines {
		if line == "" {
			ruleSection = false
			continue
		}
		if ruleSection {
			rule := strings.Split(line, "|")
			p1, err := strconv.Atoi(rule[0])
			if err != nil {
				panic(err)
			}
			p2, err := strconv.Atoi(rule[1])
			if err != nil {
				panic(err)
			}
			if rules[p1] == nil {
				rules[p1] = make([]int, 0)
			}
			rules[p1] = append(rules[p1], p2)
		} else {
			pagesStr := strings.Split(line, ",")
			pages := make([]int, len(pagesStr))
			for i, pageStr := range pagesStr {
				page, err := strconv.Atoi(pageStr)
				if err != nil {
					panic(err)
				}
				pages[i] = page
			}
			updates = append(updates, pages)
		}
	}

	fmt.Println(rules)
	fmt.Println(updates)

	total := 0
	for _, pages := range updates {
		total += check(rules, pages)
	}

	fmt.Println(total)
}

func check(rules map[int][]int, pages []int) int {
	for i := 0; i < len(pages); i++ {
		a := pages[i]

		for j := i; j < len(pages); j++ {
			b := pages[j]
			broken := contains(rules[b], a)
			if broken {
				return fix(rules, pages)
			}
		}
	}

	return 0
}

func fix(rules map[int][]int, pages []int) int {
	recheck := true

out:
	for recheck {
		for i := 0; i < len(pages); i++ {
			a := pages[i]

			for j := i + 1; j < len(pages); j++ {
				b := pages[j]
				broken := contains(rules[b], a)
				if broken {
					tmp := a
					pages[i] = b
					pages[j] = tmp
					continue out
				}
			}
		}

		recheck = false
	}

	return pages[len(pages)/2]
}

func contains(m []int, i int) bool {
	for _, v := range m {
		if v == i {
			return true
		}
	}
	return false
}
