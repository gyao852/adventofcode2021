package main

import (
	"2021advent_of_code/aoc"
	"strings"
)

// PROBLEM-SPECIFIC UTIL FUNCTIONS

// Returns a mapping between the cryptic string, and the corresponding digit
func InterpretLine(signalPattern []string) map[string]int {
	mapStrPatternToDigit := make(map[string]int)
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
			for _, c := range uniqueSignalPatterns[1] {
				if !aoc.Contains(pattern, c) {
					// Must be 2 or 5
					cntMissing := 0
					for _, c := range uniqueSignalPatterns[4] {
						if !aoc.Contains(pattern, c) {
							cntMissing++
						}
						if cntMissing == 1 {
							uniqueSignalPatterns[5] = strings.Split(pattern, "")
						} else if cntMissing == 2 {
							uniqueSignalPatterns[2] = strings.Split(pattern, "")
						} else {
							panic("Invalid logic")
						}
						continue
					}
				}
				// Must be 3
				uniqueSignalPatterns[3] = strings.Split(pattern, "")
				continue
			}
		case 6:
			// Must be 0,6,9
			for _, c := range uniqueSignalPatterns[1] {
				if !aoc.Contains(pattern, c) {
					// Must be 0 or 9
					for _, c := range uniqueSignalPatterns[4] {
						if !aoc.Contains(pattern, c) {
							// Must be 0
							uniqueSignalPatterns[0] = strings.Split(pattern, "")
						} else {
							// Must be 9
							uniqueSignalPatterns[9] = strings.Split(pattern, "")
						}
						continue
					}
				}
				// Must be 6
				uniqueSignalPatterns[6] = strings.Split(pattern, "")
				continue
			}
		default:
			continue
		}
	}

	// 0 : has len(6); Comapre with 5 stringcode, if not all chars in 5 stringCode than 0
	// 1 : if strCode == len(2)
	// 2 : has len(5); compare with 4 strCode, should be missing 2
	// 3 : has len(5); if both characters from 1 are in the string
	// 4 : if strCode == len(4)
	// 5 : has len(5); compare with 4 strCode, should be missing 1
	// 6 : has len(6); Compare with 1 strCode, should be missing 1
	// 7 : if strCode == len(3)
	// 8 : if strCode == len(7)
	// 9 : has len(6); Compare with 5 strCode. If there is one leftover than 9
	for _, pattern := range signalPattern {
		if uniqueSignalPatterns[len(pattern)] > 0 {
			mapStrPatternToDigit[pattern] = uniqueSignalPatterns[len(pattern)]
		}
	}
	return mapStrPatternToDigit
}

func Permutation(input []rune, permutations []string, start int, end int) []string {
	// permutations = append(permutations,string(input))
	aoc.Log(permutations)
	if start == end {
		for _, p := range permutations {
			if string(input) == p {
				return nil
			}
		}
		permutations = append(permutations, string(input))
		return permutations
	} else {
		for i := start; i <= end; i++ {
			input[start], input[i] = input[i], input[start]
			permutations = append(permutations, Permutation(input, permutations, start+1, end)...)
			input[start], input[i] = input[i], input[start]
		}
	}
	return permutations
}

func InterpretDigits(PatternToDigit map[string]int, signalDigits []string, part1 bool) map[int]int {
	lineResult := map[int]int{
		1: 0,
		4: 0,
		7: 0,
		8: 0,
	}

	for _, digit := range signalDigits {
		if part1 {
			switch len(digit) {
			case 2:
				lineResult[1]++
			case 4:
				lineResult[4]++
			case 3:
				lineResult[7]++
			case 7:
				lineResult[8]++
			default:
				continue
			}
		} else {
			// Get all permutations of digit

			// Wtf was I doing...
			// digitPermutations := make([]string,1)
			// digitPermutations = Permutation([]rune(digit), digitPermutations, 0, len(digit)-1)
			// aoc.Log(digitPermutations)
			// for _,d := range digitPermutations {
			// 	if PatternToDigit[d] > 0 {
			// 		lineResult[PatternToDigit[digit]]++
			// 	}
			// }
		}
	}
	return lineResult
}

// SOLVER
func Solve(inputFile string) {
	aoc.Log("Solving")
	// Read input and initialize values
	result := 0
	part1 := true
	input := aoc.ReadInput(inputFile, "\n")
	for i := 0; i < len(input); i++ {
		splitLine := strings.Split(input[i], " | ")
		signalPattern, digitOutput := strings.Split(splitLine[0], " "), strings.Split(splitLine[1], " ")
		signalPattern = signalPattern[:len(signalPattern)-1]
		interpretedLine := InterpretLine(signalPattern)
		interpretedDigits := InterpretDigits(interpretedLine, digitOutput, part1)
		aoc.Log(interpretedLine, "")
		aoc.Log(interpretedDigits)
		for _, v := range interpretedDigits {
			result += v
		}
	}
	aoc.Log(result)
}

func main() {
	// Solve("input/test8.txt")
	Solve("input/8.txt")
}
