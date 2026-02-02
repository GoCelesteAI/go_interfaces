# Go Interfaces Tutorial

Code examples from [Go Lesson 17: Interfaces](https://www.youtube.com/watch?v=YOUR_VIDEO_ID).

## Examples

### Interface Basics (`interface_basics.go`)
Learn how to define interfaces and implement them implicitly.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Rectangle implements Shape by having these methods
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

### Interface Polymorphism (`interface_polymorphism.go`)
Use interfaces to write functions that work with multiple types.

```go
func printShapeInfo(s Shape) {
    fmt.Printf("%s -> Area: %.2f\n", s.Describe(), s.Area())
}

// Works with any Shape!
printShapeInfo(Rectangle{10, 5})
printShapeInfo(Circle{7})
printShapeInfo(Triangle{8, 6})
```

### Empty Interface & Type Assertions (`interface_empty.go`)
Work with the empty interface and extract concrete types.

```go
var anything interface{} = "Hello"

// Type assertion
str := anything.(string)

// Safe assertion with ok
if s, ok := anything.(string); ok {
    fmt.Println("It's a string:", s)
}

// Type switch
switch v := anything.(type) {
case string:
    fmt.Println("String:", v)
case int:
    fmt.Println("Int:", v)
}
```

### Stringer Interface (`interface_stringer.go`)
Customize how your types are printed.

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// Now fmt.Println uses your custom format
fmt.Println(Person{"Alice", 30}) // Alice (30 years old)
```

## Running the Examples

```bash
go run interface_basics.go
go run interface_polymorphism.go
go run interface_empty.go
go run interface_stringer.go
```

## Key Takeaways

- **Interfaces** define behavior with method signatures
- **Implicit implementation** - no `implements` keyword needed
- **Empty interface** (`interface{}` or `any`) holds any type
- **Type assertions** extract the concrete type from an interface
- **Stringer** interface customizes how types print

## Subscribe

For more Go tutorials, subscribe to [CelesteAI](https://www.youtube.com/@GoCelesteAI)!
