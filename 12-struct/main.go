package main
import "fmt"

type Students struct{
  name string
  grade []int
  subject map[int]string
}
func (s Students) avgGrade(grades []int)int{
  n := len(grades)
  total := 0
  for _, grade := range grades {
    total += grade
  }
  avg := total / n
  return avg
}
func main(){
  student := Students {
    name: "John Harris",
    grade: []int{92, 82, 82},
    subject: map[int]string{
      101: "History",
      102: "Math",
      103: "English",
    },
  }
  fmt.Printf("Student: %s\n", student.name)
  fmt.Printf("Grade: %v\n", student.grade)
  fmt.Printf("Subject: %s, %s, %s\n", student.subject[101],student.subject[102], student.subject[103])
  
  avg := student.avgGrade(student.grade)
  fmt.Printf("Average grades: %d\n", avg)
}