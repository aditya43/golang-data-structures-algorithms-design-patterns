# :zap: Go: Data Structures, Algorithms and Design Patterns
Data Structures, Algorithms and Design Patterns along with AR|VR|MR landscape. In-depth internals, my personal notes, example codes and projects.

## Author
Aditya Hajare ([Linkedin](https://in.linkedin.com/in/aditya-hajare)).

## Current Status
WIP (Work In Progress)!

## License
Open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT).

----------------------------------------

## Important Notes
- `Slices` store items in an `ordered` manner.
- `Maps` store items in an `unordered` manner.

----------------------------------------

### Applications of linked list in computer science

- Implementation of stacks and queues
- Implementation of graphs : Adjacency list representation of graphs is most popular which uses linked list to store adjacent vertices.
- Dynamic memory allocation : We use linked list of free blocks.
- Maintaining directory of names
- Performing arithmetic operations on long integers
- Manipulation of polynomials by storing constants in the node of linked list
- representing sparse matrices

### Applications of linked list in real world

- `Image viewer` – Previous and next images are linked, hence can be accessed by next and previous button.
- `Previous and next page in web browser` – We can access previous and next url searched in web browser by pressing back and next button since, they are linked as linked list.
- `Music Player` – Songs in music player are linked to previous and next song. you can play songs either from starting or ending of the list.

### Applications of Circular Linked Lists

- Useful for implementation of queue. Unlike this implementation, we don’t need to maintain two pointers for front and rear if we use circular linked list. We can maintain a pointer to the last inserted node and front can always be obtained as next of last.
- Circular lists are useful in applications to repeatedly go around the list. For example, when multiple applications are running on a PC, it is common for the operating system to put the running applications on a list and then to cycle through them, giving each of them a slice of time to execute, and then making them wait while the CPU is given to another application. It is convenient for the operating system to use a circular list so that when it reaches the end of the list it can cycle around to the front of the list.
- Circular Doubly Linked Lists are used for implementation of advanced data structures like Fibonacci Heap.
- A great way to represent a deck of cards in a game.
- The browser cache which allows you to hit the BACK button (a linked list of URLs)
- Applications that have a Most Recently Used (MRU) list (a linked list of file names)
- A stack, hash table, and binary tree can be implemented using a doubly linked list.
- Undo functionality in Photoshop or Word (a linked list of state)

# Singly Linked List
Singly linked list is simply a `List` that holds two pointers:

- `head` -- Points to the head node in the list.
- `tail` -- Points to the tail node in the list.

Each node in the list holds a pointer to the next node and a value:

- `next` -- A pointer to the next connecting node in the list
- `value` -- The underlying value of the node

## Operations

Most operations in a singly linked list are `O(n)` with the exception of insertion and deletion which are both `O(1)`.

### Iteration

Iterating over a singly linked list is easy. Set your initial pointer `curr` to the `head` of the list and repeatedly access the next node in the list by accessing `curr.next` until `curr` is `nil`.

### Reversing

Reversing a singly linked list is also fairly easy to do as you're just swapping the node's pointers for `next` to the previous node as you iterate over the list (//with the initial `prev` being `nil`//). At the start of the iteration the `tail` becomes the list's `head` and once the iteration is complete the `head` becomes the last node pointed at `prev`.

----------------------------------------