package main

import (
	"fmt"
	"math/rand"
	"sort"
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
	for i := 0; array[i] <= key; i++ {
		if array[i] == key {
			return true
		}
	}
	return false
}

func binarySearch(array []int, key int) bool {

	var bot = 0
	var top = len(array) - 1
	//var top = array[len(array)-1]
	var mid = (top - bot) / 2

	for {
		if array[mid] == key {
			return true
		}
		if array[mid] < key && mid < top-1 {
			bot = mid
			mid = (top + bot) / 2
			continue
		}
		if array[mid] > key && mid > bot+1 {
			top = mid
			mid = (top + bot) / 2
			continue
		}
		break
	}
	return false
}

func generateUnsortedArray(size int) []int {
	return rand.Perm(size - 1)
}

// create a large sorted array and remove random values until
// it is of the same size as "smol"
func generateSortedArray(size int) ([]int, int) {
	rand.Seed(time.Now().UTC().UnixNano())
	smol := make([]int, size)
	large := make([]int, size+10000)
	randIndex := make([]int, len(large)-len(smol))

	for i := 0; i < len(large); i++ { //prepping large values
		large[i] = i * i
	}

	for k := range randIndex { // prepping random values
		randIndex[k] = rand.Intn(len(large))
	}

	for count := 0; len(large)-count > len(smol); count++ { // removing random values from large until equal size as smol
		large[randIndex[count]] = 0
	}

	for i, j := 0, 0; i < len(large); i, j = i+1, j+1 {
		for large[j] == 0 {
			j++
			fmt.Printf("%d\t%d\n \n___\n", j, large[j])
			continue
		}
		smol[i] = large[j]
	}
	return smol, len(large)
}

// Benchmarking function for unsorted arrays
func unsortedArrayAggregator(arraySize int, runs int64, key int, values []int) ([]time.Duration, int64) {

	lapTimes := make([]time.Duration, runs)
	var averageTime time.Duration

	for i := 0; int64(i) < runs; i++ {
		t0 := time.Now()
		searchArrayZeroToN(values, key)
		t1 := time.Now().Sub(t0)

		lapTimes[i] = t1
		averageTime += t1

		rand.Seed(time.Now().UTC().UnixNano())
	}
	return lapTimes, int64(averageTime.Nanoseconds() / runs)
}

// Benchmarking function for sorted arrays
func sortedArrayAggregator(arraySize int, runs int64, key int, values []int) ([]time.Duration, int64) {

	lapTimes := make([]time.Duration, runs)
	var averageTime time.Duration

	for i := 0; int64(i) < runs; i++ {

		t0 := time.Now()
		searchArrayZeroToKey(values, key)
		t1 := time.Now().Sub(t0)

		lapTimes[i] = t1
		averageTime += t1

		rand.Seed(time.Now().UTC().UnixNano())

	}
	return lapTimes, int64(averageTime.Nanoseconds() / runs)
}

// Benchmarking function for binary search
func binarySearchArrayAggregator(runs int64, key int, array []int) ([]time.Duration, int64) {

	lapTimes := make([]time.Duration, runs)
	var averageTime time.Duration

	for i := 0; int64(i) < runs; i++ {

		t0 := time.Now()
		binarySearch(array, key)
		t1 := time.Now().Sub(t0)

		lapTimes[i] = t1
		averageTime += t1

		rand.Seed(time.Now().UTC().UnixNano())

	}
	return lapTimes, int64(averageTime.Nanoseconds() / runs)
}

func findDuplicatesBinarySearch(runs int, alpha []int, beta []int) (int, []time.Duration) {
	duplicates := 0
	clockedTimes := make([]time.Duration, len(alpha))

	for i := range alpha {
		t0 := time.Now()
		if binarySearch(beta, alpha[i]) == true {
			t1 := time.Now().Sub(t0)
			clockedTimes[i] = t1
			duplicates++
		} else {
			continue
		}
	}
	return duplicates, clockedTimes
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	runs := 1
	arraySize := 100000000

	alpha, _ := generateSortedArray(arraySize)
	rand.Seed(time.Now().UTC().UnixNano())
	beta, _ := generateSortedArray(arraySize)

	duplicates, clockedTimes := findDuplicatesBinarySearch(runs, alpha, beta)

	sort.Slice(clockedTimes, func(i, j int) bool { // sorting the times to be able to get median
		return clockedTimes[i] < clockedTimes[j]
	})

	fmt.Printf("%d duplicates\narray length: %d \nratio: %d\nAverage duration to locate key: %d\n", duplicates, arraySize, (arraySize)/duplicates, clockedTimes[arraySize/2])

	//
	//
	//
	//
	//
	//            // a first try// 			//
	//runs := int64(10)
	//for arrSize := 100; arrSize < 16000000; arrSize += 25000 {
	//	rand.Seed(time.Now().UTC().UnixNano())
	//	//sortedArray, maxKey := generateSortedArray(arrSize)
	//	unsortedArray := generateUnsortedArray(arrSize)
	//	maxKey := arrSize
	//	key := rand.Intn(maxKey)
	//	_, avgTime := unsortedArrayAggregator(arrSize, runs, key, unsortedArray)
	//	writeDatShit("unsortedArray", arrSize, int(avgTime))
	//}
}
