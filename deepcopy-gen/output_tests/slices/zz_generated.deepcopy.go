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

package slices

import (
	bits "math/bits"
	sync "sync"
	unsafe "unsafe"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ttest) DeepCopyInto(out *Ttest) {
	*out = *in
	if in.Byte != nil {
		in, out := &in.Byte, &out.Byte
		*out = s_byteSPool.Get(len(*in))
		copy(*out, *in)
	}
	if in.Int16 != nil {
		in, out := &in.Int16, &out.Int16
		*out = s_int16SPool.Get(len(*in))
		copy(*out, *in)
	}
	if in.Int32 != nil {
		in, out := &in.Int32, &out.Int32
		*out = s_int32SPool.Get(len(*in))
		copy(*out, *in)
	}
	if in.Int64 != nil {
		in, out := &in.Int64, &out.Int64
		*out = s_int64SPool.Get(len(*in))
		copy(*out, *in)
	}
	if in.Uint8 != nil {
		in, out := &in.Uint8, &out.Uint8
		*out = s_uint8SPool.Get(len(*in))
		copy(*out, *in)
	}
	if in.Uint16 != nil {
		in, out := &in.Uint16, &out.Uint16
		*out = s_uint16SPool.Get(len(*in))
		copy(*out, *in)
	}
	if in.Uint32 != nil {
		in, out := &in.Uint32, &out.Uint32
		*out = s_uint32SPool.Get(len(*in))
		copy(*out, *in)
	}
	if in.Uint64 != nil {
		in, out := &in.Uint64, &out.Uint64
		*out = s_uint64SPool.Get(len(*in))
		copy(*out, *in)
	}
	if in.Float32 != nil {
		in, out := &in.Float32, &out.Float32
		*out = s_float32SPool.Get(len(*in))
		copy(*out, *in)
	}
	if in.Float64 != nil {
		in, out := &in.Float64, &out.Float64
		*out = s_float64SPool.Get(len(*in))
		copy(*out, *in)
	}
	if in.String != nil {
		in, out := &in.String, &out.String
		*out = s_stringSPool.Get(len(*in))
		copy(*out, *in)
	}
	if in.StringPtr != nil {
		in, out := &in.StringPtr, &out.StringPtr
		*out = s_stringp_SPool.Get(len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = stringp_Pool.Get().(*string)
				**out = **in
			}
		}
	}
	if in.StringPtrPtr != nil {
		in, out := &in.StringPtrPtr, &out.StringPtrPtr
		*out = s_stringp_p_SPool.Get(len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = stringp_p_Pool.Get().(**string)
				if **in != nil {
					in, out := *in, *out
					*out = stringp_Pool.Get().(*string)
					**out = **in
				}
			}
		}
	}
	if in.Map != nil {
		in, out := &in.Map, &out.Map
		*out = s_map_string_string_SPool.Get(len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = map_string_string_Pool.Get().(map[string]string)
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.MapPtr != nil {
		in, out := &in.MapPtr, &out.MapPtr
		*out = s_map_string_string_p_SPool.Get(len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
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
	}
	if in.Slice != nil {
		in, out := &in.Slice, &out.Slice
		*out = s_s_stringSPool.Get(len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = s_stringSPool.Get(len(*in))
				copy(*out, *in)
			}
		}
	}
	if in.SlicePtr != nil {
		in, out := &in.SlicePtr, &out.SlicePtr
		*out = s_s_stringp_SPool.Get(len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = s_stringp_Pool.Get().(*[]string)
				if **in != nil {
					in, out := *in, *out
					*out = s_stringSPool.Get(len(*in))
					copy(*out, *in)
				}
			}
		}
	}
	if in.Struct != nil {
		in, out := &in.Struct, &out.Struct
		*out = s_TtestSPool.Get(len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StructPtr != nil {
		in, out := &in.StructPtr, &out.StructPtr
		*out = s_Ttestp_SPool.Get(len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = Ttestp_Pool.Get().(*Ttest)
				(*in).DeepCopyInto(*out)
			}
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
	if in.Byte != nil {
		in := &in.Byte
		s_byteSPool.Put(*in)
	}
	if in.Int16 != nil {
		in := &in.Int16
		s_int16SPool.Put(*in)
	}
	if in.Int32 != nil {
		in := &in.Int32
		s_int32SPool.Put(*in)
	}
	if in.Int64 != nil {
		in := &in.Int64
		s_int64SPool.Put(*in)
	}
	if in.Uint8 != nil {
		in := &in.Uint8
		s_uint8SPool.Put(*in)
	}
	if in.Uint16 != nil {
		in := &in.Uint16
		s_uint16SPool.Put(*in)
	}
	if in.Uint32 != nil {
		in := &in.Uint32
		s_uint32SPool.Put(*in)
	}
	if in.Uint64 != nil {
		in := &in.Uint64
		s_uint64SPool.Put(*in)
	}
	if in.Float32 != nil {
		in := &in.Float32
		s_float32SPool.Put(*in)
	}
	if in.Float64 != nil {
		in := &in.Float64
		s_float64SPool.Put(*in)
	}
	if in.String != nil {
		in := &in.String
		s_stringSPool.Put(*in)
	}
	if in.StringPtr != nil {
		in := &in.StringPtr
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				stringp_Pool.Put(*in)
			}
		}
		s_stringp_SPool.Put(*in)
	}
	if in.StringPtrPtr != nil {
		in := &in.StringPtrPtr
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				if **in != nil {
					in := *in
					stringp_Pool.Put(*in)
					*in = nil
				}
				stringp_p_Pool.Put(*in)
			}
		}
		s_stringp_p_SPool.Put(*in)
	}
	if in.Map != nil {
		in := &in.Map
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				clear(*in)
				map_string_string_Pool.Put(*in)
			}
		}
		s_map_string_string_SPool.Put(*in)
	}
	if in.MapPtr != nil {
		in := &in.MapPtr
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				if **in != nil {
					in := *in
					clear(*in)
					map_string_string_Pool.Put(*in)
					*in = nil
				}
				map_string_string_p_Pool.Put(*in)
			}
		}
		s_map_string_string_p_SPool.Put(*in)
	}
	if in.Slice != nil {
		in := &in.Slice
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				s_stringSPool.Put(*in)
			}
		}
		s_s_stringSPool.Put(*in)
	}
	if in.SlicePtr != nil {
		in := &in.SlicePtr
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				if **in != nil {
					in := *in
					s_stringSPool.Put(*in)
					*in = nil
				}
				s_stringp_Pool.Put(*in)
			}
		}
		s_s_stringp_SPool.Put(*in)
	}
	if in.Struct != nil {
		in := &in.Struct
		for i := range *in {
			(*in)[i].ResetNoSelf()
		}
		s_TtestSPool.Put(*in)
	}
	if in.StructPtr != nil {
		in := &in.StructPtr
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				(*in).ResetNoSelf()
				Ttestp_Pool.Put(*in)
			}
		}
		s_Ttestp_SPool.Put(*in)
	}
	*in = Ttest{}
	TtestPool.Put(in)
}

