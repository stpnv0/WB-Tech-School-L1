package main

import "fmt"

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1

	for l <= r {
		m := (r + l) / 2
		if arr[m] < target {
			l = m + 1
		} else if arr[m] > target {
			r = m - 1
		} else if arr[m] == target {
			return m
		}
	}

	return -1
}

func main() {
	a := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(binarySearch(a, 7)) // 3
	fmt.Println(binarySearch(a, 2)) // -1
}
