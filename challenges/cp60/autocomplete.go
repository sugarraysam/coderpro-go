/*
Auto-complete feature using Trie
https://www.geeksforgeeks.org/auto-complete-feature-using-trie/

We are given a Trie with a set of strings stored in it. Now the user types in a prefix of his
search query, we need to give him all recommendations to auto-complete his query based on the
strings stored in the Trie. We assume that the Trie stores past searches by the users.

For example if the Trie store {“abc”, “abcd”, “aa”, “abbbaba”} and the User types in “ab” then
he must be shown {“abc”, “abcd”, “abbbaba”}.

Time Complexity:
	O(2 * n * m + k),	Creating the trie in O(n * m), where n is # of words, and m is len longest word
			  		  	Then we walk the trie for k char in prefix, than visit every child from that point,
			  			so potentially n * m again

Space Complexity:
	O(n * m), we store every char of every word as a node in the tree
				n is # of words
				m is len of longest word
*/

package cp60

import (
	"bytes"
	"fmt"
	"regexp"
)

const NilVal = ' '

// Node methods
type Node struct {
	Char     byte
	IsWord   bool
	Children []*Node
}

func NewNode(char byte, isWord bool) *Node {
	return &Node{Char: char, IsWord: isWord, Children: make([]*Node, 0)}
}

func (n *Node) findOrCreateChild(char byte, isWord bool) *Node {
	child := n.findChild(char)
	if child == nil {
		child = NewNode(char, isWord)
		n.Children = append(n.Children, child)
	}
	return child
}

func (n *Node) findChild(char byte) *Node {
	for _, child := range n.Children {
		if child.Char == char {
			return child
		}
	}
	return nil
}

// Trie method
func NewTrie(words []string) *Node {
	root := NewNode(NilVal, false) // root is empty
	for _, word := range words {
		curr := root.findOrCreateChild(word[0], false)
		lastCharIdx := len(word) - 1
		for i := 1; i < len(word); i++ { // ignore first letter
			curr = curr.findOrCreateChild(word[i], i == lastCharIdx)
		}
	}
	return root
}

func (n *Node) StringPerLevel() (string, error) {
	var b bytes.Buffer
	stack := []*Node{n}

	for len(stack) > 0 {
		levelSize := len(stack)
		for levelSize > 0 {
			// pop
			curr := stack[0]
			stack = stack[1:]

			if _, err := b.WriteString(fmt.Sprintf("%c ", curr.Char)); err != nil {
				return "", err
			}
			stack = append(stack, curr.Children...)
			levelSize--
		}
		if err := b.WriteByte('\n'); err != nil {
			return "", err
		}
	}
	return b.String(), nil
}

func (n *Node) SolveTrie(words []string, prefix string) []string {
	curr := n
	for i := 0; i < len(prefix); i++ {
		curr = curr.findChild(prefix[i])
		if curr == nil {
			return []string{}
		}
	}
	return curr.findWords(prefix)
}

func (n *Node) findWords(word string) []string {
	var res []string
	if n.IsWord {
		res = append(res, word)
	}
	for _, child := range n.Children {
		res = append(res, child.findWords(word+string(child.Char))...)
	}
	return res
}

// Also adding solve regexp for fun N profit
func SolveRegexp(words []string, prefix string) []string {
	var res []string
	re := regexp.MustCompile(fmt.Sprintf(`^%s*`, prefix))
	for _, word := range words {
		if re.MatchString(word) {
			res = append(res, word)
		}
	}
	return res
}
