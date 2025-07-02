package main

import "fmt"

func main() {
	var x float64
	var y float64
	fmt.Print("Введите первое число: ")
	fmt.Scan(&x)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&y)

	fmt.Println("Результаты операций:")
	fmt.Printf("Сложение: %.2f + %.2f = %.2f\n", x, y, x+y)
	fmt.Printf("Вычитание: %.2f - %.2f = %.2f\n", x, y, x-y)
	fmt.Printf("Умножение: %.2f * %.2f = %.2f\n", x, y, x*y)
	fmt.Printf("Делениие: %.2f * %.2f = %.2f\n", x, y, x/y)
}
