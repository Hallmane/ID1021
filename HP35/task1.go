/*
	RPN stack
	4 |   ||   ||   ||   ||   ||   |
	3 |   ||   ||   ||   ||   ||   |
	2 |   || 3 || 3 ||   ||10 ||   |
	1 | 3 || 3 || 7 ||10 || 6 ||60 |  <-- display
	    3   [e]   7    +    6    *

	hp35 € {
	    instruction array (RPN) <- type Item structA
	    instruction pointer
	    stack                   <- .Push()/.Pop()
	}

	func run (s *Stack) {
	    for ip < instruction_array.Length()
	        s.step()
	    }
	}

defer for stack?

*
*/
package main

import (
	"fmt"
)

type Item_type int

type Item struct {
	item_type  Item_type
	item_value int
}

const (
	ADD Item_type = iota
	SUB
	MUL
	DIV
	VAL
)

func makeItem(item_type Item_type, item_value int) Item {
	return Item{item_type, item_value}
}

type Static_stack struct {
	stack_capacity int
	stack          []int
	stack_pointer  int
}

type Calculator struct {
	instruction_array []Item
	s_stack           Static_stack
	ins_pointer       int
}

func makeCalculator(instructions []Item, stack Static_stack) Calculator {
	return Calculator{instructions, stack, 0}
}

func makeStaticStack(stack_capacity int) Static_stack {
	return Static_stack{stack_capacity, make([]int, stack_capacity), 0}
}

func (s *Static_stack) Push(new_value int) {
	if s.stack_pointer < s.stack_capacity {
		s.stack[s.stack_pointer] = new_value
		s.stack_pointer++
	} else {
		panic("sheiße\n")
	}
}

func (s *Static_stack) Pop() int {
	if s.stack_pointer > 0 {
		s.stack_pointer--
		return s.stack[s.stack_pointer]
	} else {
		panic("sheiße\n")
	}
}

func (calc Calculator) run() int {
	for calc.ins_pointer < len(calc.instruction_array) {
		calc.Step()
	}
	return calc.s_stack.Pop()
}

func (calc *Calculator) Step() {
	next_item := calc.instruction_array[calc.ins_pointer]
	calc.ins_pointer++

	switch next_item.item_type {
	case ADD:
		{
			fmt.Println("add")
			val_b := calc.s_stack.Pop()
			val_a := calc.s_stack.Pop()
			calc.s_stack.Push(val_a + val_b)
			break
		}
	case SUB:
		{
			fmt.Println("sub")
			val_b := calc.s_stack.Pop()
			val_a := calc.s_stack.Pop()
			calc.s_stack.Push(val_a - val_b)
			break
		}
	case MUL:
		{
			fmt.Println("mul")
			val_b := calc.s_stack.Pop()
			val_a := calc.s_stack.Pop()
			calc.s_stack.Push(val_a * val_b)
			break
		}
	case DIV:
		{
			fmt.Println("div")
			val_b := calc.s_stack.Pop()
			val_a := calc.s_stack.Pop()
			calc.s_stack.Push(val_a / val_b)
			break
		}
	case VAL:
		{
			fmt.Println(next_item.item_value)
			calc.s_stack.Push(next_item.item_value)
			break
		}
	}
}

func main() {
	item1 := makeItem(VAL, 3)
	item2 := makeItem(VAL, 10)
	item3 := makeItem(VAL, 6)
	item4 := makeItem(ADD, 0)
	item5 := makeItem(ADD, 0)

	instructions := []Item{item1, item2, item3, item4, item5}
	stack1 := makeStaticStack(4)

	calc := makeCalculator(instructions, stack1)
	fmt.Printf("%d\n", calc.run())
}
