package datt

import "sync"

// thread safe double buffer
type DoubleBuffer[T any] struct {
	backBuff  *Queue[T]
	frontBuff *Queue[T]

	// globalMutex sync.Mutex
	pushMutex sync.Mutex
	popMutex  sync.Mutex
}

func NewDoubleBuffer[T any]() DoubleBuffer[T] {
	front := NewQueue[T]()
	back := NewQueue[T]()

	return DoubleBuffer[T]{
		frontBuff: &front,
		backBuff:  &back,
	}
}

func (d *DoubleBuffer[T]) SwitchBuff() {
	d.pushMutex.Lock()
	defer d.pushMutex.Unlock()

	d.popMutex.Lock()
	defer d.popMutex.Unlock()

	temp := d.frontBuff
	d.frontBuff = d.backBuff
	d.backBuff = temp
}

func (d *DoubleBuffer[T]) Push(data T) {
	d.pushMutex.Lock()
	defer d.pushMutex.Unlock()

	d.frontBuff.Enqueue(data)
}

func (d *DoubleBuffer[T]) Pop() (T, bool) {
	d.popMutex.Lock()
	defer d.popMutex.Unlock()

	data, err := d.backBuff.Dequeue()
	if err != nil {
		var d T
		return d, false
	}

	return data, true
}
