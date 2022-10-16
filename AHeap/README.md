# A Heap
Ordering depending on a set priority, be that low/high values 
or some other field.

## Tasks
### A list of items
two linked list implementations where 
1. O(1) add and O(n) remove [ ] 
2. O(n) add and O(1) remove [ ] 

Q: Are there situations to prefer one over the other?

With O(n) remove, you traverse the list until smallest value, 
which is the last item

With O(n) add, you traverse th?


### A very special tree
    : a heap :
root of tree = highest priority element

    : balancing the tree :
keeping track of superseding items, e.g items below each item.
This makes it easy to balance the tree since new items are
always placed in the branch with fewer items.
3. Implement the head as a binary tree together with the 
previous instructions.

    : add :
* if nil, just add it to the empty tree
* else, increment size by one? (on what, the item or the heap?)
* if new val < current node -> swap the values
* if left branch = nil, add it as a new node there
* if right branch = nil, add it as a new node there
* else chose next node (left or right) based on tail length

    : remove :
* if root = nil, the heap is nil
* if lengthTail = 1 and it gets remove, the heap should 
of course be set to nil

#
__
TODO:
add new node and place accordingly, through priority 

__
Sketch




















