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

* */
package main

import ()

type Item struct {
    Item_type     item_type
    Item__value   int
}

const (
    ADD item_type = iota
    SUB
    MUL
    DIV
    VAL
)

type Static_stack struct {
    stack_capacity int
    stack []int
    stack_pointer int
}

func (s_stack Static_stack) make_static_stack() {
    s_stack.stack_capacity = 3
    s_stack.stack = make([]int, stack_capacity)
    s_stack.stack_pointer = 0
}

// check if last index == nil, can push.
func (s static_stack) Push(new_value int) () {

    if s.stack_pointer < s.stack_cap {
        &s[sp] = &s[sp-1]
        &s[sp-1] = new_value
    }
}

func (s *[]Stack) Pop(stack_pointer int) (int) {
    return_val = &s[0]
}

func run (s []Stack) {
    for ip < instruction_array.Length()  {
        s.Step() 
    }
}

func step (s Stack, ins_pointer int instruction_array int[]) {
    next_item := instruction_array[ins_pointer]
    ins_pointer ++; // problems

    switch next_item.Item_type() {
        case ADD: {
            val_b := s.Pop()                    //<-- how?
            val_a := s.Pop()                    //<-- how?
            s.Push(val_a + val_b)
            break;
        }
        case SUB: {
            val_b := s.Pop()                    
            val_a := s.Pop()                    
            s.Push(val_a - val_b)       //<-- push should recieve an int at least
            break;
        }
        case MUL: {
            val_b := s.Pop()                    
            val_a := s.Pop()                    
            s.Push(val_a * val_b)       //<-- push should recieve an int at least
            break;
        }
        case DIV: {
            val_b := s.Pop()                    
            val_a := s.Pop()                    
            s.Push(val_a / val_b)       //<-- push should recieve an int at least
            break;
        }
        case VAL: {
            s.Push(next_item.Item_value()) //<-- push should recieve an int at least
            break;
        }
    }
}
//type calculator interface {}



func main() {
    expression_slice []Item
    ins_slice [] 
    ins_pointer int
}
