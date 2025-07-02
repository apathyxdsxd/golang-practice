package main

import (
	"fmt"
	"math"
)

func main() {
	var numberP float32 = 3.1415926
	e := 2.71828
	var rad float32
	var exponent int
	fmt.Print("Программа для нахождения прощади круга/Возведение числа 'е' в нужную степень\nВведите радиус:")
	fmt.Scanln(&rad)
	fmt.Print("Введите степень:")
	fmt.Scanln(&exponent)

	result1 := numberP * rad * rad
	result2 := math.Pow(e, float64(exponent))

	fmt.Println("Площадь круга:", result1)
	fmt.Println("число е:", result2)

}
