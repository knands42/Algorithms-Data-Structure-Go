package highest_product

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HighestProduct(t *testing.T) {
	productPrices := []int{-10, -3, -1, 1, 2, 3, 2}

	result := highestProducts(productPrices)

	require.Equal(t, 30, result)
}
