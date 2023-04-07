package singly

func MergeTwoSortedLinkedListsRecursively[T int](list1 *Node[T], list2 *Node[T]) *Node[T] {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var mergedList *Node[T]

	if list1.value <= list2.value {
		mergedList = list1
		mergedList.next = MergeTwoSortedLinkedListsRecursively(list1.next, list2)
	} else {
		mergedList = list2
		mergedList.next = MergeTwoSortedLinkedListsRecursively(list1, list2.next)
	}

	return mergedList
}

func MergeTwoSortedLinkedListsIteration[T int](list1 *Node[T], list2 *Node[T]) *Node[T] {
	var newList *Node[T]

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.value <= list2.value {
		newList = list1
		list1 = list1.next
	} else {
		newList = list2
		list2 = list2.next
	}

	currentNode := newList

	for list1 != nil && list2 != nil {
		if list1.value <= list2.value {
			currentNode.next = list1
			list1 = list1.next
		} else {
			currentNode.next = list2
			list2 = list2.next
		}
		currentNode = currentNode.next

	}

	if list1 != nil {
		currentNode.next = list1
	} else if list2 != nil {
		currentNode.next = list2
	}

	return newList

}
