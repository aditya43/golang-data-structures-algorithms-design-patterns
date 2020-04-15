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

----------------------------------------

### Applications of stack

- The basic application of stack is **backtracking**. i.e. to track (keep a record) of from where we begun to where we got.
- Stacks are used in **undo** mechanism in text editor.
- Language processing:
    * space for parameters and local variables is created internally using a stack.
    * compiler's syntax check for matching braces is implemented by using stack.
    * support for recursion.
    * Expression evaluation like postfix or prefix in compilers.
- Backtracking (game playing, finding paths, exhaustive searching.
- Memory management, run-time environment for nested language features. etc

----------------------------------------