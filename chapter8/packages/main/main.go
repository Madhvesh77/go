package main

import (
	"fmt"
	"../math/mathPackage"
)

func main() {
	arr := []float64{1, 2, 3, 4}
	average := mathPackage.average(arr)
	fmt.Println("Average : ", average)
}
