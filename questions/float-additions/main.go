package main

import (
	"fmt"
	"math/big"
)

func main() {
	var total float64
	fmt.Println(total + 15.00 + 3.01)
	total2 := new(big.Float).SetFloat64(0.00)
	total2.Add(new(big.Float).SetFloat64(15.00), total2)
	total2.Add(new(big.Float).SetFloat64(3.01), total2)
	fmt.Printf("total2	 = %+v\n", total2)
}
