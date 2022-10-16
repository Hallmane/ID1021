package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//testPriorityHeap()
	//linkedHeapBenchmarks(1000, 100000)
	ah := makeArrayHeap()
	fmt.Println(ah)
	ah.Push(12)
	ah.Push(21)
	ah.Push(7)
	ah.Push(5)
	fmt.Println(ah)
	ah.Pop()
	fmt.Println(ah)
}

func linkedHeapBenchmarks(loops_per_size int, amount_operations int) {
	j := 0
	for i := 10; i < amount_operations; i = i + 250 {
		rand.Seed(2)
		linearData := make([]time.Duration, loops_per_size)
		constantData := make([]time.Duration, loops_per_size)

		for j = 0; j < loops_per_size; j++ {
			constantData[j] = constantAddBenchmark(i)
			linearData[j] = linearAddBenchmark(i)
		}
		sort.Sort(ByDuration(linearData))
		sort.Sort(ByDuration(constantData))
		writeDatShit("datfiles/linearAdd", i, linearData[j/2])
		writeDatShit("datfiles/constantAdd", i, constantData[j/2])
	}
}

func linearAddBenchmark(loops int) time.Duration {
	lh := makeLinkedHeap(5)
	lh.Constant_Add(3)
	var t0 time.Time
	var t1 time.Time
	var s0 time.Duration
	var s1 time.Duration
	for i := 0; i < loops; i++ {
		t0 = time.Now()
		lh.Linear_Add(rand.Intn(1000))
		s0 = time.Since(t0)
	}
	for i := 0; i < loops-1; i++ {
		t1 = time.Now()
		lh.Constant_Remove()
		s1 = time.Since(t1)
	}
	return s0 + s1
}

func ConstantAddBenchmark(loops int) time.Duration {
	lh := makeLinkedHeap(5)
	lh.Constant_Add(3)
	var t0 time.Time
	var t1 time.Duration
	t0 = time.Now()
	for i := 0; i < loops; i++ {
		lh.Constant_Add(rand.Intn(1000))
	}
	for i := 0; i < loops-1; i++ {
		lh.Linear_Remove()
	}
	t1 = time.Since(t0)
	return t1

}
func constantAddBenchmark(loops int) time.Duration {
	lh := makeLinkedHeap(5)
	lh.Constant_Add(3)
	var t0 time.Time
	var t1 time.Time
	var s0 time.Duration
	var s1 time.Duration
	for i := 0; i < loops; i++ {
		t0 = time.Now()
		lh.Linear_Add(rand.Intn(1000))
		s0 = time.Since(t0)
	}
	for i := 0; i < loops-1; i++ {
		t1 = time.Now()
		lh.Constant_Remove()
		s1 = time.Since(t1)
	}
	return s0 + s1
}

func testPriorityHeap() {
	h := makeHeap(3)

	h.Add(5)
	h.Add(6)
	h.Add(10)
	h.Add(7)
	h.Add(11)
	h.Add(8)
	h.Add(4)

	fmt.Println(h.root, "old root")
	h.Remove()
	h.Push(3)
	fmt.Println(h.root, "new root")

}
