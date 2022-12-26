package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TwoSum(t *testing.T) {
	target := 9
	input := []int{2, 7, 11, 15}
	require.Equal(t, []int{0, 1}, twoSum(input, target))

	target = 6
	input = []int{3, 2, 4}
	require.Equal(t, []int{1, 2}, twoSum(input, target))

	target = 6
	input = []int{3, 3}
	require.Equal(t, []int{0, 1}, twoSum(input, target))

	target = 1
	input = []int{3, 3}
	require.Equal(t, []int{}, twoSum(input, target))
}
