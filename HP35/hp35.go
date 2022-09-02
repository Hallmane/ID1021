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
	"math/rand"
	"time"
)

type ItemType int

type Item struct {
	item_type  ItemType
	item_value int
}

const (
	ADD ItemType = iota
	SUB
	MUL
	DIV
	VAL
	MOD
)

func makeItem(item_type ItemType, item_value int) Item {
	return Item{item_type, item_value}
}

type st4ck struct {
	stack_capacity int
	stack          []int
	stack_pointer  int
	lastUpdated    int
	dynamicSize    bool
}

type Calculator struct {
	instruction_array []Item
	s_stack           st4ck
	ins_pointer       int
}

func makeCalculator(instructions []Item, stack st4ck) Calculator {
	return Calculator{instructions, stack, 0}
}

func makest4ck(stack_capacity int) st4ck {
	return st4ck{stack_capacity, make([]int, stack_capacity), 0, 0, false}
}
func makeDynamicStack(stack_capacity int) st4ck {
	return st4ck{stack_capacity, make([]int, stack_capacity), 0, 0, true}
}

func (s *st4ck) Push(new_value int) {
	if s.stack_pointer == s.stack_capacity && s.dynamicSize == true {
		s.stack = s.expandDong()
		fmt.Printf("EXPAND DONG %d\n ", s.stack_pointer)
	} else {
		s.stack[s.stack_pointer] = new_value
		s.stack_pointer++
		fmt.Printf("Stack capacity: %d\n ", s.stack_capacity)
	}
}
func (s *st4ck) pop() int {
	if s.stack_pointer == 0 {
		fmt.Printf("Please don't panic\n")
		return s.stack[s.stack_pointer]
	} else {
		if s.dynamicSize == true && s.stack_pointer <= s.stack_capacity/3 && s.stack_capacity > 8 && s.lastUpdated > 5 {
			s.reduceDong()
			fmt.Printf("~reduce d0ng: %d\n ", s.stack_pointer)
		}
		s.stack_pointer--
		return s.stack[s.stack_pointer]
	}
}

func (stack *st4ck) expandDong() []int {
	newStack := make([]int, (stack.stack_capacity * 3 / 2))
	copy(newStack, stack.stack)
	stack.stack_capacity = len(newStack)
	stack.lastUpdated = 0
	return newStack
}

// remove this function
func (stack *st4ck) reduceDong() []int {
	newStack := make([]int, stack.stack_capacity*2/3)
	copy(newStack, stack.stack)
	stack.stack_capacity = len(newStack)
	stack.lastUpdated = 0
	return newStack
}

func (calc Calculator) run() int {
	for calc.ins_pointer < len(calc.instruction_array) {
		fmt.Printf("instruction pointer: %d\n", calc.ins_pointer)
		calc.step()
	}
	return calc.s_stack.pop()
}

func (calc *Calculator) step() {
	next_item := calc.instruction_array[calc.ins_pointer]
	fmt.Printf("next item: %d\n", next_item.item_type)
	fmt.Printf("stack pointer: %d\n", calc.s_stack.stack_pointer)
	calc.s_stack.lastUpdated++
	calc.ins_pointer++
	fmt.Printf("\n")

	switch next_item.item_type {
	case ADD:
		{
			fmt.Println("ADD")
			val_b := calc.s_stack.pop()
			val_a := calc.s_stack.pop()
			calc.s_stack.Push(val_a + val_b)
		}
	case SUB:
		{
			fmt.Println("SUB")
			val_b := calc.s_stack.pop()
			val_a := calc.s_stack.pop()
			calc.s_stack.Push(val_a - val_b)
		}
	case MUL:
		{
			fmt.Println("MUL")
			val_b := calc.s_stack.pop()
			val_a := calc.s_stack.pop()
			calc.s_stack.Push(val_a * val_b)
		}
	case DIV:
		{
			fmt.Println("DIV")
			val_b := calc.s_stack.pop()
			val_a := calc.s_stack.pop()
			if val_b == 0 {
				val_b = 1
			}
			calc.s_stack.Push(val_a / val_b)
		}
	case VAL:
		{
			fmt.Printf("Value added: %d\n", next_item.item_value)
			calc.s_stack.Push(next_item.item_value)
		}
	case MOD:
		{
			val_b := calc.s_stack.pop()
			val_a := calc.s_stack.pop()
			calc.s_stack.Push(val_a % val_b)
		}
	}
}
func valGen(loopSize int) []Item {
	rand.Seed(time.Now().UTC().UnixNano())

	var itemz = make([]Item, loopSize)
	for i := 0; i < loopSize; i++ {

		itemz[i] = makeItem(4, 1+rand.Intn(100))
	}
	return itemz
}

func listGen(loopSize int) []Item {
	rand.Seed(time.Now().UTC().UnixNano())

	var op ItemType = 0
	var itemz = make([]Item, loopSize)

	for i := 0; i < loopSize; i++ {
		if op == 5 {
			op = 0
		}

		itemz[i] = makeItem(op, 1+rand.Intn(100))
		op++
	}
	return itemz
}

func main() {

	randomValues := valGen(10000)
	randomItems := listGen(16500)
	items := append(randomValues, randomItems...)
	//	i1 := makeItem(VAL, 9)
	//	m1 := makeItem(VAL, 2)
	//
	//	i2 := makeItem(VAL, 6)
	//	m2 := makeItem(VAL, 1)
	//
	//	i3 := makeItem(VAL, 0)
	//	m3 := makeItem(VAL, 2)
	//
	//	i4 := makeItem(VAL, 5)
	//	m4 := makeItem(VAL, 1)
	//
	//	i5 := makeItem(VAL, 0)
	//	m5 := makeItem(VAL, 2)
	//
	//	i6 := makeItem(VAL, 4)
	//	m6 := makeItem(VAL, 1)
	//
	//	i7 := makeItem(VAL, 3)
	//	m6 := makeItem(VAL, 2)
	//
	//	i8 := makeItem(VAL, 3)
	//	m7 := makeItem(VAL, 1)
	//
	//	i9 := makeItem(VAL, 7)
	//	m8 := makeItem(VAL, 2)
	//
	// create moar data
	//item1 := makeItem(VAL, 1)
	//item2 := makeItem(VAL, 2)
	//item3 := makeItem(VAL, 3)
	//item4 := makeItem(VAL, 4)

	//instructions := []Item{item1, item2, item3, item4, item5, item6, item7, item8, item9, item10,
	//	item11, item12, item13, item14, item15, item16, item17, item18, item19, item20, item21,
	//	item22}
	stack1 := makest4ck(20000)
	stack2 := makeDynamicStack(4)
	calc1 := makeCalculator(items, stack1)
	calc2 := makeCalculator(items, stack2)

	t0 := time.Now()
	calc1.run()
	t1 := time.Now().Sub(t0)
	t2 := time.Now()
	calc2.run()
	t3 := time.Now().Sub(t2)

	fmt.Printf("Static array time: %d µs\n Dynamic array time: %d ms\n", t1.Milliseconds(), t3.Milliseconds())
}
