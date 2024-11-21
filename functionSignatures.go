package main

import "fmt"

func half(number int) (int, bool) {
  isEven := true
  if number%2 != 0 { isEven = false }
  return number/2, isEven
}

func summa(sampleSlice []int) {
  fmt.Println(sampleSlice)
}
func main() {
  fmt.Println(half(2))
  fmt.Println(half(5))
  sampleSlice := make([]int, 2)
  sampleSlice[0] = 2
  sampleSlice[1] = 4
  summa(sampleSlice)
}
