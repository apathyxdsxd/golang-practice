package main

import "fmt"

type Shape interface {
	Area() float64
}
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	//прямоугольник
	rect := Rectangle{width: 5, height: 10}
	fmt.Println("площадь прямоугольника:", rect.Area())

	//круг
	circle := Circle{radius: 7}
	fmt.Println("площадь круга:", circle.Area())

	//можно ъранить разные формы в интерфейсе
	var shape Shape
	shape = rect
	fmt.Println("площадь через интерфейс:", shape.Area())

	shape = circle
	fmt.Println("Площадь через интерфейс:", shape.Area())
}
