package cp59_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/coderpro-go/challenges/cp59"
)

func TestSolveHeap(t *testing.T) {
	cases := []struct {
		points   []cp59.Point
		k        int
		expected []cp59.Point
	}{
		{
			points:   []cp59.Point{{1, 3}, {-2, 2}},
			k:        1,
			expected: []cp59.Point{{-2, 2}},
		},
		{
			points:   []cp59.Point{{3, 3}, {5, -1}, {-2, 4}},
			k:        2,
			expected: []cp59.Point{{-2, 4}, {3, 3}},
		},
	}
	for i, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("closesToOrigin-%d", i), func(t *testing.T) {
			t.Parallel()
			require.ElementsMatch(t, cp59.SolveHeap(tc.points, tc.k), tc.expected)
		})
	}
}
