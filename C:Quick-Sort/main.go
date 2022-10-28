package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//printArrayTest(1000)
	benchmark()
}

func testArray(size int, valRange int) time.Duration {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(valRange)
	}

	t0 := time.Now()
	quicksort(array, 0, len(array)-1)
	return time.Since(t0)
}

func testArrayPerLoop(size int, valRange int, loops int) time.Duration {
	times := make([]time.Duration, size)
	array := make([]int, size)
	for i := 0; i < loops/2; i++ {
		for j := 0; j < size; j++ {
			array[j] = rand.Intn(valRange)
		}

		t0 := time.Now()
		quicksort(array, 0, len(array)-1)
		times[i] = time.Since(t0)
	}
	sort.Sort(ByDuration(times))
	return times[size/2]
}

func printArrayTest(size int) {
	array := rand.Perm(size)
	fmt.Println(array)
	quicksort(array, 0, len(array)-1)
	fmt.Println(array)
}

func testLink(size int, nodeRange int) time.Duration {
	dll := make_linkedlist(size, nodeRange)
	t0 := time.Now()
	dll.quick_sort_LL(dll.start, dll.end)
	return time.Since(t0)
}

func testLinkPerLoop(size int, nodeRange int, loops int) time.Duration {
	times := make([]time.Duration, size)
	for i := 0; i < loops/2; i++ {
		dll := make_linkedlist(size, nodeRange)
		t0 := time.Now()
		dll.quick_sort_LL(dll.start, dll.end)
		times[i] = time.Since(t0)
	}
	sort.Sort(ByDuration(times))
	return times[size/2]
}

func benchmark() {
	rand.Seed(time.Now().UnixNano())
	for i := 2; i < 100000; i = i + 2500 {
		writeDatShit("arrRedo", i, testArrayPerLoop(i, i, i))
		writeDatShit("LLRedo", i, testLinkPerLoop(i, i, i))
	}
}
