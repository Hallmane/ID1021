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
		mid = bot + (top-bot)/2
		//fmt.Printf("key: %d\n", key)
		if array[mid] == key {
			fmt.Printf("found\n")
			return true
		}
		if array[mid] < key && mid < top-1 {
			//fmt.Printf("bot, top | %d, %d\n", bot, top)
			//bot = mid
			//mid = (bot + top) / 2
			bot = mid + 1
			continue
		}
		if array[mid] > key && mid > bot+1 {
			//fmt.Printf("bot, top | %d, %d\n", bot, top)
			///top = mid
			///mid = (bot + top) / 2
			top = mid - 1
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

	for i, j := 0, 0; i < len(smol)-1; i, j = i+1, j+1 {
		for large[j] == 0 {
			j++
		}
		smol[i] = large[j]
	}
	return smol, len(large)
}

// Benchmarking function for unsorted arrays
func unsortedArrayAggregator(runs int64, key int, values []int) ([]time.Duration, int64) {

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
func sortedArrayAggregator(runs int, key int, values []int) []time.Duration {

	lapTimes := make([]time.Duration, runs)

	for i := 0; i < runs; i++ {

		t0 := time.Now()
		searchArrayZeroToKey(values, key)
		t1 := time.Now().Sub(t0)

		lapTimes[i] = t1

		rand.Seed(time.Now().UTC().UnixNano())

	}
	return lapTimes
}

// Benchmarking function for binary search
func binarySearchArrayAggregator(runs int, key int, array []int) []time.Duration {

	lapTimes := make([]time.Duration, runs)
	var averageTime time.Duration

	for i := 0; i < runs; i++ {

		t0 := time.Now()
		binarySearch(array, key)
		t1 := time.Now().Sub(t0)

		lapTimes[i] = t1
		averageTime += t1

		rand.Seed(time.Now().UTC().UnixNano())

	}
	return lapTimes
}

func findDuplicatesBinarySearch(runs int, arr1 []int, arr2 []int) []time.Duration {
	clockedTimes := make([]time.Duration, len(arr1))

	for i := range arr1 {
		t0 := time.Now()
		if binarySearch(arr2, arr1[i]) == true {
			t1 := time.Now().Sub(t0)
			clockedTimes[i] = t1
		} else {
			fmt.Printf("ee")
			continue
		}
	}
	return clockedTimes
}

func evenFurtherBeyond(runs int, piff []int, puff []int) time.Duration {
	var clockedTime time.Duration

	for i, j := 0, 0; i < (len(piff)-1) && j < (len(puff)-1); i, j = i+1, j+1 {
		t0 := time.Now()

		if piff[i] == puff[j] {
			i++
			j++
		}
		if piff[i+1] < puff[j] {
			i++
		}
		if piff[i+1] > puff[j] {
			j++
		}
		clockedTime = time.Now().Sub(t0)
	}
	return clockedTime
}
func GenerateSortedList(sliceSize int) []int {
	rand.Seed(time.Now().UTC().UnixNano())

	slice := make([]int, sliceSize)

	next := 0

	for i := 0; i < sliceSize; i++ {
		next += rand.Intn(10) + 1
		slice[i] = next
	}

	return slice
}

func arraySewing(arr1 []int, arr2 []int) time.Duration {
	var (
		i  int = 0
		j  int = 0
		t0 time.Time
	)

	t0 = time.Now()
	for i < len(arr1) && j < len(arr2) {
		if arr2[j] == arr1[i] {
			//duplicate found
			j++
			i++
			continue
		}
		if arr2[j] < arr1[i] {
			j++
			continue
		}
		if arr2[j] > arr1[i] {
			i++
			continue
		}
	}
	return time.Since(t0)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	//runs := 1
	//arraySize := 1000000
	//piff, _ := generateSortedArray(arraySize)
	//rand.Seed(time.Now().UTC().UnixNano())
	//puff, _ := generateSortedArray(arraySize)

	//clockedTime := evenFurtherBeyond(runs, piff, puff)
	////sort.Slice(clockedTimes, func(i, j int) bool { // sorting the times// to be able to get median
	////	return clockedTimes[i] < clockedTimes[j]
	////})

	//fmt.Printf("time %d\n", clockedTime)

	//arraySize := 100000
	//for arrSize := 100; arrSize < 16000000; arrSize += 25000 {
	//	alpha, _ := generateSortedArray(arrSize)
	//	rand.Seed(time.Now().UTC().UnixNano())
	//	beta, _ := generateSortedArray(arrSize)
	//	_, clockedTimes := findDuplicatesBinarySearch(runs, alpha, beta)

	//	sort.Slice(clockedTimes, func(i, j int) bool { // sorting the times// to be able to get median
	//		return clockedTimes[i] < clockedTimes[j]
	//	})

	//	fmt.Printf("array length: %d \nAverage duration to locate key: %dns\n", arrSize, clockedTimes[arrSize/2])
	//	//writeDatShit("doublyBinaryWhammy", arrSize, int(clockedTimes[arrSize/2]))
	//}

	//fmt.Printf("%d duplicates\narray length: %d \nratio: %d\nAverage duration to locate key: %dns\n", duplicates, arraySize, (arraySize)/duplicates, clockedTimes[arraySize/2])

	runs := 100
	arraySize := 1000
	piff := GenerateSortedList(arraySize)

	key := rand.Intn(arraySize)

	clockedTimes := binarySearchArrayAggregator(runs, key, piff)

	for i, v := range clockedTimes {
		fmt.Printf("i: %d\nv: %d\n\n", i, v)
	}

	// a first try// 			//
	//runs := 100
	//for arrSize := 100; arrSize < 64000000; arrSize += 250000 {
	//	rand.Seed(time.Now().UTC().UnixNano())
	//	//sortedArray, maxKey := generateSortedArray(arrSize)
	//	unsortedArray1, _ := generateSortedArray(arrSize)
	//	unsortedArray2, _ := generateSortedArray(arrSize)
	//	//maxKey := arrSize
	//	//key := rand.Intn(maxKey)
	//	//clockedTimes := sortedArrayAggregator(runs, key, unsortedArray)
	//	clockedTimes := findDuplicatesBinarySearch(runs, unsortedArray1, unsortedArray2)
	//	sort.Slice(clockedTimes, func(i, j int) bool {
	//		return clockedTimes[i] < clockedTimes[j]
	//	})
	//	//writeDatShit("binarySearch", arrSize, int(clockedTimes[runs/2]))
	//	fmt.Printf("array length: %d \tAverage duration to locate key: %dns\n", arrSize, clockedTimes[runs/2])
	//}
}

//
