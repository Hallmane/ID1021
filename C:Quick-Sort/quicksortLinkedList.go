package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Node struct {
	data int
	prev *Node
	next *Node
}
type DoublyLinkedList struct {
	start *Node
	end   *Node
}

func (dll *DoublyLinkedList) quick_sort_LL(start *Node, end *Node) {
	if start != end && start != nil &&
		end != nil && end.next != start {
		node := dll.partition_LL(start, end)
		if node != nil {
			// Recursively sort elements
			dll.quick_sort_LL(node.next, end)
			dll.quick_sort_LL(start, node.prev)
		}
	}
}

func (dll *DoublyLinkedList) partition_LL(start *Node, end *Node) *Node {
	current := start
	search := start.prev
	var temp int = 0
	for current != nil && current != end {
		if current.data <= end.data {
			if search == nil {
				search = start
			} else {
				search = search.next
			}
			temp = current.data
			current.data = search.data
			search.data = temp
		}
		current = current.next
	}
	if search != nil {
		search = search.next
	}
	temp = end.data
	if search == nil {
		search = start
	}
	end.data = search.data
	search.data = temp
	return search
}

func (dll *DoublyLinkedList) append(node *Node) {
	dll.end.next = node
	node.prev = dll.end
	dll.end = node
}

func make_node(size int) Node {
	return Node{data: rand.Intn(size), prev: nil, next: nil}
}

func make_linkedlist(amount int, size int) DoublyLinkedList {
	dll := DoublyLinkedList{}

	head := make_node(size)
	dll.start = &head
	dll.end = &head

	for i := 0; i <= amount; i++ {
		n := make_node(size)
		dll.append(&n)
	}

	return dll
}

func (dll DoublyLinkedList) String() string {
	sb := strings.Builder{}

	for iter := dll.start; iter != nil; iter = iter.next {
		sb.WriteString(fmt.Sprintf("%d ", iter.data))
	}
	sb.WriteString(fmt.Sprintf("\n"))
	return sb.String()
}
