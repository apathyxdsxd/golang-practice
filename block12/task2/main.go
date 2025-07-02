package main

import "fmt"

func long(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	longest := strings[0]

	for _, s := range strings[1:] {
		if len([]rune(s)) > len([]rune(longest)) {
			longest = s
		}
	}

	return longest
}

func main() {
	fmt.Println(long("мммммммммммммм", "солянка", "банановаякожура"))
	fmt.Println(long("один", "два", "три"))
	fmt.Println(long(""))
}
