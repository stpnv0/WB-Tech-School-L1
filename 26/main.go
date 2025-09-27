package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Введите строку: ")
	fmt.Scan(&str)

	fmt.Printf("Ответ: %t\n", checker(str))
}

func checker(str string) bool {
	runes := []rune(str)
	m := make(map[string]struct{}, len(runes))
	for _, v := range str {
		if _, ok := m[strings.ToLower(string(v))]; !ok {
			m[strings.ToLower(string(v))] = struct{}{}
		} else {
			return false
		}
	}
	return true
}
