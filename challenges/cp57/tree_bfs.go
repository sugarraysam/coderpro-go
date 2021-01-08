/*
Binary Tree Level Order Traversal
https://leetcode.com/problems/binary-tree-level-order-traversal/

Given a binary tree, return the level order traversal of its nodes' values.
(ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],

Time Complexity:
	O(n), we visit all nodes in the tree

Space Complexity:
	O(n), we keep stack of nodes to be visited next, worst case would be n
*/

package cp57

import (
	"fmt"
	"strings"
)

const NilVal = " "

// Node methods
type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

func NewNode(s string) *Node {
	return &Node{Val: s, Left: nil, Right: nil}
}

func (n *Node) String() string {
	return fmt.Sprintf("%s ", n.Val)
}

/* Breadth First Search - print tree per level
a \n
b c \n
d e f g \n
*/
func (n *Node) BFSIterative() (string, error) {
	var res strings.Builder
	stack := []*Node{n}

	for len(stack) > 0 {
		for n := len(stack); n > 0; n-- {
			// pop
			curr := stack[0]
			stack = stack[1:]

			if _, err := res.WriteString(curr.String()); err != nil {
				return "", err
			}

			// append
			if curr.Left != nil {
				stack = append(stack, curr.Left)
			}
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
		}

		if _, err := res.WriteString("\n"); err != nil {
			return "", err
		}
	}
	return res.String(), nil
}

// Tree methods
func NewTreeRecursive(vals []string) *Node {
	return newTreeHelper(vals, 0)
}

func newTreeHelper(vals []string, pos int) *Node {
	if pos >= len(vals) || vals[pos] == NilVal {
		return nil
	}
	root := NewNode(vals[pos])
	root.Left = newTreeHelper(vals, 2*pos+1)  // odd
	root.Right = newTreeHelper(vals, 2*pos+2) // even
	return root
}

func NewTreeIterative(vals []string) *Node {
	root := NewNode(vals[0])
	stack := []*Node{root}

	pos := 1
	for len(stack) > 0 && pos < len(vals) {
		// 2 Nils no pop - disgusting code :(
		if pos+1 < len(vals) && vals[pos] == NilVal && vals[pos+1] == NilVal {
			pos += 2
			continue
		}

		// pop
		curr := stack[0]
		stack = stack[1:]

		// left
		if pos < len(vals) && vals[pos] != NilVal {
			left := NewNode(vals[pos])
			curr.Left = left
			stack = append(stack, left)
		}
		pos++

		// right
		if pos < len(vals) && vals[pos] != NilVal {
			right := NewNode(vals[pos])
			curr.Right = right
			stack = append(stack, right)
		}
		pos++
	}
	return root
}