// ResetNoSelf puts the given value back into the pool.
func (in *Ttest) ResetNoSelf() {
	if in == nil {
		return
	}
	if in.Byte != nil {
		in := &in.Byte
		s_byteSPool.Put(*in)
	}
	if in.Int16 != nil {
		in := &in.Int16
		s_int16SPool.Put(*in)
	}
	if in.Int32 != nil {
		in := &in.Int32
		s_int32SPool.Put(*in)
	}
	if in.Int64 != nil {
		in := &in.Int64
		s_int64SPool.Put(*in)
	}
	if in.Uint8 != nil {
		in := &in.Uint8
		s_uint8SPool.Put(*in)
	}
	if in.Uint16 != nil {
		in := &in.Uint16
		s_uint16SPool.Put(*in)
	}
	if in.Uint32 != nil {
		in := &in.Uint32
		s_uint32SPool.Put(*in)
	}
	if in.Uint64 != nil {
		in := &in.Uint64
		s_uint64SPool.Put(*in)
	}
	if in.Float32 != nil {
		in := &in.Float32
		s_float32SPool.Put(*in)
	}
	if in.Float64 != nil {
		in := &in.Float64
		s_float64SPool.Put(*in)
	}
	if in.String != nil {
		in := &in.String
		s_stringSPool.Put(*in)
	}
	if in.StringPtr != nil {
		in := &in.StringPtr
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				stringp_Pool.Put(*in)
			}
		}
		s_stringp_SPool.Put(*in)
	}
	if in.StringPtrPtr != nil {
		in := &in.StringPtrPtr
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				if **in != nil {
					in := *in
					stringp_Pool.Put(*in)
					*in = nil
				}
				stringp_p_Pool.Put(*in)
			}
		}
		s_stringp_p_SPool.Put(*in)
	}
	if in.Map != nil {
		in := &in.Map
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				clear(*in)
				map_string_string_Pool.Put(*in)
			}
		}
		s_map_string_string_SPool.Put(*in)
	}
	if in.MapPtr != nil {
		in := &in.MapPtr
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				if **in != nil {
					in := *in
					clear(*in)
					map_string_string_Pool.Put(*in)
					*in = nil
				}
				map_string_string_p_Pool.Put(*in)
			}
		}
		s_map_string_string_p_SPool.Put(*in)
	}
	if in.Slice != nil {
		in := &in.Slice
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				s_stringSPool.Put(*in)
			}
		}
		s_s_stringSPool.Put(*in)
	}
	if in.SlicePtr != nil {
		in := &in.SlicePtr
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				if **in != nil {
					in := *in
					s_stringSPool.Put(*in)
					*in = nil
				}
				s_stringp_Pool.Put(*in)
			}
		}
		s_s_stringp_SPool.Put(*in)
	}
	if in.Struct != nil {
		in := &in.Struct
		for i := range *in {
			(*in)[i].ResetNoSelf()
		}
		s_TtestSPool.Put(*in)
	}
	if in.StructPtr != nil {
		in := &in.StructPtr
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				(*in).ResetNoSelf()
				Ttestp_Pool.Put(*in)
			}
		}
		s_Ttestp_SPool.Put(*in)
	}
	*in = Ttest{}
}

