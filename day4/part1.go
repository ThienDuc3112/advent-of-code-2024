package day4

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

var dirs = [][]int{
	{1, 0},
	{1, 1},
	{1, -1},
	{0, 1},
	{0, -1},
	{-1, 0},
	{-1, 1},
	{-1, -1},
}

func Part1() {
	input := utils.ReadInput("./day4/input")
	table := strings.Split(input, "\n")
	for i := range table {
		table[i] = strings.TrimSpace(table[i])
	}

	res := 0
	matchString := "XMAS"
	for row := range table {
		for column := range table[row] {
			for _, dir := range dirs {
				if move(table, matchString, row, column, dir[0], dir[1]) {
					res++
				}
			}
		}
	}
	fmt.Println(res)
}

func move(table []string, match string, row, column, dx, dy int) bool {
	if match == "" {
		return true
	}
	if !checkBoundary(table, row, column) || match[0] != table[row][column] {
		return false
	}

	return move(table, match[1:], row+dx, column+dy, dx, dy)
}

func checkBoundary(table []string, row, column int) bool {
	return row >= 0 && row < len(table) && column >= 0 && column < len(table[row])
}
