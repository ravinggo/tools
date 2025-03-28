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

package structs

import (
	sync "sync"

	slicepool "github.com/ravinggo/tools/deepcopy-gen/slicepool"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Inner) DeepCopyInto(out *Inner) {
	*out = *in
	if in.String != nil {
		in, out := &in.String, &out.String
		*out = s_stringSPool.Get(len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Inner.
func (in *Inner) DeepCopy() *Inner {
	if in == nil {
		return nil
	}
	out := InnerPool.Get().(*Inner)
	in.DeepCopyInto(out)
	return out
}

// Reset puts the given value back into the pool.
func (in *Inner) Reset() {
	if in == nil {
		return
	}
	if in.String != nil {
		in := &in.String
		s_stringSPool.Put(*in)
	}
	*in = Inner{}
	InnerPool.Put(in)
}

// ResetNoSelf puts the given field value back into the pool.
func (in *Inner) ResetNoSelf() {
	if in == nil {
		return
	}
	if in.String != nil {
		in := &in.String
		s_stringSPool.Put(*in)
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in *Inner) ResetOnlySelf() {
	if in == nil {
		return
	}
	*in = Inner{}
	InnerPool.Put(in)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ttest) DeepCopyInto(out *Ttest) {
	*out = *in
	if in.Inner1 != nil {
		in, out := &in.Inner1, &out.Inner1
		*out = Innerp_Pool.Get().(*Inner)
		(*in).DeepCopyInto(*out)
	}
	in.Inner2.DeepCopyInto(&out.Inner2)
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
	if in.Inner1 != nil {
		in := &in.Inner1
		(*in).ResetNoSelf()
		x := (*Inner)(*in)
		Innerp_Pool.Put(x)
	}
	in.Inner2.ResetNoSelf()
	*in = Ttest{}
	TtestPool.Put(in)
}

// ResetNoSelf puts the given field value back into the pool.
func (in *Ttest) ResetNoSelf() {
	if in == nil {
		return
	}
	if in.Inner1 != nil {
		in := &in.Inner1
		(*in).ResetNoSelf()
		x := (*Inner)(*in)
		Innerp_Pool.Put(x)
	}
	in.Inner2.ResetNoSelf()
}

// ResetOnlySelf puts the given value back into the pool.
func (in *Ttest) ResetOnlySelf() {
	if in == nil {
		return
	}
	*in = Ttest{}
	TtestPool.Put(in)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Unlock) DeepCopyInto(out *Unlock) {
	*out = *in
	if in.Progress != nil {
		in, out := &in.Progress, &out.Progress
		*out = map_int32_map_int64_int64__Pool.Get().(map[int32]map[int64]int64)
		for key, val := range *in {
			var outVal map[int64]int64
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = map_int64_int64_Pool.Get().(map[int64]int64)
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Unlocked != nil {
		in, out := &in.Unlocked, &out.Unlocked
		*out = map_int32_map_int64_bool__Pool.Get().(map[int32]map[int64]bool)
		for key, val := range *in {
			var outVal map[int64]bool
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = map_int64_bool_Pool.Get().(map[int64]bool)
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	in.DateFields.DeepCopyInto(&out.DateFields)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Unlock.
func (in *Unlock) DeepCopy() *Unlock {
	if in == nil {
		return nil
	}
	out := UnlockPool.Get().(*Unlock)
	in.DeepCopyInto(out)
	return out
}

// Reset puts the given value back into the pool.
func (in *Unlock) Reset() {
	if in == nil {
		return
	}
	if in.Progress != nil {
		in := &in.Progress
		for _, val := range *in {
			if val != nil {
				in := &val
				clear(*in)
				map_int64_int64_Pool.Put(*in)
			}
		}
		clear(*in)
		map_int32_map_int64_int64__Pool.Put(*in)
	}
	if in.Unlocked != nil {
		in := &in.Unlocked
		for _, val := range *in {
			if val != nil {
				in := &val
				clear(*in)
				map_int64_bool_Pool.Put(*in)
			}
		}
		clear(*in)
		map_int32_map_int64_bool__Pool.Put(*in)
	}
	in.DateFields.Reset()
	*in = Unlock{}
	UnlockPool.Put(in)
}

// ResetNoSelf puts the given field value back into the pool.
func (in *Unlock) ResetNoSelf() {
	if in == nil {
		return
	}
	if in.Progress != nil {
		in := &in.Progress
		for _, val := range *in {
			if val != nil {
				in := &val
				clear(*in)
				map_int64_int64_Pool.Put(*in)
			}
		}
		clear(*in)
		map_int32_map_int64_int64__Pool.Put(*in)
	}
	if in.Unlocked != nil {
		in := &in.Unlocked
		for _, val := range *in {
			if val != nil {
				in := &val
				clear(*in)
				map_int64_bool_Pool.Put(*in)
			}
		}
		clear(*in)
		map_int32_map_int64_bool__Pool.Put(*in)
	}
	in.DateFields.Reset()
}

// ResetOnlySelf puts the given value back into the pool.
func (in *Unlock) ResetOnlySelf() {
	if in == nil {
		return
	}
	*in = Unlock{}
	UnlockPool.Put(in)
}

var Innerp_Pool = sync.Pool{New: func() any { return new(Inner) }}
var InnerPool = sync.Pool{New: func() any { return new(Inner) }}
var TtestPool = sync.Pool{New: func() any { return new(Ttest) }}
var UnlockPool = sync.Pool{New: func() any { return new(Unlock) }}
var map_int32_map_int64_bool__Pool = sync.Pool{New: func() any { return make(map[int32]map[int64]bool, 16) }}
var map_int32_map_int64_int64__Pool = sync.Pool{New: func() any { return make(map[int32]map[int64]int64, 16) }}
var map_int64_bool_Pool = sync.Pool{New: func() any { return make(map[int64]bool, 16) }}
var map_int64_int64_Pool = sync.Pool{New: func() any { return make(map[int64]int64, 16) }}
var s_stringSPool = slicepool.NewSlicePool[string]()
var _ = slicepool.SlicePool[struct{}]{}
