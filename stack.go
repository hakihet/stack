package pickleslicesforall

import (
	"sync/atomic"
	"unsafe"
)

type Stack[T any] struct {
	next *element[T]
}

type element[T any] struct {
	next  *element[T]
	value T
}

func (stack *Stack[T]) Push(value T) (b bool) {
	push := element[T]{value: value, next: (*element[T])(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&stack.next))))}
	return atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&stack.next)), unsafe.Pointer(push.next), unsafe.Pointer(&push))
}

func (stack *Stack[T]) Pop() (b bool, value T) {
	if pop := (*element[T])(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&stack.next)))); pop != nil {
		if b = atomic.CompareAndSwapPointer((*unsafe.Pointer)((unsafe.Pointer)(&stack.next)), (unsafe.Pointer)(pop), (unsafe.Pointer)(pop.next)); b {
			value = pop.value
		}
	}
	return
}
