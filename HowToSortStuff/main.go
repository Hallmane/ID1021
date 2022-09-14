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
				fmt.Printf("switched %d with the smaller %d\n", arr[smallest], arr[j])
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
	//sortedarray := selectionSort(randArrayGen(100000))
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

func main() {

	unsortedArray := randArrayGen(1000)
	fmt.Print("[")
	for i := range unsortedArray {
		fmt.Printf("%d ", unsortedArray[i])
	}
	fmt.Println("]\n")

	//sortedarray := insertooor(unsortedArray)
	//sortedarray := selectionSort(unsortedArray)
	sortedarray := mergeSort(unsortedArray)

	fmt.Print("[")
	for i := range sortedarray {
		fmt.Printf("%d ", sortedarray[i])
	}
	fmt.Println("]")
}
