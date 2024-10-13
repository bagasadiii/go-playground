package main

import (
	"fmt"
	
)

func main() {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			fmt.Println("even: ", i)
		} else {
			fmt.Println("odd: ", i)
		}
	}
	days := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	for _, day := range days {
		fmt.Println(day)
		if day == "Saturday" || day ==  "Sunday"{
			fmt.Println("Weekend Time!")
		}
	}
	for i := 0; i <= 10; i++ {
		for j := i; j <= 10; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	arrs := []int{23, 46, 54, 76, 87, 32, 12, 65, 42, 78}
	n := len(arrs)
	fmt.Println(arrs)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arrs[j] > arrs[j+1]{
				arrs[j], arrs [j+1] = arrs[j+1], arrs[j]
			}
		}
	}
	fmt.Println(arrs)
}