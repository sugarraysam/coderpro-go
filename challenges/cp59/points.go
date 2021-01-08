/*
K Closest Points to Origin
https://leetcode.com/problems/k-closest-points-to-origin/

We have a list of points on the plane.  Find the K closest points to the origin (0, 0).

(Here, the distance between two points on a plane is the Euclidean distance.)

You may return the answer in any order.  The answer is guaranteed to be unique (except for the order that it is in.)

Time Complexity:
	O[ (n+k) * log(n) ], every push and pop is log(n) because we use a min heap, so we will do (n+k)* log(n) operations

Space Complexity:
	O(n + k), we store all points in a minHeap (n), and then we pop k elements in res
*/
package cp59

import (
	"container/heap"
)

// Point && PointHeap methods
type Point struct {
	X, Y int
}

// not using sqrt for quicker computation and to use ints
func (p Point) distanceFromOrigin() int {
	return p.X*p.X + p.Y*p.Y
}

// PointHeap - min heap using distance as priority (want minimum distance)
type PointHeap []Point

func (ph PointHeap) Len() int { return len(ph) }

func (ph PointHeap) Less(i, j int) bool {
	return ph[i].distanceFromOrigin() < ph[j].distanceFromOrigin()
}

func (ph PointHeap) Swap(i, j int) { ph[i], ph[j] = ph[j], ph[i] }

// Push & Pop - implement MinHeap interface
// push complexity is O(log n) , n == h.Len()
func (ph *PointHeap) Push(x interface{}) {
	*ph = append(*ph, x.(Point))
}

// pop complexity is O(log n)
func (ph *PointHeap) Pop() interface{} {
	old := *ph
	n := len(old)
	x := old[n-1]
	*ph = old[0 : n-1]
	return x
}

func SolveHeap(points []Point, k int) []Point {
	h := &PointHeap{}
	heap.Init(h)

	for _, p := range points {
		heap.Push(h, p)
	}

	// pop k min elements (by distance)
	var res []Point
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(Point))
	}
	return res
}
