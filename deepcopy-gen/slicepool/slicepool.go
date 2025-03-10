package slicepool

import (
	"math/bits"
	"sync"
	"unsafe"
)

type slice struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

type SlicePool[T any] struct {
	pools [32]sync.Pool
}

func NewSlicePool[T any]() SlicePool[T] {
	return SlicePool[T]{}
}

func index(n uint32) uint32 {
	return uint32(bits.Len32(n - 1))
}

func (p *SlicePool[T]) Get(size int) []T {
	c := size
	// Small memory allocation is too scattered. the reuse rate is relatively high, types start with len=16
	if c < 16 {
		c = 16
	}
	idx := index(uint32(c))
	if v := p.pools[idx].Get(); v != nil {
		bp := v.(unsafe.Pointer)
		x := (*int32)(bp)
		c := *x
		*x = 0
		s := &slice{
			Data: bp,
			Len:  size,
			Cap:  int(c),
		}
		return *(*[]T)(unsafe.Pointer(s))
	}
	return make([]T, size, 1<<idx)
}

func (p *SlicePool[T]) Put(value []T) {
	c := cap(value)
	if c < 16 {
		return
	}
	idx := index(uint32(c))
	// []T not obtained by Get is placed in the Pool of the previous index
	if c != 1<<idx {
		idx--
	}
	clear(value)
	slice := (*slice)(unsafe.Pointer(&value))
	x := (*int32)(slice.Data)
	*x = int32(c)
	p.pools[idx].Put(slice.Data)
}
