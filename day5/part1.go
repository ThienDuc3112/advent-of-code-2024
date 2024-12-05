package day5

import (
	"advent-of-code/utils"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func Part1() {
	input := utils.ReadInput("./day5/input")

	validOrder, updates := processInput(input)

	res := 0
	for _, update := range updates {
		if checkValid(validOrder, update) {
			mid, err := strconv.Atoi(update[len(update)/2])
			if err != nil {
				panic(err)
			}
			res += mid
		}
	}
	fmt.Println(res)
}

func checkValid(validOrder map[string]map[string]bool, update []string) bool {
	if len(update) == 1 {
		return true
	}
	valid := true
	for index := range update {
		for i := 0; i < index; i++ {
			if validOrder[update[index]][update[i]] {
				valid = false
				break
			}
		}
	}
	return valid
}

func processInput(input string) (map[string]map[string]bool, [][]string) {
	newLine := "\n"
	if runtime.GOOS == "windows" {
		newLine = "\r\n"
	}
	parts := strings.Split(input, newLine+newLine)
	orders, updatesInput := strings.Split(parts[0], newLine), strings.Split(parts[1], newLine)

	validOrder := make(map[string]map[string]bool)
	for i := range orders {
		temp := strings.Split(strings.TrimSpace(orders[i]), "|")
		if _, ok := validOrder[temp[0]]; !ok {
			validOrder[temp[0]] = make(map[string]bool)
		}
		validOrder[temp[0]][temp[1]] = true
	}

	updates := make([][]string, 0)
	for _, line := range updatesInput {
		update := strings.Split(line, ",")
		for i := range update {
			update[i] = strings.TrimSpace(update[i])
		}
		updates = append(updates, update)
	}

	return validOrder, updates
}
