package day2

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part2() {
	file, err := os.ReadFile("./day2/input")
	if err != nil {
		log.Fatal(err)
	}

	reports := strings.Split(strings.TrimSpace(string(file)), "\n")
	res := 0

	for _, report := range reports {
		report = strings.TrimSpace(report)
		rawData := strings.Split(report, " ")
		data := make([]int, 0)
		// fmt.Println(rawData)
		for _, point := range rawData {
			num, err := strconv.Atoi(point)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, num)
		}

		valid := checkReportWithDp(data, 1)

		if valid {
			res++
		}
	}
	fmt.Println(res)
}

func checkReportWithDp(data []int, attempt int) bool {
	for i := range data {
		if i == 0 {
			continue
		}
		diff := data[i] - data[i-1]

		// Either increasing or decreasing with magnitude < 3
		if diff == 0 || diff > 3 || diff < -3 {
			if attempt > 0 {
				rmFirst := checkReportWithDp(removeIndex(data, i), attempt-1)
				rmSecond := checkReportWithDp(removeIndex(data, i-1), attempt-1)
				return rmFirst || rmSecond
			}
			return false
		}

		// Check if it consistently increasing or decreasing
		if i < len(data)-1 {
			diff2 := data[i+1] - data[i]
			if (diff > 0) != (diff2 > 0) {
				if attempt <= 0 {
					return false
				}
				rmFirst := checkReportWithDp(removeIndex(data, i-1), attempt-1)
				rmSecond := checkReportWithDp(removeIndex(data, i), attempt-1)
				rmThird := checkReportWithDp(removeIndex(data, i+1), attempt-1)
				return rmFirst || rmSecond || rmThird
			}
		}
	}

	return true
}

func removeIndex[T any](data []T, index int) []T {
	return slices.Concat(data[:index], data[index+1:])
}
