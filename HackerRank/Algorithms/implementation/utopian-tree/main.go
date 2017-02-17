package main

import "fmt"

func main() {
	input := readData()
	result := []int{}
	for _, v := range input {
		result = append(result, findHeight(v))
	}

	for _, v := range result {
		fmt.Println(v)
	}
}

func findHeight(cycle int) int {
	height := 1
	for i := 0; i < cycle; i++ {
		height = nextGrow(i, height)
	}
	return height
}

func nextGrow(cycleTime, height int) int {
	if cycleTime%2 == 0 {
		return height * 2
	}
	return height + 1
}

func readData() []int {
	var size int
	var result []int

	fmt.Scanf("%d", &size)
	result = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &result[i])
	}

	return result
}
