package sorting

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_InsertionSort(t *testing.T) {
	numbers := []int{3, 5, 1, 2, 4}
	result := insertionSort(numbers)

	require.Equal(t, []int{1, 2, 3, 4, 5}, result)
}
