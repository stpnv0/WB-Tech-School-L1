package main

import (
	"fmt"
	"time"
)

func main() {
	var t int
	fmt.Println("Введите время(целое число) в секундах, на которое код должен остановиться: ")
	fmt.Scan(&t)

	fns := []func(time.Duration){
		withTime(TimerSleep),
		withTime(IterationSleep),
		withTime(SelectSleep),
		withTime(TickerSleep),
	}

	for i, fn := range fns {
		fmt.Printf("функция номер %d:\t", i+1)
		fn(time.Duration(t) * time.Second)
	}
}

func TimerSleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func IterationSleep(duration time.Duration) {
	start := time.Now()

	for time.Since(start) < duration {

	}
}

func TickerSleep(duration time.Duration) {
	ticker := time.NewTicker(duration)
	<-ticker.C
	ticker.Stop()
}

func SelectSleep(duration time.Duration) {
	select {
	case <-time.After(duration):

	}
}

func withTime(sleep func(time.Duration)) func(time.Duration) {
	return func(d time.Duration) {
		start := time.Now()
		defer func() {
			end := time.Since(start)
			fmt.Printf("закончилось ожидание. время прошло: %v, ожидалось: %v\n", end, d)
		}()

		sleep(d)
	}
}
