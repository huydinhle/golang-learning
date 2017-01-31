package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	slice1, slice2 := Scan()
	intSlice1 := ConvStrToIntSlice(slice1)
	intSlice2 := ConvStrToIntSlice(slice2)
	score1, score2 := CompareTripplet(intSlice1, intSlice2)
	fmt.Print(score1, score2)
}

// Comparing two slice on int and return 2 numers
func CompareTripplet(slice1, slice2 []int) (int, int) {
	score1 := 0
	score2 := 0
	for k, _ := range slice1 {
		if slice1[k] == slice2[k] {
			//do nothing
		} else if slice1[k] > slice2[k] {
			score1++
		} else {
			score2++
		}
	}
	return score1, score2
}

// Scan scan input from Stdin
// It will produce 2 slice of string
func Scan() ([]string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	var slice1 []string
	var slice2 []string

	scanner.Scan()
	slice1 = strings.Split(scanner.Text(), " ")
	scanner.Scan()
	slice2 = strings.Split(scanner.Text(), " ")

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return slice1, slice2
}

// ConvStrToIntSlice will convert a slice of string to
// a slice of integers
func ConvStrToIntSlice(strSlice []string) []int {
	result := []int{}
	for _, v := range strSlice {
		num, _ := strconv.ParseInt(v, 10, 64)
		result = append(result, int(num))
	}
	return result
}
