package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Stack(t *testing.T) {
	var stack Stack
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Push("d")

	require.Equal(t, stack.Length(), 4)
	require.Equal(t, stack.Empty(), false)

	stack.Pop()
	require.Equal(t, stack.Length(), 3)
	require.Equal(t, stack.Empty(), false)

	stack.Pop()
	require.Equal(t, stack.Length(), 2)
	require.Equal(t, stack.Empty(), false)

	stack.Pop()
	require.Equal(t, stack.Length(), 1)
	require.Equal(t, stack.Empty(), false)

	stack.Pop()
	require.Equal(t, stack.Length(), 0)
	require.Equal(t, stack.Empty(), true)
}

func Test_IsValid(t *testing.T) {
	stack := Stack{}

	require.Equal(t, IsValid("()", stack), true)
	require.Equal(t, IsValid("()[]{}", stack), true)
	require.Equal(t, IsValid("(]", stack), false)
	require.Equal(t, IsValid("([)]", stack), false)
	require.Equal(t, IsValid("{[]}", stack), true)
	require.Equal(t, IsValid("]", stack), false)
	require.Equal(t, IsValid("(", stack), false)
}
