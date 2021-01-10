/*
Max Consecutive Ones - https://leetcode.com/problems/max-consecutive-ones/

The problem also adds the dec -> binary conversion dimension.

Time Complexity:
	O(log n), we will divide the decimal number (n) log n times to get the binary representation, and then iterate over this
				slice of len log(n) to find the max consecutive ones

Space Complexity:
	O(log n), we store the log n bits required to represent the decimal number n
*/
package cp65

type Bits struct {
	Decimal int
	Binary  []int
}

func NewBits(d int) *Bits {
	return &Bits{Decimal: d, Binary: toBits(d)}
}

// dec is a positive integer
func toBits(dec int) []int {
	res := make([]int, 0)
	for dec > 0 {
		res = append([]int{dec % 2}, res...) // prepend
		dec /= 2
	}
	return res
}

func (b *Bits) FindLargest() int {
	res := 0
	sum := 0
	for pos := 0; pos < len(b.Binary); pos++ {
		if b.Binary[pos] == 1 {
			sum++
		} else {
			res = max(res, sum)
			sum = 0
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
