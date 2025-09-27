package main

import (
	"fmt"
)

func main() {
	var n, i, val int64
	fmt.Print("Введите число: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("ошибка ввода:", err)
		return
	}

	fmt.Print("\nВведите i-ый бит, который нужно будет поменять: ")
	if _, err := fmt.Scan(&i); err != nil {
		fmt.Println("ошибка ввода:", err)
		return
	}

	fmt.Print("\nВведите на что нужно поменять бит (0 или 1): ")
	if _, err := fmt.Scan(&val); err != nil {
		fmt.Println("ошибка ввода:", err)
		return
	}

	if i < 1 || i > 64 {
		fmt.Println("Ошибка: i должно быть от 1 до 64")
		return
	}

	mask := int64(1) << (i - 1)

	if val == 1 {
		n = n | mask
	} else if val == 0 {
		n = n &^ mask
	} else {
		fmt.Println("Ошибка: val должно быть 0 или 1")
		return
	}

	fmt.Println("Результат:", n)
}
