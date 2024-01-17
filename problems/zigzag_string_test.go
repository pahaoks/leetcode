package problems

import (
	"testing"
)

func TestZigzagString(t *testing.T) {
	s := "PAYPALISHIRING"

	r := convert(s, 3)
	if r != "PAHNAPLSIIGYIR" {
		t.Fatal("wrong")
	}

	r = convert(s, 4)
	if r != "PINALSIGYAHRPI" {
		t.Fatal("wrong")
	}

	r = convert("AB", 1)
	if r != "AB" {
		t.Fatal("wrong")
	}
}

func convert(s string, levels int) string {
	if levels == 1 {
		return s
	}

	res := make([][]byte, levels)
	level := 0
	zlevel := levels - 2

	for i := 0; i < len(s); i++ {
		if level >= 0 {
			res[level] = append(res[level], s[i])
		} else {
			res[zlevel] = append(res[zlevel], s[i])
			zlevel--
		}

		level++

		if level == levels {
			zlevel = levels - 2
			level = -zlevel
		}
	}

	result := ""
	for _, a := range res {
		result += string(a)
	}

	return result
}
