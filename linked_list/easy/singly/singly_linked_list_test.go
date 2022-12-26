package singly

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SinglyLinkedList(t *testing.T) {
	s := SinglyLinkedList[int]{}
	s.AddNode(1)
	s.AddNode(2)
	s.AddNode(3)
	s.AddNode(4)
	s.AddNode(5)
	s.AddNode(6)
	s.AddNode(7)
	s.AddNode(8)

	require.Equal(t, 1, s.head.value)
	require.Equal(t, 2, s.head.next.value)
	require.Equal(t, 3, s.head.next.next.value)
	require.Equal(t, 4, s.head.next.next.next.value)
	require.Equal(t, 5, s.head.next.next.next.next.value)
	require.Equal(t, 6, s.head.next.next.next.next.next.value)
	require.Equal(t, 7, s.head.next.next.next.next.next.next.value)
	require.Equal(t, 8, s.head.next.next.next.next.next.next.next.value)
	require.Equal(t, (*Node[int])(nil), s.head.next.next.next.next.next.next.next.next)

	s.Traverse()

	s.RemoveNode(1)
	s.RemoveNode(2)
	require.Equal(t, 3, s.head.value)
	require.Equal(t, 4, s.head.next.value)
	require.Equal(t, 5, s.head.next.next.value)
	require.Equal(t, 6, s.head.next.next.next.value)
	require.Equal(t, 7, s.head.next.next.next.next.value)
	require.Equal(t, 8, s.head.next.next.next.next.next.value)
	require.Equal(t, (*Node[int])(nil), s.head.next.next.next.next.next.next)

	s.Traverse()
}
