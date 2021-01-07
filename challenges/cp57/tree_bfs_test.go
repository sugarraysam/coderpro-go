package cp57_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/coderpro-go/challenges/cp57"
)

func TestTreeBFS(t *testing.T) {
	cases := []struct {
		vals     []string
		expected string
	}{
		{
			vals:     []string{"a", "b", "c", "d", "e", "f", "g"},
			expected: "a \nb c \nd e f g \n",
		},
		{
			vals:     []string{"a", cp57.NilVal, "b", cp57.NilVal, cp57.NilVal, "c", "d"},
			expected: "a \nb \nc d \n",
		},
		{
			vals:     []string{"a", "b", cp57.NilVal, cp57.NilVal, "c"},
			expected: "a \nb \nc \n",
		},
	}
	for i, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("treebfs-%d", i), func(t *testing.T) {
			t.Parallel()

			got, err := cp57.NewTree(tc.vals).BFSIterative()
			require.Nil(t, err)
			require.Equal(t, got, tc.expected)
		})
	}
}
