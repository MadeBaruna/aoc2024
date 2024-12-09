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
	last := len(blocks) - 1
	for i := 0; i < len(blocks); i++ {
		if blocks[i] != -1 {
			continue
		}

		for j := last; j >= 0; j-- {
			if blocks[j] == -1 {
				continue
			}
			if i >= last {
				break
			}

			blocks[i] = blocks[j]
			blocks[j] = -2
			last = j - 1
			break
		}
	}

	fmt.Println(blocks[:last+1])
	total := 0
	for i, n := range blocks[:last+1] {
		total += i * n
	}
	fmt.Println(total)
}
