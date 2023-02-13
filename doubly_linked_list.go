package datt

// type Node[T any] struct {
// 	Value T
// 	next  *Node[T]
// 	prev  *Node[T]
// 	ll    *LinkedList[T]
// }

// func (n *Node[T]) Next() *Node[T] {
// 	if p := n.next; n.ll != nil && p != &n.ll.root {
// 		return p
// 	}
// 	return nil
// }

// func (n *Node[T]) Prev() *Node[T] {
// 	if p := n.prev; n.ll != nil && p != &n.ll.root {
// 		return p
// 	}

// 	return nil
// }

// type LinkedList[T any] struct {
// 	root Node[T]
// 	len  int
// }

// func NewLinkedList[T any]() *LinkedList[T] {
// 	ll := &LinkedList[T]{}
// 	ll.root.next = &ll.root
// 	ll.root.prev = &ll.root
// 	return ll
// }

// // O(1)
// func (ll *LinkedList[T]) PushFront(v T) *Node[T] {
// 	new := &Node[T]{
// 		Value: v,
// 	}

// 	ll.insertAfter(new, &ll.root)
// 	return new
// }

// // O(1)
// func (ll *LinkedList[T]) Front() *Node[T] {
// 	if ll.len == 0 {
// 		return nil
// 	}

// 	return ll.root.next
// }

// func (ll *LinkedList[T]) Back() *Node[T] {
// 	if ll.len == 0 {
// 		return nil
// 	}

// 	return ll.root.prev
// }

// func (ll *LinkedList[T]) Len() int {
// 	return ll.len
// }

// func (ll *LinkedList[T]) Clear() {
// 	ll.root.next = &ll.root
// 	ll.root.prev = &ll.root
// 	ll.len = 0
// }

// func (ll *LinkedList[T]) insertAfter(node *Node[T], at *Node[T]) {
// 	node.prev = at
// 	node.next = at.next

// 	at.next.prev = node
// 	at.next = node
// 	node.ll = ll
// 	ll.len++
// }
