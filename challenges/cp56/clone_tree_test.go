package cp56_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/coderpro-go/challenges/cp56"
)

func TestNewTree(t *testing.T) {
	cases := []struct {
		vals           []int
		expectedStr    string
		target         int
		expectedTarget int
	}{
		{
			vals:           []int{0, 1, 2, 3, 4, 5, 6},
			expectedStr:    "3 -> 1 -> 4 -> 0 -> 5 -> 2 -> 6 -> ",
			target:         3,
			expectedTarget: 3,
		},
		{
			vals:           []int{0, 1},
			expectedStr:    "1 -> 0 -> ",
			target:         3,
			expectedTarget: cp56.NilVal, // not found
		},
		{
			vals:           []int{5, 6, 2, 3, 4},
			expectedStr:    "3 -> 6 -> 4 -> 5 -> 2 -> ",
			target:         4,
			expectedTarget: 4,
		},
		{
			vals:           []int{5, -1, 2, -1, -1, 4},
			expectedStr:    "5 -> 4 -> 2 -> ",
			target:         8,
			expectedTarget: cp56.NilVal,
		},
		{
			vals:           []int{2, 3, 4, -1, 5, -1, 7, -1, -1, 6},
			expectedStr:    "3 -> 6 -> 5 -> 2 -> 4 -> 7 -> ",
			target:         7,
			expectedTarget: 7,
		},
	}
	for i, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("newtree-%d", i), func(t *testing.T) {
			t.Parallel()

			treeA := cp56.NewTree(tc.vals)
			treeB := cp56.NewTree(tc.vals)

			// String test
			require.Equal(t, treeA.String(), tc.expectedStr)

			// Find node
			require.Equal(t, cp56.SolveRecursive(treeA, treeB, tc.target), tc.expectedTarget)
			require.Equal(t, cp56.SolveIterative(treeA, treeB, tc.target), tc.expectedTarget)
		})
	}
}
