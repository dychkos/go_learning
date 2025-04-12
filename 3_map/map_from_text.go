package main

import (
	"fmt"
	"strings"
)

// Напиши функцію, яка приймає текст (рядок) і повертає map, де ключі — це слова з тексту, а значення — кількість їх входжень.
func main() {
	result := mapFromText("hello world hello")

	fmt.Println(result)
}

func mapFromText(text string) map[string]int {
	resultMap := map[string]int{}

	words := strings.Fields(text)

	for _, word := range words {
		resultMap[word]++
	}

	return resultMap
}
