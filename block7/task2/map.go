package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	wordCount := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите текст:")

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { //энтер
			break
		}

		words := strings.Fields(line)
		for _, word := range words {
			lowerWord := strings.ToLower(word)
			lowerWord = strings.TrimRight(lowerWord, ",.!?;:\"')")
			wordCount[lowerWord]++
		}
	}

	fmt.Println("\nЧастота слов:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
