package day5

import (
	"advent-of-code/utils"
	"fmt"
)

func Part2() {
	validOrder, updates := processInput(utils.ReadInput("./day5/input"))

	res := 0
	for _, update := range updates {
		if checkValid(validOrder, update) {
			continue
		}

		res += processIncorrectUpdate(validOrder, update)
	}
	fmt.Println(res)

}

func processIncorrectUpdate(validOrder map[string]map[string]bool, update []string) int {
	if len(update) == 1 {
		return utils.Atoi(update[0])
	}

	i := 0
	for i < len(update) {
		for index := 0; index < i; index++ {
			if validOrder[update[i]][update[index]] {
				update[i], update[index] = update[index], update[i]
				i = index
				break
			}
		}
		i++
	}
	return utils.Atoi(update[len(update)/2])
}
