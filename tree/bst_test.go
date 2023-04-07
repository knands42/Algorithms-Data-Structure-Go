package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DeleteLeafNode(t *testing.T) {
	bst := BST{}
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(2)
	bst.Insert(5)
	bst.Insert(13)
	bst.Insert(22)
	bst.Insert(1)
	bst.Insert(14)

	inOrder := bst.InOrder()
	require.Equal(t, []int{1, 2, 5, 5, 10, 13, 14, 15, 22}, inOrder)
	require.True(t, bst.Search(1))
	bst.Delete(1)
	require.False(t, bst.Search(1))
	inOrder = bst.InOrder()
	require.Equal(t, []int{2, 5, 5, 10, 13, 14, 15, 22}, inOrder)
}

func Test_DeleteNodeWithOneChild(t *testing.T) {
	bst := BST{}
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(2)
	bst.Insert(5)
	bst.Insert(13)
	bst.Insert(22)
	bst.Insert(1)
	bst.Insert(14)

	inOrder := bst.InOrder()
	require.Equal(t, []int{1, 2, 5, 5, 10, 13, 14, 15, 22}, inOrder)
	require.True(t, bst.Search(2))
	bst.Delete(2)
	require.False(t, bst.Search(2))
	inOrder = bst.InOrder()
	require.Equal(t, []int{1, 5, 5, 10, 13, 14, 15, 22}, inOrder)
}

func Test_DeleteNodeWithTwoChildren(t *testing.T) {
	bst := BST{}
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(13)
	bst.Insert(22)
	bst.Insert(1)
	bst.Insert(14)

	inOrder := bst.InOrder()
	require.Equal(t, []int{1, 2, 5, 6, 10, 13, 14, 15, 22}, inOrder)
	require.True(t, bst.Search(5))
	bst.Delete(5)
	require.False(t, bst.Search(5))
	inOrder = bst.InOrder()
	require.Equal(t, []int{1, 2, 6, 10, 13, 14, 15, 22}, inOrder)
}
