package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := make(map[int][]float64)
	for _, val := range arr {
		var temp float64
		if val < 0 {
			temp = math.Ceil(val / 10)
		} else {
			temp = math.Floor(val / 10)
		}
		key := int(temp * 10)
		res[key] = append(res[key], val)
	}
	fmt.Println(res)
}
