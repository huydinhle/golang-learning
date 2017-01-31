package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	Scan()
}

// Scan scan input from Stdin
// It will produce 2 slice of string
func Scan() ([]string []string){
	scanner := bufio.NewScanner(os.Stdin)
	var slice1 []string
	var slice2 []string

	scanner.Scan()
	slice1 = strings.Split(scanner.Text(), " ")
	scanner.Scan()
	slice2 = strings.Split(scanner.Text(), " ")
	fmt.Printf("%v\n", slice1)
	fmt.Printf("%v\n", slice2)

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
		append(result, strconv.ParseInt(v, 10, 64))
	}
	return result
}
