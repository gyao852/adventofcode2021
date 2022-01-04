package main

import (
	"2021advent_of_code/aoc"
	"math"
	"regexp"
	"strings"
)

var basePath = "/Users/gyao/Documents/Personal Projects/go_projects/src/2021advent_of_code/day13/"

// PROBLEM-SPECIFIC UTIL FUNCTIONS
type sheet struct {
	paper    [][]string
	dotCount int
}

type fold struct {
	plane string
	coord int
}

type foldInstructions struct {
	instruction []fold
}

func PrettyPrintPaper(paper [][]string) {
	for _, r := range paper {
		aRow := ""
		for _, c := range r {
			aRow += c + " "
		}
		aoc.Log(aRow)
	}
}

func PopulatePaper(input []string) (sheet, foldInstructions) {
	maxX := 0
	maxY := 0
	counter := 0
	instructions := foldInstructions{make([]fold, 0)}
	// First find the max x and y coordinates
	for _, v := range input {
		coordRegex, _ := regexp.Compile("\\d+,\\d+")
		if coordRegex.MatchString(v) {
			coords := strings.Split(v, ",")
			maxX = int(math.Max(float64(maxX), float64(aoc.Atoi(coords[0]))))
			maxY = int(math.Max(float64(maxY), float64(aoc.Atoi(coords[1]))))
		} else {
			continue
		}
	}
	aPaper := sheet{make([][]string, maxY+1), 0}
	// Initialize aPaper with all empty '.'
	for r, _ := range aPaper.paper {
		aPaper.paper[r] = make([]string, maxX+1)
		for c, _ := range aPaper.paper[r] {
			aPaper.paper[r][c] = "."
		}
	}

	for _, v := range input {
		coordRegex, _ := regexp.Compile("\\d+,\\d+")
		foldRegex, _ := regexp.Compile("\\w=\\d+")
		if coordRegex.MatchString(v) {
			coords := strings.Split(v, ",")
			aPaper.paper[aoc.Atoi(coords[1])][aoc.Atoi(coords[0])] = "#"
			counter++
		} else if foldRegex.MatchString(v) {
			foldParse := strings.Split(foldRegex.FindString(v), "=")
			aFold := fold{foldParse[0], aoc.Atoi(foldParse[1])}
			instructions.instruction = append(instructions.instruction, aFold)
		} else {
			continue
		}
	}
	aPaper.dotCount = counter
	return aPaper, instructions
}

func Folder(part1 bool, aPaper sheet, instructions foldInstructions) sheet {
	return nil
}

// SOLVER
func Solve(inputFile string) {
	aoc.Log("Solving")
	// Read input and initialize values
	result := 0
	input := aoc.ReadInput(inputFile, "\n")
	aPaper, instructions := PopulatePaper(input)
	part1 := true
	// PrettyPrintPaper(aPaper.paper)
	resultPaper := Folder(part1, aPaper, instructions)
	result = resultPaper.dotCount
	PrettyPrintPaper(result.paper)
	aoc.Log("result", result)
}

func main() {
	Solve(basePath + "input/test13.txt")
	// Solve(basePath + "input/13.txt")
}
