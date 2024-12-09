package main

import (
	"fmt"
	"strconv"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

func main() {
	content, err := lib.ReadFile(true)
	if err != nil {
		panic(err)
	}

	fmt.Println(content)

	chars := strings.Split(strings.TrimSpace(content), "")
	blocks := make([]int, 0)

	id := -1
	block := true
	for _, char := range chars {
		n, _ := strconv.Atoi(char)

		numbers := make([]int, n)
		num := -1
		if block {
			id++
			num = id
		}
		for i := 0; i < n; i++ {
			numbers[i] = num
		}
		blocks = append(blocks, numbers...)

		block = !block
	}

	fmt.Println(blocks)
	fill(blocks)
}

func fill(blocks []int) {
	l := len(blocks)

	length := 0
	cur := blocks[l-1]
	min := int(^uint(0) >> 1)
	for i := l - 1; i >= 0; i-- {
		if blocks[i] == cur {
			length++
			continue
		}

		if cur == -1 {
			cur = blocks[i]
			length = 1
			continue
		}

		if cur < min {
			s := findSpace(blocks, length, i)
			fmt.Println(cur, length, s)
			if s > -1 {
				for j := s; j < s+length; j++ {
					blocks[j] = cur
				}
				for j := i + 1; j <= i+length; j++ {
					blocks[j] = -1
				}
			}
			fmt.Println(blocks)
			min = cur
		}

		cur = blocks[i]
		length = 1
	}

	fmt.Println(blocks)
	total := 0
	for i, n := range blocks {
		if n == -1 {
			continue
		}
		total += i * n
	}
	fmt.Println(total)
}

func findSpace(blocks []int, length int, limit int) int {
	space := 0
	start := 0
	cur := blocks[0]
	for i := 0; i < limit; i++ {
		if blocks[i] != cur {
			cur = blocks[i]
			start = i
			space = 0
		}

		if blocks[i] == -1 {
			space++
			if space == length {
				return start
			}
		}
	}

	return -1
}
