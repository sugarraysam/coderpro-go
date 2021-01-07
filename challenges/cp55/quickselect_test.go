package cp55_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/coderpro-go/challenges/cp55"
)

func TestQuickselect(t *testing.T) {
	cases := []struct {
		numbers  []int
		k        int
		expected int
	}{
		{numbers: []int{3, 2, 1, 5, 6, 4}, k: 2, expected: 5},
		{numbers: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4, expected: 4},
		{numbers: []int{4, 3, 5, 2, 0, 1}, k: 3, expected: 3},
	}

	for i, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("quickselect-%d", i), func(t *testing.T) {
			t.Parallel()

			q1 := cp55.NewQuickselect(copyNumbers(tc.numbers), tc.k)
			require.Equal(t, q1.SolveSort(), tc.expected)

			q2 := cp55.NewQuickselect(copyNumbers(tc.numbers), tc.k)
			require.Equal(t, q2.SolveQS(), tc.expected)
		})
	}
}

func copyNumbers(numbers []int) []int {
	res := make([]int, len(numbers))
	copy(res, numbers)
	return res
}
