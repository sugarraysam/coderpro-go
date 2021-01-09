package cp60_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/coderpro-go/challenges/cp60"
)

func TestTrieBuildingAndPrinting(t *testing.T) {
	cases := []struct {
		words       []string
		prefix      string
		expectedStr string
		expected    []string
	}{
		{
			words:    []string{"dog", "door", "cat", "rat"},
			prefix:   "do",
			expected: []string{"dog", "door"},
		},
		{
			words:    []string{"pepwi", "pepsi", "pepzi", "pepti"},
			prefix:   "pep",
			expected: []string{"pepwi", "pepsi", "pepzi", "pepti"},
		},
		{
			words:    []string{"aa", "bb", "cc"},
			prefix:   "zz",
			expected: []string{},
		},
	}
	for i, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("triePrinting-%d", i), func(t *testing.T) {
			t.Parallel()
			require.ElementsMatch(t, cp60.NewTrie(tc.words).SolveTrie(tc.words, tc.prefix), tc.expected)
			require.ElementsMatch(t, cp60.SolveRegexp(tc.words, tc.prefix), tc.expected)
		})
	}
}
