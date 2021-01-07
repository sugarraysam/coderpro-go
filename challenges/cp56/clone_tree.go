/*
Find a Corresponding Node of a Binary Tree in a Clone of That Tree
https://leetcode.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/

Given two binary trees original and cloned and given a reference to a node target in the original tree.

The cloned tree is a copy of the original tree.

Return a reference to the same node in the cloned tree.

Note that you are not allowed to change any of the two trees or the target node and the answer must be a reference to a node in the cloned tree.

Follow up: Solve the problem if repeated values on the tree are allowed.

Time Complexity:
	O(n), worst case is we traverse the whole tree (n Nodes), because we traverse in sync

Space Complexity:
	O(n), recursive case -> it is not tail recursive, we pile function calls
		  iterative case -> we store a maximum of n nodePairs in a stack
*/

package cp56

import (
	"fmt"
	"strings"
)

const NilVal = -1

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func NewNode(v int) *Node {
	return &Node{Val: v, Left: nil, Right: nil}
}

// NewTree - build a tree recursively, vals are >= 0 , if val == -1, this means nil child
func NewTree(vals []int) *Node {
	return newTreeHelper(vals, 0)
}

// Recursively build a tree
func newTreeHelper(vals []int, i int) *Node {
	if i >= len(vals) || vals[i] == NilVal {
		return nil
	}
	root := NewNode(vals[i])
	root.Left = newTreeHelper(vals, 2*i+1)  // odd for left childs
	root.Right = newTreeHelper(vals, 2*i+2) // even for right childs
	return root
}

// Print tree with InOrder traversal
func (n *Node) String() string {
	var res strings.Builder
	stringInOrder(n, &res)
	return res.String()
}

func stringInOrder(root *Node, acc *strings.Builder) {
	if root == nil {
		return
	}
	stringInOrder(root.Left, acc)
	fmt.Fprintf(acc, "%d -> ", root.Val)
	stringInOrder(root.Right, acc)
}

// returns -1 if node not found
func SolveRecursive(a *Node, b *Node, target int) int {
	if a.Val == target {
		return b.Val
	}
	if a.Left != nil && b.Left != nil {
		found := SolveRecursive(a.Left, b.Left, target)
		if found != NilVal {
			return found
		}
	}
	if a.Right != nil && b.Right != nil {
		found := SolveRecursive(a.Right, b.Right, target)
		if found != NilVal {
			return found
		}
	}
	return NilVal
}

// helper type
type nodePair struct {
	A *Node
	B *Node
}

func SolveIterative(a *Node, b *Node, target int) int {
	stack := []nodePair{{a, b}}
	for len(stack) > 0 {
		np := stack[0]
		stack = stack[1:]
		if np.A.Val == target {
			return np.B.Val
		}
		if np.A.Left != nil && np.B.Left != nil {
			stack = append(stack, nodePair{np.A.Left, np.B.Left})
		}
		if np.A.Right != nil && np.B.Right != nil {
			stack = append(stack, nodePair{np.A.Right, np.B.Right})
		}
	}
	return NilVal
}
