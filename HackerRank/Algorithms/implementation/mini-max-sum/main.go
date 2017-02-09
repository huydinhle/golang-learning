package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	slice, error := readData()
	if error != nil {
		log.Panic("no no no")
	}

	min, max := minMaxSum(slice)
	fmt.Print(min, " ")
	fmt.Print(max)
}

func minMaxSum(nums []int) (int, int) {
	sort.Ints(nums)

	total := 0
	for i := 1; i < len(nums)-1; i++ {
		total += nums[i]
	}
	return nums[0] + total, nums[len(nums)-1] + total
}

func readData() ([]int, error) {

	data := make([]int, 5)

	for i := range data {
		_, err := fmt.Scanf("%d", &data[i])
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}
