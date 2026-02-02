package main

import (
  "fmt"
  "math"
)

// Shape interface defines behavior for shapes
type Shape interface {
  Area() float64
  Perimeter() float64
}

// Rectangle implements Shape
type Rectangle struct {
  Width  float64
  Height float64
}

func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
  return 2 * (r.Width + r.Height)
}

// Circle implements Shape
type Circle struct {
  Radius float64
}

func (c Circle) Area() float64 {
  return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
  return 2 * math.Pi * c.Radius
}

func main() {
  fmt.Println("=== Interface Basics ===")

  rect := Rectangle{Width: 10, Height: 5}
  circle := Circle{Radius: 7}

  // Both types implement Shape
  fmt.Println("\nRectangle:")
  fmt.Printf("  Area: %.2f\n", rect.Area())
  fmt.Printf("  Perimeter: %.2f\n", rect.Perimeter())

  fmt.Println("\nCircle:")
  fmt.Printf("  Area: %.2f\n", circle.Area())
  fmt.Printf("  Perimeter: %.2f\n", circle.Perimeter())
}
