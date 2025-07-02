package main

import "fmt"

func main() {
	firstNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	secondNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for num := 0; num < len(firstNums); num++ {

		for i := 0; i < 9; i++ {
			fmt.Printf("%dx%d=%d ", firstNums[num], secondNums[i], firstNums[num]*secondNums[i])
		}
		fmt.Println()
	}
}
