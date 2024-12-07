package day6

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

var dirs = map[rune][2]int{
	'v': {1, 0},
	'>': {0, 1},
	'<': {0, -1},
	'^': {-1, 0},
}

func Part1() {
	input := utils.ReadInput("./day6/input")
	matrix := utils.Map(strings.Split(input, "\n"), func(row string, i int) string {
		return strings.TrimSpace(row)
	})

	startingPos := [2]int{-1, -1}
	for i, row := range matrix {
		for j, char := range row {
			if _, ok := dirs[char]; ok {
				startingPos[0] = i
				startingPos[1] = j
				break
			}
		}
	}
	if startingPos[0] == -1 {
		panic("Cannot find the position")
	}

	pos := startingPos
	for {
		// time.Sleep(200 * time.Millisecond)
		dir := dirs[rune(matrix[pos[0]][pos[1]])]
		curHead := matrix[pos[0]][pos[1]]
		newHead := ' '
		newPos := [2]int{
			pos[0] + dir[0],
			pos[1] + dir[1],
		}
		if !utils.CheckBoundary(matrix, newPos[0], newPos[1]) {
			fmt.Println("Outside")
			matrix[pos[0]] = utils.ReplaceAtIndex(matrix[pos[0]], rune('X'), pos[1])
			break
		}
		if matrix[newPos[0]][newPos[1]] == '#' {
			if curHead == '^' {
				newHead = '>'
			} else if curHead == '>' {
				newHead = 'v'
			} else if curHead == 'v' {
				newHead = '<'
			} else {
				newHead = '^'
			}
			matrix[pos[0]] = utils.ReplaceAtIndex(matrix[pos[0]], rune(newHead), pos[1])
			fmt.Println("Rotate")
			continue
		}
		matrix[pos[0]] = utils.ReplaceAtIndex(matrix[pos[0]], rune('X'), pos[1])
		pos = newPos
		matrix[pos[0]] = utils.ReplaceAtIndex(matrix[pos[0]], rune(curHead), pos[1])
	}
	res := 0
	for i, row := range matrix {
		for j := range row {
			if matrix[i][j] == 'X' {
				res++
			}
		}
	}
	fmt.Println(res)
}
