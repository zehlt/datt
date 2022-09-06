# datt - Go Data Structure Compilation

**DATT** is compilation of various useful *data structures* using *go generics*.

## Current Structures

Below some examples of currently implemented structures.

### Linked List

	ll := datt.NewLinkedList[string]()

	ll.PushHead("a")
	ll.PushHead("b")
	ll.PushTail("c")

	log.Println(ll) // [b] -> [a] -> [c] -> nil

	ll.PopTail()
	log.Println(ll) // [b] -> [a] -> nil

	ll.PopHead()
	log.Println(ll) // [a] -> nil

### Bitset

	a := datt.NewBitset(1) // 2 -> initial byte capacity
	b := datt.NewBitset(1)

	a.Set(1, true)
	a.Set(3, true)

	b.Set(0, true)
	b.Set(1, true)
	b.Set(7, true)

	log.Println(a) // 00001010
	log.Println(b) // 10000011

	a.Or(b)
	log.Println(a) // 10001011

	log.Println(a.Get(0)) // true
	log.Println(a.Get(6)) // false

### Queue

	queue := datt.NewQueue[byte]()
	queue.Enqueue(5)
	queue.Enqueue(255)
	queue.Enqueue(21)
	queue.Enqueue(125)

	log.Println(queue.Dequeue()) // 5
	log.Println(queue.Dequeue()) // 255
	log.Println(queue.Dequeue()) // 21
	log.Println(queue.Dequeue()) // 125


### Stack

	stack := datt.NewStack[int]()

	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	log.Println(stack.Pop()) // 1
	log.Println(stack.Pop()) // 2
	log.Println(stack.Pop()) // 3

## Set

	a := datt.NewSet[string]()
	b := datt.NewSet[string]()

	a.Append("henry")
	a.Append("marie")
	a.Append("john")

	b.Append("tyler")
	b.Append("alice")
	b.Append("henry")

	c := a.Union(b) // doesn't mutate a or b, create brand new set
	log.Println(c)  // (tyler, henry, marie, john, alice)

	d := b.Intersection(a)
	log.Println(d) // henry

	log.Println(a.IsSubset(b)) // false

### Binary Search Tree

	type Human struct {
		Age int
	}

	func CompareOrderedHuman(current Human, other Human) datt.CompareResult {
		if current.Age > other.Age {
			return datt.HIGHER
		} else if current.Age < other.Age {
			return datt.LOWER
		} else {
			return datt.EQUAL
		}
	}

	func main() {
		bst := datt.NewBinarySearchTree(CompareOrderedHuman)
		bst.Insert(Human{Age: 15})
		bst.Insert(Human{Age: 54})
		bst.Insert(Human{Age: 21})
		bst.Insert(Human{Age: 9})

		bst.Iterate(func(h Human) {
			log.Println(h) // 9, 15, 21, 54
		})

		log.Println(bst.Contains(Human{Age: 341})) // false
		log.Println(bst.Contains(Human{Age: 54}))  // true

		// also works with go types
		other := datt.NewBinarySearchTree(datt.CompareOrdered[string])
		other.Insert("henry")
	}

### Priority Queue

	pq := datt.NewPriorityQueue(datt.CompareOrdered[uint16])

	pq.Push(15)
	pq.Push(0)
	pq.Push(7)
	pq.Push(9)

	log.Println(pq.Pop()) // 15
	log.Println(pq.Pop()) // 9
	log.Println(pq.Pop()) // 7
	log.Println(pq.Pop()) // 0

### Trie

	trie := datt.NewTrie()

	trie.Insert("car")
	trie.Insert("cartridge")
	trie.Insert("plane")
	trie.Insert("planer")

	ca := trie.AutoComplete("ca")
	log.Println(ca) // [r, rtridge]

	p := trie.AutoComplete("p")
	log.Println(p) // [lane, laner]

	doc := trie.AutoComplete("doc")
	log.Println(doc) // []

### Binary Heap

	bheap := datt.NewBinaryHeap(datt.CompareOrdered[float64])

	bheap.Push(36.6)
	bheap.Push(-14.687)
	bheap.Push(0.1415)
	bheap.Push(99.95)

	log.Println(bheap.PeekFront()) // 99.95
	log.Println(bheap.PeekBack())  // -14.687

	log.Println(bheap.Pop()) // 99.95
	log.Println(bheap.Pop()) // 36.6
	log.Println(bheap.Pop()) // 0.1415

### Generational Arena

	arena := datt.NewArena[string]()

	henry := arena.Create("henry data")
	marie := arena.Create("marie data")

	henryValue, _ := arena.Get(henry) // error if wrong key
	log.Println(henryValue) // henry data

	arena.Remove(marie)
	arena.Remove(marie) // fail cannot delete twice

### Tuple
 
	tuple := datt.Tuple[int, string]{
		First:  36,
		Second: "Henry",
	}
	log.Println(tuple.First)  // 36
	log.Println(tuple.Second) // Henry

	// there is also
	// datt.Triple
	// datt.Quadruple

### Dispatcher

exemple in progress

### Double Buffer Sync

exemple in progress

### In Progess

- Weight Graph
- Ordered Array
- Add iterate to missing struct
- Add tables showing O(N) best - average - worst