package main

import "fmt"

func main() {
	fruits := []string{
		"Banana",
		"Apple",
		"Dragonfruit",
		"Papaya",
		"Lemon",
		"Lime",
		"Pineapple",
	}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
	fruit := fruits[0:2]
	fmt.Println(fruits, len(fruits))
	fmt.Println(fruit, cap(fruit))
	newFruits := append(fruits, "Mango")
	fmt.Println(newFruits)

}