package day2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {
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

		increasing := true
		valid := true
		for i := range data {
			if i == 0 {
				continue
			}
			diff := data[i] - data[i-1]

			// Either increasing or decreasing with magnitude < 3
			if diff == 0 || diff > 3 || diff < -3 {
				valid = false
				break
			}

			// Check if it consistently increasing or decreasing
			if i == 1 {
				increasing = diff > 0
			} else {
				valid = (diff > 0) == increasing
				if !valid {
					break
				}
			}
		}

		if valid {
			res++
		}
	}
	fmt.Println(res)
}
