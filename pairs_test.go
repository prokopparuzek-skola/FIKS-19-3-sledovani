package main

import (
	"fmt"
	"math/big"
	"testing"
)

type pairsTest struct {
	citisens uint64
	numPairs uint64
}

func TestPairs(t *testing.T) {
	var testInput []pairsTest = []pairsTest{
		{0,	0},
		{1, 0},
		{4, 3},
		{25, 17},
		{8, 5},
		{17, 11},
		{20, 13},
		{5, 3},
		{9, 6},
		{7, 5},
		{13, 9},
	}
	for _, in := range testInput {
		result := pairs(big.NewInt(int64(in.citisens)))
		if result.Uint64() != in.numPairs {
			msg := fmt.Sprintf("Pro %d mělo být %d ne %d", in.citisens, in.numPairs, result)
			t.Error(msg)
		}
	}
}
