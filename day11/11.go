package main

import (
	"2021advent_of_code/aoc"
	"strings"
)

var basePath = "/Users/gyao/Documents/Personal Projects/go_projects/src/2021advent_of_code/day11/"

// PROBLEM-SPECIFIC UTIL FUNCTIONS
type octopus struct {
	x, y, energy int
	flashing     bool
}

func PrettyPrint(inputMatrix [][]octopus) {
	for _, row := range inputMatrix {
		rowString := ""
		for _, c := range row {
			rowString += aoc.Itoa(c.energy) + "|"
		}
		aoc.Log(rowString)
	}
}

// Check e is in a
func Contains(a []octopus, e octopus) bool {
	for _, v := range a {
		if v == e {
			return true
		}
	}
	return false
}

// Check e is in a and has not yet been seen
func FindReadyOctopuses(a [][]octopus, e int) int {
	length := len(a)
	width := len(a[0])
	result := 0
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if a[i][j].energy > e && !a[i][j].flashing {
				result++
			}
		}
	}
	return result
}

func flash(i int, j int, inputMatrix [][]octopus, counter int) ([][]octopus, int) {
	if inputMatrix[i][j].flashing == true {
		return inputMatrix, counter
	}
	counter++
	inputMatrix[i][j].flashing = true
	length := len(inputMatrix[0])
	height := len(inputMatrix)
	if i > 0 {
		// check above
		inputMatrix[i-1][j].energy++
	}
	if i < (height - 1) {
		// check below digit
		inputMatrix[i+1][j].energy++
	}
	if j > 0 {
		// check left digit
		inputMatrix[i][j-1].energy++
	}
	if j < (length - 1) {
		// check right digit
		inputMatrix[i][j+1].energy++
	}
	if i > 0 && j > 0 {
		// check above left digit
		inputMatrix[i-1][j-1].energy++
	}
	if i > 0 && j < (length-1) {
		// check above right digit
		inputMatrix[i-1][j+1].energy++
	}
	if i < (height-1) && j > 0 {
		// check below left digit
		inputMatrix[i+1][j-1].energy++
	}
	if i < (height-1) && j < (length-1) {
		// check below right digit
		inputMatrix[i+1][j+1].energy++
	}
	return inputMatrix, counter
}

func resetOctopus(inputMatrix [][]int) [][]int {
	length := len(inputMatrix)
	width := len(inputMatrix[0])
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if inputMatrix[i][j] > 9 {
				inputMatrix[i][j] = 0
			}
		}
	}
	return inputMatrix
}

func countFlashesInStep(inputMatrix [][]octopus) ([][]octopus, int) {
	// Increase energy level of all octopus by 1
	length := len(inputMatrix)
	width := len(inputMatrix[0])
	counter := 0
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			inputMatrix[i][j].energy++
		}
	}
	// Count the number of current octopus that can flash
	readyOctopus := FindReadyOctopuses(inputMatrix, 9)

	// aoc.Log("Octopus that can flash:", readyOctopus)
	// Find an octopus with val > 9, and increase
	// all adjacent octopuse's values by 1.
	for readyOctopus != 0 {
		for i := 0; i < length; i++ {
			for j := 0; j < width; j++ {
				if inputMatrix[i][j].energy > 9 {
					inputMatrix, counter = flash(i, j, inputMatrix, counter)

					readyOctopus = FindReadyOctopuses(inputMatrix, 9)
				}
			}
		}
	}
	// Finally, reset the status of each octopus as at this point
	// all octopus should've flashed already
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if inputMatrix[i][j].energy > 9 {
				inputMatrix[i][j].energy = 0
			}
			inputMatrix[i][j].flashing = false
		}
	}
	return inputMatrix, counter
}

// SOLVER
func Solve(inputFile string) {
	aoc.Log("Solving")
	// Read input and initialize values
	// var part1 = true
	result := 0
	input := aoc.ReadInput(inputFile, "\n")
	length := len(input[0])
	height := len(input)
	inputMatrix := make([][]octopus, height)
	steps := 100 //100

	// Generate a 2d slice of input
	for i := 0; i < height; i++ {
		numbersInLine := strings.Split(input[i], "")
		inputMatrix[i] = make([]octopus, length)
		for j := 0; j < length; j++ {
			currDigit := aoc.Atoi(numbersInLine[j])
			inputMatrix[i][j] = octopus{i, j, currDigit, false}
		}
	}
	aoc.Log("Before any steps:")
	PrettyPrint(inputMatrix)
	// Run n number of steps:
	for i := 0; i < steps; i++ {
		flashes := 0
		aoc.Log("After Step", i+1, ":")
		inputMatrix, flashes = countFlashesInStep(inputMatrix)
		PrettyPrint(inputMatrix)
		result += flashes
	}
	aoc.Log(result)
}

func main() {
	// Solve(basePath + "input/t.txt")
	// Solve(basePath + "input/test11.txt")
	Solve(basePath + "input/11.txt")
}
