package tasks

import "strings"

func isVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, r)
}

// CountVowels подсчитывает количество гласных в строке
func CountVowels(str string) int {
	count := 0

	for _, char := range str {
		if isVowel(char) {
			count++
		}
	}

	return count
}
