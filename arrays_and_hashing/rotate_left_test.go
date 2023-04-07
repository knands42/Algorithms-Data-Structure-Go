package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RotLeft(t *testing.T) {
	var tests = []struct {
		a    []int32
		d    int32
		want []int32
	}{
		{[]int32{1, 2, 3, 4, 5}, 4, []int32{5, 1, 2, 3, 4}},
		{[]int32{1, 2, 3, 4, 5}, 2, []int32{3, 4, 5, 1, 2}},
	}

	for _, test := range tests {
		got := RotLeft(test.a, test.d)
		require.Equal(t, test.want, got)
	}
}
