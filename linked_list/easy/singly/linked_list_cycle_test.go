package singly

import "testing"

func Test_LinkedListCycle(t *testing.T) {
	list := SinglyLinkedList[int]{}
	list.AddNode(1)
	list.AddNode(2)
	list.AddNode(3)
	list.AddNode(4)
	list.AddNode(5)

	if LinkedListCycle(list.head) {
		t.Errorf("Expected false, got true")
	}

	list.head.next.next.next.next.next = list.head.next.next

	if !LinkedListCycle(list.head) {
		t.Errorf("Expected true, got false")
	}
}

func Test_LinkedListCycleTwoPointers(t *testing.T) {

	list := SinglyLinkedList[int]{}
	list.AddNode(1)
	list.AddNode(2)
	list.AddNode(3)

	if LinkedListCycleTwoPointers(list.head) {
		t.Errorf("Expected false, got true")
	}

	list.head.next.next.next = list.head.next

	if !LinkedListCycleTwoPointers(list.head) {
		t.Errorf("Expected true, got false")
	}
}
