package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

type Block struct {
	id     int
	length int
	moved  bool
}

func main() {
	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	// fmt.Println(content)

	chars := strings.Split(strings.TrimSpace(content), "")
	blocks := make([]*Block, 0)

	id := -1
	block := true
	for _, char := range chars {
		n, _ := strconv.Atoi(char)

		num := -1
		if block {
			id++
			num = id
		}
		blocks = append(blocks, &Block{num, n, false})

		block = !block
	}

	printBlocks(blocks)
	fmt.Println()
	move(blocks)
}

func move(blocks []*Block) {
	_blocks := make([]*Block, len(blocks))
	copy(_blocks, blocks)

	for i := len(_blocks) - 1; i >= 0; i-- {
		_b := _blocks[i]

		if _b.id == -1 || _b.moved {
			continue
		}

		s := findSpace(blocks, _b)
		// fmt.Println(i, "-> space index", s, "for", _b.id)
		if s == -1 {
			continue
		}

		blocks[s].length -= _b.length
		bCopy := *_b

		_b.moved = true
		blocks = slices.Insert(blocks, s, &bCopy)
		_b.id = -1
		// printBlocks(blocks)
	}

	printBlocks(blocks)

	total := 0
	i := 0
	for _, b := range blocks {
		for j := 0; j < b.length; j++ {
			if b.id > -1 {
				total += i * b.id
			}
			i++
		}
	}
	fmt.Println()
	fmt.Println(total)
}

func findSpace(blocks []*Block, b *Block) int {
	for i := range blocks {
		block := blocks[i]
		if block.id == b.id {
			return -1
		}
		if block.id != -1 {
			continue
		}
		if block.length >= b.length {
			return i
		}
	}

	return -1
}

func printBlocks(blocks []*Block) {
	for _, b := range blocks {
		if b.id == -1 {
			fmt.Printf("[. %d]", b.length)
		} else {
			fmt.Printf("[%d %d]", b.id, b.length)
		}
	}
	fmt.Println()
}
