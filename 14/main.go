package main

import (
	"fmt"
	"reflect"
)

func main() {
	//просто захардкодил значения, чтобы проверить четыре типа из условия
	i := 1
	s := "s"
	b := true
	c := make(chan interface{})

	determineType(i)
	determineType(s)
	determineType(b)
	determineType(c)
}

func determineType(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("это string")
	case int:
		fmt.Println("это int")
	case bool:
		fmt.Println("это bool")
	default:
		//проверим на любой тип канала
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			fmt.Println("это chan")
		}
		fmt.Println("неизвестный тип")
	}
}
