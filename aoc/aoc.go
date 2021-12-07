package aoc

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	// "regexp"
	// "math"
)

// UTIL FUNCTIONS
func Check(e error) {
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

// Min returns the smaller of x or y.
func Absolute(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

// Convert string to int
func Atoi(number string) int {
	val, err := strconv.Atoi(strings.Split(number, "\n")[0])
	Check(err)
	return val
}

// Convert int to string
func Itoa(number int) string {
	val := strconv.Itoa(number)
	return val
}

// Map function f to all values in the slice
func MapSliceIntToInt(slice *[]int, f func(int) int) *[]int {
	for i, v := range *slice {
		(*slice)[i] = f(v)
	}
	return slice
}

// Reduce function f to all values in the slice
func ReduceSliceIntToInt(slice []int, f func(int) int) int {
	result := 0
	for _, v := range slice {
		result = f(v)
	}
	return result
}

// Map function f to all values in the map
func MapMapIntToInt(aMap *map[int]int, f func(int) int) *map[int]int {
	for i, v := range *aMap {
		(*aMap)[i] = f(v)
	}
	return aMap
}

// Reduce function f to all values in the slice
func ReduceMapIntToInt(aMap *map[int]int, f func(int) int) int {
	result := 0
	for _, v := range *aMap {
		result = f(v)
	}
	return result
}

// Log inputs
func Log(x ...interface{}) {
	fmt.Println(x...)
}

// Read input split by pass-in delimiter
func ReadInput(input string, delimiter string) []string {
	readInput, err := os.ReadFile(input)
	Check(err)
	inputSplit := strings.Split(string(readInput), delimiter)
	for _, val := range inputSplit {
		if val == "" {
			return inputSplit[:len(inputSplit)-1]
		}
	}
	return inputSplit
}
