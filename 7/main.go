package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

func (s *SafeMap) Set(k string, v int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[k] = v
}

func (s *SafeMap) Get(k string) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.m[k]
	return v, ok
}

func (s *SafeMap) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.m)
}

func main() {
	sm := NewSafeMap()
	var n int
	fmt.Println("Введите количество воркеров, которые будут писать в мапу: ")
	fmt.Scan(&n)

	var wg sync.WaitGroup
	wg.Add(n)

	// Конкурентные записи в одну map
	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				key := fmt.Sprintf("g%d-%d", id, i)
				sm.Set(key, i)
			}
		}(i)
	}

	wg.Wait()

	fmt.Printf("len(map) = %d (ожидалось %d)\n", sm.Len(), n*10)
}
