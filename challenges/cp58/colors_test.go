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
			gotString, err := g.String()
			require.Nil(t, err)
			require.Greater(t, len(gotString), 0)

			// make sure iterative & recursive give same result
			gotIterative := g.SolveIterative()
			g.Reset()
			gotRecursive := g.SolveRecursive()
			require.Equal(t, gotIterative, gotRecursive)

			// Testing by visual inspection
			// fmt.Println(gotString)
			// fmt.Println("iterative: ", gotIterative)
			// fmt.Println("recursive: ", gotRecursive)
			// require.True(t, false) // force test runner output
		})
	}
}
