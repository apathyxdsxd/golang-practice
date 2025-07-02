package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var myString string

	reader := bufio.NewReader(os.Stdin)
	//чтение
	fmt.Print("Введите строку: ")
	myString, _ = reader.ReadString('\n')
	myString = strings.TrimSpace(myString)

	mySuperNewString := strings.Split(myString, " ")
	for _, substring := range mySuperNewString {
		fmt.Println(substring)

	}

}
