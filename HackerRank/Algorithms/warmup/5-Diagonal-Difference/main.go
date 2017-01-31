package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	size, matrix := Scan()
	// fmt.Printf("size = %+v\n", size)
	// fmt.Printf("matrix = %+v\n", matrix)
	fmt.Println(DiagonalDifference(matrix, size))
}

func DiagonalDifference(matrix [][]int, size int) int {
	diag1, diag2 := 0, 0
	for i := 0; i < size; i++ {
		diag1 = matrix[i][i]
	}
	for i := 0; i < size; i++ {
		diag2 = matrix[size-1-i][i]
	}
	fmt.Printf("diag1 = %+v\n", diag1)
	fmt.Printf("diag2 = %+v\n", diag2)
	return int(math.Abs(float64(diag1 - diag2)))
}

// Scan just scan a bunch of numbers
func Scan() (int, [][]int) {
	scanner := bufio.NewScanner(os.Stdin)
	result := [][]int{}

	scanner.Scan()
	sizeString := scanner.Text()
	size, _ := strconv.ParseInt(sizeString, 10, 64)

	for scanner.Scan() {
		slice := strings.Split(scanner.Text(), " ")
		result = append(result, ConvStrToSlice(slice))
	}
	return int(size), result
}

// ConvStrToIntSlice will convert a slice of string to
// a slice of integers
func ConvStrToSlice(strSlice []string) []int {
	result := []int{}
	for _, v := range strSlice {
		num, _ := strconv.ParseInt(v, 10, 64)
		result = append(result, int(num))
	}
	return result
}
