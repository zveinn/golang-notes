package main

import "fmt"

func slices() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original slice: %v, length: %d, capacity: %d\n", numbers, len(numbers), cap(numbers))

	middle := numbers[1:4]
	fmt.Printf("Middle slice: %v, length: %d, capacity: %d\n", middle, len(middle), cap(middle))

	numbers = append(numbers, 6, 7)
	fmt.Printf("After append: %v, length: %d, capacity: %d\n", numbers, len(numbers), cap(numbers))

	numbers[0] = 10
	fmt.Printf("After modification: %v\n", numbers)

	fullSlice := numbers[:cap(numbers)]
	fmt.Printf("Full slice: %v, length: %d, capacity: %d\n", fullSlice, len(fullSlice), cap(fullSlice))

	zeroLenFullCap := numbers[:0:cap(numbers)]
	fmt.Printf("Zero length, full capacity: %v, length: %d, capacity: %d\n", zeroLenFullCap, len(zeroLenFullCap), cap(zeroLenFullCap))

	copy(middle, []int{8, 9, 10})
	fmt.Printf("After copy: %v\n", numbers)

	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Printf("After removing element: %v\n", numbers)

	numbers = append(numbers[:2], append([]int{11}, numbers[2:]...)...)
	fmt.Printf("After inserting element: %v\n", numbers)
}
