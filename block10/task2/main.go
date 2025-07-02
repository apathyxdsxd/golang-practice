package main

import (
	"fmt"
)

func leap(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func main() {
	var year int
	fmt.Print("Введите год: ")
	fmt.Scan(&year)

	if leap(year) {
		fmt.Printf("%d год - високосный\n", year)
	} else {
		fmt.Printf("%d год - не високосный\n", year)
	}
}
