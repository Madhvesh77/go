package main

import "fmt"

type Shape interface {
  calculateArea() int
}

type Circle struct {
  radius int
}

type Rectangle struct {
  width int
  length int
}

func (rectangle *Rectangle) calculateArea() int {
  return rectangle.width * rectangle.length
}

func (circle *Circle) calculateArea() int {
  return int(3 * circle.radius * circle.radius)
}

func calculateAreaForAllShapes(shapes ...Shape) int {
  var total int
  for _, shape := range shapes {
    total += shape.calculateArea()
  }
  return total
}

func main() {
    rectangle := &Rectangle{4, 6}
    circle := &Circle{10}
    circleArea := circle.calculateArea()
    rectangleArea := rectangle.calculateArea()
    fmt.Println("Circle Area = ", circleArea, "\nRectangle Area = ", rectangleArea)
    totalArea := calculateAreaForAllShapes(circle, rectangle)
    fmt.Println("Total area of all the shapes = ", totalArea)
}
