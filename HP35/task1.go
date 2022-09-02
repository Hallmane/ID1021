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
	"time"
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
	lastUpdated    int
	dynamicSize    bool
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
	return staticStack{stack_capacity, make([]int, stack_capacity), 0, 0, false}
}
func makeDynamicStack(stack_capacity int) staticStack {
	return staticStack{stack_capacity, make([]int, stack_capacity), 0, 0, true}
}

func (s *staticStack) Push(new_value int) {
	if s.stack_pointer == s.stack_capacity && s.dynamicSize == true {
		s.stack = s.expandDong()
		fmt.Printf("EXPAND DONG %d\n ", s.stack_pointer)
	} else {
		s.stack[s.stack_pointer] = new_value
		s.stack_pointer++
		fmt.Printf("Stack capacity: %d\n ", s.stack_capacity)
	}
}
func (s *staticStack) Pop() int {
	if s.stack_pointer == -1 {
		fmt.Printf("Please don't panic\n")
		return s.stack[s.stack_pointer]
		//panic("Underflow.\n")
	} else {
		if s.stack_pointer <= s.stack_capacity/3 && s.stack_capacity > 4 && s.lastUpdated > 5 {
			s.reduceDong()
			fmt.Printf("~reduce d0ng: %d\n ", s.stack_pointer)
		}
		s.stack_pointer--
		return s.stack[s.stack_pointer]
	}
}

func (stack *staticStack) expandDong() []int {
	newStack := make([]int, (stack.stack_capacity * 2))
	copy(newStack, stack.stack)
	stack.lastUpdated = 0
	stack.stack_capacity = len(newStack)
	return newStack
}

/*
logic for this stack is not sound
*/
func (stack *staticStack) reduceDong() []int {
	newStack := make([]int, stack.stack_capacity*2/3)
	copy(newStack, stack.stack)
	stack.stack_capacity = len(newStack)
	stack.lastUpdated = 0
	return newStack
}

func (calc Calculator) run() int {
	for calc.ins_pointer < len(calc.instruction_array) {
		fmt.Printf("instruction pointer: %d\n", calc.ins_pointer)
		calc.Step()
	}
	return calc.s_stack.Pop()
}

func (calc *Calculator) Step() {
	next_item := calc.instruction_array[calc.ins_pointer]
	fmt.Printf("next item: %d\n", next_item.item_type)
	fmt.Printf("stack pointer: %d\n", calc.s_stack.stack_pointer)
	calc.s_stack.lastUpdated++
	calc.ins_pointer++
	fmt.Printf("\n")

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
			fmt.Printf("Value added: %d\n", next_item.item_value)
			calc.s_stack.Push(next_item.item_value)
			break
		}
	}
}

func main() {

	// create moar data
	item1 := makeItem(VAL, 1)
	item2 := makeItem(VAL, 2)
	item3 := makeItem(VAL, 3)
	item4 := makeItem(VAL, 4)
	item5 := makeItem(VAL, 5)
	item6 := makeItem(VAL, 6)
	item7 := makeItem(VAL, 7)
	item8 := makeItem(VAL, 8)
	item9 := makeItem(VAL, 9)
	item10 := makeItem(VAL, 9)
	item11 := makeItem(ADD, 0)
	item12 := makeItem(VAL, 9)
	item13 := makeItem(ADD, 0)
	item14 := makeItem(VAL, 9)
	item15 := makeItem(ADD, 0)
	item16 := makeItem(VAL, 9)
	item17 := makeItem(ADD, 0)
	item18 := makeItem(ADD, 0)
	item19 := makeItem(ADD, 0)
	item20 := makeItem(ADD, 0)
	item21 := makeItem(ADD, 0)
	item22 := makeItem(ADD, 0)

	instructions := []Item{item1, item2, item3, item4, item5, item6, item7, item8, item9, item10,
		item11, item12, item13, item14, item15, item16, item17, item18, item19, item20, item21,
		item22}
	stack1 := makeStaticStack(16)
	stack2 := makeDynamicStack(4)
	calc1 := makeCalculator(instructions, stack1)
	calc2 := makeCalculator(instructions, stack2)

	t0 := time.Now()
	fmt.Printf("%d\n", calc1.run())
	t1 := time.Now().Sub(t0)
	t2 := time.Now()
	fmt.Printf("%d\n", calc2.run())
	t3 := time.Now().Sub(t2)

	fmt.Printf("Static array time: %d µs\n Dynamic array time: %d µs\n", t1.Microseconds(), t3.Microseconds())
}
