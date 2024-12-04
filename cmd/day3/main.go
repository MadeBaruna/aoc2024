package main

import (
	"fmt"
	"strconv"

	"baruna.me/aoc2024/internal/lib"
)

var enabled bool
var start bool
var sum int

func main() {
	enabled = true
	start = false
	sum = 0

	content, err := lib.ReadFile(false)
	if err != nil {
		panic(err)
	}

	acc(content)
	fmt.Println(sum)
}

func isNumber(n rune) bool {
	return n >= '0' && n <= '9'
}

func acc(str string) {
	res := 0
	mul := ""

	tmp := ""
	for i := 0; i < len(str); i++ {
		tmp += string(str[i])
		if enabled && len(tmp) >= 4 && tmp[len(tmp)-4:] == "mul(" {
			tmp = ""
			start = true
			continue
		}
		if len(tmp) >= 4 && tmp[len(tmp)-4:] == "do()" {
			tmp = ""
			enabled = true
		}
		if len(tmp) >= 7 && tmp[len(tmp)-7:] == "don't()" {
			tmp = ""
			enabled = false
		}

		fmt.Println(string(str[i]), start, enabled)

		if start {
			char := rune(str[i])
			if char == ')' {
				_res, err := strconv.Atoi(mul)
				if err != nil {
					panic(err)
				}
				res *= _res
				sum += res
				start = false
				mul = ""
				res = 0
				tmp = ""
			} else if isNumber(char) {
				mul += string(str[i])
			} else if char == ',' {
				_res, err := strconv.Atoi(mul)
				if err != nil {
					panic(err)
				}
				res = _res
				mul = ""
			} else {
				start = false
				mul = ""
				res = 0
				tmp = ""
			}
		}
	}
}
