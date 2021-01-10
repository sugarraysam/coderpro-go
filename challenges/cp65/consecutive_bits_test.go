package cp65_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/coderpro-go/challenges/cp65"
)

func TestConsecutiveBits(t *testing.T) {
	cases := []struct {
		dec             int
		expectedBin     []int
		expectedLargest int
	}{
		{
			dec:             242,
			expectedBin:     []int{1, 1, 1, 1, 0, 0, 1, 0},
			expectedLargest: 4,
		},
		{
			dec:             155,
			expectedBin:     []int{1, 0, 0, 1, 1, 0, 1, 1},
			expectedLargest: 2,
		},
		{
			dec:             12345,
			expectedBin:     []int{1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1},
			expectedLargest: 3,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("bits-%d", tc.dec), func(t *testing.T) {
			t.Parallel()
			bits := cp65.NewBits(tc.dec)
			require.Equal(t, bits.Binary, tc.expectedBin)
			require.Equal(t, bits.FindLargest(), tc.expectedLargest)
		})
	}
}
