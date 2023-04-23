package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	sorted := MergeSort(arr)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, sorted)
}
