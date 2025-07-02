package main

import (
	"fmt"
)

func main() {
	indexx()
}
func indexx() {
	var index int
	Slice := []string{"please", "delete", "me!!"}

	fmt.Print("Введите индекс: ")
	fmt.Scanln(&index)

	switch index {
	case (0):
		Slice = append(Slice[:index], Slice[index+1:]...)
		fmt.Println(Slice)
	case (1):
		Slice = append(Slice[:index], Slice[index+1:]...)
		fmt.Println(Slice)
	case (2):
		Slice = append(Slice[:index], Slice[index+1:]...)
		fmt.Println(Slice)
	default:
		fmt.Println("Вы вышли за предел индексов")
	}
}
