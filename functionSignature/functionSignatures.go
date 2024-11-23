package main

import "fmt"

//returns half of the number and a boolean representiung whether the given number is true
func Half(number int) (int, bool) {
  isEven := true
  if number%2 != 0 { isEven = false }
  return number/2, isEven
}

func summa(sampleSlice []int) {
  fmt.Println(sampleSlice)
}
func main() {
  fmt.Println(Half(2))
  fmt.Println(Half(5))
  sampleSlice := make([]int, 2)
  sampleSlice[0] = 2
  sampleSlice[1] = 4
  summa(sampleSlice)
}
