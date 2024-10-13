package main

import (
	"fmt"
	"math"
	"strings"
)

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y

}
func mul(x, y int) int {
	return x * y
}
func div(x, y int) int {
	return x / y
}
//Multiple values
func circle(d float64)(float64, float64){
	area := math.Pi * math.Pow(d / 2, 2)
	circumference := math.Pi * d
	return area, circumference
}
//Variadic function
func total(nums...int)int{
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func concatStrings(separator string, text...string)string {
	return strings.Join(text, separator)
}
func main() {
	x, y := 10, 5
	var d float64 = 15
	fmt.Println(add(x,y))
	fmt.Println(sub(x,y))
	fmt.Println(mul(x,y))
	fmt.Println(div(x,y))
	area, circumference := circle(d)
	fmt.Printf("Area of circle D: %.2f\n", area)
	fmt.Printf("Circumference of circle D: %.2f",circumference)
	//Closure
	minMax := func (nums []int) (int,int)  {

		min := nums[0]
		max := nums[0]
		for _, num := range nums {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
		return min, max
	}
	nums := []int{9,3,14,5,6,1,2}
	fmt.Println(nums)
	minNum, maxNum := minMax(nums)
	fmt.Println(minNum, maxNum)
	result := total(1,2,3,4,5,6,7,8,9)
	fmt.Println(result)

	word := concatStrings("","B","A","G","A","S")
	sentence := concatStrings(" ","Hello", "world", "Go", "stands", "for", "GOAT")
	fmt.Println(word)
	fmt.Println(sentence)
}