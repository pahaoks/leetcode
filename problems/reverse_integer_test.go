package problems

import (
	"math"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	var i int = 123

	r := reverseInteger(i)
	if r != 321 {
		t.Fatal("wrong")
	}

	r = reverseInteger(-123)
	if r != -321 {
		t.Fatal("wrong")
	}

	r = reverseInteger(1534236469)
	if r != 0 {
		t.Fatal("wrong")
	}
}

func reverseInteger(x int) int {
	var last int
	var tmp int = x
	var reversed int

	for tmp != 0 {
		last = tmp % 10
		tmp /= 10
		reversed = reversed*10 + last
	}

	if reversed > math.MaxInt32 ||
		reversed < -math.MaxInt32 {
		return 0
	}

	return reversed
}
