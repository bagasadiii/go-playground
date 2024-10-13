package main

import "fmt"

func main() {
	fruits := [5]string{
		"Banana",
		"Apple",
		"Pear",
		"Rambutans",
		"Lemon",
	}
	planets := [8]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	for _, fruit := range fruits {
		fmt.Printf("%s, ", fruit)
	}
	fmt.Printf("\n")

	for _, planet := range planets {
		fmt.Printf("%s, ", planet)
	}
	fmt.Printf("\n")


	lenFruits := len(fruits)
	lenPlanets := len(planets)

	fmt.Println("Total Fruits length: ", lenFruits)
	fmt.Println("Total Planets length: ", lenPlanets)
	
	outerBelts := planets[4:]
	fmt.Println("Planet outside the asteroid belts", outerBelts)
}