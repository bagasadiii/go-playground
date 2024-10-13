package main

import (
	"fmt"
	"math"
)
type calculate interface {
	add()int
	sub()int
	mul()int
	div()int
}
type nums struct {
	x, y int
}
func (n nums) add()int{
	return n.x + n.y
}
func (n nums) sub()int{
	return n.x - n.y
}
func (n nums) mul()int{
	return n.x * n.y
}
func (n nums) div()int{
	return n.x / n.y
}

type shapes interface {
	area() float64
	circumference() float64
}
type circle struct {
	diameter float64
}
type rectangle struct {
	width float64
}
func (r rectangle) area () float64 {
	return math.Pow(r.width, 2)
}
func (r rectangle) circumference () float64 {
	return r.width * 4
}
func (c circle) radius () float64 {
	return c.diameter / 2
}
func (c circle) area () float64 {
	return math.Pi * math.Pow(c.radius(), 2)
}
func (c circle) circumference () float64 {
	return math.Pi * c.diameter
}
func main() {
	number := nums {x:10, y:5}
	var x calculate = number
	fmt.Println(x.add())
	fmt.Println(x.sub())
	fmt.Println(x.mul())
	fmt.Println(x.div())

	//New Interface

	var rect shapes
	var circ shapes

	rect = rectangle{10.00}
	circ = circle{20.00}
	fmt.Printf("Circle area: %.2f\nCircle circumference: %.2f\n", circ.area(), circ.circumference())
	fmt.Printf("Rectangle area: %.2f\nRectangle circumference: %.2f\n", rect.area(), rect.circumference())
}