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

	//times := l.append_timer(100)
	//fmt.Print(l)
	//fmt.Printf("\nlinked list: %d, array %d\n", times[0], times[1])

	listA := make_linked_list(10)
	listB := make_linked_list(10)
	fmt.Print(listA)
	listA.append_node(listB.head.data)
	fmt.Print(listA)
	fmt.Printf("\n")

}

func make_linked_list(size int) linkedList {
	ll := linkedList{}
	node := new(node)
	node.data = 1337
	for i := 0; i < size; i++ {
		ll.append_node(i)
	}
	return ll
}

func append_arrays_and_time(amount_elements int) []time.Duration {
	times := make([]time.Duration, 2)

	first_list := make_linked_list(amount_elements)
	second_list := make_linked_list(1000)

	t0 := time.Now()
	first_list.append_node(second_list.head.data)
	times[0] = time.Since(t0)

	arr := make([]int, ll.size)
	t0 = time.Now()
	arr_newlen := make([]int, len(arr)+amount_elements)

	for i := 0; i < len(arr)-1; i++ {
		arr_newlen[i] = arr[i]
	}
	for j := 0; j < amount_elements; j++ {
		arr_newlen[j] = data_to_add
		fmt.Printf("%d\n", arr_newlen[j])
	}
	times[1] = time.Since(t0)

	for i := range arr_newlen {
		fmt.Printf("%d\n", arr_newlen[i])
	}

	return times

// takes a linked list and appends varying amounts of elements

func (ll *linkedList) append_node(data int) {
	newNode := new(node)
	newNode.data = data
	//newNode.next = nil

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
