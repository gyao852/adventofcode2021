package main

import (
	"2021advent_of_code/aoc"
	"strings"
)

var basePath = "/Users/gyao/Documents/Personal Projects/go_projects/src/2021advent_of_code/day12/"

var COUNT = 0

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
		allPaths := ""
		for _, j := range v.adjacentCaves {
			allPaths += j.name + ","
		}
		aoc.Log(k, "->", allPaths)
	}
	aoc.Log(s)
}

func IsBigCave(caveName string) bool {
	return caveName == strings.ToUpper(string(caveName))
}

func CanExploreSmallCave(part1 bool, smallCavesVisited []cave, currCave cave) bool {
	if part1 {
		return !Contains(smallCavesVisited, currCave)
	}
	return false
}
func DFS(currPos string, currPath string, adjMap map[string]cave, smallCavesVisited []cave, part1 bool) {
	if currPos == "end" {
		aoc.Log(currPath)
		COUNT++
	}

	// For each adjacent cave, we try to explore each of them
	for _, adjCave := range adjMap[currPos].adjacentCaves {
		// if we haven't explored an adjacent cave yet, try exploring it
		if IsBigCave(adjCave.name) {
			DFS(adjCave.name, currPath+" -> "+adjCave.name, adjMap, smallCavesVisited, part1)
		} else if CanExploreSmallCave(part1, smallCavesVisited, adjCave) {
			DFS(adjCave.name, currPath+" -> "+adjCave.name, adjMap, append(smallCavesVisited, adjCave), part1)
		}
	}
}

func PopulateAdjMap(input []string) map[string]cave {
	numOfCaves := len(input)
	adjMap := make(map[string]cave, numOfCaves)
	adjMap["start"] = cave{"start", false, make([]cave, 0)}

	for i := 0; i < numOfCaves; i++ {
		line := strings.Split(input[i], "-")
		head := cave{line[0], IsBigCave(line[0]), make([]cave, 0)}
		tail := cave{line[1], IsBigCave(line[1]), make([]cave, 0)}

		// Add {a : b} mapping
		if tail.name != "start" && head.name != "end" {
			if val, ok := adjMap[head.name]; ok {
				if !Contains(val.adjacentCaves, tail) {
					val.adjacentCaves = append(val.adjacentCaves, tail)
					adjMap[head.name] = val
				}
			} else {
				head.adjacentCaves = append(head.adjacentCaves, tail)
				adjMap[head.name] = head
			}
		}

		// Add {b : a} mapping
		if head.name != "start" && tail.name != "end" {
			if val, ok := adjMap[tail.name]; ok {
				if !Contains(val.adjacentCaves, head) {
					val.adjacentCaves = append(val.adjacentCaves, head)
					adjMap[tail.name] = val
				}
			} else {
				tail.adjacentCaves = append(tail.adjacentCaves, head)
				adjMap[tail.name] = tail
			}
		}

	}
	return adjMap
}

// SOLVER
func Solve(inputFile string) {
	aoc.Log("Solving")
	// Read input and initialize values
	result := 0
	input := aoc.ReadInput(inputFile, "\n")
	adjMap := PopulateAdjMap(input)
	part1 := true

	PrettyPrintCaveMap(adjMap)
	DFS("start", "start", adjMap, make([]cave, 0), part1)
	result = COUNT
	aoc.Log("result", result)
}

func main() {
	// Solve(basePath + "input/t.txt")
	// Solve(basePath + "input/t2.txt")
	// Solve(basePath + "input/t3.txt")
	Solve(basePath + "input/12.txt")
}
