package main

import (
	"fmt"
	"strings"
	"time"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

type benchmark struct {
	times []time.Duration
}

func main() {

	//times := l.append_timer(100)
	//fmt.Print(l)

	//listA := make_linked_list(10)
	//listB := make_linked_list(10)
	//fmt.Print(listA)
	//listA.append_node(listB.head.data)
	//fmt.Print(listA)
	//fmt.Printf("\n")

	//times_task_1 := append_arrays_task_1(100000)
	//fmt.Printf("\nlinked list: %d, array %d\n", times_task_1[0], times_task_1[1])
	//times_task_2 := append_arrays_task_2(100000)
	//fmt.Printf("\nlinked list: %d, array %d\n", times_task_2[0], times_task_2[1])

	//fmt.Printf("\nlinked list: %d, array %d\n", times[0], times[1])
	//fmt.Printf("benchmarkscores: %d\n", bs_scores)
	benchmarker()

}

func benchmarker() {
	max_size := 100_000
	for j, i := 0, 1; i < max_size; i, j = i+50, j+1 {
		for t1, t2 := range append_arrays_task_1(i) {
			fmt.Printf("[%d, %d]\n", t1, t2)
		}
		//fmt.Printf("%d\n", bs_scores[i])
	}

}

func make_linked_list(size int) linkedList {
	ll := linkedList{}
	data_to_add := 1337357
	for i := 0; i < size; i++ {
		ll.append_node(data_to_add)
	}
	return ll
}
func append_arrays_task_1(amount_elements int) []time.Duration {
	times := make([]time.Duration, 2)

	//____linked lists____//
	first_list := make_linked_list(1000)
	second_list := make_linked_list(amount_elements) // vary this
	t0 := time.Now()
	first_list.append_node(second_list.head.data)
	times[0] = time.Since(t0)

	//____arrays____//
	arr := make([]int, amount_elements)
	data_to_add := 8008135
	t0 = time.Now()
	arr_newlen := make([]int, len(arr)+amount_elements)

	for i := 0; i < len(arr)-1; i++ {
		arr_newlen[i] = arr[i]
	}
	for j := 0; j < amount_elements-1; j++ {
		arr_newlen[j] = data_to_add
	}
	times[1] = time.Since(t0)

	return times
}
func append_arrays_task_2(amount_elements int) []time.Duration {
	times := make([]time.Duration, 2)

	//____linked lists____//
	first_list := make_linked_list(amount_elements) // vary this
	second_list := make_linked_list(1000)
	t0 := time.Now()
	first_list.append_node(second_list.head.data)
	times[0] = time.Since(t0)

	//____arrays___//
	arr := make([]int, amount_elements)
	data_to_add := 8008135
	t0 = time.Now()
	arr_newlen := make([]int, len(arr)+amount_elements)

	for i := 0; i < len(arr)-1; i++ {
		arr_newlen[i] = arr[i]
	}
	for j := 0; j < amount_elements-1; j++ {
		arr_newlen[j] = data_to_add
	}
	times[1] = time.Since(t0)

	return times
}

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
