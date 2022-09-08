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

// Search until value of array[i] <= key [sorted array]
func searchArrayZeroToKey(array []int, key int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == key {
			return true
		}
		if array[i] > key {
			break
		}
	}
	return false
}

func binarySearch(array []int, key int) bool {

	bot := 0
	top := len(array) - 1
	mid := (top - bot) / 2

	for {
		if array[mid] == key {
			return true
		}
		if array[mid] < key && mid < top {
			bot = mid
			mid = (top + bot) / 2
			continue
		}
		if array[mid] > key && mid > bot {
			top = mid
			mid = (top + bot) / 2
			continue
		}
		return false
	}

}

//	func oldgenerateUnsortedArray(size int) []int {
//		array := make([]int, size)
//
//		for i := 0; i < size; i++ {
//			array[i] = rand.Intn(size - 1)
//			//fmt.Printf("[%d], ", array[i])
//		}
//		return array
//	}

func generateUnsortedArray(size int) []int {
	return rand.Perm(size - 1)
}

func generateSortedArray(size int) []int {
	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = i
	}
	return array
}

// Benchmarking function for arrays of size arraySize
func unsortedArrayAggregator(arraySize int, runs int64, key int, values []int) ([]time.Duration, int64, []int) {

	lapTimes := make([]time.Duration, runs)
	var averageTime time.Duration
	keys := make([]int, arraySize)

	for i := 0; int64(i) < runs; i++ {
		t0 := time.Now()
		searchArrayZeroToN(values, key)
		t1 := time.Now().Sub(t0)
		//fmt.Printf("{%d}", key)

		lapTimes[i] = t1
		averageTime += t1
		keys[i] = key

		rand.Seed(time.Now().UTC().UnixNano())
	}
	return lapTimes, int64(averageTime.Nanoseconds() / runs), keys
}
func sortedArrayAggregator(arraySize int, runs int64, key int, values []int) ([]time.Duration, int64, []int) {

	lapTimes := make([]time.Duration, runs)
	var averageTime time.Duration
	keys := make([]int, arraySize)

	for i := 0; int64(i) < runs; i++ {

		t0 := time.Now()
		searchArrayZeroToKey(values, key)
		t1 := time.Now().Sub(t0)
		//fmt.Printf("{%d}", key)

		lapTimes[i] = t1
		averageTime += t1
		keys[i] = key

		rand.Seed(time.Now().UTC().UnixNano())

	}
	return lapTimes, int64(averageTime.Nanoseconds() / runs), keys
}
func binarySearchArrayAggregator(arraySize int, runs int64, key int, values []int) ([]time.Duration, int64, []int) {

	lapTimes := make([]time.Duration, runs)
	var averageTime time.Duration
	keys := make([]int, arraySize)

	for i := 0; int64(i) < runs; i++ {

		t0 := time.Now()
		binarySearch(values, key)
		t1 := time.Now().Sub(t0)
		//fmt.Printf("{%d}", key)

		lapTimes[i] = t1
		averageTime += t1
		keys[i] = key

		rand.Seed(time.Now().UTC().UnixNano())

	}
	return lapTimes, int64(averageTime.Nanoseconds() / runs), keys
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	runs := int64(1000)
	sizeOfArray := 16000000
	key := rand.Intn(sizeOfArray)

	//arrayUnsorted := generateUnsortedArray(sizeOfArray)
	arraySorted := generateSortedArray(sizeOfArray)

	//for k, v := range array {
	//	fmt.Printf("[index: %d | value: %d]\n", k, v)
	//}
	//fmt.Printf("%d\n", key)
	//searchArrayZeroToKey(array, key)

	//testUnsorted, avgUnsorted, keysUnsorted := unsortedArrayAggregator(sizeOfArray, runs, key, arrayUnsorted)
	//_, avgUnsorted, _ := unsortedArrayAggregator(sizeOfArray, runs, key, arrayUnsorted)
	//testSorted, avgSorted, keysSorted := sortedArrayAggregator(sizeOfArray, runs, key, arraySorted)
	_, avgSorted, _ := binarySearchArrayAggregator(sizeOfArray, runs, key, arraySorted)

	//fmt.Printf("________unsorted array________\n")
	//for k, v := range testUnsorted {
	//	fmt.Printf("Key[%d] found after %s\n", keysUnsorted[k], v)
	//}
	//fmt.Printf("For an array with size %d, the average duration for %d runs was %dns\n", sizeOfArray, runs, avgUnsorted)
	//fmt.Printf("_______________________________________________\n")

	//for k, v := range testSorted {
	//	fmt.Printf("Key[%d] found after %s\n", keysSorted[k], v)
	//}
	fmt.Printf("For a sorted array with size %d, the average duration for %d runs was %dns\n", sizeOfArray, runs, avgSorted)
	fmt.Printf("________________________END____________________________\n")
	//fmt.Printf("For an unsorted array with size %d, the average duration for %d runs was %dns\n", sizeOfArray, runs, avgUnsorted)
	//fmt.Printf("________________________END____________________________\n")
}
