package main

import (
  "os"
  "fmt"
  "strconv"
  "strings"
  "regexp"
  // "math"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func crossValue(currentMove int, allBingoBoards *map[int][][]int) *map[int][][]int {
  for _, board := range *allBingoBoards {
    for i, row := range board {
      for j := len(row)-1; j >= 0; j-- {
        if row[j] == currentMove {
          if j != 0 {
            // fmt.Println(row, "->")
            row = append(row[:j],row[j+1:]...)
          } else {
            row = row[1:]
          }
          board[i] = row
        }
      }
    }
  }
  return allBingoBoards
}

// func checkWinner(allBingoBoards *map[int][][]int) int {
//   for i, board := range *allBingoBoards {
//     for _, row := range board {
//       if len(row) == 0 {
//         return i
//       }
//     }
//   }
//   return -1
// }

func checkForWinners(allBingoBoards *map[int][][]int) []int {
  var winners []int
  for i, board := range *allBingoBoards {
    for _, row := range board {
      if len(row) == 0 {
        winners = append(winners,i)
        break
      }
    }
  }
  return winners
}

func sumBoard(board [][]int) int {
  sum := 0
  fmt.Println("Winning Board:", board)
  for i := 0; i < len(board) / 2; i++ {
    for _, r := range board[i] {
      sum += r
    }
  }
  fmt.Println("Sum:", sum)
  return sum
}

// func runBingo(moveSet []string, allBingoBoards *map[int][][]int) int {
//   sumOfWinningBoard := 0
//   for _, move := range moveSet {
//     currentMove, err := strconv.Atoi(move)
//     check(err)
//     allBingoBoards = crossValue(currentMove, allBingoBoards)
//     fmt.Println("Current Move", currentMove)
//     winner_index := checkWinner(allBingoBoards)
//     if winner_index >= 0 {
//       fmt.Println("Winner is board:", winner_index)
//       sumOfWinningBoard = sumBoard((*allBingoBoards)[winner_index])
//       return currentMove * sumOfWinningBoard
//     }
//   }
//   panic("No winners were found!")
// }

func runBingo(moveSet []string, allBingoBoards *map[int][][]int) int {
  sumOfWinningBoard := 0
  sum := 0
  for _, move := range moveSet {
    if len((*allBingoBoards)) == 0 {
      break
    }
    currentMove, err := strconv.Atoi(move)
    check(err)
    allBingoBoards = crossValue(currentMove, allBingoBoards)
    fmt.Println("Current Move", currentMove)
    winners := checkForWinners(allBingoBoards)
    fmt.Println("winners", winners)
    if len(winners) >= 0 {
      for _,winner_index := range winners {
        fmt.Println("Last Winning board:", winner_index)
        sumOfWinningBoard = sumBoard((*allBingoBoards)[winner_index])
        sum = currentMove * sumOfWinningBoard
        fmt.Println("Removing board", winner_index)
        delete((*allBingoBoards),winner_index)
      }
    }
  }
  if sum == 0 {
    panic("No last winners were found!")
  }
  return sum
}

func parseBingoBoards(rawBingoBoards []string) map[int][][]int {
  parsedBoards := make(map[int][][]int)
  // There are some extra whitespaces for some reason
  space := regexp.MustCompile(`\s+`)
  for index, board := range rawBingoBoards {
    rows := strings.Split(string(board), "\n")
    lenOfRow := len(strings.Split(space.ReplaceAllString(rows[0], " ")," "))
    for _ , v := range strings.Split(space.ReplaceAllString(rows[0], " ")," ") {
      if v == "" {
        lenOfRow -= 1
      }
    }
    singleBoard := make([][]int, 2 * lenOfRow)

    for r, row := range rows {
      if row == "" {
        continue
      }
      var a_row []int
      // Clean up double spaces for each row
      filteredRow := strings.Split(space.ReplaceAllString(row, " ")," ")
      for i , v := range filteredRow {
        if v == "" {
          if i != 0 {
            filteredRow = append(filteredRow[i-1:i],filteredRow[i+1:]...)
          } else {
            filteredRow = filteredRow[1:]
          }

        }
      }

      // Push each row's value in a_row and in respective column
      for c, v := range filteredRow {
        bingo_val, err := strconv.Atoi(v)
        check(err)
        a_row = append(a_row, bingo_val)
        singleBoard[lenOfRow + c] = append(singleBoard[c + lenOfRow],bingo_val)
      }
      singleBoard[r] = a_row
    }
    parsedBoards[index] = singleBoard
  }
  return parsedBoards
}

func main() {
    fmt.Println("Solving!")
    input, err := os.ReadFile("/Users/gyao/Downloads/4.txt")
    check(err)
    result := 0
    in := strings.Split(string(input),"\n\n")
    moveSet := strings.Split(in[0],",")
    rawBingoBoards := in[1:len(in)]
    allBingoBoards := parseBingoBoards(rawBingoBoards)
    fmt.Println("All Moves")
    fmt.Println(moveSet)
    fmt.Println("All Bingo Boards")
    fmt.Println(allBingoBoards)
    result = runBingo(moveSet, &allBingoBoards)
    fmt.Println(result)
}
