package main

import "fmt"

func main() {
 go sayHello()
 var input string
 fmt.Scanln(&input)
}

func sayHello() {
	fmt.Println("Hello...")
}