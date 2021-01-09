package cp62_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/coderpro-go/challenges/cp62"
)

func TestRoman(t *testing.T) {
	cases := []struct {
		roman    string
		expected int
	}{
		{roman: "III", expected: 3},
		{roman: "IV", expected: 4},
		{roman: "IX", expected: 9},
		{roman: "LVIII", expected: 58},
		{roman: "MCMXCIV", expected: 1994},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.roman, func(t *testing.T) {
			t.Parallel()
			r := cp62.NewRoman(tc.roman)
			require.Equal(t, r.SolveIterative(), tc.expected)
			require.Equal(t, r.SolveRecursive(), tc.expected)
		})
	}
}
