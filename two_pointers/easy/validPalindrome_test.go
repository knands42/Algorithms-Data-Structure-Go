package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Palindrome(t *testing.T) {
	require.Equal(t, IsPalindrome("A man, a plan, a canal: Panama"), true)
	require.Equal(t, IsPalindrome("race a car"), false)
	require.Equal(t, IsPalindrome(""), true)
}
