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
		expectedStr     string
		expectedLargest int
	}{
		{
			dec:             242,
			expectedStr:     "11110010",
			expectedLargest: 4,
		},
		{
			dec:             155,
			expectedStr:     "10011011",
			expectedLargest: 2,
		},
		{
			dec:             12345,
			expectedStr:     "11000000111001",
			expectedLargest: 3,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("bits-%d", tc.dec), func(t *testing.T) {
			t.Parallel()
			bits := cp65.NewBits(tc.dec)
			bitsStr, err := bits.String()
			require.Nil(t, err)
			require.Equal(t, bitsStr, tc.expectedStr)
			require.Equal(t, bits.FindLargest(), tc.expectedLargest)
			require.Equal(t, cp65.SolveShifting(tc.dec), tc.expectedLargest)
		})
	}
}
