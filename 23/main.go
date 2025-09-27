package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var lenArr = rand.Intn(100)
	if lenArr == 0 {
		lenArr = 1
	}

	arr := make([]int, lenArr)
	for i := 0; i < lenArr; i++ {
		arr[i] = rand.Intn(100)
	}

	index := rand.Intn(len(arr))
	fmt.Printf("Удаляем элемент с индексом %d\n", index)
	fmt.Printf("Исходный слайс: %v\n", arr)

	copy(arr[index:], arr[index+1:])
	arr = arr[:len(arr)-1]

	fmt.Printf("После удаления: %v\n", arr)
}
