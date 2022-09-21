package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type ByDuration []time.Duration

func (a ByDuration) Len() int           { return len(a) }
func (a ByDuration) Less(i, j int) bool { return a[i].Nanoseconds() < a[j].Nanoseconds() }
func (a ByDuration) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func main() {
	benchmarker(100_000)
}

// writes median times of linked list and arrays to individual files
func benchmarker(max_size int) {
	go array_benchmarker(max_size)
	linked_list_benchmarker(max_size)
}

func linked_list_benchmarker(max_size int) {
	for i := 50; i < max_size; i += 100 {
		writeDatShit("linkedListBenchmark", i, linked_list_for_each_size(i))
	}
}

func array_benchmarker(max_size int) {
	for i := 50; i < max_size; i += 100 {
		writeDatShit("arrayBenchmarker", i, arrays_for_each_size(i))
	}
}

// ____linked lists____//
func linked_list_for_each_size(amount_elements int) time.Duration {
	times_per_i := make([]time.Duration, 3)

	first_list := make_linked_list(1000)
	second_list := make_linked_list(amount_elements)
	for i := 0; i < 3; i++ {
		t0 := time.Now()
		first_list.append_node(second_list.head.data)
		times_per_i[i] = time.Since(t0)
	}
	sort.Sort(ByDuration(times_per_i))
	return times_per_i[1]

}

// ____arrays____//
func arrays_for_each_size(amount_elements int) time.Duration {
	times_per_i := make([]time.Duration, 100)
	data_to_add := 8008135

	for k := 0; k < 100; k++ {
		arr := make([]int, amount_elements)
		t0 := time.Now()
		arr_newlen := make([]int, len(arr)+amount_elements)

		for i := 0; i < len(arr)-1; i++ {
			arr_newlen[i] = arr[i]
		}
		for j := 0; j < amount_elements-1; j++ {
			arr_newlen[j] = data_to_add
		}
		times_per_i[k] = time.Since(t0)
	}

	sort.Sort(ByDuration(times_per_i))
	return times_per_i[50]

}

// old
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

func make_linked_list(size int) linkedList {
	ll := linkedList{}
	data_to_add := 1337357
	for i := 0; i < size; i++ {
		ll.append_node(data_to_add)
	}
	return ll
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
