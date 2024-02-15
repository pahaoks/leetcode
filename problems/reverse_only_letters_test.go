package problems

import "testing"

func TestReverseOnlyLetters(t *testing.T) {
	r := reverseOnlyLetters("aBc-xYz")

	if r != "zYx-cBa" {
		t.Fatal("wrong")
	}

	r = reverseOnlyLetters("a-Bc-x-Yz")

	if r != "z-Yx-c-Ba" {
		t.Fatal("wrong")
	}
}

func reverseOnlyLetters(s string) string {
	b := []byte(s)
	left := 0
	right := len(s) - 1
	isLeftLetter := false
	isRightLetter := false

	for left < right {
		isLeftLetter = (b[left] >= 65 && b[left] <= 90) || (b[left] >= 97 && b[left] <= 122)
		isRightLetter = (b[right] >= 65 && b[right] <= 90) || (b[right] >= 97 && b[right] <= 122)

		if !isLeftLetter {
			left++
			continue
		}

		if !isRightLetter {
			right--
			continue
		}

		b[left], b[right] = b[right], b[left]
		left++
		right--
	}

	return string(b)
}
