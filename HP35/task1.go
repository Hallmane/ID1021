/*
 RPN stack

 4 |   ||   ||   ||   ||   ||   |
 3 |   ||   ||   ||   ||   ||   |
 2 |   || 3 || 3 ||   ||10 ||   |
 1 | 3 || 3 || 7 ||10 || 6 ||60 |  <-- display
     3   [e]   7    +    6    *

hp35 € {
    instruction array (RPN) <- type Item struct
    instruction pointer     
    stack                   <- .Push()/.Pop()
}

    func run (s *Stack) {
        for ip < instruction_array.Length()  {
            s.step() 
        }
    }
const (
    ADD iota
    SUB 
    MUL
    DIV
    VAL
)

    func step (s Stack, ins_pointer int) {
        Item next_item = instruction_array[ins_pointer++]

        switch next_item.Item_type() {
            case ADD: {
                val_a := s.Pop()  <-- how?
                val_b := s.Pop()  <-- how?
                s.Push(val_a + val_b)
                break;
            }
            case SUB: {
                val_a := s.Pop()  <-- how?
                val_b := s.Pop()  <-- how?
                s.Push(val_a - val_b) <-- push should recieve an int at least
                break;
            }
            .
            .
            case VAL: {
                s.Push(next_item.Item_value()) <-- push should recieve an int at least
                break;
            }
        }
    }

* */
package main

import ()

type Item struct {
    ItemType    type //how to implement type
    ItemValue   int
}

type Stack struct {
    stack []Item 
}

func (s *Stack) stackOperation(type_of_item Item.itemType) {
    switch type_of_item {
        case 
    }
     
}

type calculator interface {}



func main() {
    expression_slice Item[]
    ins_slice [] 
    ins_pointer int
}
