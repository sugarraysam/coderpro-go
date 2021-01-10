/*
Max Consecutive Ones - https://leetcode.com/problems/max-consecutive-ones/

The problem also adds the dec -> binary conversion dimension.

Time Complexity:
	[Singly Linked List]
	O(log n), we will divide the decimal number (n) log n times to get the binary representation, and then iterate over this
				slice of len log(n) to find the max consecutive ones

	[Shifting]
	O(log n), we take binary and and right shift the integer n, log n times

Space Complexity:
	[Snigly Linked List]
	O(log n), we store the log n bits required to represent the decimal number n in a singly linked list

	[Shifting]
	O(1), we dont store anything
*/
package cp65

import "strings"

type Bits struct {
	Dec int
	Bin *Node
}

func NewBits(d int) *Bits {
	return &Bits{Dec: d, Bin: toBits(d)}
}

// dec is a positive integer
// storing bits as a singly linked list
func toBits(dec int) *Node {
	var lastRoot *Node // nil
	for dec > 0 {
		curr := NewNode(getByte(dec), lastRoot)
		lastRoot = curr
		dec /= 2
	}
	return lastRoot
}

func getByte(d int) byte {
	if d%2 == 1 {
		return '1'
	}
	return '0'
}

func (b *Bits) FindLargest() int {
	return b.Bin.FindLargest()
}

func (b *Bits) String() (string, error) {
	return b.Bin.String()
}

// Singly Linked List
type Node struct {
	Bit  byte
	Next *Node
}

func NewNode(b byte, n *Node) *Node {
	return &Node{Bit: b, Next: n}
}

func (n *Node) String() (string, error) {
	var b strings.Builder
	if err := stringHelper(n, &b); err != nil {
		return "", err
	}
	return b.String(), nil
}

func stringHelper(root *Node, b *strings.Builder) error {
	if root == nil {
		return nil
	}
	if err := b.WriteByte(root.Bit); err != nil {
		return err
	}
	return stringHelper(root.Next, b)
}

func (n *Node) FindLargest() int {
	res := 0
	sum := 0
	curr := n
	for curr != nil {
		if curr.Bit == '1' {
			sum++
		} else {
			res = max(res, sum)
			sum = 0
		}
		curr = curr.Next
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Shifting and & binary operators solution
func SolveShifting(d int) int {
	res := 0
	sum := 0
	for d > 0 {
		if d&1 == 1 {
			sum++
		} else {
			res = max(res, sum)
			sum = 0
		}
		d >>= 1 // divide by 2
	}
	return max(res, sum)
}
