package main

import (
	"fmt"
	"log"
)

func main() {
	slice, err := readData()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("slice = %+v\n", slice)

	a, b, c := plusMinus(slice)
	fmt.Printf("%f\n", a)
	fmt.Printf("%f\n", b)
	fmt.Printf("%f\n", c)
}

func plusMinus(nums []int) (float64, float64, float64) {
	var positive int
	var negative int
	var zero int

	for _, v := range nums {
		if v > 0 {
			positive++
		} else if v < 0 {
			negative++
		} else {
			zero++
		}
	}
	// fmt.Printf("positive = %+v\n", positive)
	// fmt.Printf("negative = %+v\n", negative)
	// fmt.Printf("zero = %+v\n", zero)
	return float64(positive) / float64(len(nums)), float64(negative) / float64(len(nums)), float64(zero) / float64(len(nums))
}

func readData() ([]int, error) {
	var length int

	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return nil, err
	}

	data := make([]int, length)
	for i := range data {
		_, err = fmt.Scanf("%d", &data[i])
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}
