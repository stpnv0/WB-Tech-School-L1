package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var n int
	ch1, ch2 := make(chan int), make(chan int)
	wg := &sync.WaitGroup{}

	fmt.Scan(&n)

	arr := make([]int, n)

	for i := range arr {
		arr[i] = rand.Int()
	}

	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := range arr {
			ch1 <- arr[i]
		}
		close(ch1)
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for x := range ch1 {
			ch2 <- x * 2
		}
		close(ch2)
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for el := range ch2 {
			fmt.Println(el)
		}
	}(wg)

	wg.Wait()
}
