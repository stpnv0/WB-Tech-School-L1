package main

import (
	"fmt"
	"math/big"
)

func Multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func Division(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func Add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func main() {
	var a, b big.Int
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Printf("Сумма: %s\n", Add(&a, &b).String())
	fmt.Printf("Умножение: %s\n", Multiply(&a, &b).String())
	fmt.Printf("Разность: %s\n", Sub(&a, &b).String())
	fmt.Printf("Деление: %s\n", Division(&a, &b).String())
}
