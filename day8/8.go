package main

import (
	"2021advent_of_code/aoc"
	"strings"
	"math"
)

// PROBLEM-SPECIFIC UTIL FUNCTIONS

// Returns a mapping between the cryptic string, and the corresponding digit
func InterpretLine(signalPattern []string) map[string]int {
	uniqueSignalPatterns := map[int][]string{
		0: make([]string, 0),
		1: make([]string, 0),
		2: make([]string, 0),
		3: make([]string, 0),
		4: make([]string, 0),
		5: make([]string, 0),
		6: make([]string, 0),
		7: make([]string, 0),
		8: make([]string, 0),
		9: make([]string, 0),
	}
	// First find digits that have unique properties
	for _, pattern := range signalPattern {
		switch len(pattern) {
		case 2:
			uniqueSignalPatterns[1] = strings.Split(pattern, "")
		case 4:
			uniqueSignalPatterns[4] = strings.Split(pattern, "")
		case 3:
			uniqueSignalPatterns[7] = strings.Split(pattern, "")
		case 7:
			uniqueSignalPatterns[8] = strings.Split(pattern, "")
		default:
			continue
		}
	}

	// Now use deductive reasoning for the remaining digits
	for _, pattern := range signalPattern {
		switch len(pattern) {
		case 5:
			// Must be 2,3,5
			cntMissing := 0
			for _, c := range uniqueSignalPatterns[1] {
				if aoc.Contains(pattern, c) {
					cntMissing++
				}
			}
			if cntMissing == len(uniqueSignalPatterns[1]) {
				// Must be 3 since 2 and 5 both should
				uniqueSignalPatterns[3] = strings.Split(pattern, "")
				break
			} else {
				// Must be 2 or 5
				cntMissing = 0
				for _, c := range uniqueSignalPatterns[4] {
					if aoc.Contains(pattern, c) {
						cntMissing++
					}
				}
					// This isn't working on 2...
					if cntMissing == 3 {
						uniqueSignalPatterns[5] = strings.Split(pattern, "")
					} else if cntMissing == 2 {
						uniqueSignalPatterns[2] = strings.Split(pattern, "")
						break
					} else {
						panic("Invalid logic!")
					}
			}



		continue
		case 6:
			// Must be 0,6,9
			if aoc.Contains(pattern,uniqueSignalPatterns[1][0]) &&
			aoc.Contains(pattern,uniqueSignalPatterns[1][1]) {
				// Must be 0 or 9
				is4Subset := true
				for _, c := range uniqueSignalPatterns[4] {
					if !aoc.Contains(pattern, c) {

						is4Subset = false
					}
				}
				if is4Subset {
					uniqueSignalPatterns[9] = strings.Split(pattern, "")
				} else {
					uniqueSignalPatterns[0] = strings.Split(pattern, "")
				}
			} else {
				uniqueSignalPatterns[6] = strings.Split(pattern, "")
			}
		default:
			continue
		}
	}
	mapStrPatternToDigit := make(map[string]int,len(uniqueSignalPatterns))
	for k,v := range(uniqueSignalPatterns) {
		mapStrPatternToDigit[strings.Join(v,"")] = k
	}
	return mapStrPatternToDigit
}

func InterpretDigits(pattternToDiggit map[string]int, signalDigits []string, part1 bool) int {
	signalVal := 0
	for i,strDigit := range signalDigits {
		for k,v := range pattternToDiggit {
			if len(k) == len(strDigit){
				isPermutation := true
				for _,c := range strings.Split(strDigit,"") {
					if !aoc.Contains(k,c) {
						isPermutation = false
					}
				}
				if isPermutation {
					signalVal += int(math.Pow(10,float64(len(signalDigits)-i-1)) * float64(v))
					break
				}
			}
			continue
		}
	}
	return signalVal
}

// SOLVER
func Solve(inputFile string) {
	aoc.Log("Solving")
	// Read input and initialize values
	result := 0
	part1 := true
	input := aoc.ReadInput(inputFile, "\n")
	for i := 0; i < len(input); i++ {
		splitLine := strings.Split(input[i], "| ")
		signalPattern, digitOutput := strings.Split(splitLine[0], " "), strings.Split(splitLine[1], " ")
		signalPattern = signalPattern[:len(signalPattern)-1]
		interpretedLine := InterpretLine(signalPattern)
		interpretedDigits := InterpretDigits(interpretedLine, digitOutput, part1)
		aoc.Log(signalPattern, digitOutput, "=>")
		aoc.Log(interpretedLine)
		// aoc.Log(interpretedDigits)
		result += interpretedDigits
	}
	aoc.Log(result)
}

func main() {
	// Solve("input/test8.txt")
	Solve("input/8.txt")
}
