package main

import "fmt"

func main() {
  simpleMap := make(map[string]int)
  simpleMap["abc"] = 2
  name, ok := simpleMap["a"]
  fmt.Println("name ", name, "ok ", ok)
}
