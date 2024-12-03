package day3

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
)

func Part1() {
	str := utils.ReadInput("./day3/input")
	mulRegex, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		panic(err)
	}
	digitRegex, err := regexp.Compile(`\d+`)
	if err != nil {
		panic(err)
	}
	calls := mulRegex.FindAllString(str, -1)

	res := 0
	for _, call := range calls {
		digitsStr := digitRegex.FindAllString(call, -1)
		first, err := strconv.Atoi(digitsStr[0])
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(digitsStr[1])
		if err != nil {
			panic(err)
		}

		res += first * second
	}

	fmt.Println(res)
}
