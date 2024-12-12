package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"baruna.me/aoc2024/internal/lib"
)

var total = 0

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
	for i := 0; i < length; i++ {
		split([]int{nums[i]}, 25)
	}
	fmt.Println(total)
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
	total += len(nums)
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
