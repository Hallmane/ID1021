package main

func quicksort(arr []int, start int, end int) []int {
	if start < end {
		p := partition(arr, start, end)
		quicksort(arr, start, p-1)
		quicksort(arr, p+1, end)
	}
	return arr
}

func partition(array []int, start int, end int) int {
	pivot := array[end]
	li := start
	ri := end

	for li < ri {
		for array[li] < pivot && li < end {
			li++
		}
		for array[ri] >= pivot && ri > 0 {
			ri--
		}
		if li < ri {
			array[li], array[ri] = array[ri], array[li]
		} else {
			array[li], array[end] = array[end], array[li]
		}
	}
	return li
}
