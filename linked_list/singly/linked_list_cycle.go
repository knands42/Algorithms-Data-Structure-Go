package singly

func LinkedListCycle[T int](list *Node[T]) bool {

	nodeRef := make(map[*Node[T]]bool)

	currentNode := list
	for currentNode.next != nil {
		if _, ok := nodeRef[currentNode]; ok {
			return true
		} else {
			nodeRef[currentNode] = true
			currentNode = currentNode.next
		}
	}

	return false
}

func LinkedListCycleTwoPointers[T int](list *Node[T]) bool {
	slow := list
	fast := list

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}

	return false
}
