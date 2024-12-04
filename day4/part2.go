package day4

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func Part2() {
	input := utils.ReadInput("./day4/input")
	table := strings.Split(input, "\n")
	for i := range table {
		table[i] = strings.TrimSpace(table[i])
	}

	res := 0
	for row := range table {
		for column := range table[row] {
			if checkCross(table, row, column) {
				res++
			}
		}
	}
	fmt.Println(res)
}

func checkCross(table []string, row, column int) bool {
	if !(checkBoundary(table, row-1, column-1) && checkBoundary(table, row+1, column+1)) {
		return false
	}
	if table[row][column] != 'A' {
		return false
	}

	tiltLeft := make(map[byte]bool)
	tiltLeft[table[row-1][column-1]] = true
	tiltLeft[table[row+1][column+1]] = true

	tiltRight := make(map[byte]bool)
	tiltRight[table[row-1][column+1]] = true
	tiltRight[table[row+1][column-1]] = true

	return tiltRight['M'] && tiltRight['S'] && tiltLeft['M'] && tiltLeft['S']
}
