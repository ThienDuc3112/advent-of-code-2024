package main

import (
	"advent-of-code/day7"
	"fmt"
	"time"
)

func main() {
	defer exeTime("main")()
	day7.Part1()
}

// func to calculate and print execution time
func exeTime(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s execution time: %v\n", name, time.Since(start))
	}
}
