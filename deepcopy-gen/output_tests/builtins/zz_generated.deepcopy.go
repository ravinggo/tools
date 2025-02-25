//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package builtins

import (
	bits "math/bits"
	sync "sync"
	unsafe "unsafe"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ttest) DeepCopyInto(out *Ttest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ttest.
func (in *Ttest) DeepCopy() *Ttest {
	if in == nil {
		return nil
	}
	out := TtestPool.Get().(*Ttest)
	in.DeepCopyInto(out)
	return out
}

// Reset puts the given value back into the pool.
func (in *Ttest) Reset() {
	if in == nil {
		return
	}
	*in = Ttest{}
	TtestPool.Put(in)
}

// ResetNoSelf puts the given value back into the pool.
func (in *Ttest) ResetNoSelf() {
	if in == nil {
		return
	}
	*in = Ttest{}
}

var TtestPool = sync.Pool{New: func() any { return new(Ttest) }}

type slice struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

type slicePool[T any] struct {
	pools [32]sync.Pool
}

func newSlicePool[T any]() slicePool[T] {
	return slicePool[T]{}
}

func index(n uint32) uint32 {
	return uint32(bits.Len32(n - 1))
}

func (p *slicePool[T]) Get(size int) []T {
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

func (p *slicePool[T]) Put(value []T) {
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
