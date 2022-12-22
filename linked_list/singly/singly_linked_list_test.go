package singly

import "testing"

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

	s.Traverse()

	s.RemoveNode(1)
	s.RemoveNode(2)

	s.Traverse()
}
