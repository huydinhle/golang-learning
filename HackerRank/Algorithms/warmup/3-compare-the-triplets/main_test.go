package main

import (
	"testing"
)

func TestComareTripplets(t *testing.T) {
	score1, score2 := CompareTripplet([]int{5, 6, 7}, []int{3, 6, 10})
	if score1 != 1 {
		t.Error("Expecting 1 but got", score1)
	}
	if score2 != 1 {
		t.Error("Expecting 1 but got", score1)
	}
}
