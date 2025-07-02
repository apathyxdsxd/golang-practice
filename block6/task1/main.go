package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var mySlice = []string{"Hello", "World"}
	fmt.Println("Базовый срез:", mySlice)

	fmt.Println("Добавьте первый элемент:")
	if scanner.Scan() {
		mySlice = append(mySlice, scanner.Text())
	}
	fmt.Println("Добавьте второй элемент:")
	if scanner.Scan() {
		mySlice = append(mySlice, scanner.Text())
	}
	fmt.Println("Добавьте третий элемент:")
	if scanner.Scan() {
		mySlice = append(mySlice, scanner.Text())
	}
}
