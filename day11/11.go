package main

import (
	"2021advent_of_code/aoc"
	"strings"
)


// PROBLEM-SPECIFIC UTIL FUNCTIONS
type octopus struct {
  x, y int
}

func PrettyPrint(inputMatrix [][]int) {
	for _,row := range inputMatrix {
		rowString := ""
		for _,c := range row {
			rowString += aoc.Itoa(c)
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

// Check e is in a
func ContainsCount(a [][]int, e int, alreadySeen []octopus) int {
	length := len(a)
	width := len(a[0])
	result := 0
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if a[i][j] > e && !Contains(alreadySeen,octopus{i,j}){
				result += 1
			}
		}
	}
  return result
}

func flash(i int, j int, inputMatrix [][]int) [][]int {
	length := len(inputMatrix[0])
	height := len(inputMatrix)
	if i > 0 {
		// check above
		inputMatrix[i-1][j]++
	}
	if i < (height - 1) {
		// check below digit
	  inputMatrix[i+1][j]++
	}
	if j > 0 {
		// check left digit
		inputMatrix[i][j-1]++
	}
	if j < (length - 1) {
		// check right digit
		inputMatrix[i][j+1]++
	}
	if i > 0 && j > 0 {
		// check above left digit
		inputMatrix[i-1][j-1]++
	}
	if i > 0 && j < (length - 1) {
		// check above right digit
		inputMatrix[i-1][j+1]++
	}
	if i < (height - 1) && j > 0 {
		// check below left digit
		inputMatrix[i-1][j-1]++
	}
	if i < (height - 1) && j < (length - 1) {
		// check below right digit
		inputMatrix[i-1][j+1]++
	}
	return inputMatrix
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

func countFlashesInStep(inputMatrix [][]int) ([][]int, int) {
	// Increase energy level of all octopus by 1
	length := len(inputMatrix)
	width := len(inputMatrix[0])
	alreadyFlashed := make([]octopus,0)
	// Count the number of new octopus that can flash
	counter := ContainsCount(inputMatrix,9,alreadyFlashed)

	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			inputMatrix[i][j] ++
		}
	}

	// Find an octopus with val > 9, and increase
	// all adjacent octopuse's values by 1.
  for counter != 0 {
		for i := 0; i < length; i++ {
			for j := 0; j < width; j++ {
				if inputMatrix[i][j] > 9 {
						aoc.Log("hi")
						inputMatrix = flash(i,j,inputMatrix)
						PrettyPrint(inputMatrix)
				}
			}
		}
		counter = ContainsCount(inputMatrix,9,alreadyFlashed)
	}
	inputMatrix = resetOctopus(inputMatrix)
	return inputMatrix, len(alreadyFlashed)
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
	inputMatrix := make([][]int,height)
	steps := 2 //100

	// Generate a 2d slice of input
	for i := 0; i < height; i ++ {
		numbersInLine := strings.Split(input[i], "")
		inputMatrix[i] = make([]int, length)
		for j:=0; j < length; j++ {
      currDigit := aoc.Atoi(numbersInLine[j])
			inputMatrix[i][j] = currDigit
		}
	}
	aoc.Log("Step 0:")
	PrettyPrint(inputMatrix)
	// Run n number of steps:
	for i := 0; i < steps; i ++ {
		flashes := 0
		inputMatrix,flashes = countFlashesInStep(inputMatrix)
		aoc.Log("Step",i+1,":")
		PrettyPrint(inputMatrix)
		result += flashes
	}
	aoc.Log(result)
}

func main() {
	Solve("input/t.txt")
	// Solve("input/test11.txt")
	// Solve("input/11.txt")
}
