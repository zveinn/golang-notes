package main

import "fmt"

func loops() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	slice := []int{1, 2, 3, 3, 4, 5, 6}
	for i, v := range slice {
		fmt.Println(i, v)
	}
}
