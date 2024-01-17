package problems

import (
	"math"
	"testing"
)

func TestMedianOfSortedArrays(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	r := median(a, b)
	if r != 5.5 {
		t.Fatal("Error")
	}

	a = []int{1, 2}
	b = []int{3, 4}

	r = median(a, b)
	if r != 2.5 {
		t.Fatal("Error")
	}
}

func median(a []int, b []int) float64 {
	n1 := len(a)
	n2 := len(b)

	if n1 > n2 {
		return median(b, a)
	}

	n := n1 + n2
	low := 0
	high := n1
	left := (n1 + n2 + 1) / 2

	for low <= high {
		mid1 := (low + high) / 2
		mid2 := left - mid1
		l1, l2, r1, r2 := math.MinInt, math.MinInt, math.MaxInt, math.MaxInt

		if mid1 < n1 {
			r1 = a[mid1]
		}
		if mid2 < n2 {
			r2 = b[mid2]
		}
		if mid1-1 >= 0 {
			l1 = a[mid1-1]
		}
		if mid2-1 >= 0 {
			l2 = b[mid2-1]
		}

		if l1 <= r2 && l2 <= r1 {
			if n%2 == 1 {
				return float64(max(l1, l2))
			} else {
				return float64(max(l1, l2)+min(r1, r2)) / 2
			}
		} else if l1 < r2 {
			low = mid1 + 1
		} else {
			high = mid1 - 1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
