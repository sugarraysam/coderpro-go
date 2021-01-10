package revtree_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/coderpro-go/challenges/revtree"
)

func TestReverseBinTree(t *testing.T) {
	cases := []struct {
		vals     []int
		expected string
	}{
		{
			vals:     []int{0, 1, 2, 3, 4, 5},
			expected: "3 4 5 \n1 2 \n0 \n",
		},
		{
			vals:     []int{0, 1, 2},
			expected: "1 2 \n0 \n",
		},
		{
			vals:     []int{0},
			expected: "0 \n",
		},
		{
			vals:     []int{0, 1, 2, 3, 4, 5, 6, 7},
			expected: "7 4 5 6 \n3 1 2 \n0 \n",
		},
	}
	for i, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("rtree-%d", i), func(t *testing.T) {
			t.Parallel()

			tree := revtree.NewTree(tc.vals)
			rtree := revtree.Invert(tree)
			got, err := rtree.DFSString()
			require.Nil(t, err)
			require.Equal(t, got, tc.expected)
		})
	}
}
