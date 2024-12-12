package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

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

	length := len(nums)
	half := length / 2
	a := make([]int, len(nums[:half]))
	copy(a, nums[:half])
	b := make([]int, len(nums[half:]))
	copy(b, nums[half:])
	_a := split(a, 25)
	fmt.Println(len(_a))
	_b := split(b, 25)
	fmt.Println(len(_b))

	fmt.Println(len(_a) + len(_b))
	// fmt.Println(len(nums))
}

func split(nums []int, remaining int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		// fmt.Println("->", nums)

		n := nums[i]
		if n == 0 {
			nums[i] = 1
			// fmt.Println("zero", nums)
			continue
		}

		digits := countDigit(n)
		if digits%2 == 0 {
			first, second := splitDigits(n, digits)
			nums[i] = first
			nums = append(nums[:i+1], append([]int{second}, nums[i+1:]...)...)
			length++
			i++
			// fmt.Println("split", first, second, nums)
			continue
		}

		nums[i] *= 2024
		// fmt.Println("multiplied", nums)
	}

	// fmt.Println(nums)

	if remaining > 1 {
		return split(nums, remaining-1)
	}
	return nums
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
