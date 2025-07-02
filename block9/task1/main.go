package main

import "fmt"

var tryes int = 0 //попытки

func calc() {
	var x, y int
	var operation int
	fmt.Println("Введите первое число:")
	fmt.Scan(&x)
	fmt.Println("Введите второе число:")
	fmt.Scan(&y)
	fmt.Println("Выберите операцию:\n1.Сложение\n2.Вычитание\n3.Умножение\n4.Деление")
	fmt.Scan(&operation)

	switch operation {
	case 1:
		fmt.Printf("Результат: %d\n", x+y)
	case 2:
		fmt.Printf("Результат: %d\n", x-y)
	case 3:
		fmt.Printf("Результат: %d\n", x*y)
	case 4:
		if y != 0 {
			fmt.Printf("Результат: %.2f\n", float64(x)/float64(y))
		} else {
			fmt.Println("На ноль делить нельзя")
		}
	default:
		fmt.Println("Неверная операция")
	}
}

func main() {
	goodbye := "До скорого!"

	for {
		tryes++
		fmt.Printf("Добро пожаловать! Вы используете калькулятор %d раз\n", tryes)

		var choice int
		fmt.Println("Выберите действие:\n1.Калькулятор\n2.Выход")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			calc()
		case 2:
			fmt.Println(goodbye)
			return
		default:
			fmt.Println("Неверный выбор")
		}
	}
}
