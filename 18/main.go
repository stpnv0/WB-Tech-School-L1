package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu  sync.Mutex
	val int
}

func (c *Counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func (c *Counter) decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val--
}

func main() {
	numOfGoroutines := 1000
	counter := Counter{val: 0}
	wg := &sync.WaitGroup{}
	for i := 0; i < numOfGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.increment()
		}()
	}

	wg.Wait()
	fmt.Println(counter.val) //1000
}
