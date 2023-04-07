// https://leetcode.com/problems/contains-duplicate/
package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ContainsDuplicate(t *testing.T) {
	input := []int{1, 2, 3, 1}
	require.True(t, containsDuplicate(input))

	input = []int{1, 2, 3, 4}
	require.False(t, containsDuplicate(input))

	input = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	require.True(t, containsDuplicate(input))
}
