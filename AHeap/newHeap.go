package main

import "fmt"

type Heap struct {
	root *HeapNode
}

type HeapNode struct {
	priority, children int
	left, right        *HeapNode
}

func (heap *Heap) Add(priority int) {

	current := heap.root
	var higher_number int
	new_node := makeHeapNode(priority)

	if current == nil {
		heap.root = new_node
		return
	}
	if current.priority > priority {
		higher_number = current.priority
		current.priority = priority
		new_node.priority = higher_number
	}

	heap.root.RecursiveAdd(new_node)
}

func (node *HeapNode) RecursiveAdd(new_node *HeapNode) {
	node.children += new_node.children + 1
	if node.left == nil {
		node.left = new_node
		return
	}
	if node.right == nil {
		node.right = new_node
		return
	}
	target := node._Least_children()

	if target.priority > new_node.priority {
		higher_number := target.priority
		target.priority = new_node.priority
		new_node.priority = higher_number
	}
	target.RecursiveAdd(new_node)
}

func (heap *HeapNode) RecursiveCheck(new_priority int) {

}

func (heap *Heap) Push(increment int) {
	heap.root.priority += increment
	steps := 0
	var target *HeapNode
	if heap.root.left == nil {
		target = heap.root.right
	} else if heap.root.right == nil {
		target = heap.root.left
	} else {
		if heap.root.left.priority <= heap.root.right.priority {
			target = heap.root.left
		} else {
			target = heap.root.right
		}
	}
	if target == nil {
		steps++
		return
	}
	if target.priority < heap.root.priority {
		temp := target.priority
		target.priority = heap.root.priority
		heap.root.priority = temp
	}
	target.push_down(&steps)
}

func (node *HeapNode) push_down(steps *int) *int {
	*steps++
	var target *HeapNode
	if node.left == nil {
		target = node.right
	} else if node.right == nil {
		target = node.left
	} else {
		if node.left.priority <= node.right.priority {
			target = node.left
		} else {
			target = node.right
		}
	}
	if target == nil {
		return steps
	}
	if target.priority < node.priority {
		temp := target.priority
		target.priority = node.priority
		node.priority = temp
	}
	return target.push_down(steps)
}

func swap_priority(node_a *HeapNode, node_b *HeapNode) {
	temp_val_a := node_a.priority
	node_a.priority = node_b.priority
	node_b.priority = temp_val_a
}

func (heap *Heap) Remove() *HeapNode {
	var new_root *HeapNode
	if heap.root == nil {
		return nil
	}

	if heap.root.left == nil && heap.root.right == nil {
		return heap.root
	}
	if heap.root.left == nil && heap.root.right != nil {
		new_root = heap.root.right
		old_root := heap.root
		heap.root = new_root
		return old_root
	}
	if heap.root.left != nil && heap.root.right == nil {
		new_root = heap.root.left
		old_root := heap.root
		heap.root = new_root
		return old_root
	} else {
		new_root = heap.root.MostImportant()
		branch_to_fix := heap.root.LeastImportant()
		old_root := heap.root
		heap.root = new_root
		new_root.RecursiveAdd(branch_to_fix)
		return old_root
	}
}

func (node *HeapNode) _NextAvailable() *HeapNode {
	if node.left != nil {
		fmt.Println("(_NextAvailable, if) current = ", node)
		return node.left
	} else {
		if node.right != nil {
			fmt.Println("(_NextAvailable, else) current = ", node)
			return node.right
		}
		return node
	}
}
