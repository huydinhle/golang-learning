package main

import (
	"fmt"
	"log"
)

// Square is a compose type that has a list of square
// integers and the horizontal length and vertical length
// of the original matrix
type Square struct {
	squares [][]int
	h, v    int
}

func main() {

	matrix, times := readData()

	result := rorateMatrix(matrix, times)

	printMatrix(result)
}

func printMatrix(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
func readData() ([][]int, int) {
	var row int
	var col int
	var times int

	_, err := fmt.Scanf("%d", &row)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Scanf("%d", &col)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Scanf("%d", &times)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("row = %+v\n", row)
	// fmt.Printf("col = %+v\n", col)
	result := make([][]int, row)
	for i := 0; i < row; i++ {
		result[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// fmt.Printf("i = %+v\n", i)
			// fmt.Printf("j = %+v\n", j)
			fmt.Scanf("%d", &result[i][j])
		}
	}

	return result, times
}

func rorateMatrix(matrix [][]int, times int) [][]int {
	squares := converMatrixToSquare(matrix)
	rotateSquareLefT(squares, times)
	return convertSquaresToMatrix(squares)
}
func convertSquaresToMatrix(square Square) [][]int {
	startingIndex := 0
	h := square.h
	v := square.v
	result := make([][]int, v)
	for i := 0; i < v; i++ {
		result[i] = make([]int, h)
	}
	index := 0
	for h > 0 && v > 0 {
		// fmt.Printf("startingIndex = %+v\n", startingIndex)
		// fmt.Printf("h = %+v\n", h)
		// fmt.Printf("v = %+v\n", v)
		// fmt.Printf("square.squares = %+v\n", square.squares)
		applySquare(startingIndex, h, v, result, square.squares[index])
		h -= 2
		v -= 2
		startingIndex++
		index++
	}
	return result
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
	return square
}

func applySquare(startIndex, h, v int, matrix [][]int, square []int) {
	index := 0
	// first side, from a(i,i) to a(i,i+k)
	for i := 0; i < h; i++ {
		matrix[startIndex][startIndex+i] = square[index]
		index++
	}

	// second side, from a(i,i) to a(i,i+k)
	for i := 1; i < v; i++ {
		matrix[startIndex+i][startIndex+h-1] = square[index]
		index++
	}

	// third side, from a(i,i) to a(i,i+k)
	if v > 1 {
		for i := 1; i < h; i++ {
			matrix[startIndex+v-1][startIndex+h-1-i] = square[index]
			index++
		}
	}

	// fourth side, from a(i,i) to a(i,i+k)
	if h > 1 {
		for i := 1; i < v-1; i++ {
			matrix[startIndex+v-1-i][startIndex] = square[index]
			index++
		}
	}
	// fmt.Printf("result = %+v\n", result)
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

func rotateSquareLefT(square Square, n int) {
	for k := range square.squares {
		square.squares[k] = rotateLeft(square.squares[k], n)
	}
}

func rotateLeft(slice []int, times int) []int {
	times = times % len(slice)
	// leftslice := slice[0 : len(slice)-times]
	leftslice := slice[0:times]
	// fmt.Printf("leftslice = %+v\n", leftslice)
	// rightslice := slice[len(slice)-times:]
	rightslice := slice[times:]
	// fmt.Printf("rightslice = %+v\n", rightslice)
	return append(rightslice, leftslice...)
}
