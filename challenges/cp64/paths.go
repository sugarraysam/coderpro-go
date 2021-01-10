/*
Simplify Path - https://leetcode.com/problems/simplify-path/

Given an absolute path for a file (Unix-style), simplify it. Or in other words, convert it to the canonical path.
In a UNIX-style file system, a period '.' refers to the current directory. Furthermore, a double period '..' moves
the directory up a level.

Note that the returned canonical path must always begin with a slash '/', and there must be only a single slash '/'
between two directory names. The last directory name (if it exists) must not end with a trailing '/'. Also, the
canonical path must be the shortest string representing the absolute path.

Time Complexity:
	O(n), where n is the number of chars, we scan over all the chars multiple times, to split and replace "//" -> "/"

Space Complexity:
	O(m), where m is the # of parts in the raw string, we store them both in a stack and as raw parts
*/
package cp64

import (
	"fmt"
	"strings"
)

type Path struct {
	Raw      string
	Absolute string
}

func NewPath(raw string) *Path {
	return &Path{Raw: raw, Absolute: ""}
}

func (p *Path) Abs() string {
	if p.Absolute != "" {
		return p.Absolute
	}
	parts := p.getParts()
	stack := make([]string, 0)
	for _, part := range parts {
		switch part {
		case "..":
			// pop from the end
			n := len(stack)
			if n > 0 {
				stack = stack[:n-1]
			}
		case ".":
			continue
		default:
			// push
			stack = append(stack, part)
		}
	}
	res := p.joinStack(stack)
	p.Absolute = res
	return res
}

func (p *Path) getParts() []string {
	p.Raw = strings.ReplaceAll(p.Raw, "//", "/")
	p.Raw = strings.TrimPrefix(p.Raw, "/")
	p.Raw = strings.TrimSuffix(p.Raw, "/")
	return strings.Split(p.Raw, "/")
}

func (p *Path) joinStack(stack []string) string {
	if len(stack) == 0 {
		return "/"
	}
	return fmt.Sprintf("/%s/", strings.Join(stack, "/"))
}
