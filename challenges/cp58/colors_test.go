package cp58_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/coderpro-go/challenges/cp58"
)

func TestFindLargestColorPatch(t *testing.T) {
	n := 3

	for i := 0; i < n; i++ {
		t.Run(fmt.Sprintf("gridString-%d", i), func(t *testing.T) {
			t.Parallel()
			g := cp58.NewRandomGrid()
			got, err := g.String()
			require.Nil(t, err)
			require.Greater(t, len(got), 0)

			largest := g.FindLargestColorPatch()
			require.NotEqual(t, largest, -1)

			// Testing by visual inspection
			// fmt.Println(got)
			// fmt.Println(largest)
			// require.True(t, false) // force test runner output
		})
	}
}
