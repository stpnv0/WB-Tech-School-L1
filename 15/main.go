package main

import "strings"

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	var b strings.Builder
	cnt := 0
	for _, r := range v {
		if cnt == 100 {
			break
		}
		b.WriteRune(r)
		cnt++
	}
	justString = b.String()
}

func main() {
	someFunc()
}
