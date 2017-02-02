package main

import (
	"fmt"
	// "log"
)

func main() {
	// slice := rotateRight([]int{3, 7, 8, 9, 1, 2, 0}, 18)
	// fmt.Printf("slice = %+v\n", slice)
	_, rotationNums, _, slice, indexes := processData()
	slice = rotateRight(slice, rotationNums)
	for _, k := range indexes {
		fmt.Println(slice[k])
	}
}

func rotateRight(slice []int, times int) []int {
	times = times % len(slice)
	leftslice := slice[0 : len(slice)-times]
	// fmt.Printf("leftslice = %+v\n", leftslice)
	rightslice := slice[len(slice)-times:]
	// fmt.Printf("rightslice = %+v\n", rightslice)
	return append(rightslice, leftslice...)
}

func processData() (int, int, int, []int, []int) {
	var size int
	var rotationNums int
	var elementNums int
	var slice []int
	var indexes []int

	fmt.Scanf("%d", &size)
	fmt.Scanf("%d", &rotationNums)
	fmt.Scanf("%d", &elementNums)

	slice = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &slice[i])
	}

	indexes = make([]int, elementNums)
	for i := 0; i < elementNums; i++ {
		fmt.Scanf("%d", &indexes[i])
	}
	//print debuggin
	// fmt.Printf("size = %+v\n", size)
	// fmt.Printf("rotationNums = %+v\n", rotationNums)
	// fmt.Printf("elementNums = %+v\n", elementNums)
	// fmt.Printf("slice = %+v\n", slice)
	// fmt.Printf("indexes = %+v\n", indexes)
	return size, rotationNums, elementNums, slice, indexes
}
