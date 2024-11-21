package main

import "fmt"

func main() {
  defer func ()  {
    errorMessage := recover()
    fmt.Println(errorMessage)
    fmt.Println("Program ended gracefully! :)")
  }()
  sampleSlice := []int{1, 2, 3}
  sampleSlice[4] = 5
  fmt.Println(sampleSlice)
}
