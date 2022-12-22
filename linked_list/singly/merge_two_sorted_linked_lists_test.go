package singly

import "testing"

func Test_MergeTwoSortedLinkedLists(t *testing.T) {
	list1 := SinglyLinkedList[int]{}
	list1.AddNode(1)
	list1.AddNode(3)
	list1.AddNode(5)

	list2 := SinglyLinkedList[int]{}
	list2.AddNode(2)
	list2.AddNode(4)

	mergedList := MergeTwoSortedLinkedLists(list1.head, list2.head)

	if mergedList.value != 1 {
		t.Errorf("Expected 1, got %d", mergedList.value)
	}

	if mergedList.next.value != 2 {
		t.Errorf("Expected 2, got %d", mergedList.next.value)
	}

	if mergedList.next.next.value != 3 {
		t.Errorf("Expected 3, got %d", mergedList.next.next.value)
	}

	if mergedList.next.next.next.value != 4 {

		t.Errorf("Expected 4, got %d", mergedList.next.next.next.value)
	}

	if mergedList.next.next.next.next.value != 5 {
		t.Errorf("Expected 5, got %d", mergedList.next.next.next.next.value)
	}

}
