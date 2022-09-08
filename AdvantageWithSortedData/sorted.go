package main

import (
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
			continue
		}
	}
	return false
}

func binarySearch(array []int, key int) bool {

	var bot = 0
	var top = len(array) - 1
	var mid = (top - bot) / 2

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

// Benchmarking function for unsorted arrays
func unsortedArrayAggregator(arraySize int, runs int64, key int, values []int) ([]time.Duration, int64) {

	lapTimes := make([]time.Duration, runs)
	var averageTime time.Duration
	//keys := make([]int, arraySize)

	for i := 0; int64(i) < runs; i++ {
		t0 := time.Now()
		searchArrayZeroToN(values, key)
		t1 := time.Now().Sub(t0)
		//fmt.Printf("{%d}", key)

		lapTimes[i] = t1
		averageTime += t1
		//keys[i] = key

		rand.Seed(time.Now().UTC().UnixNano())
	}
	return lapTimes, int64(averageTime.Nanoseconds() / runs)
}

// Benchmarking function for sorted arrays
func sortedArrayAggregator(arraySize int, runs int64, key int, values []int) ([]time.Duration, int64) {

	lapTimes := make([]time.Duration, runs)
	var averageTime time.Duration
	//keys := make([]int, arraySize)

	for i := 0; int64(i) < runs; i++ {

		t0 := time.Now()
		searchArrayZeroToKey(values, key)
		t1 := time.Now().Sub(t0)
		//fmt.Printf("{%d}", key)

		lapTimes[i] = t1
		averageTime += t1
		//keys[i] = key

		rand.Seed(time.Now().UTC().UnixNano())

	}
	return lapTimes, int64(averageTime.Nanoseconds() / runs)
}

// Benchmarking function for binary search
func binarySearchArrayAggregator(arraySize int, runs int64, key int, values []int) ([]time.Duration, int64) {

	lapTimes := make([]time.Duration, runs)
	var averageTime time.Duration
	//keys := make([]int, arraySize)

	for i := 0; int64(i) < runs; i++ {

		t0 := time.Now()
		binarySearch(values, key)
		t1 := time.Now().Sub(t0)
		//fmt.Printf("{%d}", key)

		lapTimes[i] = t1
		averageTime += t1
		//keys[i] = key

		rand.Seed(time.Now().UTC().UnixNano())

	}
	return lapTimes, int64(averageTime.Nanoseconds() / runs)
}

func main() {
	arraySize := 10000

	runs := int64(100)

	rand.Seed(time.Now().UTC().UnixNano())
	key := rand.Intn(arraySize * arraySize)
	sortedArray := generateSortedArray(arraySize)
	_, avgTime := binarySearchArrayAggregator(arraySize, runs, key, sortedArray)
	writeDatShit("BBBBBBBBBBBBBinsortedArraySearch", arraySize, int(avgTime))

	//for arrSize := 32; arrSize < 1600; arrSize = arrSize * 2 {
	//	rand.Seed(time.Now().UTC().UnixNano())
	//	key := rand.Intn(arrSize * arrSize)
	//	sortedArray := generateSortedArray(arrSize)
	//	_, avgTime := binarySearchArrayAggregator(arrSize, runs, key, sortedArray)
	//	writeDatShit("BBBBBBBBBBBBBinsortedArraySearch", arrSize, int(avgTime))
	//}
}
