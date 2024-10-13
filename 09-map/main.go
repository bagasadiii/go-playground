package main

import "fmt"

func main(){
  lists := map[int]string{
    1001: "Alpha",
    1002: "Beta",
    1003: "Charlie",
    1004: "Delta",
	1005: "Epsilon",
  }
  for id, name := range lists{
    fmt.Printf("ID: %d\nName: %s\n", id, name)
  }
  students := map[string][]int{
    "Alex": {78, 82, 82},
    "Bobby": {89, 78, 76},
    "Charles": {82, 68, 89},
    "Debbie": {76, 62, 76},
    "Edgar": {89, 62, 98},
  }
  for name, scores := range students{
    fmt.Printf("Name: %s\nScores: %v\n", name, scores)
  }
}