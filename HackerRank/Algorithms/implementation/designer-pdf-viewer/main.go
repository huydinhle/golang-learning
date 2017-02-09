package main

import "fmt"

func main() {
	// num := calculateHighlightedArea("abc", []int{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5})
	// fmt.Printf("num	 = %+v\n", num)
	slice, str := readData()
	result := calculateHighlightedArea(str, slice)
	fmt.Print(result)
}

func calculateHighlightedArea(str string, nums []int) int {
	m := generateMap(nums)
	height := 0
	for _, v := range str {
		if m[v] > height {
			height = m[v]
		}
	}
	return area(height, len(str))
}
func generateMap(input []int) map[rune]int {
	m := map[rune]int{}
	for k, v := range input {
		key := rune(k + 97)
		m[key] = v
	}
	return m
}

func area(a, b int) int {
	return a * b
}

func readData() ([]int, string) {
	result := make([]int, 26)
	var str string
	for i := 0; i < 26; i++ {
		fmt.Scanf("%d", &result[i])
	}
	fmt.Scanf("%s", &str)
	return result, str
}
