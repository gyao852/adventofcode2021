package main

import (
  "os"
  "fmt"
  "strconv"
  "strings"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main() {
    fmt.Println("Solving!")
    input, err := os.ReadFile("input/test2.txt")
    check(err)
    result := 0
    horizontal_pos := 0
    depth_pos := 0
    aim_magnitude := 0
    const (
    	FOWARD string = "forward"
    	DOWNWARD      = "down"
    	UPWARD        = "up"
    )

    vectors := strings.Split(string(input),"\n")
    for i := 0; i < len(vectors); i++ {
      vector := strings.Split(string(vectors[i]), " ")
      direction := vector[0]
      if len(direction) == 0 {
        // Split creates an empty last value
        continue
      }
      magnitude, err := strconv.Atoi(vector[1])
      check(err)
      switch direction {
      case FOWARD:
        horizontal_pos += magnitude
        depth_pos += (aim_magnitude * magnitude)
      case DOWNWARD:
        aim_magnitude += magnitude
      case UPWARD:
        aim_magnitude -= magnitude
      }
    }
  fmt.Println(horizontal_pos)
  fmt.Println(depth_pos)
  result = horizontal_pos * depth_pos
  fmt.Println(result)
}

// func main() {
//     fmt.Println("Solving!")
//     input, err := os.ReadFile("input/2.txt")
//     check(err)
//     result := 0
//     forward_magnitude := 0
//     downward_magnitude := 0
//     const (
//     	FOWARD string = "forward"
//     	DOWNWARD      = "down"
//     	UPWARD        = "up"
//     )
//
//     vectors := strings.Split(string(input),"\n")
//     for i := 0; i < len(vectors); i++ {
//       vector := strings.Split(string(vectors[i]), " ")
//       direction := vector[0]
//       if len(direction) == 0 {
//         // Split creates an empty last value
//         continue
//       }
//       magnitude, err := strconv.Atoi(vector[1])
//       check(err)
//       switch direction {
//       case FOWARD:
//         forward_magnitude += magnitude
//       case DOWNWARD:
//         downward_magnitude += magnitude
//       case UPWARD:
//         downward_magnitude -= magnitude
//       }
//     }
//   result = forward_magnitude * downward_magnitude
//   fmt.Println(result)
// }
