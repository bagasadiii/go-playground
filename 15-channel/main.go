package main

import (
	"fmt"
	"sync"
	"time"
)

func sendData(ch chan string){
	students := map[int]string{
		1: "Alice",
		2: "Bobby",
		3: "Chaplin",
		4: "Dennis",
		5: "Edward",
		6: "Fanny",
	}
	for _, student := range students {
		fmt.Println("Sending student data...")
		ch <- student
		time.Sleep(2 * time.Second)
	}
	close(ch)
}
func receiveData(ch chan string, wg *sync.WaitGroup){
	defer wg.Done()
	
	for data := range ch {
		fmt.Println("Received: ",data)
	}
}
func main() {
	ch:= make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go sendData(ch)
	go receiveData(ch, &wg)
	wg.Wait()
}