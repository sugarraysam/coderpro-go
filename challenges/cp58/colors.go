/*
Largest connected component on a grid
https://www.geeksforgeeks.org/largest-connected-component-on-a-grid/

Given a grid with different colors in a different cell, each color represented by a different number.
The task is to find out the largest connected component on the grid. Largest component grid refers to
a maximum set of cells such that you can move from any cell to any other cell in this set by only moving
between side-adjacent cells from the set.

My version uses Color vs Not Colored simplification (2 classes)

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

func NewRandomGrid() Grid {
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
	return Grid{Data: data, Visited: make(map[int]bool), Rows: nRows, Cols: nCols}
}

func (g Grid) String() (string, error) {
	var b strings.Builder
	for _, row := range g.Data {
		if _, err := b.WriteString(fmt.Sprintf("%s\n", row)); err != nil {
			return "", err
		}
	}
	return b.String(), nil
}

func (g Grid) FindLargestColorPatch() int {
	largest := -1

	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			pos := g.toPos(i, j)
			if g.isColored(pos) && !g.Visited[pos] {
				patch := g.BFS(pos)
				if patch > largest {
					largest = patch
				}
			}
		}
	}
	return largest
}

func (g Grid) BFS(pos int) int {
	patch := 0
	neighbors := []int{pos}

	for len(neighbors) > 0 {
		// pop
		curr := neighbors[0]
		neighbors = neighbors[1:]

		// visit?
		if !g.Visited[curr] {
			g.Visited[curr] = true
			patch++

			// up, down, left, right
			if up := g.Up(curr); up != -1 && g.isColored(up) {
				neighbors = append(neighbors, up)
			}
			if down := g.Down(curr); down != -1 && g.isColored(down) {
				neighbors = append(neighbors, down)
			}
			if left := g.Left(curr); left != -1 && g.isColored(left) {
				neighbors = append(neighbors, left)
			}
			if right := g.Right(curr); right != -1 && g.isColored(right) {
				neighbors = append(neighbors, right)
			}
		}
	}
	return patch
}

// Navigation helper methods
// toCoord, toPos, Up, Down, Left, Right
func (g Grid) isColored(pos int) bool {
	i, j := g.toCoord(pos)
	return g.Data[i][j] == Color
}

func (g Grid) toCoord(pos int) (int, int) {
	i := pos / g.Cols // integer division
	j := pos - (i * g.Cols)
	return i, j
}

func (g Grid) toPos(i, j int) int {
	return (i * g.Cols) + j
}

func (g Grid) Up(pos int) int {
	i, j := g.toCoord(pos)
	i -= 1
	if i >= 0 {
		return g.toPos(i, j)
	}
	return -1
}

func (g Grid) Down(pos int) int {
	i, j := g.toCoord(pos)
	i += 1
	if i < g.Rows {
		return g.toPos(i, j)
	}
	return -1
}

func (g Grid) Left(pos int) int {
	i, j := g.toCoord(pos)
	j -= 1
	if j >= 0 {
		return g.toPos(i, j)
	}
	return -1
}

func (g Grid) Right(pos int) int {
	i, j := g.toCoord(pos)
	j += 1
	if j < g.Cols {
		return g.toPos(i, j)
	}
	return -1
}
