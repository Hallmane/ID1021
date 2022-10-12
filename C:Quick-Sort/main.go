package main

import (
	"fmt"
	"math"
	"math/rand"
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

func benchmark() {
	rand.Seed(time.Now().UnixNano())
	for i := 2; i < 10000000; i = i + 25000 {
		writeDatShit("arrTest", i, testArray(i, int(math.Floor(math.Sqrt(float64(i))))))
		writeDatShit("llTest", i, testLink(i, int(math.Floor(math.Sqrt(float64(i))))))

	}
}
