package main

import (
	"fmt"
	"strings"
)

func filter(words ...string)string{
	for _, word := range words {
		lowerWord := strings.ToLower(word)
		switch lowerWord {
		case "fuck", "shit":
			return "..."
		default:
			return word
		}
	}
	return ""
}
func wordFilter(text string, filter func(...string) string) string{
	words := strings.Split(text, " ")
	for i, word := range words{
		words[i] = filter(word)
	}
	return strings.Join(words, " ")
}
func name(n string, say func (string)){
	say(n)
}
func sayHello(n string){
	fmt.Printf("Hello %s!\n", n)
}
func sayBye(n string){
	fmt.Printf("Bye, %s\n", n)
}
func main() {
	name("Bagas", sayHello)
	comment1 := "I'm new to golang"
	comment2 := "Golang is shit kinda fuck"
	result1 := wordFilter(comment1, filter)
	result2 := wordFilter(comment2, filter)
	fmt.Println(result1)
	fmt.Println(result2)
	name("Bagas", sayBye)
}