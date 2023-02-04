package main

import "fmt"

// DECLARE variables and do not assign value
// Scope of these variable is inside package
var a int
var b string
var c bool

// CRATE own type
type myType int

// DECLARE a variable with myType. More detail: https://go.dev/ref/spec#Types
var m myType
var n int

// iota
const (
	year1 = 2023  - iota 
	year2 = 2023 - iota
	year3 = 2023 - iota
	year4 = 2023 - iota
)

func main() {
	// DECLARE variables and assign value
	// Scope of these variables are inside main function. (Local scope)
	x := 42
	y := "Thang"
	z := true

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	fmt.Println("z = ", z)

	// Print format. More detail: https://pkg.go.dev/fmt
	fmt.Printf("x = %d, y = %s, z = %t\n", x, y, z)

	fmt.Printf("Default value of a = %d, b = %s, c = %t\n", a, b, c)
	a = 1
	b = "ThangLQ1"
	c = false

	fmt.Printf("a = %d, b = %s , c = %t\n", a, b, c)

	fmt.Println("Default value of m = ", m)
	fmt.Printf("Type of m is %T\n", m)
	m = 100
	fmt.Println("New value of m = ", m)

	// Casting type: cast myType to int
	n = int(m)
	fmt.Println("n = ", n)

	// DECLARE const variable
	const number = 42
	// Print a number in decimal, binary, hex
	fmt.Printf("%d\t%b\t%#x\n", number, number, number)

	// Raw string literal
	const rawString = `This is
	a raw string
	literal "You can see it "
	right?`
	fmt.Println(rawString)

	// Print iota variable
	fmt.Println(year1, year2, year3, year4)
}
