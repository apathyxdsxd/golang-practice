package main

import "fmt"

func main() {
	x := 10
	fmt.Println("исходное значение x:", x)
	changeByValue(x)
	fmt.Println("значение x после changeByValue:", x)

	y := 20
	fmt.Println("\nисходное значение y:", y)
	changeByPointer(&y)
	fmt.Println("значение y после changeByPointer:", y)
}
func changeByValue(val int) {
	val = 50
	fmt.Println("значение val внутри changeByValue:", val)
}
func changeByPointer(ptr *int) {
	*ptr = 80
	fmt.Println("значение *ptr внутри changeByPointer:", *ptr)
}
