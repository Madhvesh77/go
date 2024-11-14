package main

import ("fmt"; "math")

type Circle struct{
  radius float64
}

func main() {
  c := &Circle{float64(70.0)}
  fmt.Println("Area = ", calculateArea(c))
  fmt.Println("Radius = ", c.radius)
  fmt.Println("Area using method = ", c.calculateAreaMethod())
}

func calculateArea(circle *Circle) (area float64) {
  area = math.Pi * circle.radius * circle.radius
  circle.radius = 100
  return
}

func (circle *Circle) calculateAreaMethod() (area float64) {
  area = calculateArea(circle)
  return
}
