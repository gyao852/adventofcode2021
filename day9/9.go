package main

import (
	"2021advent_of_code/aoc"
	"strings"
)

// PROBLEM-SPECIFIC UTIL FUNCTIONS
type seaCoordinate struct {
  depthValue, x, y int
}

// Check e is in a
func Contains(a []seaCoordinate, e seaCoordinate) bool {
  for _, v := range a {
      if v == e {
          return true
      }
  }
  return false
}

func GetBasinFromLowPoint(lowPoint seaCoordinate, seaMap [][]int) []seaCoordinate {
	var queue []seaCoordinate
	var seenCoordinates []seaCoordinate
	height := len(seaMap)
	width := len(seaMap[0])
	queue = append(queue,lowPoint)
	for len(queue) > 0 {
		currentPoint := queue[0]
		queue = queue[1:]
		if !Contains(seenCoordinates,currentPoint) {
			seenCoordinates = append(seenCoordinates,currentPoint)
		}
		if currentPoint.x > 0 {
			// check above
			aboveDigit := seaMap[currentPoint.x-1][currentPoint.y]
			basinPoint := seaCoordinate{aboveDigit,currentPoint.x-1,currentPoint.y}
			if aboveDigit != 9 && !Contains(seenCoordinates,basinPoint) {
				queue = append(queue,basinPoint)
			}
		}
		if currentPoint.x < (height - 1) {
			// check below
			belowDigit := seaMap[currentPoint.x+1][currentPoint.y]
			basinPoint := seaCoordinate{belowDigit,currentPoint.x+1,currentPoint.y}
			if belowDigit != 9 && !Contains(seenCoordinates,basinPoint) {
				queue = append(queue,basinPoint)
			}
		}
		if currentPoint.y > 0 {
			// check left
			leftDigit := seaMap[currentPoint.x][currentPoint.y-1]
			basinPoint := seaCoordinate{leftDigit,currentPoint.x,currentPoint.y-1}
			if leftDigit != 9 && !Contains(seenCoordinates,basinPoint) {
				queue = append(queue,basinPoint)
			}
		}
		if currentPoint.y < (width - 1) {
			// check right
			rightDigit := seaMap[currentPoint.x][currentPoint.y+1]
			basinPoint := seaCoordinate{rightDigit,currentPoint.x,currentPoint.y+1}
			if rightDigit != 9 && !Contains(seenCoordinates,basinPoint) {
				queue = append(queue,basinPoint)
			}
		}
	}
	return seenCoordinates
}

func CalculateNLargestBasins(lowPoints []seaCoordinate, seaMap [][]int, n int) int {
	largestBasins := make([][]seaCoordinate,0)
	result := 1
	smallestLargeBasin := 0
	for _, lowPoint := range lowPoints {
		basinMap := GetBasinFromLowPoint(lowPoint, seaMap)
		if len(largestBasins) < n {
			largestBasins = append(largestBasins, basinMap)
			for i := 0; i < len(largestBasins) - 1; i++ {
				if len(largestBasins[i]) > len(largestBasins[i+1]) {
					smallestLargeBasin = i
				}
			}
		} else {
			if len(largestBasins[smallestLargeBasin]) < len(basinMap) {
				largestBasins[smallestLargeBasin] = basinMap
			}
			for i := 0; i < len(largestBasins) - 1; i++ {
				if len(largestBasins[i]) > len(largestBasins[i+1]) {
					smallestLargeBasin = i
				}
			}
		}
	}
	for _, basin := range largestBasins {
		aoc.Log("Size of basin:",len(basin),"\nBasin:",basin)
		result *= len(basin)
	}
	return result
}

func CalculateRiskLevel(lowPoints []int) int {
	result := 0
	for _, val := range lowPoints {
		result += val + 1
	}
	return result
}

func FindLowPoints(inputMatrix [][]int,length int, height int) []seaCoordinate {
	lowPoints := make([]seaCoordinate, 0)
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
			newLowPoint := seaCoordinate{currDigit,i,j}
      lowPoints = append(lowPoints,newLowPoint)
    }
    }
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

	// Generate a 2d slice of input
	for i := 0; i < height; i ++ {
		numbersInLine := strings.Split(input[i], "")
		inputMatrix[i] = make([]int, length)
		for j:=0; j < length; j++ {
      currDigit := aoc.Atoi(numbersInLine[j])
			inputMatrix[i][j] = currDigit
		}
	}

	// First find all the lowest points in the sea
  lowpoints := FindLowPoints(inputMatrix,length, height)

	// Part1 just do some math
	// Part2, need to do BPF to look for basin map
  // result = CalculateRiskLevel(lowpoints) // part1
	result = CalculateNLargestBasins(lowpoints, inputMatrix, 3) // part2
	aoc.Log(result)
}

func main() {
	// Solve("input/test9.txt")
	Solve("input/9.txt")
}
