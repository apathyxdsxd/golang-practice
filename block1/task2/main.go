package main

import (
	"fmt"
	"time"
)

func main() {
	time_data := time.Now()
	formatted_data := time_data.Format("15:04:05, 2006-01-02")
	var name string
	fmt.Print("Введите имя: ")
	fmt.Scanln(&name)
	fmt.Println("Добро пожаловать,", name+".", "\nТекущая дата:", formatted_data)
}
