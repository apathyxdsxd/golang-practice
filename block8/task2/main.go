package main

import "fmt"

func main() {
	first()
}

func first() {

	for k := 2; k <= 100; k++ {
		isPrime := true
		for i := 2; i < k; i++ {
			if k%i == 0 {
				isPrime = false
				break
			}

		}
		if isPrime {

			fmt.Println(k, "Число простое -", isPrime)
		}
	}
}

func second() {
	firstNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	secondNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for num := 0; num < len(firstNums); num++ {
		for i := 1; i < 9; i++ {
			fmt.Printf("%dx%d=%d ", firstNums[num], secondNums[i], firstNums[num]*secondNums[i])
		}
		fmt.Println()
	}

}
