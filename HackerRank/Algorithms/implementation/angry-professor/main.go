package main

import "fmt"

type TestCase struct {
	times []int
	n     int
	k     int
}

func main() {
	input := readData()
	for _, v := range input {
		check := cancel(v)
		if check {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func cancel(t TestCase) bool {
	count := 0
	for i := 0; i < t.n; i++ {
		if t.times[i] <= 0 {
			count++
		}
	}
	// fmt.Printf("count = %+v\n", count)
	// fmt.Printf("t.k = %+v\n", t.k)
	return count < t.k
}

func readData() []TestCase {
	var count int
	tc := []TestCase{}

	fmt.Scanf("%d", &count)

	for i := 0; i < count; i++ {
		var n, k int
		var slice []int
		fmt.Scanf("%d", &n)
		fmt.Scanf("%d", &k)

		slice = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &slice[i])
		}

		tc = append(tc, TestCase{slice, n, k})
	}
	return tc
}
