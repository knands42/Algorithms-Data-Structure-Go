package singly

func MergeTwoSortedLinkedLists[T int](list1 *Node[T], list2 *Node[T]) *Node[T] {
	var newList *Node[T]

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	currentList1 := list1
	currentList2 := list2
	for currentList1.next != nil && currentList2.next != nil {
		if currentList1.value > currentList2.value {
			newList.next = currentList2
			currentList2 = currentList2.next
		} else {
			newList.next = currentList1
			currentList1 = currentList1.next
		}
	}

	for currentList1.next != nil {
		newList.next = currentList1.next
		currentList1 = currentList1.next
	}

	for currentList2.next != nil {
		newList.next = currentList2.next
		currentList2 = currentList2.next
	}

	return newList
}
