package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	ch := make(chan int, 1)
	var n time.Duration
	fmt.Println("Введите n - количество секунд работы программы: ")
	fmt.Scan(&n)

	timer := time.After(n * time.Second)
	for {
		select {
		case <-timer:
			fmt.Println("Время прошло, завершаю программу...")
			close(ch)
			os.Exit(0)
		case ch <- rand.Int():
		case val := <-ch:
			fmt.Println(val)
		}
	}
}
