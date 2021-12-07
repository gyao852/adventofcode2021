package main

import (
  "os"
  "fmt"
  "strconv"
  "strings"
  "math"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main() {
    fmt.Println("Solving!")
    input, err := os.ReadFile("/Users/gyao/Downloads/3.txt")
    check(err)
    result := 0
    oxygen_generator_rating := 0
    co2_scrubber_rating := 0
    bit_criteria := 0
    bit_criteria_pos := 0
    inverse_bit_criteria := 0
    var bit_cnt = make(map[int]int) // (holds bit_position : cnt of 1s)
    var elgible_bits = make(map[string]string) // holds elgible bits to consider
    bits := strings.Split(string(input),"\n")
    total_bits := len(bits)

    // Map of each (bit_position : cnt) of 1s in that position
    for i := 0; i < len(bits[0]); i++ {
      bit_cnt[i] = 0
    }

    // Populating the above map. This will help determine the
    // first bit criteria
    for i := 0; i < total_bits; i++ {
      bit_array := strings.Split(string(bits[i]), "")
      if len(bit_array) == 0 {
        // Split creates an empty last value
        continue
      }
      val,err := strconv.Atoi(bit_array[bit_criteria_pos])
      check(err)
      bit_cnt[bit_criteria_pos] += val
    }

    // Fetch first bit_criteria
    fmt.Println(bit_cnt)
    if bit_cnt[bit_criteria_pos] < (total_bits / 2) - 1 {
      bit_criteria = 0
      inverse_bit_criteria = 1
    } else {
      bit_criteria = 1
      inverse_bit_criteria = 0
    }

    // Generating bitset of eligible bits
    for i := 0; i < total_bits; i++ {
      bit_array := strings.Split(string(bits[i]), "")
      if len(bit_array) == 0 {
        // Split creates an empty last value
        continue
      }
      elgible_bits[string(bits[i])] = "t"
    }

    // Calculating oxygen_generator_rating
    for len(elgible_bits) > 1 {
      // Based on bit_criteria, remove all non-legal bits
      fmt.Println(bit_criteria)
      fmt.Println(elgible_bits)
      fmt.Println("->")
      for bit, _ := range elgible_bits {
        val,err := strconv.Atoi(string(bit[bit_criteria_pos]))
        check(err)
        if val != bit_criteria {
          delete(elgible_bits,bit)
        }
      }
      fmt.Println(elgible_bits)

      // Break early if we found the bit
      if len(elgible_bits) <= 1 {
        break
      }

      // Update bit_criteria to next
      bit_criteria_pos += 1
      for bit, _ := range elgible_bits {
        bit_array := strings.Split(string(bit), "")
        if len(bit_array) == 0 {
          // Split creates an empty last value
          continue
        }
        val,err := strconv.Atoi(bit_array[bit_criteria_pos])
        check(err)
        bit_cnt[bit_criteria_pos] += val
      }
      if bit_cnt[bit_criteria_pos] < int(math.Round(float64(len(elgible_bits))/ 2.0)) {
        bit_criteria = 0
      } else {
        bit_criteria = 1
      }
    }

    fmt.Println("Only one value left for oxygen_generator_rating:")
    for bit_string, _ := range elgible_bits {
      val, err := strconv.ParseInt(bit_string, 2, 64);
      check(err)
      oxygen_generator_rating = int(val)
    }
    fmt.Println(oxygen_generator_rating)

    // Generating bitset of eligible bits
    bit_criteria_pos = 0
    for i := 0; i < len(bits[0]); i++ {
      bit_cnt[i] = 0
    }
    for i := 0; i < total_bits; i++ {
      bit_array := strings.Split(string(bits[i]), "")
      if len(bit_array) == 0 {
        // Split creates an empty last value
        continue
      }
      elgible_bits[string(bits[i])] = "t"
    }

    // Calculating co2_scrubber_rating
    for len(elgible_bits) > 1 {
      // Based on bit_criteria, remove all non-legal bits
      fmt.Println(inverse_bit_criteria)
      fmt.Println(elgible_bits)
      fmt.Println("->")
      for bit, _ := range elgible_bits {
        val,err := strconv.Atoi(string(bit[bit_criteria_pos]))
        check(err)
        if val != inverse_bit_criteria {
          delete(elgible_bits,bit)
        }
      }
      fmt.Println(elgible_bits)

      // Break early if we found the bit
      if len(elgible_bits) <= 1 {
        break
      }

      // Update bit_criteria to next
      bit_criteria_pos += 1
      for bit, _ := range elgible_bits {
        bit_array := strings.Split(string(bit), "")
        if len(bit_array) == 0 {
          // Split creates an empty last value
          continue
        }
        val,err := strconv.Atoi(bit_array[bit_criteria_pos])
        check(err)
        bit_cnt[bit_criteria_pos] += val
      }

      if bit_cnt[bit_criteria_pos] >= int(math.Round(float64(len(elgible_bits))/ 2.0)) {
        inverse_bit_criteria = 0
      } else {
        inverse_bit_criteria = 1
      }
    }

    fmt.Println("Only one value left for co2_scrubber_rating:")
    for bit_string, _ := range elgible_bits {
      val, err := strconv.ParseInt(bit_string, 2, 64);
      check(err)
      co2_scrubber_rating = int(val)

    }

    fmt.Println(co2_scrubber_rating)
    result = oxygen_generator_rating * co2_scrubber_rating
    fmt.Println(result)
}

// func main() {
//     fmt.Println("Solving!")
//     input, err := os.ReadFile("/Users/gyao/Downloads/3.txt")
//     check(err)
//     result := 0
//     gamma_rate := 0
//     epsilon_rate := 0
//     var bit_cnt = make(map[int]int)
//     bits := strings.Split(string(input),"\n")
//     total_bits := len(bits)
//     for i := 0; i < len(bits[0]); i++ {
//       bit_cnt[i] = 0
//     }
//
//     for i := 0; i < total_bits; i++ {
//       bit_array := strings.Split(string(bits[i]), "")
//       if len(bit_array) == 0 {
//         // Split creates an empty last value
//         continue
//       }
//       for j := 0; j < len(bit_array); j++ {
//         val,err := strconv.Atoi(bit_array[j])
//         check(err)
//         bit_cnt[j] += val
//       }
//     }
//
//     for i := 0; i < len(bits[0]); i++ {
//       if bit_cnt[i] < total_bits / 2 {
//         bit_cnt[i] = 0
//       } else {
//         bit_cnt[i] = 1
//       }
//     }
//
//     for i := len(bits[0]) - 1; i >= 0; i-- {
//       gamma_rate += (bit_cnt[i] * int(math.Pow(2, float64(len(bits[0])-1-i))))
//       if bit_cnt[i] == 0 {
//         epsilon_rate += (1 * int(math.Pow(2, float64(len(bits[0])-1-i))))
//       } else {
//         epsilon_rate += 0
//       }
//     }
//   fmt.Println(gamma_rate)
//   fmt.Println(epsilon_rate)
//   result = gamma_rate * epsilon_rate
//   fmt.Println(result)
// }
