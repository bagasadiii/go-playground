package main

import "fmt"

func main() {
	var name string = "Bagas Adi Susetyo"
	var age int8 = 24
	var currentYear int16 = 2024
	var isGraduated bool = true
	var gpa float32 = 3.46
	var pi float64 = 3.1415926535897932384626433832795028841971
	fmt.Printf("My name is %s, I'm %d years old.\nThis project is made on %d.\n I'm graduated? %t\n with %.2f of GPA\n", name, age, currentYear, isGraduated, gpa)
	fmt.Printf("The number of 40 first digit of pi is %.40f", pi)
}