package problems

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	r := []int{0, 0, 1}

	moveZeroes(r)

	fmt.Println(r)
}

func moveZeroes(nums []int) {
	pos := 0

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != 0 {
			continue
		}

		pos = i
		for pos < len(nums)-1 || nums[pos] != 0 {
			nums[pos], nums[pos+1] = nums[pos+1], nums[pos]
			pos++
		}
	}
}
