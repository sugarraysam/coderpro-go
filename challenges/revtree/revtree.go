/*
Given a binary tree, create a reversed binary tree.

Given:
	  0
	/   \
   1    2
  / \  / \
 3  4 5  6


 Return:

 3  4 5  6
 \  / \ /
  1    2
   \  /
    0
*/
package revtree

import (
	"fmt"
	"strings"
)

// Normal tree
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func NewNode(v int) *Node {
	return &Node{Val: v, Left: nil, Right: nil}
}

func NewTree(vals []int) *Node {
	return newTreeHelper(vals, 0)
}

func newTreeHelper(vals []int, i int) *Node {
	if i >= len(vals) {
		return nil
	}
	root := NewNode(vals[i])
	root.Left = newTreeHelper(vals, 2*i+1)  // odd == left
	root.Right = newTreeHelper(vals, 2*i+2) // even == right
	return root
}

// Reversed tree
type RNode struct {
	Val   int
	Child *RNode
}

func NewRNode(v int, c *RNode) *RNode {
	return &RNode{Val: v, Child: c}
}

type RTree struct {
	Roots []*RNode
}

func (rt *RTree) appendRoot(rn *RNode) {
	rt.Roots = append(rt.Roots, rn)
}

// Invert
func Invert(root *Node) *RTree {
	res := &RTree{}
	res.invertHelper(root, nil)
	return res
}

func (rt *RTree) invertHelper(curr *Node, prev *RNode) {
	rn := NewRNode(curr.Val, prev)
	if curr.Left == nil && curr.Right == nil {
		rt.appendRoot(rn)
	}
	if curr.Left != nil {
		rt.invertHelper(curr.Left, rn)
	}
	if curr.Right != nil {
		rt.invertHelper(curr.Right, rn)
	}
}

// Print
func (rt *RTree) DFSString() (string, error) {
	var res strings.Builder
	var stack []*RNode
	inStack := make(map[int]bool) // works only if tree has unique values
	for _, rn := range rt.Roots {
		inStack[rn.Val] = true
		stack = append(stack, rn)
	}
	for len(stack) > 0 {
		levelSize := len(stack)
		for levelSize > 0 {
			// pop
			rn := stack[0]
			stack = stack[1:]
			if _, err := res.WriteString(fmt.Sprintf("%d ", rn.Val)); err != nil {
				return "", err
			}
			if rn.Child != nil && !inStack[rn.Child.Val] {
				inStack[rn.Child.Val] = true
				stack = append(stack, rn.Child)
			}
			levelSize--
		}
		if err := res.WriteByte('\n'); err != nil {
			return "", err
		}
	}
	return res.String(), nil
}
