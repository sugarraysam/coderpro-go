package cp63_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/coderpro-go/challenges/cp63"
)

func TestSubarray(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		{
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2,
		},
		{
			nums:     []int{1, 3, 2, 5, 7, 2},
			k:        14,
			expected: 2,
		},
		{
			nums:     []int{1, 3, 2, 5, 6, 8},
			k:        27,
			expected: 0,
		},
	}
	for i, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("subarray-%d", i), func(t *testing.T) {
			t.Parallel()
			require.Equal(t, cp63.SolveBruteforce(tc.nums, tc.k), tc.expected)
			require.Equal(t, cp63.SolvePointers(tc.nums, tc.k), tc.expected)
			require.Equal(t, cp63.SolveHashmap(tc.nums, tc.k), tc.expected)
		})
	}
}
