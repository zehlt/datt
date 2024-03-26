package datt

type btNode[T any] struct {
	value T
	left  *btNode[T]
	right *btNode[T]
}

type BinarySearchTree[T any] struct {
	root    *btNode[T]
	len     int
	compare CompareFunc[T]
}

func NewBinarySearchTree[T any](compare CompareFunc[T]) *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		compare: compare,
	}
}

// Average O(log(N))
// Worst  O(N)
func (b *BinarySearchTree[T]) Insert(v T) {
	if b.root == nil {
		b.root = &btNode[T]{value: v}
		b.len++
		return
	}

	b.insert(v, b.root)
}

// Average O(log(N))
// Worst  O(N)
func (b *BinarySearchTree[T]) Remove(v T) bool {
	if b.root == nil {
		return false
	}

	if !b.Has(v) {
		return false
	}

	b.len--
	b.root = b.remove(v, b.root)
	return true
}

// O(1)
func (b *BinarySearchTree[T]) Len() int {
	return b.len
}

// Average O(log(N))
// Worst  O(N)
func (b *BinarySearchTree[T]) Has(v T) bool {
	return b.search(v, b.root) != nil
}

// O(1)
func (b *BinarySearchTree[T]) Clear() {
	b.root = nil
	b.len = 0
}

// func (b *BinarySearchTree[T]) String() string {
// 	var builder strings.Builder
// 	count := 0

// 	bst.levelOrderTraversal(bst.root, func(bn *btNode[T]) {
// 		count++
// 		builder.WriteString(fmt.Sprintf("[%v]", bn.Value))
// 	})

// 	return builder.String()
// }

func (b *BinarySearchTree[T]) search(v T, node *btNode[T]) *btNode[T] {
	if node == nil {
		return nil
	}

	result := b.compare(v, node.value)

	switch result {
	case HIGHER:
		return b.search(v, node.right)
	case LOWER:
		return b.search(v, node.left)
	case EQUAL:
		return node
	default:
		panic("should the right value for CompareResult")
	}
}

func (b *BinarySearchTree[T]) insert(v T, node *btNode[T]) {
	result := b.compare(v, node.value)

	switch result {
	case HIGHER:
		if node.right == nil {
			node.right = &btNode[T]{value: v}
			b.len++
			return
		}
		b.insert(v, node.right)
		return

	case LOWER:
		if node.left == nil {
			node.left = &btNode[T]{value: v}
			b.len++
			return
		}
		b.insert(v, node.left)
		return

	case EQUAL:
		return
		// TODO: handle case when equal
	default:
		panic("should the right value for CompareResult")
	}
}

func (bst *BinarySearchTree[T]) remove(v T, node *btNode[T]) *btNode[T] {
	result := bst.compare(v, node.value)

	switch result {
	case LOWER:
		node.left = bst.remove(v, node.left)

	case HIGHER:
		node.right = bst.remove(v, node.right)

	case EQUAL:
		if node.left == nil {
			return node.right

		} else if node.right == nil {
			return node.left

		} else {
			left := bst.digLeft(node.right)
			node.value = left.value
			node.right = bst.remove(node.value, node.right)
			return node
		}
	}

	return node
}

func (bst *BinarySearchTree[T]) digLeft(node *btNode[T]) *btNode[T] {
	if node.left != nil {
		return bst.digLeft(node.left)
	} else {
		return node
	}
}

func (bst *BinarySearchTree[T]) Do(order OrderTraversal, f func(v T)) {
	switch order {
	case PreOrder:
		bst.preorder(f, bst.root)
	case InOrder:
		bst.inorder(f, bst.root)
	case PostOrder:
		bst.postorder(f, bst.root)
	case LevelOrder:
		bst.levelorder(f)
	default:
		panic("must pick the right order traversal constant")
	}
}

func (bst *BinarySearchTree[T]) preorder(f func(v T), node *btNode[T]) {
	if node == nil {
		return
	}

	f(node.value)
	bst.preorder(f, node.left)
	bst.preorder(f, node.right)
}

func (bst *BinarySearchTree[T]) inorder(f func(v T), node *btNode[T]) {
	if node == nil {
		return
	}

	bst.inorder(f, node.left)
	f(node.value)
	bst.inorder(f, node.right)
}

func (bst *BinarySearchTree[T]) postorder(f func(v T), node *btNode[T]) {
	if node == nil {
		return
	}

	bst.postorder(f, node.left)
	bst.postorder(f, node.right)
	f(node.value)
}

// TODO: check if correct
func (bst *BinarySearchTree[T]) levelorder(f func(v T)) {
	if bst.root == nil {
		return
	}

	q := QueueCircularArray[*btNode[T]]{}
	q.Enqueue(bst.root)

	for q.Len() > 0 {
		v, _ := q.Dequeue()
		f(v.value)

		if v.left != nil {
			q.Enqueue(v.left)
		}
		if v.right != nil {
			q.Enqueue(v.right)
		}
	}
}
