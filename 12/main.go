package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	var res []string
	set := make(map[string]struct{})
	for _, w := range arr {
		if _, ok := set[w]; !ok {
			set[w] = struct{}{}
			res = append(res, w)
		}
	}

	fmt.Println(res)
}
