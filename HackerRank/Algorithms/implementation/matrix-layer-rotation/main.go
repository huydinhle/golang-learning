package main

import "fmt"

type Square struct {
	squares [][]int
	h, v    int
}

func main() {
	// fmt.Println("vim-go")
	matrix0 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix1 := [][]int{
		{1, 2, 3, 4, 5, 6, 100},
		{7, 8, 9, 10, 11, 12, 200},
		{13, 14, 15, 16, 17, 18, 300},
		{19, 20, 21, 22, 23, 24, 400},
		{25, 26, 27, 28, 29, 30, 500},
		{31, 32, 33, 34, 35, 36, 600}}
	matrix2 := [][]int{
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 10, 11, 12},
		{13, 14, 15, 16, 17, 18}}
	matrix4 := [][]int{{1, 2, 3, 4}, {4, 5, 6, 7}}
	m1 := converMatrixToSquare(matrix0)
	m2 := converMatrixToSquare(matrix1)
	m3 := converMatrixToSquare(matrix2)
	m4 := converMatrixToSquare(matrix4)
	fmt.Printf("m1 = %+v\n", m1)
	fmt.Printf("m2 = %+v\n", m2)
	fmt.Printf("m3 = %+v\n", m3)
	fmt.Printf("m4 = %+v\n", m4)
}

func convertSquaresToMatrix(squares [][]int) [][]int {

	return nil
}
func converMatrixToSquare(matrix [][]int) Square {
	result := [][]int{}
	startingIndex := 0
	h := len(matrix[0])
	v := len(matrix)
	for h > 0 && v > 0 {
		// fmt.Printf("h = %+v\n", h)
		// fmt.Printf("v = %+v\n", v)
		// fmt.Printf("startingIndex = %+v\n", startingIndex)
		result = append(result, getSquare(startingIndex, h, v, matrix))
		h -= 2
		v -= 2
		startingIndex++
	}
	square := Square{result, len(matrix[0]), len(matrix)}
	return Square
}

func getSquare(startIndex, h, v int, slice [][]int) []int {
	result := []int{}
	// first side, from a(i,i) to a(i,i+k)
	for i := 0; i < h; i++ {
		result = append(result, slice[startIndex][startIndex+i])
	}

	// second side, from a(i,i) to a(i,i+k)
	for i := 1; i < v; i++ {
		result = append(result, slice[startIndex+i][startIndex+h-1])
	}

	// third side, from a(i,i) to a(i,i+k)
	if v > 1 {
		for i := 1; i < h; i++ {
			result = append(result, slice[startIndex+v-1][startIndex+h-1-i])
		}
	}

	// fourth side, from a(i,i) to a(i,i+k)
	if h > 1 {
		for i := 1; i < v-1; i++ {
			result = append(result, slice[startIndex+v-1-i][startIndex])
		}
	}
	// fmt.Printf("result = %+v\n", result)
	return result
}
