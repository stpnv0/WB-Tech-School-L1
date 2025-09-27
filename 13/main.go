package main

import "fmt"

func main() {
	var x, y int

	fmt.Println("Введите первое число x")
	fmt.Scan(&x)
	fmt.Println("Введите второе число y")
	fmt.Scan(&y)

	//использую двойной XOR
	x = x ^ y
	y = x ^ y
	x = x ^ y

	fmt.Printf("Теперь x = %d, y = %d", x, y)
}
