/*
Fibonacci Number - https://leetcode.com/problems/fibonacci-number/

The Fibonacci numbers, commonly denoted F(n) form a sequence,
called the Fibonacci sequence, such that each number is the sum
of the two preceding ones, starting from 0 and 1.

F(n) = F(n-1) + F(n-2)
F(0) = 0
F(1) = 1

Time complexity:
	O(n), in both cases (iterative, recursive) we compute fib(n) only once for each i in n

Space complexity:
	O(1) for iterative version (does not store anything)
	O(n) for recursive version, we use a cache

*/
package cp61

type fib struct {
	Cache map[int]int
}

func NewFibonacci() *fib {
	return &fib{Cache: make(map[int]int)}
}

func (f *fib) Reset() {
	f.Cache = make(map[int]int)
}

//					n: 0, 1, 2, 3, 4, 5, 6, 7  ...
// Fibonacci sequence: 0, 1, 1, 2, 3, 5, 8, 13 ...
func (f *fib) SolveIterative(n int) int {
	if n < 2 {
		return n
	}

	n1 := 0
	n2 := 1
	for i := 1; i < n; i++ {
		n1, n2 = n2, n1+n2
	}
	return n2
}

func (f *fib) SolveRecursive(n int) int {
	if n < 2 {
		return n
	}
	if v, ok := f.Cache[n]; ok {
		return v
	}
	v := f.SolveRecursive(n-1) + f.SolveRecursive(n-2)
	f.Cache[n] = v
	return v
}
