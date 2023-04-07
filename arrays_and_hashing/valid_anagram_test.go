package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ValidAnagram(t *testing.T) {
	a := "anagram"
	b := "nagaram"
	require.True(t, isAnagram(a, b))

	a = "rat"
	b = "car"
	require.False(t, isAnagram(a, b))

	a = "aaa"
	b = "bbbb"
	require.False(t, isAnagram(a, b))
}
