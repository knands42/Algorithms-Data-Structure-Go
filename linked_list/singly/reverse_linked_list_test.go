package singly

import "testing"

func Test_ReverserSinglyLinkedList(t *testing.T) {
	list := SinglyLinkedList[int]{}
	list.AddNode(1)
	list.AddNode(2)
	list.AddNode(3)
	list.AddNode(4)
	list.AddNode(5)
	list.AddNode(6)
	list.AddNode(7)
	list.AddNode(8)

	list.head = ReverseLinkedList(list.head)

	if list.head.value != 8 {
		t.Errorf("Expected 8, got %v", list.head.value)
	}
	if list.head.next.value != 7 {
		t.Errorf("Expected 7, got %v", list.head.next.value)
	}
	if list.head.next.next.value != 6 {
		t.Errorf("Expected 6, got %v", list.head.next.next.value)
	}
	if list.head.next.next.next.value != 5 {
		t.Errorf("Expected 5, got %v", list.head.next.next.next.value)
	}
	if list.head.next.next.next.next.value != 4 {
		t.Errorf("Expected 4, got %v", list.head.next.next.next.next.value)
	}
	if list.head.next.next.next.next.next.value != 3 {
		t.Errorf("Expected 3, got %v", list.head.next.next.next.next.next.value)
	}
	if list.head.next.next.next.next.next.next.value != 2 {
		t.Errorf("Expected 2, got %v", list.head.next.next.next.next.next.next.value)
	}
	if list.head.next.next.next.next.next.next.next.value != 1 {
		t.Errorf("Expected 1, got %v", list.head.next.next.next.next.next.next.next.value)
	}
}
