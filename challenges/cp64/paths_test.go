package cp64_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/coderpro-go/challenges/cp64"
)

func TestAbsolutPath(t *testing.T) {
	cases := []struct {
		raw      string
		expected string
	}{
		{
			raw:      "/users/tech/docs/.././desk/.././",
			expected: "/users/tech/",
		},
		{
			raw:      "/home/",
			expected: "/home/",
		},
		{
			raw:      "/../",
			expected: "/",
		},
		{
			raw:      "/home//foo/",
			expected: "/home/foo/",
		},
		{
			raw:      "/a/./b/../../c/",
			expected: "/c/",
		},
	}
	for i, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("path-%d", i), func(t *testing.T) {
			t.Parallel()
			require.Equal(t, cp64.NewPath(tc.raw).Abs(), tc.expected)
		})
	}
}
