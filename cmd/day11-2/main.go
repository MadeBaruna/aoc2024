package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

type Key struct {
	n, remaining int
}

var memo = make(map[Key]int)

func main() {
	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	_nums := strings.Split(strings.TrimSpace(content), " ")
	nums := make([]int, len(_nums))

	for i, n := range _nums {
		r, _ := strconv.Atoi(n)
		nums[i] = r
	}

	fmt.Println(nums)

	run(nums)
}

func run(nums []int) {
	loop := 75

	total := 0
	for _, n := range nums {
		total += split(n, loop)
	}
	fmt.Println(total)
}

func split(n, remaining int) int {
	// fmt.Println(n, remaining)
	if remaining == 0 {
		return 1
	}

	v, ok := memo[Key{n, remaining}]
	if ok {
		// fmt.Println("found", n, v)
		return v
	}

	if n == 0 {
		return split(1, remaining-1)
	}

	digits := countDigit(n)
	if digits%2 == 0 {
		first, second := splitDigits(n, digits)
		result := split(first, remaining-1) + split(second, remaining-1)
		memo[Key{n, remaining}] = result
		return result
	}

	return split(n*2024, remaining-1)
}

func countDigit(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}

	return count
}

func splitDigits(num, digits int) (int, int) {
	divisor := int(math.Pow10(digits / 2))

	part1 := num / divisor
	part2 := num % divisor

	return part1, part2
}
