package datt

type color byte

const (
	red color = iota
	black
)

type rbNode[K any, V any] struct {
	key   K
	val   V
	color color
	left  *rbNode[K, V]
	right *rbNode[K, V]
}

// Left Leaning Red Black Tree
type RedBlackTree[K any, V any] struct {
	root    *rbNode[K, V]
	len     int
	compare CompareFunc[K]
}

func NewRedBlackTree[K any, V any](compare CompareFunc[K]) *RedBlackTree[K, V] {
	return &RedBlackTree[K, V]{
		compare: compare,
		root:    nil,
		len:     0,
	}
}

func (rb *RedBlackTree[K, V]) Insert(key K, val V) {
	rb.len++

	rb.root = rb.insert(rb.root, key, val)
	rb.root.color = black
}

func (rb *RedBlackTree[K, V]) Get(key K) (V, bool) {
	node := rb.search(rb.root, key)

	if node == nil {
		var d V
		return d, false
	}

	return node.val, true
}

// func (rb *RedBlackTree[K, V]) Remove(key K) bool {
// 	if !rb.Has(key) {
// 		return false
// 	}

// 	// maybe change the color of the root to red

// 	rb.root = rb.remove(rb.root, key)
// 	// reswitch the color of the root to black
// 	return true
// }

func (rb *RedBlackTree[K, V]) Len() int {
	return rb.len
}

func (rb *RedBlackTree[K, V]) Has(key K) bool {
	return rb.search(rb.root, key) != nil
}

// func (rb *RedBlackTree[K, V]) remove(node *rbNode[K, V], key K) *rbNode[K, V] {
// 	cmp := rb.compare(key, node.key)

// 	switch cmp {
// 	case EQUAL:
// 		return node

// 	case LOWER:
// 		return rb.search(node.left, key)

// 	case HIGHER:
// 		return rb.search(node.right, key)

// 	default:
// 		panic("search impossible")
// 	}

// 	return nil
// }

func (rb *RedBlackTree[K, V]) search(node *rbNode[K, V], key K) *rbNode[K, V] {
	if node == nil {
		return nil
	}

	cmp := rb.compare(key, node.key)

	switch cmp {
	case EQUAL:
		return node

	case LOWER:
		return rb.search(node.left, key)

	case HIGHER:
		return rb.search(node.right, key)
	default:
		panic("search impossible")
	}
}

func (rb *RedBlackTree[K, V]) insert(node *rbNode[K, V], key K, val V) *rbNode[K, V] {
	if node == nil {
		return &rbNode[K, V]{key: key, val: val, color: red}
	}

	cmp := rb.compare(key, node.key)

	switch cmp {
	case LOWER:
		node.left = rb.insert(node.left, key, val)
	case HIGHER:
		node.right = rb.insert(node.right, key, val)
	case EQUAL:
		node.val = val
	default:
		panic("insert panic")
	}

	if rb.isRed(node.right) && !rb.isRed(node.left) {
		node = rb.rotateLeft(node)
	}

	if rb.isRed(node.left) && rb.isRed(node.left.left) {
		node = rb.rotateRight(node)
	}

	if rb.isRed(node.left) && rb.isRed(node.right) {
		rb.flipColors(node)
	}

	return node
}

func (rb *RedBlackTree[K, V]) rotateLeft(h *rbNode[K, V]) *rbNode[K, V] {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = red
	return x
}

func (rb *RedBlackTree[K, V]) rotateRight(h *rbNode[K, V]) *rbNode[K, V] {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = red
	return x
}

func (rb *RedBlackTree[K, V]) flipColors(node *rbNode[K, V]) {
	node.color = red
	node.left.color = black
	node.right.color = black
}

func (rb *RedBlackTree[K, V]) isRed(node *rbNode[K, V]) bool {
	if node == nil {
		return false
	}

	return node.color == red
}
