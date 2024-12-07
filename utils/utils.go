package utils

import (
	"os"
	"strconv"
)

func ReadInput(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func Atoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func Map[T, R any](slice []T, tranform func(element T, index int) R) []R {
	res := make([]R, len(slice))
	for i, element := range slice {
		res[i] = tranform(element, i)
	}
	return res
}

func CheckBoundary(table []string, row, column int) bool {
	return row >= 0 && row < len(table) && column >= 0 && column < len(table[row])
}

func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
