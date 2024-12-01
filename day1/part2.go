package day1

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part2() {
	file, err := os.ReadFile("./day1/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)
	lines := strings.Split(input, "\n")
	firstList := make([]int, 0)
	secondList := make(map[int]int, 0)
	for _, line := range lines {
		in := strings.Split(strings.TrimSpace(line), "   ")
		// fmt.Println(in)
		first, err := strconv.Atoi(in[0])
		if err != nil {
			log.Fatal(err)
		}
		firstList = append(firstList, first)
		second, err := strconv.Atoi(in[1])
		if err != nil {
			log.Fatal(err)
		}
		secondList[second]++
	}

	sort.Slice(firstList, func(i, j int) bool {
		return firstList[i] < firstList[j]
	})

	// fmt.Println("======== First list ========\n", firstList, "\n")
	// fmt.Println("======== Second list ========\n", secondList, "\n")
	res := 0
	for _, num := range firstList {
		diff := num * secondList[num]
		res += diff
	}
	fmt.Println(res)
}
