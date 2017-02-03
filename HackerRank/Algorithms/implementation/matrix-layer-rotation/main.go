package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix := [][]int{
		{1, 2, 3, 4, 5, 6, 100},
		{7, 8, 9, 10, 11, 12, 200},
		{13, 14, 15, 16, 17, 18, 300},
		{19, 20, 21, 22, 23, 24, 400},
		{25, 26, 27, 28, 29, 30, 500},
		{31, 32, 33, 34, 35, 36, 600}}
	for i := range matrix {
		for j := range matrix {
			fmt.Printf("matrix[%+v][%+v] = %+v\n", i, j, matrix[i][j])
		}
	}
	getSquare(0, 7, 6, matrix)
	// getSquare(1, 4, matrix)
	// getSquare(2, 2, matrix)
	// getSquare(3, 0, matrix)
}

func converMatrixToSquare(matrix [][]int) {
}

func getSquare(startIndex, h, v int, slice [][]int) {

	result := []int{}
	// first side, from a(i,i) to a(i,i+k)
	for i := 0; i < h; i++ {
		result = append(result, slice[startIndex][startIndex+i])
	}
	// fmt.Printf("result = %+v\n", result)

	// second side, from a(i,i) to a(i,i+k)
	for i := 1; i < v; i++ {
		result = append(result, slice[startIndex+i][startIndex+h-1])
	}

	// third side, from a(i,i) to a(i,i+k)
	for i := 1; i < h; i++ {
		result = append(result, slice[startIndex+v-1][startIndex+h-1-i])
	}
	// fmt.Printf("result = %+v\n", result)

	// fourth side, from a(i,i) to a(i,i+k)
	for i := 1; i < v; i++ {
		result = append(result, slice[startIndex+v-1-i][startIndex])
	}
	fmt.Printf("result = %+v\n", result)
}
