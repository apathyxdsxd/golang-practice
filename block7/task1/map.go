package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var solution int
	var surname string
	var score int
	scanner := bufio.NewScanner(os.Stdin)
	studentMap := map[string]int{
		"Иванов":    4,
		"Петров":    2,
		"Кузнецов":  4,
		"Рыбаков":   5,
		"Николенко": 2,
	}
	fmt.Println("1. Добавить студента\n2. Удалить студента")
	if scanner.Scan() {
		solution, _ = strconv.Atoi(scanner.Text())
	}
	switch solution {
	case 1:
		fmt.Println("Введите фамилию:")
		if scanner.Scan() {
			surname = scanner.Text()
		}

		fmt.Println("Введите оценку:")
		if scanner.Scan() {
			score, _ = strconv.Atoi(scanner.Text())
		}

		studentMap[surname] = score
		fmt.Println("новая карта:", studentMap)

	case 2:
		fmt.Println("Введите фамилию для удаления:")
		if scanner.Scan() {
			surname = scanner.Text()
			delete(studentMap, surname)
			fmt.Println("новая карта:", studentMap)
		}
	}
}
