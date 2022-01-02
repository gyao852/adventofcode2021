package main

import (
	"2021advent_of_code/aoc"
	"strings"
)

var basePath = "/Users/gyao/Documents/Personal Projects/go_projects/src/2021advent_of_code/day12/"

// PROBLEM-SPECIFIC UTIL FUNCTIONS
type cave struct {
	name          string
	isBigCave     bool
	adjacentCaves []cave
}

// Check e is in s
func Contains(s []cave, e cave) bool {
	for _, v := range s {
		if v.name == e.name {
			return true
		}
	}
	return false
}

func PrettyPrintCaveMap(s map[string]cave) {
	for k, v := range s {
		aoc.Log(k)
		for i, j := range v.adjacentCaves {
			aoc.Log(i+1, ":", j.name)
		}
	}
}

func IsBigCave(caveName string) bool {
	return caveName == strings.ToUpper(string(caveName))
}

func DFS(currPos string, currPath string, adjMap map[string]cave, smallCavesVisited []cave, numOfPaths int) int {
	if currPos == "end" {
		aoc.Log(currPath)
		return 1
	}

	// For each adjacent cave, we try to explore each of them
	for _, adjCave := range adjMap[currPos].adjacentCaves {
		// if we haven't explored an adjacent cave yet, try exploring it
		if IsBigCave(adjCave.name) {
			numOfPaths += DFS(adjCave.name, currPath+" -> "+adjCave.name, adjMap, smallCavesVisited, numOfPaths)
		} else if !Contains(smallCavesVisited, adjCave) {
			numOfPaths += DFS(adjCave.name, currPath+" -> "+adjCave.name, adjMap, append(smallCavesVisited, adjCave), numOfPaths)
		}
	}
	return numOfPaths
}

// SOLVER
func Solve(inputFile string) {
	aoc.Log("Solving")
	// Read input and initialize values
	result := 0
	input := aoc.ReadInput(inputFile, "\n")
	numOfCaves := len(input)
	adjMap := make(map[string]cave, numOfCaves)
	// part1 := false

	for i := 0; i < numOfCaves; i++ {
		line := strings.Split(input[i], "-")
		head := cave{line[0], IsBigCave(line[0]), make([]cave, 0)}
		tail := cave{line[1], IsBigCave(line[1]), make([]cave, 0)}

		// Populate head or tail if not already
		if val, ok := adjMap[head.name]; ok {
			if !Contains(val.adjacentCaves, tail) {
				val.adjacentCaves = append(val.adjacentCaves, tail)
				adjMap[head.name] = val
			}
		} else {
			head.adjacentCaves = append(head.adjacentCaves, tail)
			adjMap[head.name] = head
		}

		if val, ok := adjMap[tail.name]; ok {
			if head.name == "start" || tail.name == "end" {
				continue
			}
			if !Contains(val.adjacentCaves, head) {
				val.adjacentCaves = append(val.adjacentCaves, head)
				adjMap[tail.name] = val
			}
		} else {
			if head.name == "start" || tail.name == "end" {
				continue
			}
			tail.adjacentCaves = append(tail.adjacentCaves, head)
			adjMap[tail.name] = tail
		}
	}
	PrettyPrintCaveMap(adjMap)
	result = DFS("start", "start", adjMap, make([]cave, 0), 0)
	// PrettyPrintCaves(adjMap)
	aoc.Log("result", result)
}

func main() {
	Solve(basePath + "input/t.txt")
	// Solve(basePath + "input/t2.txt")
	// Solve(basePath + "input/t3.txt")
	// Solve(basePath + "input/test12.txt")
	// Solve(basePath + "input/12.txt")
}
