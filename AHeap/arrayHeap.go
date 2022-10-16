package main

type ArrayHeap struct {
	heap []int
}

func makeArrayHeap() *ArrayHeap {
	return &ArrayHeap{heap: make([]int, 0)}
}

func left_child_index(index int) int {
	return 2*index + 1
}

func right_child_index(index int) int {
	return 2*index + 2
}

func parent_index(index int) int {
	return (index - 1) / 2
}

func (ah *ArrayHeap) Swap(i, j int) {
	ah.heap[i], ah.heap[j] = ah.heap[j], ah.heap[i]
}

func (ah *ArrayHeap) Push(priority int) {
	ah.heap = append(ah.heap, priority)
	ah.Bubble(len(ah.heap) - 1)
}

func (ah *ArrayHeap) Pop() int {
	if len(ah.heap) == 0 {
		panic("heap is empty")
	}
	to_be_popped := ah.heap[0]
	ah.Swap(0, len(ah.heap)-1)
	ah.heap = ah.heap[:len(ah.heap)-1]
	ah.Sink(0)
	return to_be_popped
}

func (ah *ArrayHeap) Sink(index int) {
	l, r := left_child_index(index), right_child_index(index)
	highest := index
	if r < len(ah.heap) && ah.Compare(highest, r) {
		highest = r
	}
	if l < len(ah.heap) && ah.Compare(index, l) {
		highest = l
	}
	if highest != index {
		ah.Swap(index, highest)
		ah.Sink(highest)
	}
}

func (ah ArrayHeap) Compare(a, b int) bool {
	return ah.heap[a] > ah.heap[b]
}

func (ah *ArrayHeap) Bubble(index int) {
	right_place := false
	for !right_place {
		if ah.heap[parent_index(index)] <= ah.heap[index] {
			right_place = true
		} else {
			ah.Swap(parent_index(index), index)
			index = parent_index(index)
			ah.Bubble(index)
			return
		}
	}
}
