package day7

import (
	"advent-of-code/utils"
	"fmt"
)

func Part2() {
	input := utils.ReadInput("./day7/input")
	equations := ProcessInput(input)

	res := 0

	for _, eq := range equations {
		if test2(eq) {
			res += eq.res
		}
	}
	fmt.Println(res)
}

func test2(input equation) bool {
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
		concatNum := utils.Atoi(fmt.Sprintf("%d%d", cur, curNum))
		return dp(newEq, concatNum) || dp(newEq, cur*curNum) || dp(newEq, cur+curNum)
	}

	return dp(equation{res: input.res, in: input.in[1:]}, input.in[0])
}
