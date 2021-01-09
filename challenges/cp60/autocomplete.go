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
	"fmt"
	"regexp"
)

const NilVal = ' '

// Node methods
type Node struct {
	Char     byte
	IsWord   bool
	Children map[byte]*Node
}

func NewNode(char byte) *Node {
	return &Node{Char: char, IsWord: false, Children: make(map[byte]*Node)}
}

func (n *Node) findOrCreateChild(char byte) *Node {
	child := n.findChild(char)
	if child == nil {
		child = NewNode(char)
		n.Children[char] = child
	}
	return child
}

func (n *Node) findChild(char byte) *Node {
	res, ok := n.Children[char]
	if !ok {
		return nil
	}
	return res
}

// Trie method
func NewTrie(words []string) *Node {
	root := NewNode(NilVal) // root is empty
	for _, word := range words {
		curr := root.findOrCreateChild(word[0])
		for i := 1; i < len(word); i++ { // ignore first letter
			curr = curr.findOrCreateChild(word[i])
		}
		curr.IsWord = true
	}
	return root
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
