/*
Largest connected component on a grid
https://www.geeksforgeeks.org/largest-connected-component-on-a-grid/

Given a grid with different colors in a different cell, each color represented by a different number.
The task is to find out the largest connected component on the grid. Largest component grid refers to
a maximum set of cells such that you can move from any cell to any other cell in this set by only moving
between side-adjacent cells from the set.

My version uses Color vs Not Colored simplification (2 classes).

Time Complexity:
	O(n * m), where n is # of rows, and m is # of cols, we never visit a node twice because we keep track
				of a visited map

Space Complexity:
	O(2 * n * m), the visited map will hold a value for each square of the grid, also, we keep track of the neighbors
				to be visited, which can potentially be the whole grid (n * m)
*/
package cp58

import (
	"fmt"
	"math/rand"
	"strings"
)

// Square methods
type Square int

const (
	Color Square = iota
	NoColor
)

func (s Square) String() string {
	return [...]string{"X", " "}[s]
}

// Grid methods
type Grid struct {
	Data    [][]Square
	Visited map[int]bool
	Rows    int
	Cols    int
}

func NewRandomGrid() *Grid {
	nRows := 3 + rand.Intn(8) // rows between [3, 10]
	nCols := 3 + rand.Intn(8) // cols between [3, 10]

	var data [][]Square
	for i := 0; i < nRows; i++ {
		var row []Square
		for j := 0; j < nCols; j++ {
			if rand.Intn(3) == 0 { // 33% chance Color
				row = append(row, Color)
			} else {
				row = append(row, NoColor)
			}
		}
		data = append(data, row)
	}
	return &Grid{Data: data, Visited: make(map[int]bool), Rows: nRows, Cols: nCols}
}

func (g *Grid) String() (string, error) {
	var b strings.Builder
	for _, row := range g.Data {
		if _, err := b.WriteString(fmt.Sprintf("%s\n", row)); err != nil {
			return "", err
		}
	}
	return b.String(), nil
}

//
// ---------- Iterative solution ----------
// Uses helper methods for Up,Down,Left,Right, and translates i,j -> pos for hashing
//
func (g *Grid) SolveIterative() int {
	largest := -1
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			pos := g.toPos(i, j)
			if g.isColored(pos) && !g.Visited[pos] {
				largest = max(largest, g.BFSIterative(pos))
			}
		}
	}
	return largest
}

// CAREFUL ! we can have a same node twice in the neighbors queue
func (g *Grid) BFSIterative(pos int) int {
	res := 0
	neighbors := []int{pos}

	for len(neighbors) > 0 {
		// pop
		curr := neighbors[0]
		neighbors = neighbors[1:]

		// visit?
		if !g.Visited[curr] {
			g.Visited[curr] = true
			res++
			for _, n := range g.Neighbors(curr) {
				if g.isColored(n) && !g.Visited[n] {
					neighbors = append(neighbors, n)
				}
			}
		}
	}
	return res
}

// Navigation helper methods
// toCoord, toPos, Up, Down, Left, Right
func (g *Grid) isColored(pos int) bool {
	i, j := g.toCoord(pos)
	return g.Data[i][j] == Color
}

func (g Grid) toCoord(pos int) (int, int) {
	i := pos / g.Cols // integer division
	j := pos - (i * g.Cols)
	return i, j
}

func (g *Grid) toPos(i, j int) int {
	return (i * g.Cols) + j
}

func (g *Grid) Up(pos int) int {
	i, j := g.toCoord(pos)
	i -= 1
	if i >= 0 {
		return g.toPos(i, j)
	}
	return -1
}

func (g *Grid) Down(pos int) int {
	i, j := g.toCoord(pos)
	i += 1
	if i < g.Rows {
		return g.toPos(i, j)
	}
	return -1
}

func (g *Grid) Left(pos int) int {
	i, j := g.toCoord(pos)
	j -= 1
	if j >= 0 {
		return g.toPos(i, j)
	}
	return -1
}

func (g *Grid) Right(pos int) int {
	i, j := g.toCoord(pos)
	j += 1
	if j < g.Cols {
		return g.toPos(i, j)
	}
	return -1
}

//
// ---------- Recursive Solution ----------
//
//
func (g *Grid) SolveRecursive() int {
	largest := -1
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			pos := g.toPos(i, j)
			if !g.Visited[pos] && g.isColored(pos) {
				largest = max(largest, g.BFSRecursive(pos))
			}
		}
	}
	return largest
}

func (g *Grid) BFSRecursive(pos int) int {
	// base case
	if g.Visited[pos] || !g.isColored(pos) {
		return 0
	}
	// visit
	g.Visited[pos] = true
	res := 1
	for _, n := range g.Neighbors(pos) {
		res += g.BFSRecursive(n)
	}
	return res
}

// Find up, down, left, right neighbors
func (g *Grid) Neighbors(pos int) []int {
	var res []int
	if up := g.Up(pos); up != -1 {
		res = append(res, up)
	}
	if down := g.Down(pos); down != -1 {
		res = append(res, down)
	}
	if left := g.Left(pos); left != -1 {
		res = append(res, left)
	}
	if right := g.Right(pos); right != -1 {
		res = append(res, right)
	}
	return res
}

// Reset - reset the visited grid
func (g *Grid) Reset() {
	g.Visited = make(map[int]bool)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
