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

package pointer

import (
	bits "math/bits"
	sync "sync"
	unsafe "unsafe"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ttest) DeepCopyInto(out *Ttest) {
	*out = *in
	if in.Builtin != nil {
		in, out := &in.Builtin, &out.Builtin
		*out = stringp_Pool.Get().(*string)
		**out = **in
	}
	if in.Ptr != nil {
		in, out := &in.Ptr, &out.Ptr
		*out = stringp_p_Pool.Get().(**string)
		if **in != nil {
			in, out := *in, *out
			*out = stringp_Pool.Get().(*string)
			**out = **in
		}
	}
	if in.Map != nil {
		in, out := &in.Map, &out.Map
		*out = map_string_string_p_Pool.Get().(*map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = map_string_string_Pool.Get().(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Slice != nil {
		in, out := &in.Slice, &out.Slice
		*out = s_stringp_Pool.Get().(*[]string)
		if **in != nil {
			in, out := *in, *out
			*out = s_stringSPool.Get(len(*in))
			copy(*out, *in)
		}
	}
	if in.MapPtr != nil {
		in, out := &in.MapPtr, &out.MapPtr
		*out = map_string_string_p_p_Pool.Get().(**map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = map_string_string_p_Pool.Get().(*map[string]string)
			if **in != nil {
				in, out := *in, *out
				*out = map_string_string_Pool.Get().(map[string]string)
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.SlicePtr != nil {
		in, out := &in.SlicePtr, &out.SlicePtr
		*out = s_stringp_p_Pool.Get().(**[]string)
		if **in != nil {
			in, out := *in, *out
			*out = s_stringp_Pool.Get().(*[]string)
			if **in != nil {
				in, out := *in, *out
				*out = s_stringSPool.Get(len(*in))
				copy(*out, *in)
			}
		}
	}
	if in.Struct != nil {
		in, out := &in.Struct, &out.Struct
		*out = Ttestp_Pool.Get().(*Ttest)
		(*in).DeepCopyInto(*out)
	}
	if in.StructPtr != nil {
		in, out := &in.StructPtr, &out.StructPtr
		*out = Ttestp_p_Pool.Get().(**Ttest)
		if **in != nil {
			in, out := *in, *out
			*out = Ttestp_Pool.Get().(*Ttest)
			(*in).DeepCopyInto(*out)
		}
	}
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
	if in.Builtin != nil {
		in := &in.Builtin
		x := (*string)(*in)
		stringp_Pool.Put(x)
	}
	if in.Ptr != nil {
		in := &in.Ptr
		if **in != nil {
			in := *in
			stringp_Pool.Put(*in)
			*in = nil
		}
		x := (**string)(*in)
		stringp_p_Pool.Put(x)
	}
	if in.Map != nil {
		in := &in.Map
		if **in != nil {
			in := *in
			clear(*in)
			map_string_string_Pool.Put(*in)
			*in = nil
		}
		x := (*map[string]string)(*in)
		map_string_string_p_Pool.Put(x)
	}
	if in.Slice != nil {
		in := &in.Slice
		if **in != nil {
			in := *in
			s_stringSPool.Put(*in)
			*in = nil
		}
		x := (*[]string)(*in)
		s_stringp_Pool.Put(x)
	}
	if in.MapPtr != nil {
		in := &in.MapPtr
		if **in != nil {
			in := *in
			if **in != nil {
				in := *in
				clear(*in)
				map_string_string_Pool.Put(*in)
				*in = nil
			}
			map_string_string_p_Pool.Put(*in)
			*in = nil
		}
		x := (**map[string]string)(*in)
		map_string_string_p_p_Pool.Put(x)
	}
	if in.SlicePtr != nil {
		in := &in.SlicePtr
		if **in != nil {
			in := *in
			if **in != nil {
				in := *in
				s_stringSPool.Put(*in)
				*in = nil
			}
			s_stringp_Pool.Put(*in)
			*in = nil
		}
		x := (**[]string)(*in)
		s_stringp_p_Pool.Put(x)
	}
	if in.Struct != nil {
		in := &in.Struct
		(*in).ResetNoSelf()
		x := (*Ttest)(*in)
		Ttestp_Pool.Put(x)
	}
	if in.StructPtr != nil {
		in := &in.StructPtr
		if **in != nil {
			in := *in
			(*in).ResetNoSelf()
			Ttestp_Pool.Put(*in)
			*in = nil
		}
		x := (**Ttest)(*in)
		Ttestp_p_Pool.Put(x)
	}
	*in = Ttest{}
	TtestPool.Put(in)
}

// ResetNoSelf puts the given value back into the pool.
func (in *Ttest) ResetNoSelf() {
	if in == nil {
		return
	}
	if in.Builtin != nil {
		in := &in.Builtin
		x := (*string)(*in)
		stringp_Pool.Put(x)
	}
	if in.Ptr != nil {
		in := &in.Ptr
		if **in != nil {
			in := *in
			stringp_Pool.Put(*in)
			*in = nil
		}
		x := (**string)(*in)
		stringp_p_Pool.Put(x)
	}
	if in.Map != nil {
		in := &in.Map
		if **in != nil {
			in := *in
			clear(*in)
			map_string_string_Pool.Put(*in)
			*in = nil
		}
		x := (*map[string]string)(*in)
		map_string_string_p_Pool.Put(x)
	}
	if in.Slice != nil {
		in := &in.Slice
		if **in != nil {
			in := *in
			s_stringSPool.Put(*in)
			*in = nil
		}
		x := (*[]string)(*in)
		s_stringp_Pool.Put(x)
	}
	if in.MapPtr != nil {
		in := &in.MapPtr
		if **in != nil {
			in := *in
			if **in != nil {
				in := *in
				clear(*in)
				map_string_string_Pool.Put(*in)
				*in = nil
			}
			map_string_string_p_Pool.Put(*in)
			*in = nil
		}
		x := (**map[string]string)(*in)
		map_string_string_p_p_Pool.Put(x)
	}
	if in.SlicePtr != nil {
		in := &in.SlicePtr
		if **in != nil {
			in := *in
			if **in != nil {
				in := *in
				s_stringSPool.Put(*in)
				*in = nil
			}
			s_stringp_Pool.Put(*in)
			*in = nil
		}
		x := (**[]string)(*in)
		s_stringp_p_Pool.Put(x)
	}
	if in.Struct != nil {
		in := &in.Struct
		(*in).ResetNoSelf()
		x := (*Ttest)(*in)
		Ttestp_Pool.Put(x)
	}
	if in.StructPtr != nil {
		in := &in.StructPtr
		if **in != nil {
			in := *in
			(*in).ResetNoSelf()
			Ttestp_Pool.Put(*in)
			*in = nil
		}
		x := (**Ttest)(*in)
		Ttestp_p_Pool.Put(x)
	}
	*in = Ttest{}
}

var s_stringp_p_Pool = sync.Pool{New: func() any { return new(*[]string) }}
var Ttestp_p_Pool = sync.Pool{New: func() any { return new(*Ttest) }}
var map_string_string_p_p_Pool = sync.Pool{New: func() any { return new(*map[string]string) }}
var stringp_p_Pool = sync.Pool{New: func() any { return new(*string) }}
var s_stringp_Pool = sync.Pool{New: func() any { return new([]string) }}
var Ttestp_Pool = sync.Pool{New: func() any { return new(Ttest) }}
var map_string_string_p_Pool = sync.Pool{New: func() any { return new(map[string]string) }}
var stringp_Pool = sync.Pool{New: func() any { return new(string) }}
var TtestPool = sync.Pool{New: func() any { return new(Ttest) }}
var map_string_string_Pool = sync.Pool{New: func() any { return make(map[string]string, 16) }}
var s_stringSPool = newSlicePool[string]()

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