var stringp_p_Pool = sync.Pool{New: func() any { return new(*string) }}
var s_stringp_Pool = sync.Pool{New: func() any { return new([]string) }}
var Ttestp_Pool = sync.Pool{New: func() any { return new(Ttest) }}
var map_string_string_p_Pool = sync.Pool{New: func() any { return new(map[string]string) }}
var stringp_Pool = sync.Pool{New: func() any { return new(string) }}
var TtestPool = sync.Pool{New: func() any { return new(Ttest) }}
var map_string_string_Pool = sync.Pool{New: func() any { return make(map[string]string, 16) }}
var s_stringp_p_SPool = newSlicePool[**string]()
var s_s_stringp_SPool = newSlicePool[*[]string]()
var s_Ttestp_SPool = newSlicePool[*Ttest]()
var s_map_string_string_p_SPool = newSlicePool[*map[string]string]()
var s_stringp_SPool = newSlicePool[*string]()
var s_s_stringSPool = newSlicePool[[]string]()
var s_byteSPool = newSlicePool[byte]()
var s_float32SPool = newSlicePool[float32]()
var s_float64SPool = newSlicePool[float64]()
var s_TtestSPool = newSlicePool[Ttest]()
var s_int16SPool = newSlicePool[int16]()
var s_int32SPool = newSlicePool[int32]()
var s_int64SPool = newSlicePool[int64]()
var s_map_string_string_SPool = newSlicePool[map[string]string]()
var s_stringSPool = newSlicePool[string]()
var s_uint16SPool = newSlicePool[uint16]()
var s_uint32SPool = newSlicePool[uint32]()
var s_uint64SPool = newSlicePool[uint64]()
var s_uint8SPool = newSlicePool[uint8]()

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
	idx := index(uint32(c))
	if c != 1<<idx { // 不是Get获取的[]byte，放在前一个索引的Pool里面
		idx--
	}
	clear(value)
	slice := (*slice)(unsafe.Pointer(&value))
	x := (*int32)(slice.Data)
	*x = int32(c)
	p.pools[idx].Put(slice.Data)
}
