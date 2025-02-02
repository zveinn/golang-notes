package main

import (
	"fmt"
	"unsafe"
)

func main() {
	zero_aloc_maps()
	slices_as_reference()
}

func zero_aloc_maps() {
	fmt.Println("")

	zero := make(map[int]struct{})
	zero[1] = struct{}{}

	fmt.Println("zero size:", unsafe.Sizeof(zero[1]), " bytes")

	alloc := make(map[int]int32)
	alloc[1] = 1

	fmt.Println("alloc size:", unsafe.Sizeof(alloc[1]), " bytes")

	pointer := make(map[int]*int)
	var one int = 1
	pointer[1] = &one

	fmt.Println("pointer size:", unsafe.Sizeof(pointer[1]), " bytes")
}

func slices_as_reference() {
	fmt.Println("")

	slice1 := make([]int, 4, 20)
	failureToModifySlice(slice1)

	fmt.Printf("Failure after return: %p %v \n", slice1, slice1)

	fmt.Println("")

	slice2 := make([]int, 4, 20)
	modifySlice(&slice2)

	fmt.Printf("mod after return: %p %v \n", slice2, slice2)

	fmt.Println("")

	slice3 := make([]int, 4)
	modifySingleEntry(slice3)

	fmt.Printf("mod(single) after return: %p %v \n", slice3, slice3)
}

func failureToModifySlice(s []int) {
	fmt.Printf("Failure before append: %p %v \n", s, s)
	s = append(s, 777)
	fmt.Printf("Failure after append: %p %v \n", s, s)
}

func modifySlice(s *[]int) {
	fmt.Printf("mod before adding: %p %v \n", s, s)
	*s = append(*s, 777)
	fmt.Printf("mod after adding: %p %v \n", s, s)
}

func modifySingleEntry(s []int) {
	fmt.Printf("mod(single) before adding: %p %v \n", s, s)
	s[0] = 1
	fmt.Printf("mod(single) after adding: %p %v \n", s, s)
}
