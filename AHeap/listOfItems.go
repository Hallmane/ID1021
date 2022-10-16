package main

import (
	"fmt"
	"strings"
)

type LNode struct {
	priority int
	next     *LNode
}

type LinkedHeap struct {
	entry *LNode
}

func makeLinkedHeap(priority int) *LinkedHeap {
	entry_node := makeNode(priority)
	return &LinkedHeap{entry: entry_node}
}

func makeNode(node int) *LNode {
	return &LNode{priority: node, next: nil}
}

// "O(1) add & O(n) remove"
func (lh *LinkedHeap) Constant_Add(node int) {
	new_node := makeNode(node)

	if lh.entry == nil {
		lh.entry = new_node
	} else {
		new_node.next = lh.entry
		lh.entry = new_node
	}
}

func (lh *LinkedHeap) Constant_Remove() int {
	var returnVal int
	if lh.entry.next == nil {
		returnVal = lh.entry.priority
		lh.entry = nil
		return returnVal
	}
	returnVal = lh.entry.priority
	lh.entry = lh.entry.next
	return returnVal
}

func (lh *LinkedHeap) Linear_Add(new_prio int) {
	current := lh.entry
	prev := current
	new_node := makeNode(new_prio)
	higher_prio := false
	first := false

	for current.next != nil {
		if current.priority > new_prio {
			higher_prio = true
			if current == lh.entry {
				first = true
			}
			break
		}
		prev = current
		current = current.next
	}
	if first {
		new_node.next = current
		lh.entry = new_node
		return
	}
	if higher_prio {
		new_node.next = current
		prev.next = new_node
	} else if new_prio < current.priority {
		prev.next = new_node
		new_node.next = current
		return
	} else {
		current.next = new_node
	}
}

func (lh *LinkedHeap) Linear_Remove() int {
	current := lh.entry
	lowest := lh.entry
	var node_before_removed *LNode

	for ; current.next != nil; current = current.next {
		if current.next.priority < lowest.priority {
			node_before_removed = current
			lowest = current.next
		}
	}
	if node_before_removed != nil {
		if node_before_removed.next != nil {
			if node_before_removed.next.next != nil {
				node_before_removed.next = node_before_removed.next.next
				lowest.next = nil
			}
		} else {
			node_before_removed.next = nil
		}
	}
	return lowest.priority
}

func (lh LinkedHeap) String() string {
	sb := strings.Builder{}

	for iter := lh.entry; iter != nil; iter = iter.next {
		sb.WriteString(fmt.Sprintf("%d ", iter.priority))
	}

	sb.WriteString(fmt.Sprintf("\n"))
	return sb.String()
}
