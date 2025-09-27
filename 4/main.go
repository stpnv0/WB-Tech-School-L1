package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

// 1 вариант с каналом
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

	wg := &sync.WaitGroup{}
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, ch, wg)
	}

	//создаем сигнальный канал на корректного завершения
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

loop:
	for {
		select {
		case <-signalChan:
			fmt.Println("Получен сигнал, завершаю программу...")
			close(ch)
			break loop
		default:
			ch <- rand.Int()
		}
	}

	wg.Wait()
}

func worker(id int, job <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for res := range job {
		fmt.Printf("Worker %d; result %d\n", id, res)
	}
}

//2 вариант с контекстом
//func main() {
//	if len(os.Args) != 2 {
//		fmt.Println("При запуске программы передайте количество рабочих воркеров")
//		os.Exit(1)
//	}
//	workerCount, err := strconv.Atoi(os.Args[1])
//	if err != nil || workerCount <= 0 {
//		fmt.Println("Некорректное количество воркеров")
//		os.Exit(1)
//	}
//	ch := make(chan int)
//
//	wg := &sync.WaitGroup{}
//	for i := 1; i <= workerCount; i++ {
//		wg.Add(1)
//		go worker(i, ch, wg)
//	}
//	//создаем контекст с отменой на корректного завершения
//	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
//	defer cancel()
//
//loop:
//	for {
//		select {
//		case <-ctx.Done():
//			fmt.Println("Получен сигнал, завершаю программу...")
//			close(ch)
//			break loop
//		case ch <- rand.Int():
//		}
//	}
//
//	wg.Wait()
//}
