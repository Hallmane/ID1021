package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randArrayGen(size int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Perm(size - 1)
}

// done
func selectionSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		smallest := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[smallest] {
				smallest = j
			}
		}
		if smallest != i {
			copy := arr[smallest]
			arr[smallest] = arr[i]
			arr[i] = copy
		}
	}
	return arr
}

func insertooor(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			copy := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = copy
			j--
		}
	}
	return arr
}

func mergeSort(arr []int) []int {

	if len(arr) == 1 {
		return arr
	}

	arr1 := arr[:(len(arr))/2]
	arr2 := arr[(len(arr))/2:]

	arr1 = mergeSort(arr1)
	arr2 = mergeSort(arr2)

	return mörge(arr1, arr2)
}
func mörge(arra []int, arrb []int) []int {

	fffinalll := make([]int, len(arra)+len(arrb))
	a, b, m := 0, 0, 0

	for ; a < len(arra) && b < len(arrb); m++ {
		if arra[a] < arrb[b] {
			fffinalll[m] = arra[a]
			a++
		} else {
			fffinalll[m] = arrb[b]
			b++
		}
	}
	for len(arra) > a {
		fffinalll[m] = arra[a]
		a++
		m++
	}
	for len(arrb) > b {
		fffinalll[m] = arrb[b]
		b++
		m++
	}
	return fffinalll
}

func timer(array_size int, which int) time.Duration {
	unsortedArray := randArrayGen(array_size)

	var t1 time.Duration

	if which == 1 {
		t0 := time.Now()
		selectionSort(unsortedArray)
		t1 = time.Now().Sub(t0)
		writeDatShit("selection_sort_1", array_size, int(t1))
	}
	if which == 2 {
		t0 := time.Now()
		insertooor(unsortedArray)
		t1 = time.Now().Sub(t0)
		writeDatShit("insertion_sort_1", array_size, int(t1))
	}
	if which == 3 {
		t0 := time.Now()
		mergeSort(unsortedArray)
		t1 = time.Now().Sub(t0)
		writeDatShit("merge_sort_3", array_size, int(t1))
	}
	return t1
}

func bencher(loop_incrementer int) []time.Duration {
	res := make([]time.Duration, 1000)
	for i := 1000; i < 1_000_000; i += loop_incrementer {
		timer(i, 3) //ugly
	}
	for i := range res {
		fmt.Printf("Time: %d\n", i)
	}

	return res
}

func main() {
	//array_size := 1000
	//which_algo := 1
	//result := timer(array_size, which_algo)

	bencher(1000)
	//fmt.Printf("Time (size %d): %d\n", array_size, result)

}
