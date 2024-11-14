package main

import "fmt"

func makeEvenGenerator() func() uint {
  localVariable := uint(0)
  return func() (returnValue uint) {
    returnValue = localVariable
    localVariable += 2
    return
  }
}
func main() {
  generateEven := makeEvenGenerator()
  fmt.Println(generateEven())
  fmt.Println(generateEven())
  fmt.Println(generateEven())
}
