package main

import "fmt"

func main() {
  fmt.Println("=== Empty Interface & Type Assertions ===")

  // Empty interface can hold any type
  var anything interface{}

  anything = "Hello, Go!"
  fmt.Println("String:", anything)

  anything = 42
  fmt.Println("Int:", anything)

  anything = 3.14
  fmt.Println("Float:", anything)

  anything = true
  fmt.Println("Bool:", anything)

  // Type assertion - extract the concrete type
  fmt.Println("\n=== Type Assertions ===")

  var value interface{} = "Hello"

  // Basic assertion (panics if wrong type)
  str := value.(string)
  fmt.Println("Asserted string:", str)

  // Safe assertion with ok pattern
  if s, ok := value.(string); ok {
    fmt.Println("It's a string:", s)
  }

  if i, ok := value.(int); ok {
    fmt.Println("It's an int:", i)
  } else {
    fmt.Println("Not an int!")
  }

  // Type switch - check multiple types
  fmt.Println("\n=== Type Switch ===")

  checkType := func(v interface{}) {
    switch t := v.(type) {
    case string:
      fmt.Printf("String: %q (len=%d)\n", t, len(t))
    case int:
      fmt.Printf("Int: %d\n", t)
    case float64:
      fmt.Printf("Float: %.2f\n", t)
    case bool:
      fmt.Printf("Bool: %v\n", t)
    default:
      fmt.Printf("Unknown type: %T\n", t)
    }
  }

  checkType("Go")
  checkType(100)
  checkType(2.718)
  checkType(false)
  checkType([]int{1, 2, 3})
}
