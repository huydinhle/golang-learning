package main

import (
	"testing"
)

func TestMinMaxSum(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	min, max := minMaxSum(slice)
	if min != 10 {
		t.Error("min is supposed to be 10, but it is ", min)
	}
	if max != 14 {
		t.Error("min is supposed to be 14, but it is ", max)
	}

}
