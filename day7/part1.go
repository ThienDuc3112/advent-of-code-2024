package day7

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

type equation struct {
	res int
	in  []int
}

func Part1() {
	input := utils.ReadInput("./day7/input")
	equations := ProcessInput(input)

	res := 0

	for _, eq := range equations {
		if test(eq) {
			res += eq.res
		}
	}
	fmt.Println(res)
}

func test(input equation) bool {
	if len(input.in) == 0 {
		return false
	}
	var dp func(equation, int) bool
	dp = func(e equation, cur int) bool {
		if len(e.in) == 0 {
			return cur == e.res
		}
		if cur > e.res {
			return false
		}
		curNum := e.in[0]
		newEq := equation{res: e.res, in: e.in[1:]}
		return dp(newEq, cur*curNum) || dp(newEq, cur+curNum)
	}

	return dp(equation{res: input.res, in: input.in[1:]}, input.in[0])
}

func ProcessInput(input string) []equation {
	return utils.Map(strings.Split(input, "\r\n"), func(element string, index int) equation {
		equ := strings.Split(strings.TrimSpace(element), ":")
		res := utils.Atoi(strings.TrimSpace(equ[0]))
		in := make([]int, 0)
		for _, num := range strings.Split(strings.TrimSpace(equ[1]), " ") {
			in = append(in, utils.Atoi(strings.TrimSpace(num)))
		}
		return equation{
			res: res,
			in:  in,
		}
	})

}
