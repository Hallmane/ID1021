package main

import (
	"fmt"
	"math/rand"
	"time"
)

// searches through an array 0->n
func searchArrayZeroToN(array []int, key int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == key {
			return true
		}
	}
	return false
}

func generateRandomArray(size int) []int {
	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = rand.Intn(size - 1)
	}
	return array
}

// Benchmarking function for arrays of size arraySize
func unsortedArrayAggregator(arraySize int, runs int64) ([]time.Duration, int64, []int) {

	lapTimes := make([]time.Duration, runs)
	var averageTime time.Duration
	keys := make([]int, arraySize)

	for i := 0; int64(i) < runs; i++ {
		key := rand.Intn(arraySize)
		array := generateRandomArray(arraySize)

		t0 := time.Now()
		searchArrayZeroToN(array, key)
		t1 := time.Now().Sub(t0)

		lapTimes[i] = t1
		averageTime += t1
		keys[i] = key

		rand.Seed(time.Now().UTC().UnixNano())

	}
	return lapTimes, int64(averageTime.Microseconds() / runs), keys
}

func main() {

	runs := int64(50)
	sizeOfArray := 1000
	test, avg, keys := unsortedArrayAggregator(sizeOfArray, runs)
	for k, v := range test {
		fmt.Printf("Key[%d] found after %s\n", keys[k], v)
	}
	fmt.Printf("For an array with size %d, the average duration for %d runs was %dµs\n", sizeOfArray, runs, avg)
	fmt.Printf("_______________________________________________\n")
	sizeOfArray = 10000
	test, avg, keys = unsortedArrayAggregator(sizeOfArray, runs)
	for k, v := range test {
		fmt.Printf("Key[%d] found after %s\n", keys[k], v)
	}
	fmt.Printf("For an array with size %d, the average duration for %d runs was %dµs\n", sizeOfArray, runs, avg)
	fmt.Printf("_______________________________________________\n")
	sizeOfArray = 100000
	test, avg, keys = unsortedArrayAggregator(sizeOfArray, runs)
	for k, v := range test {
		fmt.Printf("Key[%d] found after %s\n", keys[k], v)
	}
	fmt.Printf("For an array with size %d, the average duration for %d runs was %dµs\n", sizeOfArray, runs, avg)
	fmt.Printf("_______________________________________________\n")
	sizeOfArray = 1000000
	test, avg, keys = unsortedArrayAggregator(sizeOfArray, runs)
	for k, v := range test {
		fmt.Printf("Key[%d] found after %s\n", keys[k], v)
	}
	fmt.Printf("For an array with size %d, the average duration for %d runs was %dµs\n", sizeOfArray, runs, avg)
	fmt.Printf("_______________________________________________\n")
	sizeOfArray = 10000000
	test, avg, keys = unsortedArrayAggregator(sizeOfArray, runs)
	for k, v := range test {
		fmt.Printf("Key[%d] found after %s\n", keys[k], v)
	}
	fmt.Printf("For an array with size %d, the average duration for %d runs was %dµs\n", sizeOfArray, runs, avg)
	fmt.Printf("_______________________________________________\n")
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
