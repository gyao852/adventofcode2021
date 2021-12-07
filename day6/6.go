package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	// "regexp"
	// "math"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func solve(inputFile string) {
	fmt.Println("Solving!")
	input, err := os.ReadFile(inputFile)
	check(err)
	result := 0
	initialFishes := strings.Split(string(input), ",")
	allFishes := make(map[int]int)
	for i := 0; i < 9; i++ {
		allFishes[i] = 0
	}
	for _, fish := range initialFishes {
		fishVal, err := strconv.Atoi(strings.Split(fish, "\n")[0])
		check(err)
		allFishes[fishVal] += 1
	}

	dayLimit := 256 // 80
	currentVal := 0
	fmt.Println("Inital Fishes: ", allFishes)
	for day := 1; day <= dayLimit; day++ {
		currentVal = allFishes[8]
		for i := 8; i > 0; i-- {
			prevVal := allFishes[i-1]
			allFishes[i-1] = currentVal
			currentVal = prevVal
		}
		allFishes[6] += currentVal
		allFishes[8] = currentVal
		fmt.Println("Day", day, ":", allFishes)
	}
	for _, fishVal := range allFishes {
		result += fishVal
	}
	fmt.Println("All fishes", allFishes)
	fmt.Println(result)
}

func main() {
	solve("../downloads/test6.txt")
	solve("../downloads/6.txt")
}
