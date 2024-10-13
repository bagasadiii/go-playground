package main

import "fmt"

func main() {
	x := 10
	y := 20
	fmt.Printf("x = %d y = %d\n", x, y)
	add := x + y
	sub := x - y
	mul := x * y
	div := x / y
	mod := x % y
	fmt.Printf("10 + 20 = %d\n10 - 20 = %d\n10 x 20 = %d\n10 / 20 = %d\n10 %% 20 = %d\n", add, sub, mul, div, mod)
	
	// Comparison operator

	fmt.Println("x == y ?", x == y)
	fmt.Println("x != y ?", x != y)
	fmt.Println("x <= y ?", x <= y)
	fmt.Println("x >= y ?", x >= y)
	fmt.Println("x > y ?", x > y)
	fmt.Println("x < y ?", x < y)

}