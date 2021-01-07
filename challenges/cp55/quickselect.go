package cp55

import (
	"math/rand"
	"sort"
)

type Quickselect struct {
	Numbers []int
	K       int
}

func NewQuickselect(numbers []int, k int) *Quickselect {
	return &Quickselect{Numbers: numbers, K: k}
}

/* SolveSort - sorts then returns kth largest element

Time Complexity: O(n log n) because we are sorting

Space Complexity: O(1), sorting in-place
*/
func (q *Quickselect) SolveSort() int {
	sort.Ints(q.Numbers)
	return q.Numbers[len(q.Numbers)-q.K]
}

func (q *Quickselect) SolveQS() int {
	pivotIdx := 0
	left := 0
	right := len(q.Numbers)
	k := right - q.K // final position of kth largest in a sorted array
	for {
		pivotIdx = q.partition(left, right)

		if pivotIdx == k {
			return q.Numbers[pivotIdx]
		} else if pivotIdx < k {
			left = pivotIdx + 1
		} else {
			right = pivotIdx
		}
	}
}

func (q *Quickselect) partition(left, right int) int {
	pivotIdx := left + rand.Intn(right-left)
	pivotVal := q.Numbers[pivotIdx]
	q.swap(pivotIdx, right-1) // place pivot at the end

	storeIdx := left
	for i := left; i < right; i++ {
		if q.Numbers[i] < pivotVal {
			q.swap(i, storeIdx)
			storeIdx++
		}
	}
	q.swap(right-1, storeIdx) // store pivot at storeIdx
	return storeIdx
}

func (q *Quickselect) swap(i, j int) {
	q.Numbers[i], q.Numbers[j] = q.Numbers[j], q.Numbers[i]
}
