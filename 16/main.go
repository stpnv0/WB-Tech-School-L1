package main

import "fmt"

// так как в задании было написано написать функцию quickSort([]int) []int, тогда
// для сортировки необходимо будет создавать массивы, левее и правее pivot
// если бы не было этой формулировки, то правильнее было написать функцию
// quickSort(arr []int, l, r int)
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]
	var left, right []int

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}

	sorted := append(quickSort(left), pivot)
	return append(sorted, quickSort(right)...)
}

func main() {
	arr := []int{1, 4, 2, 25, 7, 5}
	fmt.Println(quickSort(arr))
}
