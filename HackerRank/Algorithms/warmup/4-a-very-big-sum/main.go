package main

import "fmt"

func main() {
	slice, _ := readData()
	fmt.Println(Addition(slice))

}

// Addition get the sum of all numers
func Addition(nums []int) int {
	total := 0

	for _, v := range nums {
		total += v
	}
	return total
}

func readData() ([]int, error) {
	var length int

	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return nil, err
	}

	data := make([]int, length)

	for i := range data {
		_, err := fmt.Scanf("%d", &data[i])
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}
