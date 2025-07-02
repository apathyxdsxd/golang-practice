package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//создаем ридер
	reader := bufio.NewReader(os.Stdin)

	//вводим нашу строку
	var myString string
	var mySubString string
	var solution string

	fmt.Print("Введите строку:")
	myString, _ = reader.ReadString('\n')
	myString = strings.TrimSpace(myString) //удаляем лишние символы

	//какую субстроку находим
	fmt.Print("Введите подстроку:")
	mySubString, _ = reader.ReadString('\n')
	mySubString = strings.TrimSpace(mySubString) //удаляем лишние символы

	//обработка строки
	fmt.Println("длина строки:", len(myString)) //длинастроки
	fmt.Println("количество подстрок:", strings.Contains(myString, mySubString))

	fmt.Print("выберите регистр(lower/upper): ")
	fmt.Scanln(&solution)
	switch solution {
	case "upper":
		fmt.Println(strings.ToUpper(myString))
	case "lower":
		fmt.Println(strings.ToLower(myString))
	}
}
