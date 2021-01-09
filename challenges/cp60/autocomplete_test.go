package cp60_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/coderpro-go/challenges/cp60"
)

func TestTrieBuildingAndPrinting(t *testing.T) {
	cases := []struct {
		words               []string
		prefix              string
		expectedStr         string
		expectedCompletions []string
	}{
		{
			words:               []string{"dog", "door", "cat", "rat"},
			prefix:              "do",
			expectedStr:         "  \nd c r \no a a \ng o t t \nr \n",
			expectedCompletions: []string{"dog", "door"},
		},
		{
			words:               []string{"pepwi", "pepsi", "pepzi", "pepti"},
			prefix:              "pep",
			expectedStr:         "  \np \ne \np \nw s z t \ni i i i \n",
			expectedCompletions: []string{"pepwi", "pepsi", "pepzi", "pepti"},
		},
		{
			words:               []string{"aa", "bb", "cc"},
			prefix:              "zz",
			expectedStr:         "  \na b c \na b c \n",
			expectedCompletions: []string{},
		},
	}
	for i, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("triePrinting-%d", i), func(t *testing.T) {
			t.Parallel()

			// string
			trie := cp60.NewTrie(tc.words)
			got, err := trie.StringPerLevel()
			require.Nil(t, err)
			require.Equal(t, got, tc.expectedStr)

			// completions
			require.ElementsMatch(t, trie.SolveTrie(tc.words, tc.prefix), tc.expectedCompletions)
			require.ElementsMatch(t, cp60.SolveRegexp(tc.words, tc.prefix), tc.expectedCompletions)
		})
	}
}
