package main

import (
  "os"
  "fmt"
  "strconv"
  "strings"
  // "regexp"
  // "math"
)

type seaCoordinate struct {
  x, y int
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Max returns the larger of x or y.
func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
    if x > y {
        return y
    }
    return x
}
func main() {
    fmt.Println("Solving!")
    input, err := os.ReadFile("/Users/gyao/Downloads/5.txt")
    check(err)
    result := 0
    lines := strings.Split(string(input),"\n")
    lineCoordinates := make([]seaCoordinate,1)
    // First, find the largest x/y coordinates in input
    // to initilaize a 2d list of that size
    maxXYCoordinate := 0
    for _,line := range(lines) {
      if line == "" {
        continue
      }
      endPoints := strings.Split(line," -> ")
      start := strings.Split(endPoints[0], ",")
      end := strings.Split(endPoints[1], ",")
      for _,v := range start {
        val, err := strconv.Atoi(v)
        check(err)
        if maxXYCoordinate < val {
          maxXYCoordinate = val
        }
      }
      for _,v := range end {
        val, err := strconv.Atoi(v)
        check(err)
        if maxXYCoordinate < val {
          maxXYCoordinate = val
        }
      }

      x1, _ := strconv.Atoi(start[0])
      y1, _ := strconv.Atoi(start[1])
      x2, _ := strconv.Atoi(end[0])
      y2, _ := strconv.Atoi(end[1])
      fmt.Println(x1,x2,y1,y2)
      fmt.Println(x1 <= x2 && y1 > y2)
      // Fetch DiagonalLineCoordinates
      if x1 != x2 && y1 != y2 {
        if x1 <= x2 && y1 < y2 {
          fmt.Println("diagonal left up",x1,y1,x2,y2)
          i := x1
          j := y1
          for i <= x2 {
            lineCoordinates = append(lineCoordinates,seaCoordinate{i,j})
            fmt.Println(i,j)
            i += 1
            j += 1
          }
        } else if x1 <= x2 && y2 < y1 {
          fmt.Println("diagonal right up",x1,y1,x2,y2)
          i := x1
          j := y1
          for i <= x2 {
            lineCoordinates = append(lineCoordinates,seaCoordinate{i,j})
            fmt.Println(i,j)
            i += 1
            j -= 1
          }
        } else if x1 > x2 && y2 > y1 {
          fmt.Println("diagonal right down",x1,y1,x2,y2)
          i := x2
          j := y2
          for i <= x1 {
            lineCoordinates = append(lineCoordinates,seaCoordinate{i,j})
            fmt.Println(i,j)
            i += 1
            j -= 1
          }
        } else if x1 > x2 && y2 < y1 {
          fmt.Println("diagonal left up",x1,y1,x2,y2)
          i := x2
          j := y2
          for i <= x1 {
            lineCoordinates = append(lineCoordinates,seaCoordinate{i,j})
            fmt.Println(i,j)
            i += 1
            j += 1
          }
        }
      } else {
        // Fetch lineCoordinates
        for i := Min(x1,x2); i <= Max(x1,x2); i++ {
          for j := Min(y1,y2); j <= Max(y1,y2); j++ {
            lineCoordinates = append(lineCoordinates,seaCoordinate{i,j})
          }
        }
      }


    }
    fmt.Println("Largest coordinate:",maxXYCoordinate)
    fmt.Println("Creating 2D array and populating of size ",maxXYCoordinate,"x",maxXYCoordinate)
    fmt.Println("Fetched Line coordinates: ",lineCoordinates)
    seaMap := make([][]int, maxXYCoordinate + 1)
    for i := range seaMap {
        seaMap[i] = make([]int, maxXYCoordinate + 1)
    }
    // Now populate seaMap with coordinates from above
    for _,coord := range lineCoordinates {
      seaMap[coord.x][coord.y] += 1
    }


    for _,row := range seaMap {
      for _,col := range row {
        if col > 1 {
          result += 1
        }
      }
    }
    fmt.Println("Final Sea Map:",seaMap)
    fmt.Println(result)
}

// func main() {
//     fmt.Println("Solving!")
//     input, err := os.ReadFile("/Users/gyao/Downloads/5.txt")
//     check(err)
//     result := 0
//     lines := strings.Split(string(input),"\n")
//     lineCoordinates := make([]seaCoordinate,1)
//     // First, find the largest x/y coordinates in input
//     // to initilaize a 2d list of that size
//     maxXYCoordinate := 0
//     for _,line := range(lines) {
//       if line == "" {
//         continue
//       }
//       endPoints := strings.Split(line," -> ")
//       start := strings.Split(endPoints[0], ",")
//       end := strings.Split(endPoints[1], ",")
//       for _,v := range start {
//         val, err := strconv.Atoi(v)
//         check(err)
//         if maxXYCoordinate < val {
//           maxXYCoordinate = val
//         }
//       }
//       for _,v := range end {
//         val, err := strconv.Atoi(v)
//         check(err)
//         if maxXYCoordinate < val {
//           maxXYCoordinate = val
//         }
//       }
//
//
//       x1, _ := strconv.Atoi(start[0])
//       y1, _ := strconv.Atoi(start[1])
//       x2, _ := strconv.Atoi(end[0])
//       y2, _ := strconv.Atoi(end[1])
//
//       if x1 != x2 && y1 != y2 {
//         continue
//       }
//       // Fetch lineCoordinates
//       for i := Min(x1,x2); i <= Max(x1,x2); i++ {
//         for j := Min(y1,y2); j <= Max(y1,y2); j++ {
//           lineCoordinates = append(lineCoordinates,seaCoordinate{i,j})
//         }
//       }
//
//     }
//     fmt.Println("Largest coordinate:",maxXYCoordinate)
//     fmt.Println("Creating 2D array and populating of size ",maxXYCoordinate,"x",maxXYCoordinate)
//     seaMap := make([][]int, maxXYCoordinate + 1)
//     for i := range seaMap {
//         seaMap[i] = make([]int, maxXYCoordinate + 1)
//     }
//     // Now populate seaMap with coordinates from above
//     for _,coord := range lineCoordinates {
//       seaMap[coord.x][coord.y] += 1
//     }
//
//
//     for _,row := range seaMap {
//       for _,col := range row {
//         if col > 1 {
//           result += 1
//         }
//       }
//     }
//     fmt.Println(seaMap)
//     fmt.Println(result)
// }
