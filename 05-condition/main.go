package main

import "fmt"

func main() {
	minGPA := 3.00
	name := "Bagas Adi Susetyo"
	urGPA := 3.46
	if urGPA > minGPA {
		fmt.Printf("Congratulations! %s is graduated!\n", name)
	} else {
		fmt.Println("You failed")
	}
	switch {
	case (urGPA >= 3.00) && (urGPA < 3.50	):
		fmt.Printf("%s graduated with grade A and GPA of %.2f\n", name, urGPA)
	case (urGPA >= 3.50) && (urGPA <= 3.80):
		fmt.Printf("%s graduated with grade A+ and GPA of %.2f\n", name, urGPA)
	case (urGPA >= 4.00):
		fmt.Println("You're genius!")
	}

	goGrade := 4

	switch {
	case goGrade > 8:
		fmt.Println("You are perfect!")
	case goGrade >= 3 && goGrade <= 7:
		fmt.Println("Nice one")
		fallthrough
	case goGrade <= 2:
		fmt.Println("You need to learn more")
	default:
		fmt.Println("You can learn if you want to")
	}
}