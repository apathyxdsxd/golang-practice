package main

import "fmt"

func swapper(b, d *int) {
	*b, *d = *d, *b
}

func main() {
	b := 66
	d := 99
	fmt.Printf("До обмена: b = %d, d = %d\n", b, d)
	swapper(&b, &d)
	fmt.Printf("После обмена: b = %d, d = %d\n", b, d)
}
