package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	fmt.Println("введите предложение для переворачивания: ")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении:", err)
		return
	}
	str = str[:len(str)-1]
	fmt.Printf("перевернутое предложение: %s", reverseSent(str))
}

func reverseSent(s string) string {
	runes := []rune(s)
	// Переворачиваем все символы
	reverseRunes(runes)

	//проходимся по каждому слову и меняем в нем буквы
	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			reverseRunes(runes[start:i])
			start = i + 1
		}
	}
	return string(runes)
}

func reverseRunes(runes []rune) {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}
