package main

import "fmt"

func main() {
	var firstName string = "Bagas Adi"
	var lastName string = "Susetyo"
	var age int = 23
	var university string = "Hasanuddin University"
	var gpa float32 = 3.46
	var isGraduated bool = true

	fmt.Printf("Hello world, my name is: %s %s\n", firstName, lastName)
	fmt.Printf("As of , I'm %d years old\n", age)
	if isGraduated {
		fmt.Printf("I'm already graduated!\n")
	} else {
		fmt.Printf("I'm not graduated yet\n")
	}
	fmt.Printf("I'm graduated from %s with the GPA of %.2f\n", university, gpa)

	// Or with gopher function!

	address := "Bogor"
	job := "Driver"
	rating := 4.98

	fmt.Printf("I'm currently stay in %s, working as %s with rating %.2f", address, job, rating)

	// Or the simplified version

	friends1, friends2, gf := "Fajar", "Bayu", "Yaya"

	fmt.Printf("I have some friends named: %s, %s and a beautiful girlfriend named: %s", friends1, friends2, gf)
}