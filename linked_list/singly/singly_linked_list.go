package singly

import "fmt"

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type SinglyLinkedList[T comparable] struct {
	head *Node[T]
}

func (s *SinglyLinkedList[T]) AddNode(value T) {
	newNode := &Node[T]{
		value: value,
		next:  nil,
	}

	if s.head == nil {
		s.head = newNode
	} else {
		currentNode := s.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
}
func (s *SinglyLinkedList[T]) RemoveNode(value T) {
	if s.head == nil {
		return
	}

	if s.head.value == value {
		s.head = s.head.next
	} else {
		currentNode := s.head
		nextNode := currentNode.next

		for s.head.next != nil {
			if nextNode.value == value {
				currentNode.next = nextNode.next
				return
			}

			currentNode = nextNode
			nextNode = nextNode.next
		}
	}
}
func (s *SinglyLinkedList[T]) Traverse() {
	if s.head == nil {
		return
	}

	currentNode := s.head
	for currentNode != nil {
		fmt.Printf("%v ", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Println()
}
