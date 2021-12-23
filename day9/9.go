package main

import (
	"2021advent_of_code/aoc"
	"strings"
)

// PROBLEM-SPECIFIC UTIL FUNCTIONS

func CalculateRiskLevel(lowPoints []int) int {
	// aoc.Log(lowPoints)
	result := 0
	for _, val := range lowPoints {
		result += val + 1
	}
	return result
}

func FindLowPoints(inputMatrix [][]int,length int, height int) []int {
	lowPoints := make([]int, 0)
	aoc.Log()
	result := 0
  for i := 0; i < height; i ++ {
    for j:=0; j < length; j++ {
      currDigit := inputMatrix[i][j]
			belowDigit := 10
			aboveDigit := 10
			leftDigit := 10
			rightDigit := 10
      if i > 0 {
        // check above
        aboveDigit = inputMatrix[i-1][j]
        if currDigit >= aboveDigit {
          continue
        }
      }
      if i < (height - 1) {
        // check below digit
        belowDigit = inputMatrix[i+1][j]
        if currDigit >= belowDigit {
          continue
        }
      }
      if j > 0 {
        // check left digit
        leftDigit = inputMatrix[i][j-1]
        if currDigit >= leftDigit {
          continue
        }
      }
      if j < (length - 1) {
        // check right digit
        rightDigit = inputMatrix[i][j+1]
        if currDigit >= rightDigit {
          continue
        }
      }
			aoc.Log(currDigit,aboveDigit,rightDigit, belowDigit,leftDigit)
      lowPoints = append(lowPoints,currDigit)
			result += currDigit + 1
    }
    }
		aoc.Log(result)
    return lowPoints
  }

// SOLVER
func Solve(inputFile string) {
	aoc.Log("Solving")
	// Read input and initialize values
	result := 0
	input := aoc.ReadInput(inputFile, "\n")
  length := len(input[0])
  height := len(input)
	inputMatrix := make([][]int,height)
	for i := 0; i < height; i ++ {
		numbersInLine := strings.Split(input[i], "")
		inputMatrix[i] = make([]int, length)
		for j:=0; j < length; j++ {
      currDigit := aoc.Atoi(numbersInLine[j])
			inputMatrix[i][j] = currDigit
		}
	}
  lowpoints := FindLowPoints(inputMatrix,length, height)
  result = CalculateRiskLevel(lowpoints)
	aoc.Log(result)
}

func main() {
	Solve("input/test9.txt")
	Solve("input/9.txt")
}
