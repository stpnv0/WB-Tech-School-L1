package main

import (
	"fmt"
	"sync"
)

func main() {
	var arr = []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for i := range arr {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			fmt.Printf("%d squared is %d\n", arr[i], arr[i]*arr[i])
		}(i)
	}

	wg.Wait()
}
