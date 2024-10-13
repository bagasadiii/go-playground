package main

import (
	"fmt"
	"time"
)
func Fizz(n int){
	for i := 0; i < n; i++ {
		fmt.Printf("%d (from fizz)\n", i)
		if i % 3 == 0 && n != 0{
			fmt.Println("Fizz")
		}
		time.Sleep(2 * time.Second)
	}
}
func Buzz(n int){
	for i := 0; i < n; i++ {
		fmt.Printf("%d (from buzz)\n", i)
		if i % 5 == 0 && n != 0 {
			fmt.Println("Buzz")
		}
		time.Sleep(2 * time.Second)
	}
}
func main(){
	go Fizz(10)
	go Buzz(10)
	time.Sleep(20 * time.Second)
}