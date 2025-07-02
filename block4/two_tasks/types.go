package main

import (
	"fmt"
	"unsafe"
)

var type_ string

func main() {
	//task1. демонстрирование использования всех базовых типов данных с примерами значений
	var number1 int = 12332
	var number2 int8 = 120
	var number3 int16 = 4343
	var number4 int32 = 99999999
	var number5 int64 = 1234567890987654321
	var number6 uint8 = 255
	var number7 uint16 = 4096
	var float32Var float32 = 33332.223
	var float64Var float64 = 433434343434.222332233223
	var boolean bool = true
	var string1 string = "Hello world!"
	var byteVar byte = 255

	//task2. Обновленаня программа, которая показывает размер каждого типа данных в байтах, используя unsafe.Sizeof()
	fmt.Print("Введите тип переменной для вывода (int, int8, int16, int32, int64, uint8, uint16, float32, float64, bool, string, byte): ")
	fmt.Scanln(&type_)

	fmt.Println("\nРезультат:")
	switch type_ {
	case "int":
		fmt.Println(unsafe.Sizeof(number1))
	case "int8":
		fmt.Println(unsafe.Sizeof(number2))
	case "int16":
		fmt.Println(unsafe.Sizeof(number3))
	case "int32":
		fmt.Println(unsafe.Sizeof(number4))
	case "int64":
		fmt.Println(unsafe.Sizeof(number5))
	case "uint8":
		fmt.Println(unsafe.Sizeof(number6))
	case "uint16":
		fmt.Println(unsafe.Sizeof(number7))
	case "float32":
		fmt.Println(unsafe.Sizeof(float32Var))
	case "float64":
		fmt.Println(unsafe.Sizeof(float64Var))
	case "bool":
		fmt.Println(unsafe.Sizeof(boolean))
	case "string":
		fmt.Println(unsafe.Sizeof(string1))
	case "byte":
		fmt.Println(unsafe.Sizeof(byteVar))
	default:
		fmt.Println("Неизвестный тип переменной")
	}
}
