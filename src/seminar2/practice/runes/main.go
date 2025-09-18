package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// type rune = int32

func checkRune(r rune) {
	fmt.Printf("Руна: %c\n", r)
	fmt.Printf("IsLetter: %t\n", unicode.IsLetter(r))
	fmt.Printf("IsDigit: %t\n", unicode.IsDigit(r))
	fmt.Printf("IsSpace: %t\n", unicode.IsSpace(r))
	fmt.Printf("IsLower: %t\n", unicode.IsLower(r))
	fmt.Printf("ToUpper: %c\n", unicode.ToUpper(r))
	fmt.Printf("ToLower: %c\n", unicode.ToLower(r))
	fmt.Println()
}

func utf8Info(s string) {
	fmt.Printf("Строка: %s\n", s)
	fmt.Printf("Длина в байтах: %d\n", len(s))
	fmt.Printf("Длина в рунах: %d\n", utf8.RuneCountInString(s))

	// Декодирование первой руны
	r, size := utf8.DecodeRuneInString(s)
	fmt.Printf("Первая руна: %c, размер: %d байт\n", r, size)

	// Проверка валидности UTF-8
	fmt.Printf("Валидная UTF-8: %t\n", utf8.ValidString(s))
	fmt.Println()
}

func runeToBytes(r rune) []byte {
	buf := make([]byte, utf8.UTFMax)
	n := utf8.EncodeRune(buf, r)
	return buf[:n]
}

func bytesToRune(b []byte) (rune, int) {
	return utf8.DecodeRune(b)
}

func main() {
	s := "Привет"
	fmt.Println(len(s))

	s = "Hello, 世界"
	runes := []rune(s)
	fmt.Printf("Строка: %s\n", s)
	fmt.Printf("Длина в байтах: %d\n", len(s))
	fmt.Printf("Длина в рунах: %d\n", len(runes))
	fmt.Printf("Руны: %v\n", runes)

	s = "Привет"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i]) // Будет выводить ерунду
	}

	fmt.Println()

	s = "Привет"
	for _, r := range s {
		fmt.Printf("%c ", r) // П р и в е т
	}

	fmt.Println()

	// Литералы рун
	r1 := 'A'          // руна 'A'
	r2 := '世'          // руна '世'
	r3 := '\u4e16'     // руна '世' через Unicode-escape
	r4 := '\U0001f600' // руна '😀' (смайлик)

	fmt.Printf("%c %c %c %c\n", r1, r2, r3, r4)

	checkRune(r1)
	checkRune(r2)
	checkRune(r3)
	checkRune(r4)

	utf8Info("Hello")
	utf8Info("世界")
	utf8Info("H€llo")

	s = "café" // "café" с акцентом (4 символа, 5 байт)

	// Неправильно - работа с байтами
	fmt.Println("Байтовый подход:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("Байт %d: %c\n", i, s[i])
	}
	fmt.Printf("Длина в байтах: %d\n", len(s))

	// Правильно - работа с рунами
	fmt.Println("\nРунный подход:")
	for i, r := range s {
		fmt.Printf("Руна %d: %c\n", i, r)
	}
	fmt.Printf("Длина в символах: %d\n", utf8.RuneCountInString(s))

	r := '世'
	bytes := runeToBytes(r)
	fmt.Printf("Руна %c в байтах: %v\n", r, bytes)

	restoredRune, size := bytesToRune(bytes)
	fmt.Printf("Байты %v в руну: %c (размер: %d байт)\n",
		bytes, restoredRune, size)
}
