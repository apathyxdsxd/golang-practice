package main

import (
	"fmt"
	"sort"
	"strings"
)

func contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

func filter(slice []string, condition func(string) bool) []string {
	var result []string
	for _, item := range slice {
		if condition(item) {
			result = append(result, item)
		}
	}
	return result
}

func sortStrings(slice []string) []string {
	sorted := make([]string, len(slice))
	copy(sorted, slice)
	sort.Strings(sorted)
	return sorted
}

func sortStringsDesc(slice []string) []string {
	sorted := sortStrings(slice)
	sort.Sort(sort.Reverse(sort.StringSlice(sorted)))
	return sorted
}

func main() {
	foods := []string{"Молоко", "Хлеб", "Каша", "Какао", "Арбуз", "Кокос"}

	fmt.Println("Содержит 'Хлеб':", contains(foods, "Хлеб"))
	fmt.Println("Содержит 'Дыня':", contains(foods, "Дыня"))

	filtered := filter(foods, func(item string) bool {
		return strings.HasPrefix(item, "К")
	})
	fmt.Println("Продукты на 'К':", filtered)

	sortedAsc := sortStrings(foods)
	fmt.Println("По возрастанию:", sortedAsc)

	sortedDesc := sortStringsDesc(foods)
	fmt.Println("По убыванию:", sortedDesc)

	fmt.Println("Оригинальный список:", foods)
}
