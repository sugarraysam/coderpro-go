/*
Roman to Integer - https://leetcode.com/problems/roman-to-integer/

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII,
which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four
is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it
making four. The same principle applies to the number nine, which is written as IX. There are six instances
where subtraction is used:

    I can be placed before V (5) and X (10) to make 4 and 9.
    X can be placed before L (50) and C (100) to make 40 and 90.
    C can be placed before D (500) and M (1000) to make 400 and 900.

Given a roman numeral, convert it to an integer.

Time Complexity:
	O(n), iterating over the roman literal

Space Complexity:
	O(1), constant, not storing anything
*/
package cp62

var RomanToInt = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// Base methods
type Roman struct {
	Val string
}

func NewRoman(val string) *Roman {
	return &Roman{val}
}

func (r *Roman) toInt(i int) int {
	if i >= r.Len() || i < 0 {
		return -1
	}
	return RomanToInt[r.Val[i]]
}

func (r *Roman) Len() int {
	return len(r.Val)
}

// Solve Methods
func (r *Roman) SolveIterative() int {
	res := 0
	for i := 0; i < r.Len(); {
		val := r.toInt(i)
		nextVal := r.toInt(i + 1)

		if nextVal > val { // nextVal == -1 if outbound
			res += nextVal - val
			i += 2
		} else {
			res += val
			i += 1
		}
	}
	return res
}

func (r *Roman) SolveRecursive() int {
	return r.helper(0)
}

func (r *Roman) helper(pos int) int {
	if pos >= r.Len() {
		return 0
	}
	res := r.toInt(pos)
	nextVal := r.toInt(pos + 1)
	nextPos := pos + 1
	if nextVal > res { // nextVal == -1 if outbound
		res = nextVal - res
		nextPos++
	}
	res += r.helper(nextPos)
	return res
}

func (r *Roman) SolveReversedIteration() int {
	res := 0
	prev := -1
	for i := r.Len() - 1; i >= 0; i-- {
		val := r.toInt(i)
		if val < prev {
			res -= val
		} else {
			res += val
		}
		prev = val
	}
	return res
}
