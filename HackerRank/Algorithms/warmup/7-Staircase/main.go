package main

import (
	"fmt"
	"log"
)

func main() {
	size, err := readFile()
	if err != nil {
		log.Fatal(err)
	}
	printStairCase(size)
}

func printStairCase(n int) {
	for i := 1; i <= n; i++ {
		printSpace(i, n)
		for j := 0; j < i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func printSpace(index, n int) {
	for i := 0; i <= n-index-1; i++ {
		fmt.Print(" ")
	}
}

func readFile() (int, error) {
	var size int

	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		return 0, err
	}
	return size, nil
}
