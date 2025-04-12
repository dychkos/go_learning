package main

import "fmt"

// Напиши функцію, яка приймає слайс рядків і повертає новий слайс,
// що містить тільки унікальні елементи (без дублікатів). Використай map для відстеження вже зустрінутих значень.
func main() {
	textSlice := []string{
		"hello", "serhii", "you", "are", "you", "serhii", "the", "the", "best",
	}

	result := findUnique(textSlice)

	fmt.Println(result, len(result))
}

func findUnique(textSlice []string) []string {
	uniqueMap := make(map[string]bool)
	var result []string

	for _, elem := range textSlice {
		_, exists := uniqueMap[elem]

		if !exists {
			uniqueMap[elem] = true
			result = append(result, elem)
		}
	}

	return result
}
