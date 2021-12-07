package main

import (
	"aoc"
)

// PROBLEM-SPECIFIC UTIL FUNCTIONS
func findCrabEnergy(delta int, part1 bool) int {
	val := 0
	for i := 1; i <= delta; i++ {
		if part1 {
			val += 1
		} else {
			val += i
		}
	}
	return val
}

func populateAllCrabs(initialCrabPos []string, allCrabs *map[int]int, part1 bool) (int, map[int]int) {
	highestEnergy := 0
	for _, crabPos := range initialCrabPos {
		if crabPos == "" {
			continue
		}
		crabPosVal := aoc.Atoi(crabPos)
		for k, _ := range *allCrabs {
			delta := aoc.Absolute(crabPosVal, k)
			(*allCrabs)[k] += findCrabEnergy(delta, part1)
			highestEnergy = aoc.Max(highestEnergy, (*allCrabs)[k])
		}
	}
	return highestEnergy, (*allCrabs)
}

func findLowestEnergy(highestEnergy int, allCrabs *map[int]int) (int, int) {
	lowestEnergy := highestEnergy
	position := 0
	for pos, energy := range *allCrabs {
		if energy < lowestEnergy {
			lowestEnergy = energy
			position = pos
		}
	}
	return position, lowestEnergy
}

// SOLVER
func Solve(inputFile string) {
	aoc.Log("Solving")
	// Read input and initialize values
	initialCrabPos := aoc.ReadInput(inputFile, ",")
	maxCrabPos, highestEnergy, position, result := 0, 0, 0, 0
	allCrabs := make(map[int]int) //

	// Find the largest crab position. This will be our upper bound
	for _, crabPos := range initialCrabPos {
		crabPosVal := aoc.Atoi(crabPos)
		maxCrabPos = aoc.Max(maxCrabPos, crabPosVal)
	}

	// With the largest crab position, we create a map with k:v pair
	// {position: totalEnergy}, from 0 to maxCrabPos
	for i := 0; i <= maxCrabPos; i++ {
		allCrabs[i] = 0
	}

	// Populate allCrabs map with the delta of each crab from that key
	// which is the index/position [0,maxCrabPos]
	highestEnergy, allCrabs = populateAllCrabs(initialCrabPos, &allCrabs, false)

	// Find lowest energy in allCrabs and set that as the result
	position, result = findLowestEnergy(highestEnergy, &allCrabs)
	aoc.Log(position, result)
}

func main() {
	Solve("input/test7.txt")
	Solve("input/7.txt")
}
