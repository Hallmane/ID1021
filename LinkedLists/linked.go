package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	size int
}

func main() {

	l := linkedList{}
	l.append_node(1)
	l.append_node(2)
	l.append_node(3)
	l.append_node(4)
	l.append_node(5)
	l.append_node(6)
	l.append_node(7)
	l.append_node(8)
	l.append_node(9)
	fmt.Print(l)
	fmt.Print(l.get_node(7).data)
	l.remove_node(7)
	fmt.Print(l)
}

// takes a linked list and appends varying amounts of elements
//
//	This is for comparing arrays with linked lists
func (ll *linkedList) append_timer(amount_elements int) []time.Duration {
	times := make([]time.Duration, 2)
	data_to_add := rand.Intn(1000)
	ll_elems, arr_elems := amount_elements, amount_elements

	t0 := time.Now()
	for ; amount_elements > 0; amount_elements-- {
		ll.append_node(data_to_add)
	}
	times[0] = time.Now().Sub(t0)

	arr := make([]int, ll.size)
	t0 = time.Now()
	arr_newlen := make([]int, len(arr)+amount_elements)
	for i := 0; i < len(arr_newlen); i++ {

	}

}

func (ll *linkedList) append_node(data int) {
	newNode := new(node)
	newNode.data = data
	//newNode.next = nil // by defualt nil

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for ; current.next != nil; current = current.next {
		}
		current.next = newNode
	}
	ll.size += 1
}

func (ll linkedList) get_node(search_data int) *node {
	for iter := ll.head; iter != nil; iter = iter.next {
		if iter.data == search_data {
			return iter
		}
	}
	return nil
}

func (ll *linkedList) remove_node(search_data int) {
	var reti *node

	for iter := ll.head; iter != nil; iter = iter.next {
		if iter.data == search_data {
			reti.next = iter.next
			return
		}

		reti = iter
	}
}

func (ll linkedList) String() string {
	sb := strings.Builder{}

	for iter := ll.head; iter != nil; iter = iter.next {
		sb.WriteString(fmt.Sprintf("%d-> ", iter.data))
	}
	return sb.String()
}
