package main

import (
	"2021advent_of_code/aoc"
	"regexp"
	"strings"
)

var basePath = "/Users/gyao/Documents/Personal Projects/go_projects/src/2021advent_of_code/day13/"

// PROBLEM-SPECIFIC UTIL FUNCTIONS
type Node struct {
	data string
	next *Node
	prev *Node
}
type Polymer struct {
	head *Node
	tail *Node
}

type InstructionSet struct {
	instruction map[string]string
}

func PrettyPrintPolymer(polymer Polymer) {
	node := polymer.head
	polymerString := ""
	for node != nil {
		polymerString += node.data + "->"
		node = node.next
	}
	aoc.Log(polymerString)
}

func parseInput(input []string) (Polymer, InstructionSet) {
	var instructions InstructionSet
	polymer Polymer = Polymer{}
	for _, v := range input {
		instructionRegex, _ := regexp.Compile("\\w+->\\w+")
		if instructionRegex.MatchString(v) {
			anInstruction := strings.Split(v, " -> ")
			instructions.instruction[anInstruction[0]] = anInstruction[1]
		} else {
			basePolymerRegex, _ := regexp.Compile("\\w+")
			if basePolymerRegex.MatchString(v) {
				// polymer &Node{"X",nil,nil}
			} else {
				continue
			}
		}
	}
	return polymer, instructions
}

// SOLVER
func Solve(inputFile string) {
	aoc.Log("Solving")
	// Read input and initialize values
	result := 0
	input := aoc.ReadInput(inputFile, "\n")
	aPaper, instructions := parseInput(input)
	part1 := false
	resultPaper := Folder(part1, aPaper, instructions)
	PrettyPrintPaper(resultPaper.paper)
	result = resultPaper.dotCount
	aoc.Log("result", result)
}

func main() {
	// Solve(basePath + "input/test13.txt")
	Solve(basePath + "input/13.txt")
}
