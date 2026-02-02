package main

import (
  "fmt"
  "math"
)

// Shape interface
type Shape interface {
  Area() float64
  Describe() string
}

// Rectangle
type Rectangle struct {
  Width, Height float64
}

func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

func (r Rectangle) Describe() string {
  return fmt.Sprintf("Rectangle %.1fx%.1f", r.Width, r.Height)
}

// Circle
type Circle struct {
  Radius float64
}

func (c Circle) Area() float64 {
  return math.Pi * c.Radius * c.Radius
}

func (c Circle) Describe() string {
  return fmt.Sprintf("Circle r=%.1f", c.Radius)
}

// Triangle
type Triangle struct {
  Base, Height float64
}

func (t Triangle) Area() float64 {
  return 0.5 * t.Base * t.Height
}

func (t Triangle) Describe() string {
  return fmt.Sprintf("Triangle b=%.1f h=%.1f", t.Base, t.Height)
}

// Function accepting interface - polymorphism!
func printShapeInfo(s Shape) {
  fmt.Printf("%s -> Area: %.2f\n", s.Describe(), s.Area())
}

// Function accepting slice of interfaces
func totalArea(shapes []Shape) float64 {
  total := 0.0
  for _, s := range shapes {
    total += s.Area()
  }
  return total
}

func main() {
  fmt.Println("=== Interface Polymorphism ===")

  // Different types, same interface
  rect := Rectangle{Width: 10, Height: 5}
  circle := Circle{Radius: 7}
  triangle := Triangle{Base: 8, Height: 6}

  // Pass different types to same function
  fmt.Println("\nShapes:")
  printShapeInfo(rect)
  printShapeInfo(circle)
  printShapeInfo(triangle)

  // Slice of interfaces
  shapes := []Shape{rect, circle, triangle}

  fmt.Println("\nIterating over shapes:")
  for i, s := range shapes {
    fmt.Printf("  %d. %s\n", i+1, s.Describe())
  }

  fmt.Printf("\nTotal area: %.2f\n", totalArea(shapes))
}
