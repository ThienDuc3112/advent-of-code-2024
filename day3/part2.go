package day3

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
)

func Part2() {
	str := utils.ReadInput("./day3/input")
	callsRegex, err := regexp.Compile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		panic(err)
	}
	doRegex, err := regexp.Compile(`do\(\)`)
	if err != nil {
		panic(err)
	}
	dontRegex, err := regexp.Compile(`don't\(\)`)
	if err != nil {
		panic(err)
	}

	calls := callsRegex.FindAllString(str, -1)
	enable := true
	res := 0
	for _, call := range calls {
		if doRegex.MatchString(call) {
			enable = true
		} else if dontRegex.MatchString(call) {
			enable = false
		} else if enable {
			res += ExtractMul(call)
		}
	}

	fmt.Println(res)
}

func ExtractMul(call string) int {
	digitRegex, err := regexp.Compile(`\d+`)
	if err != nil {
		panic(err)
	}

	digitsStr := digitRegex.FindAllString(call, -1)
	first, err := strconv.Atoi(digitsStr[0])
	if err != nil {
		panic(err)
	}
	second, err := strconv.Atoi(digitsStr[1])
	if err != nil {
		panic(err)
	}

	return first * second
}
