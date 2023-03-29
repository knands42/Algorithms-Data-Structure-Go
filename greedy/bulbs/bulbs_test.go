package bulbs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Bulbs(t *testing.T) {
	bulbs := []int{0, 1, 0, 1, 1, 1, 1}

	result := lightBulbs(bulbs)

	require.Equal(t, 4, result)

}
