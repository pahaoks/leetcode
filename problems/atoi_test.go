package problems

import (
	"math"
	"testing"
)

func TestAtoi(t *testing.T) {
	s := "42"

	r := myAtoi(s)
	if r != 42 {
		t.Fatal("wrong")
	}

	r = myAtoi("-42")
	if r != -42 {
		t.Fatal("wrong")
	}

	r = myAtoi("    -42")
	if r != -42 {
		t.Fatal("wrong")
	}

	r = myAtoi("-42 sfsd fsd")
	if r != -42 {
		t.Fatal("wrong")
	}

	r = myAtoi("0000042 sfsd fsd")
	if r != 42 {
		t.Fatal("wrong")
	}

	r = myAtoi("words and 987")
	if r != 0 {
		t.Fatal("wrong")
	}

	r = myAtoi("+-42")
	if r != 0 {
		t.Fatal("wrong")
	}

	r = myAtoi("00000-42a1234")
	if r != 0 {
		t.Fatal("wrong")
	}

	r = myAtoi("9223372036854775808")
	if r != 2147483647 {
		t.Fatal("wrong")
	}

	r = myAtoi("    -88827   5655  U")
	if r != -88827 {
		t.Fatal("wrong")
	}

	r = myAtoi("  +  413")
	if r != 0 {
		t.Fatal("wrong")
	}

	r = myAtoi("18446744073709551617")
	if r != 2147483647 {
		t.Fatal("wrong")
	}

	r = myAtoi(" ")
	if r != 0 {
		t.Fatal("wrong")
	}
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	var i, res, l int
	var sign = 1
	l = len(s)

	for i < l && s[i] == byte(' ') {
		i++
	}

	if i >= l {
		return res
	}

	if i < l && s[i] == byte('-') || s[i] == byte('+') {
		if s[i] == byte('-') {
			sign = -1
		}
		if s[i] == byte('+') {
			sign = 1
		}
		i++
	}

	for i < l {
		if s[i] >= 48 && s[i] <= 57 {
			res = res*10 + int(uint8(s[i])-48)
		} else {
			return res * sign
		}

		if res > math.MaxInt32 {
			if sign > 0 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		i++
	}

	return res * sign
}

/*
if(s.empty())return 0; //Base Case
        long i=0, sign=1, sum=0, res=0;
        while(i<s.length() && s[i]==' ')++i; //Remove Blank Space
        if(s[i] == '-' || s[i]=='+'){
            s[i++]=='-' ? sign = -1 : sign = 1; //handle Pos/Neg
        }

        while(i<s.length()){
            //if current value is digit then add it, else exit;
            if(s[i]>=48 && s[i]<=57) res = res *10 + (s[i++]-'0');
            else return res*sign;

            //Handle overflow
            if(res > INT_MAX)return sign== -1 ? INT_MIN : INT_MAX;

        }
        return res*sign;
*/
