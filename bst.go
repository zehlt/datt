package datt

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNotFoundBST       = errors.New("data not found in BST")
	ErrDuplicateValueBST = errors.New("duplicate value not found in BST")
)

func NewBinarySearchTree[T any](compare func(current T, other T) CompareResult) *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		compare: compare,
	}
}

type btNode[T any] struct {
	Value T
	Left  *btNode[T]
	Right *btNode[T]
}

type BinarySearchTree[T any] struct {
	compare   func(current T, other T) CompareResult
	root      *btNode[T]
	nodeCount int
}

func (bst *BinarySearchTree[T]) Insert(value T) error {
	if bst.root == nil {
		bst.root = &btNode[T]{
			Value: value,
		}
		bst.nodeCount++
		return nil
	}

	err := bst.insert(value, bst.root)
	if err != nil {
		return err
	}

	bst.nodeCount++
	return nil
}

func (bst *BinarySearchTree[T]) Iterate(callback func(value T)) {
	bst.inOrderTraversal(bst.root, func(bn *btNode[T]) {
		callback(bn.Value)
	})
}

func (bst *BinarySearchTree[T]) Contains(value T) bool {
	node := bst.search(value, bst.root)
	return node != nil
}

func (bst *BinarySearchTree[T]) Remove(value T) error {
	isContained := bst.Contains(value)
	if !isContained {
		return ErrNotFoundBST
	}

	bst.remove(value, bst.root)
	bst.nodeCount--
	return nil
}

func (bst *BinarySearchTree[T]) Clear() {
	bst.root = nil
	bst.nodeCount = 0
}

func (bst *BinarySearchTree[T]) Length() int {
	return bst.nodeCount
}

func (bst *BinarySearchTree[T]) String() string {
	var builder strings.Builder
	count := 0

	bst.levelOrderTraversal(bst.root, func(bn *btNode[T]) {
		count++
		builder.WriteString(fmt.Sprintf("[%v]", bn.Value))
	})

	return builder.String()
}

func (bst *BinarySearchTree[T]) search(value T, node *btNode[T]) *btNode[T] {
	if node == nil {
		return nil
	}

	result := bst.compare(value, node.Value)

	switch result {
	case HIGHER:
		return bst.search(value, node.Right)
	case LOWER:
		return bst.search(value, node.Left)
	}

	return node
}

func (bst *BinarySearchTree[T]) levelOrderTraversal(node *btNode[T], callback func(*btNode[T])) {
	if node == nil {
		return
	}

	queue := NewQueue[*btNode[T]]()
	queue.Enqueue(bst.root)

	for !queue.IsEmpty() {
		node, _ := queue.Dequeue()
		callback(node)
		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}
}

func (bst *BinarySearchTree[T]) inOrderTraversal(node *btNode[T], callback func(*btNode[T])) {
	if node == nil {
		return
	}

	bst.inOrderTraversal(node.Left, callback)
	callback(node)
	bst.inOrderTraversal(node.Right, callback)
}

func (bst *BinarySearchTree[T]) insert(value T, node *btNode[T]) error {
	result := bst.compare(value, node.Value)

	switch result {
	case HIGHER:
		if node.Right == nil {
			node.Right = &btNode[T]{
				Value: value,
			}
		} else {
			return bst.insert(value, node.Right)
		}
	case LOWER:
		if node.Left == nil {
			node.Left = &btNode[T]{
				Value: value,
			}
		} else {
			return bst.insert(value, node.Left)
		}
	case EQUAL:
		return ErrDuplicateValueBST
	}

	return nil
}

func (bst *BinarySearchTree[T]) remove(value T, node *btNode[T]) *btNode[T] {
	if node == nil {
		return nil
	}

	result := bst.compare(value, node.Value)

	switch result {
	case LOWER:
		node.Left = bst.remove(value, node.Left)
	case HIGHER:
		node.Right = bst.remove(value, node.Right)
	case EQUAL:
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		} else {
			smallestInRight := bst.digLeft(node.Right)
			node.Value = smallestInRight.Value
			node.Right = bst.remove(node.Value, node.Right)
		}
	}

	return node
}

func (bst *BinarySearchTree[T]) digLeft(node *btNode[T]) *btNode[T] {
	if node.Left != nil {
		return bst.digLeft(node.Left)
	} else {
		return node
	}
}
