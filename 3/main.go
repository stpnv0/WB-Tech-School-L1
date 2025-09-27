package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("При запуске программы передайте количество рабочих воркеров")
		os.Exit(1)
	}
	workerCount, err := strconv.Atoi(os.Args[1])
	if err != nil || workerCount <= 0 {
		fmt.Println("Некорректное количество воркеров")
		os.Exit(1)
	}
	ch := make(chan int)

	for i := 1; i <= workerCount; i++ {
		go worker(i, ch)
	}

	for {
		ch <- rand.Int()
	}
}

func worker(id int, job <-chan int) {
	for res := range job {
		fmt.Printf("Worker %d; result %d\n", id, res)
	}
}
