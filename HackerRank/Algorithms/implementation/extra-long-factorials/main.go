package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)
	result := fact(n)
	fmt.Print(result.String())
}

func fact(n int) big.Int {
	var mul big.Int
	mul.SetInt64(1)
	for i := 1; i <= n; i++ {
		var num big.Int
		num.SetInt64(int64(i))
		mul.Mul(&mul, &num)
	}
	return mul
}
