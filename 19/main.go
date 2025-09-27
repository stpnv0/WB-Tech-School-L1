package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Println("введите строку для переворачивания: ")
	fmt.Scan(&str)
	runes := []rune(str)
	reverseRunes(runes)
	fmt.Printf("перевернутая строка: %s\n", string(runes))
}

func reverseRunes(s []rune) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
