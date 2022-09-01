/*
RPN stack
4 |   ||   ||   ||   ||   ||   |
3 |   ||   ||   ||   ||   ||   |
2 |   || 3 || 3 ||   ||10 ||   |
1 | 3 || 3 || 7 ||10 || 6 ||60 |  <-- display

	3   [e]   7    +    6    *
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

type staticStack struct {
	stack_capacity int
	stack          []int
	stack_pointer  int
	stackDynamic   bool
}

type Calculator struct {
	instruction_array []Item
	s_stack           staticStack
	ins_pointer       int
}

func makeCalculator(instructions []Item, stack staticStack) Calculator {
	return Calculator{instructions, stack, 0}
}

func makeStaticStack(stack_capacity int) staticStack {
	return staticStack{stack_capacity, make([]int, stack_capacity), 0, false}
}

func (s *staticStack) Push(new_value int) {
	if s.stack_pointer == s.stack_capacity {
		s.stack = s.expandDong()
		//fmt.Printf("STACK POINTER %d\n ", s.stack_pointer)
		//panic("push sheiße\n")
	} else {
		s.stack[s.stack_pointer] = new_value
		s.stack_pointer++
		//fmt.Printf("CAP %d\n ", s.stack_capacity)
	}
}

//func (s *staticStack) Push(new_value int) {
//	if s.stack_pointer < s.stack_capacity {
//		s.stack[s.stack_pointer] = new_value
//		s.stack_pointer++
//	} else {
//		s.stack = s.expandDong()
//		panic("push sheiße\n")
//	}
//}

func (stack *staticStack) expandDong() []int {
	newStack := make([]int, (stack.stack_capacity * 3 / 2))
	copy(newStack, stack.stack)
	stack.stack_capacity = len(newStack)
	return newStack
}

// logic for this stack
func (stack *staticStack) reduceDong() []int {
	newStack := make([]int, stack.stack_capacity*2)
	copy(newStack, stack.stack)
	stack.stack_capacity = len(newStack)
	return newStack
}

func (s *staticStack) Pop() int {
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
			//fmt.Println("add")
			val_b := calc.s_stack.Pop()
			val_a := calc.s_stack.Pop()
			calc.s_stack.Push(val_a + val_b)
			break
		}
	case SUB:
		{
			//fmt.Println("sub")
			val_b := calc.s_stack.Pop()
			val_a := calc.s_stack.Pop()
			calc.s_stack.Push(val_a - val_b)
			break
		}
	case MUL:
		{
			//fmt.Println("mul")
			val_b := calc.s_stack.Pop()
			val_a := calc.s_stack.Pop()
			calc.s_stack.Push(val_a * val_b)
			break
		}
	case DIV:
		{
			//fmt.Println("div")
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
	item4 := makeItem(VAL, 7)
	item5 := makeItem(VAL, 8)
	item6 := makeItem(VAL, 9)
	item7 := makeItem(ADD, 0)
	item8 := makeItem(ADD, 0)
	item9 := makeItem(ADD, 0)

	instructions := []Item{item1, item2, item3, item4, item5, item6, item7, item8, item8, item9}
	stack1 := makeStaticStack(4)

	calc := makeCalculator(instructions, stack1)
	fmt.Printf("%d\n", calc.run())
}
