package main

import "fmt"

// Stringer interface (from fmt package)
// type Stringer interface {
//     String() string
// }

// Person with custom String method
type Person struct {
  Name string
  Age  int
}

// Implement Stringer interface
func (p Person) String() string {
  return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// Temperature with unit
type Temperature struct {
  Value float64
  Unit  string
}

func (t Temperature) String() string {
  return fmt.Sprintf("%.1f %s", t.Value, t.Unit)
}

// Money with currency
type Money struct {
  Amount   float64
  Currency string
}

func (m Money) String() string {
  symbols := map[string]string{
    "USD": "$",
    "EUR": "E",
    "GBP": "P",
  }
  symbol := symbols[m.Currency]
  if symbol == "" {
    symbol = m.Currency + " "
  }
  return fmt.Sprintf("%s%.2f", symbol, m.Amount)
}

func main() {
  fmt.Println("=== Stringer Interface ===")

  // Person uses custom String()
  alice := Person{Name: "Alice", Age: 30}
  bob := Person{Name: "Bob", Age: 25}

  fmt.Println("Person:", alice)
  fmt.Println("Person:", bob)

  // Temperature
  fmt.Println("\n=== Custom Formatting ===")
  temp1 := Temperature{Value: 25.5, Unit: "C"}
  temp2 := Temperature{Value: 98.6, Unit: "F"}

  fmt.Println("Temperature:", temp1)
  fmt.Println("Temperature:", temp2)

  // Money
  fmt.Println("\n=== Money Formatting ===")
  usd := Money{Amount: 99.99, Currency: "USD"}
  eur := Money{Amount: 85.50, Currency: "EUR"}

  fmt.Println("Price:", usd)
  fmt.Println("Price:", eur)

  // Works with Printf %v
  fmt.Printf("\nUsing Printf: %v and %v\n", alice, temp1)
}
