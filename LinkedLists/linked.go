package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type ByDuration []time.Duration

// sorting
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
	ll := linkedList{}
	ll.push(1337)
	ll.push(2)
	ll.push(3)
	ll.push(4)
	ll.push(5)
	ll.push(6)
	ll.push(7)
	ll.push(8)
	ll.push(9)
	ll.push(10)
	ll.push(11)
	fmt.Print(ll)
	ll.pop()
	fmt.Println("")
	fmt.Print(ll)
	ll.pop()
	fmt.Println("")
	fmt.Print(ll)
	ll.pop()
	fmt.Println("")
	fmt.Print(ll)
	ll.pop()
	fmt.Println("")
	fmt.Print(ll)
	ll.pop()
	fmt.Println("")
	fmt.Print(ll)
	ll.pop()
	fmt.Println("")
	fmt.Print(ll)
}

// writes median times of linked list and arrays to individual files
func benchmarker(max_size int) {
	//array_benchmarker(max_size)
	linked_list_benchmarker(max_size)
}

func linked_list_benchmarker(max_size int) {
	for i := 50; i < max_size; i += 1000 {
		writeDatShit("ass", i, linked_list_for_each_size(i))
		//linked_list_benchmarker(i)
		//fmt.Printf("%d\n", linked_list_for_each_size_task_1(i))
	}
}

func array_benchmarker(max_size int) {
	for i := 50; i < max_size; i += 100 {
		writeDatShit("arrayBenchmarkerTwo", i, arrays_for_each_size(i))
	}
}

// ____linked lists____//
func linked_list_for_each_size(amount_elements int) time.Duration {
	times_per_i := make([]time.Duration, 3)
	n := make_node(1337)

	for i := 0; i < 3; i++ {
		//_second_list := make_linked_list(make_node(i), 1000)
		t0 := time.Now()
		make_linked_list(n, amount_elements)
		times_per_i[i] = time.Since(t0)
		//first_list.append_data(second_list.head.data)
	}
	sort.Sort(ByDuration(times_per_i))
	return times_per_i[1]
}

// ____arrays____ //
func arrays_for_each_size(amount_elements int) time.Duration {
	times_per_i := make([]time.Duration, 100)
	data_to_add := 8008135

	for k := 0; k < 100; k++ {

		arr := make([]int, amount_elements)
		for j := 0; j < len(arr)-1; j++ {
			arr[j] = data_to_add
		}

		t0 := time.Now()
		arr_newlen := make([]int, len(arr)+amount_elements)
		for i := 0; i < len(arr)-1; i++ {
			arr_newlen[i] = arr[i]
		}
		for j := 0; j < amount_elements-1; j++ {
			arr_newlen[j] = data_to_add
		}
		times_per_i[k] = time.Since(t0)
		//fmt.Printf("%d\n", arr_newlen[len(arr_newlen)/2])
	}
	sort.Sort(ByDuration(times_per_i))
	return times_per_i[50]
}

func make_linked_list(n node, size int) linkedList {
	ll := linkedList{}
	for i := 0; i < size; i++ {
		ll.append_node(n)
	}
	return ll
}
func _old_make_linked_list(size int) linkedList {
	ll := linkedList{}
	data_to_add := 1337357
	for i := 0; i < size; i++ {
		ll.append_data(data_to_add)
	}
	return ll
}

func make_node(data int) node {
	return node{data: data, next: nil}
}

func (ll *linkedList) pop() int {
	var returnVal = ll.head.data

	if ll.head.next != nil {
		newHead := ll.head.next
		returnVal := ll.head.data
		ll.head = newHead
		return returnVal
	}
	return returnVal
}

func (ll *linkedList) push(data int) {
	newNode := new(node)
	newNode.data = data

	if ll.head == nil {
		ll.head = newNode
	} else {
		oldNext := ll.head.next
		ll.head.next = newNode
		newNode.next = oldNext

	}
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

// takes a linked list and appends varying amounts of elements
func (ll *linkedList) append_node(n node) {

	if ll.head == nil {
		ll.head = &n
	} else {
		current := ll.head
		for ; current.next != nil; current = current.next {
		}
		current.next = &n
	}
}

func (ll *linkedList) append_data(data int) {
	newNode := new(node)
	newNode.data = data

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

func (ll linkedList) String() string {
	sb := strings.Builder{}

	for iter := ll.head; iter != nil; iter = iter.next {
		sb.WriteString(fmt.Sprintf("%d-> ", iter.data))
	}
	return sb.String()
}
