package main

import "fmt"

func main() {
	a, b := count()
	fmt.Println(a)
	fmt.Println(b)
}

func count() (int, int) {
	var tomStart int
	var tomEnd int

	var appleOrigin int
	var orangeOrigin int

	apples := []int{}
	oranges := []int{}

	tomStart, tomEnd, appleOrigin, orangeOrigin, apples, oranges = readData()

	// tomStart = 7
	// tomEnd = 11

	// appleOrigin = 5
	// orangeOrigin = 5

	// apples := []int{-2, 2, 1}
	// oranges := []int{5, -6}

	acount := 0
	ocount := 0
	for _, v := range apples {
		position := appleOrigin + v
		// fmt.Printf("position = %+v\n", position)
		// fmt.Printf("tomStart = %+v\n", tomStart)
		// fmt.Printf("tomEnd = %+v\n", tomEnd)
		if position >= tomStart && position <= tomEnd {
			acount++
		}
	}

	for _, v := range oranges {
		position := orangeOrigin + v
		if position >= tomStart && position <= tomEnd {
			ocount++
		}
	}
	return acount, ocount
}

func readData() (int, int, int, int, []int, []int) {

	var a, b, c, d int
	var size1, size2 int
	var e, f []int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &d)
	fmt.Scanf("%d", &size1)
	fmt.Scanf("%d", &size2)

	e = make([]int, size1)
	f = make([]int, size2)

	for i := range e {
		fmt.Scanf("%d", &e[i])
	}

	for i := range f {
		fmt.Scanf("%d", &f[i])
	}

	return a, b, c, d, e, f
}
