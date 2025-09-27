package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	var res1 []int
	var res []int

	m := make(map[int]struct{})
	for _, val := range a {
		m[val] = struct{}{}
	}

	for _, val := range b {
		if _, ok := m[val]; ok {
			res1 = append(res1, val)
		}
	}

	fmt.Println(res1)

	//решение 2 через два указателя, оптимальнее по памяти из-за отсутствия мапы
	sort.Ints(a)
	sort.Ints(b)

	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			res = append(res, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	fmt.Println(res)
}
