package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// func calculateDepthWindow(i int) {
//     return i
// }
func main() {
	fmt.Println("Solving!")
	input, err := os.ReadFile("/Users/gyao/Downloads/1.txt")
	check(err)
	result := 0
	depths := strings.Split(string(input), "\n")
	var last_depth_window int
	for i := 0; i < 3; i++ {
		depth, err := strconv.Atoi(depths[i])
		check(err)
		last_depth_window += depth
	}

	for i := 1; i < len(depths)-3; i++ {
		var depth_window int
		for j := i; j < i+3; j++ {
			depth, err := strconv.Atoi(depths[j])
			check(err)
			depth_window += depth
		}
		if depth_window > last_depth_window {
			result += 1
		}
		last_depth_window = depth_window
	}
	fmt.Println(result)
}

//
// func main() {
//     fmt.Println("Solving!")
//     input, err := os.ReadFile("/Users/gyao/Downloads/1.txt")
//     check(err)
//     result := 0
//     depths := strings.Split(string(input),"\n")
//     last_depth, err := strconv.Atoi(depths[0])
//     check(err)
//     for index, depth := range depths {
//       if index == 0 || len(depths) == 0 {
//         continue
//       }
//       d, err := strconv.Atoi(depth)
//       fmt.Println(d)
//       if err != nil || d == 0{
//         fmt.Println("Empty value found, skipping")
//       }
//       if d > last_depth {
//         result += 1
//       }
//       last_depth = d
//     }
//   fmt.Println(result)
// }
