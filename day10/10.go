package main

import (
	"2021advent_of_code/aoc"
	"sort"
	"strings"
)

var basePath = "/Users/gyao/Documents/Personal Projects/go_projects/src/2021advent_of_code/day10/"

// PROBLEM-SPECIFIC UTIL FUNCTIONS
var chunkCloses = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var syntaxErrorScore = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var syntaxCorrectionScore = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

func checkForSyntax(line string, part1 bool) int {
	var stack []string
	var syntaxError = false
	var characters = strings.Split(line, "")
	var points = 0
	for _, c := range characters {
		if _, result := chunkCloses[c]; result {
			stack = append(stack, c)
		} else {
			// Check if the last item we pushed matches current
			n := len(stack) - 1 // Top element
			lastElement := stack[n]
			stack = stack[:n]
			if chunkCloses[lastElement] != c {
				aoc.Log("Syntax Error Found on line:", line, ". Expected", chunkCloses[lastElement], "but found", c)
				if !part1 {
					syntaxError = true
					break
				}
				points += syntaxErrorScore[c]
				break
			}
		}
	}
	if !part1 && !syntaxError {
		for len(stack) > 0 {
			n := len(stack) - 1 // Top element
			lastElement := stack[n]
			stack = stack[:n]
			points *= 5
			points += syntaxCorrectionScore[lastElement]
		}
	}
	return points
}

// SOLVER
func Solve(inputFile string) {
	aoc.Log("Solving")
	// Read input and initialize values
	var part1 = false
	var part2Points []int
	result := 0
	input := aoc.ReadInput(inputFile, "\n")
	for _, line := range input {
		if part1 {
			result += checkForSyntax(line, part1)
		} else {
			points := checkForSyntax(line, part1)
			if points > 0 {
				part2Points = append(part2Points, points)
			}
		}
	}
	if !part1 {
		sort.Slice(part2Points, func(i, j int) bool {
			return part2Points[i] > part2Points[j]
		})
		aoc.Log(part2Points)
		result = part2Points[len(part2Points)/2]
	}

	aoc.Log(result)
}

func main() {
	// Solve("input/test10.txt")
	Solve("input/10.txt")
}
