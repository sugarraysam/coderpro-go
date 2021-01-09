package cp61_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/coderpro-go/challenges/cp61"
)

func TestFibonacci(t *testing.T) {
	cases := []struct {
		n        int
		expected int
	}{
		{n: 0, expected: 0},
		{n: 1, expected: 1},
		{n: 2, expected: 1},
		{n: 3, expected: 2},
		{n: 6, expected: 8},
		{n: 7, expected: 13},
		{n: 10, expected: 55},
		{n: 27, expected: 196418},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("fib-%d", tc.n), func(t *testing.T) {
			t.Parallel()

			fib := cp61.NewFibonacci()
			require.Equal(t, fib.SolveIterative(tc.n), tc.expected)
			require.Equal(t, fib.SolveRecursive(tc.n), tc.expected)
		})
	}
}
