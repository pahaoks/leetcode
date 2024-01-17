package problems

import "testing"

func TestLongestPalindrome(t *testing.T) {
	r := longestPalindrome("abracadabrarba")

	if len(r) != 7 {
		t.Fatal("wring")
	}

	r = longestPalindrome("a")

	if len(r) != 1 {
		t.Fatal("wring")
	}

	r = longestPalindrome("bb")

	if len(r) != 2 {
		t.Fatal("wring")
	}

	r = longestPalindrome("bbb")

	if len(r) != 3 {
		t.Fatal("wring")
	}

	r = longestPalindrome("bbbsss")

	if len(r) != 3 {
		t.Fatal("wring")
	}

	r = longestPalindrome("abccba")

	if len(r) != 6 {
		t.Fatal("wring")
	}
}

func TestBB(t *testing.T) {
	r := longestPalindrome("bbb")

	if len(r) != 3 {
		t.Fatal("wring")
	}
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	minLeft := 0
	maxRight := 0

	for i := range s {
		left := i
		right := i + 1
		matchRight := 0
		matchLeft := 0

		for left >= 0 && right < len(s) && s[left] == s[right] {
			matchLeft = left
			matchRight = right
			left--
			right++
		}

		if maxRight-minLeft < matchRight-matchLeft {
			minLeft = matchLeft
			maxRight = matchRight
		}

		left = i - 1
		right = i + 1

		for left >= 0 && right < len(s) && s[left] == s[right] {
			matchLeft = left
			matchRight = right
			left--
			right++
		}

		if maxRight-minLeft < matchRight-matchLeft {
			minLeft = matchLeft
			maxRight = matchRight
		}
	}

	res := make([]byte, 0, maxRight-minLeft)
	for i := minLeft; i <= maxRight; i++ {
		res = append(res, s[i])
	}

	return string(res)
}
