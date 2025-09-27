package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fakeTime := 3 * time.Second
	wg := &sync.WaitGroup{}

	wg.Add(5)
	//горутина №1
	stop := make(chan struct{})
	go workerChan(stop, wg)

	//горутина №2
	//можно аналогично использовать context.WithCancel
	ctx, cancel := context.WithTimeout(context.Background(), fakeTime)
	defer cancel()
	go workerCtx(ctx, wg)

	//горутина №3
	go workerGoexit(wg)

	//горутина №4
	timer := time.NewTimer(fakeTime)
	go workerWithTimer(timer, wg)

	//горутина №5
	go workerWithCondition(wg)

	time.Sleep(fakeTime)
	close(stop)
	wg.Wait()
}

func workerChan(ch <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Горутина 1 работает")

	<-ch
	fmt.Println("Горутина 1 остановилась через канал")
	return
}

func workerCtx(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Горутина 2 работает")

	<-ctx.Done()
	fmt.Println("Горутина 2 остановилась через контекст")
	return
}

func workerGoexit(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Горутина 3 работает")

	for i := 0; i < 20; i++ {
		//работа чтобы не сразу закрывался
	}

	fmt.Println("Горутина 3 остановилась через runtime.Goexit")
	runtime.Goexit()
}

func workerWithTimer(timer *time.Timer, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Горутина 4 работает")

	<-timer.C
	fmt.Println("Горутина 4 остановилась по таймеру")
	return
}

func workerWithCondition(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Горутина 5 работает")

	for i := 0; i < 20; i++ {
		//типо работа
		if i == 19 {
			fmt.Println("Горутина 5 остановилась по условию")
			return
		}
	}
}
