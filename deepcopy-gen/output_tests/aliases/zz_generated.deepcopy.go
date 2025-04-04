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

package aliases

import (
	sync "sync"

	slicepool "github.com/ravinggo/tools/deepcopy-gen/slicepool"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AliasInterfaceMap) DeepCopyInto(out *AliasInterfaceMap) {
	{
		in := &in
		*out = map_string_AliasInterface_Pool.Get().(map[string]AliasInterface)
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = val.DeepCopyAliasInterface()
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasInterfaceMap.
func (in AliasInterfaceMap) DeepCopy() AliasInterfaceMap {
	if in == nil {
		return nil
	}
	out := map_string_AliasInterface_Pool.Get().(map[string]AliasInterface)
	in.DeepCopyInto((*AliasInterfaceMap)(&out))
	return out
}

// Reset puts the given value back into the pool.
func (in AliasInterfaceMap) Reset() {
	{
		in := &in
		if in == nil {
			return
		}
		for _, val := range *in {
			if val != nil {
				val.ResetAliasInterface()
			}
		}
		clear(*in)
		x := (map[string]AliasInterface)(*in)
		map_string_AliasInterface_Pool.Put(x)
	}
}

// ResetNoSelf puts the given field value back into the pool.
func (in AliasInterfaceMap) ResetNoSelf() {
	{
		in := &in
		if in == nil {
			return
		}
		for _, val := range *in {
			if val != nil {
				val.ResetAliasInterface()
			}
		}
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in AliasInterfaceMap) ResetOnlySelf() {
	if in == nil {
		return
	}
	clear(in)
	x := (map[string]AliasInterface)(in)
	map_string_AliasInterface_Pool.Put(x)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AliasInterfaceSlice) DeepCopyInto(out *AliasInterfaceSlice) {
	{
		in := &in
		*out = s_AliasInterfaceSPool.Get(len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				(*out)[i] = (*in)[i].DeepCopyAliasInterface()
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasInterfaceSlice.
func (in AliasInterfaceSlice) DeepCopy() AliasInterfaceSlice {
	if in == nil {
		return nil
	}
	out := s_AliasInterfaceSPool.Get(len(in))
	in.DeepCopyInto((*AliasInterfaceSlice)(&out))
	return out
}

// Reset puts the given value back into the pool.
func (in AliasInterfaceSlice) Reset() {
	{
		in := &in
		if in == nil {
			return
		}
		for i := range *in {
			if (*in)[i] != nil {
				(*in)[i].ResetAliasInterface()
			}
		}
		x := ([]AliasInterface)(*in)
		s_AliasInterfaceSPool.Put(x)
	}
}

// ResetNoSelf puts the given field value back into the pool.
func (in AliasInterfaceSlice) ResetNoSelf() {
	{
		in := &in
		if in == nil {
			return
		}
		for i := range *in {
			if (*in)[i] != nil {
				(*in)[i].ResetAliasInterface()
			}
		}
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in AliasInterfaceSlice) ResetOnlySelf() {
	if in == nil {
		return
	}
	x := ([]AliasInterface)(in)
	s_AliasInterfaceSPool.Put(x)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AliasMap) DeepCopyInto(out *AliasMap) {
	{
		in := &in
		*out = map_string_int_Pool.Get().(map[string]int)
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasMap.
func (in AliasMap) DeepCopy() AliasMap {
	if in == nil {
		return nil
	}
	out := map_string_int_Pool.Get().(map[string]int)
	in.DeepCopyInto((*AliasMap)(&out))
	return out
}

// Reset puts the given value back into the pool.
func (in AliasMap) Reset() {
	{
		in := &in
		if in == nil {
			return
		}
		clear(*in)
		x := (map[string]int)(*in)
		map_string_int_Pool.Put(x)
	}
}

// ResetNoSelf puts the given field value back into the pool.
func (in AliasMap) ResetNoSelf() {
	{
		in := &in
		if in == nil {
			return
		}
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in AliasMap) ResetOnlySelf() {
	if in == nil {
		return
	}
	clear(in)
	x := (map[string]int)(in)
	map_string_int_Pool.Put(x)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AliasSlice) DeepCopyInto(out *AliasSlice) {
	{
		in := &in
		*out = s_intSPool.Get(len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasSlice.
func (in AliasSlice) DeepCopy() AliasSlice {
	if in == nil {
		return nil
	}
	out := s_intSPool.Get(len(in))
	in.DeepCopyInto((*AliasSlice)(&out))
	return out
}

// Reset puts the given value back into the pool.
func (in AliasSlice) Reset() {
	{
		in := &in
		if in == nil {
			return
		}
		x := ([]int)(*in)
		s_intSPool.Put(x)
	}
}

// ResetNoSelf puts the given field value back into the pool.
func (in AliasSlice) ResetNoSelf() {
	{
		in := &in
		if in == nil {
			return
		}
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in AliasSlice) ResetOnlySelf() {
	if in == nil {
		return
	}
	x := ([]int)(in)
	s_intSPool.Put(x)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasStruct) DeepCopyInto(out *AliasStruct) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasStruct.
func (in *AliasStruct) DeepCopy() *AliasStruct {
	if in == nil {
		return nil
	}
	out := AliasStructPool.Get().(*AliasStruct)
	in.DeepCopyInto(out)
	return out
}

// Reset puts the given value back into the pool.
func (in *AliasStruct) Reset() {
	if in == nil {
		return
	}
	*in = AliasStruct{}
	AliasStructPool.Put(in)
}

// ResetNoSelf puts the given field value back into the pool.
func (in *AliasStruct) ResetNoSelf() {
	if in == nil {
		return
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in *AliasStruct) ResetOnlySelf() {
	if in == nil {
		return
	}
	*in = AliasStruct{}
	AliasStructPool.Put(in)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Foo) DeepCopyInto(out *Foo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Foo.
func (in *Foo) DeepCopy() *Foo {
	if in == nil {
		return nil
	}
	out := FooPool.Get().(*Foo)
	in.DeepCopyInto(out)
	return out
}

// Reset puts the given value back into the pool.
func (in *Foo) Reset() {
	if in == nil {
		return
	}
	*in = Foo{}
	FooPool.Put(in)
}

// ResetNoSelf puts the given field value back into the pool.
func (in *Foo) ResetNoSelf() {
	if in == nil {
		return
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in *Foo) ResetOnlySelf() {
	if in == nil {
		return
	}
	*in = Foo{}
	FooPool.Put(in)
}

// DeepCopyAliasAliasInterface is an autogenerated deepcopy function, copying the receiver, creating a new AliasAliasInterface.
func (in *Foo) DeepCopyAliasAliasInterface() AliasAliasInterface {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyAliasInterface is an autogenerated deepcopy function, copying the receiver, creating a new AliasInterface.
func (in *Foo) DeepCopyAliasInterface() AliasInterface {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Interface.
func (in *Foo) DeepCopyInterface() Interface {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// ResetAliasAliasInterface is an autogenerated deepcopy function, copying the receiver, creating a new AliasAliasInterface.
func (in *Foo) ResetAliasAliasInterface() {
	in.Reset()
}

// ResetAliasInterface is an autogenerated deepcopy function, copying the receiver, creating a new AliasInterface.
func (in *Foo) ResetAliasInterface() {
	in.Reset()
}

// ResetInterface is an autogenerated deepcopy function, copying the receiver, creating a new Interface.
func (in *Foo) ResetInterface() {
	in.Reset()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FooAlias) DeepCopyInto(out *FooAlias) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FooAlias.
func (in *FooAlias) DeepCopy() *FooAlias {
	if in == nil {
		return nil
	}
	out := FooAliasPool.Get().(*FooAlias)
	in.DeepCopyInto(out)
	return out
}

// Reset puts the given value back into the pool.
func (in *FooAlias) Reset() {
	if in == nil {
		return
	}
	*in = FooAlias{}
	FooAliasPool.Put(in)
}

// ResetNoSelf puts the given field value back into the pool.
func (in *FooAlias) ResetNoSelf() {
	if in == nil {
		return
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in *FooAlias) ResetOnlySelf() {
	if in == nil {
		return
	}
	*in = FooAlias{}
	FooAliasPool.Put(in)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FooMap) DeepCopyInto(out *FooMap) {
	{
		in := &in
		*out = map_string_Foo_Pool.Get().(map[string]Foo)
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FooMap.
func (in FooMap) DeepCopy() FooMap {
	if in == nil {
		return nil
	}
	out := map_string_Foo_Pool.Get().(map[string]Foo)
	in.DeepCopyInto((*FooMap)(&out))
	return out
}

// Reset puts the given value back into the pool.
func (in FooMap) Reset() {
	{
		in := &in
		if in == nil {
			return
		}
		clear(*in)
		x := (map[string]Foo)(*in)
		map_string_Foo_Pool.Put(x)
	}
}

// ResetNoSelf puts the given field value back into the pool.
func (in FooMap) ResetNoSelf() {
	{
		in := &in
		if in == nil {
			return
		}
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in FooMap) ResetOnlySelf() {
	if in == nil {
		return
	}
	clear(in)
	x := (map[string]Foo)(in)
	map_string_Foo_Pool.Put(x)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FooSlice) DeepCopyInto(out *FooSlice) {
	{
		in := &in
		*out = s_FooSPool.Get(len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FooSlice.
func (in FooSlice) DeepCopy() FooSlice {
	if in == nil {
		return nil
	}
	out := s_FooSPool.Get(len(in))
	in.DeepCopyInto((*FooSlice)(&out))
	return out
}

// Reset puts the given value back into the pool.
func (in FooSlice) Reset() {
	{
		in := &in
		if in == nil {
			return
		}
		for i := range *in {
			(*in)[i].ResetNoSelf()
		}
		x := ([]Foo)(*in)
		s_FooSPool.Put(x)
	}
}

// ResetNoSelf puts the given field value back into the pool.
func (in FooSlice) ResetNoSelf() {
	{
		in := &in
		if in == nil {
			return
		}
		for i := range *in {
			(*in)[i].ResetNoSelf()
		}
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in FooSlice) ResetOnlySelf() {
	if in == nil {
		return
	}
	x := ([]Foo)(in)
	s_FooSPool.Put(x)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Map) DeepCopyInto(out *Map) {
	{
		in := &in
		*out = map_string_int_Pool.Get().(map[string]int)
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Map.
func (in Map) DeepCopy() Map {
	if in == nil {
		return nil
	}
	out := map_string_int_Pool.Get().(map[string]int)
	in.DeepCopyInto((*Map)(&out))
	return out
}

// Reset puts the given value back into the pool.
func (in Map) Reset() {
	{
		in := &in
		if in == nil {
			return
		}
		clear(*in)
		x := (map[string]int)(*in)
		map_string_int_Pool.Put(x)
	}
}

// ResetNoSelf puts the given field value back into the pool.
func (in Map) ResetNoSelf() {
	{
		in := &in
		if in == nil {
			return
		}
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in Map) ResetOnlySelf() {
	if in == nil {
		return
	}
	clear(in)
	x := (map[string]int)(in)
	map_string_int_Pool.Put(x)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Slice) DeepCopyInto(out *Slice) {
	{
		in := &in
		*out = s_intSPool.Get(len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Slice.
func (in Slice) DeepCopy() Slice {
	if in == nil {
		return nil
	}
	out := s_intSPool.Get(len(in))
	in.DeepCopyInto((*Slice)(&out))
	return out
}

// Reset puts the given value back into the pool.
func (in Slice) Reset() {
	{
		in := &in
		if in == nil {
			return
		}
		x := ([]int)(*in)
		s_intSPool.Put(x)
	}
}

// ResetNoSelf puts the given field value back into the pool.
func (in Slice) ResetNoSelf() {
	{
		in := &in
		if in == nil {
			return
		}
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in Slice) ResetOnlySelf() {
	if in == nil {
		return
	}
	x := ([]int)(in)
	s_intSPool.Put(x)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct) DeepCopyInto(out *Struct) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct.
func (in *Struct) DeepCopy() *Struct {
	if in == nil {
		return nil
	}
	out := StructPool.Get().(*Struct)
	in.DeepCopyInto(out)
	return out
}

// Reset puts the given value back into the pool.
func (in *Struct) Reset() {
	if in == nil {
		return
	}
	*in = Struct{}
	StructPool.Put(in)
}

// ResetNoSelf puts the given field value back into the pool.
func (in *Struct) ResetNoSelf() {
	if in == nil {
		return
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in *Struct) ResetOnlySelf() {
	if in == nil {
		return
	}
	*in = Struct{}
	StructPool.Put(in)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ttest) DeepCopyInto(out *Ttest) {
	*out = *in
	if in.Slice != nil {
		in, out := &in.Slice, &out.Slice
		*out = s_intSPool.Get(len(*in))
		copy(*out, *in)
	}
	if in.Pointer != nil {
		in, out := &in.Pointer, &out.Pointer
		*out = intp_Pool.Get().(*int)
		**out = **in
	}
	if in.PointerAlias != nil {
		in, out := &in.PointerAlias, &out.PointerAlias
		*out = Builtinp_Pool.Get().(*Builtin)
		**out = **in
	}
	out.Struct = in.Struct
	if in.Map != nil {
		in, out := &in.Map, &out.Map
		*out = map_string_int_Pool.Get().(map[string]int)
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SliceSlice != nil {
		in, out := &in.SliceSlice, &out.SliceSlice
		*out = s_SliceSPool.Get(len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = s_intSPool.Get(len(*in))
				copy(*out, *in)
			}
		}
	}
	if in.MapSlice != nil {
		in, out := &in.MapSlice, &out.MapSlice
		*out = map_string_Slice_Pool.Get().(map[string]Slice)
		for key, val := range *in {
			var outVal []int
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = s_intSPool.Get(len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	out.FooAlias = in.FooAlias
	if in.FooSlice != nil {
		in, out := &in.FooSlice, &out.FooSlice
		*out = s_FooSPool.Get(len(*in))
		copy(*out, *in)
	}
	if in.FooPointer != nil {
		in, out := &in.FooPointer, &out.FooPointer
		*out = Foop_Pool.Get().(*Foo)
		**out = **in
	}
	if in.FooMap != nil {
		in, out := &in.FooMap, &out.FooMap
		*out = map_string_Foo_Pool.Get().(map[string]Foo)
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AliasSlice != nil {
		in, out := &in.AliasSlice, &out.AliasSlice
		*out = s_intSPool.Get(len(*in))
		copy(*out, *in)
	}
	if in.AliasPointer != nil {
		in, out := &in.AliasPointer, &out.AliasPointer
		*out = intp_Pool.Get().(*int)
		**out = **in
	}
	out.AliasStruct = in.AliasStruct
	if in.AliasMap != nil {
		in, out := &in.AliasMap, &out.AliasMap
		*out = map_string_int_Pool.Get().(map[string]int)
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AliasInterface != nil {
		out.AliasInterface = in.AliasInterface.DeepCopyAliasInterface()
	}
	if in.AliasAliasInterface != nil {
		out.AliasAliasInterface = in.AliasAliasInterface.DeepCopyAliasAliasInterface()
	}
	if in.AliasInterfaceMap != nil {
		in, out := &in.AliasInterfaceMap, &out.AliasInterfaceMap
		*out = map_string_AliasInterface_Pool.Get().(map[string]AliasInterface)
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = val.DeepCopyAliasInterface()
			}
		}
	}
	if in.AliasInterfaceSlice != nil {
		in, out := &in.AliasInterfaceSlice, &out.AliasInterfaceSlice
		*out = s_AliasInterfaceSPool.Get(len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				(*out)[i] = (*in)[i].DeepCopyAliasInterface()
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
	if in.Slice != nil {
		in := &in.Slice
		s_intSPool.Put(*in)
	}
	if in.Pointer != nil {
		in := &in.Pointer
		x := (*int)(*in)
		intp_Pool.Put(x)
	}
	if in.PointerAlias != nil {
		in := &in.PointerAlias
		x := (*Builtin)(*in)
		Builtinp_Pool.Put(x)
	}
	if in.Map != nil {
		in := &in.Map
		clear(*in)
		x := (map[string]int)(*in)
		map_string_int_Pool.Put(x)
	}
	if in.SliceSlice != nil {
		in := &in.SliceSlice
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				s_intSPool.Put(*in)
			}
		}
		s_SliceSPool.Put(*in)
	}
	if in.MapSlice != nil {
		in := &in.MapSlice
		for _, val := range *in {
			if cap(val) != 0 {
				in := &val
				s_intSPool.Put(*in)
			}
		}
		clear(*in)
		map_string_Slice_Pool.Put(*in)
	}
	if in.FooSlice != nil {
		in := &in.FooSlice
		for i := range *in {
			(*in)[i].ResetNoSelf()
		}
		s_FooSPool.Put(*in)
	}
	if in.FooPointer != nil {
		in := &in.FooPointer
		x := (*Foo)(*in)
		Foop_Pool.Put(x)
	}
	if in.FooMap != nil {
		in := &in.FooMap
		clear(*in)
		x := (map[string]Foo)(*in)
		map_string_Foo_Pool.Put(x)
	}
	if in.AliasSlice != nil {
		in := &in.AliasSlice
		s_intSPool.Put(*in)
	}
	if in.AliasPointer != nil {
		in := &in.AliasPointer
		x := (*int)(*in)
		intp_Pool.Put(x)
	}
	if in.AliasMap != nil {
		in := &in.AliasMap
		clear(*in)
		x := (map[string]int)(*in)
		map_string_int_Pool.Put(x)
	}
	if in.AliasInterface != nil {
		in.AliasInterface.ResetAliasInterface()
	}
	if in.AliasAliasInterface != nil {
		in.AliasAliasInterface.ResetAliasAliasInterface()
	}
	if in.AliasInterfaceMap != nil {
		in := &in.AliasInterfaceMap
		for _, val := range *in {
			if val != nil {
				val.ResetAliasInterface()
			}
		}
		clear(*in)
		x := (map[string]AliasInterface)(*in)
		map_string_AliasInterface_Pool.Put(x)
	}
	if in.AliasInterfaceSlice != nil {
		in := &in.AliasInterfaceSlice
		for i := range *in {
			if (*in)[i] != nil {
				(*in)[i].ResetAliasInterface()
			}
		}
		s_AliasInterfaceSPool.Put(*in)
	}
	*in = Ttest{}
	TtestPool.Put(in)
}

// ResetNoSelf puts the given field value back into the pool.
func (in *Ttest) ResetNoSelf() {
	if in == nil {
		return
	}
	if in.Slice != nil {
		in := &in.Slice
		s_intSPool.Put(*in)
	}
	if in.Pointer != nil {
		in := &in.Pointer
		x := (*int)(*in)
		intp_Pool.Put(x)
	}
	if in.PointerAlias != nil {
		in := &in.PointerAlias
		x := (*Builtin)(*in)
		Builtinp_Pool.Put(x)
	}
	if in.Map != nil {
		in := &in.Map
		clear(*in)
		x := (map[string]int)(*in)
		map_string_int_Pool.Put(x)
	}
	if in.SliceSlice != nil {
		in := &in.SliceSlice
		for i := range *in {
			if (*in)[i] != nil {
				in := &(*in)[i]
				s_intSPool.Put(*in)
			}
		}
		s_SliceSPool.Put(*in)
	}
	if in.MapSlice != nil {
		in := &in.MapSlice
		for _, val := range *in {
			if cap(val) != 0 {
				in := &val
				s_intSPool.Put(*in)
			}
		}
		clear(*in)
		map_string_Slice_Pool.Put(*in)
	}
	if in.FooSlice != nil {
		in := &in.FooSlice
		for i := range *in {
			(*in)[i].ResetNoSelf()
		}
		s_FooSPool.Put(*in)
	}
	if in.FooPointer != nil {
		in := &in.FooPointer
		x := (*Foo)(*in)
		Foop_Pool.Put(x)
	}
	if in.FooMap != nil {
		in := &in.FooMap
		clear(*in)
		x := (map[string]Foo)(*in)
		map_string_Foo_Pool.Put(x)
	}
	if in.AliasSlice != nil {
		in := &in.AliasSlice
		s_intSPool.Put(*in)
	}
	if in.AliasPointer != nil {
		in := &in.AliasPointer
		x := (*int)(*in)
		intp_Pool.Put(x)
	}
	if in.AliasMap != nil {
		in := &in.AliasMap
		clear(*in)
		x := (map[string]int)(*in)
		map_string_int_Pool.Put(x)
	}
	if in.AliasInterface != nil {
		in.AliasInterface.ResetAliasInterface()
	}
	if in.AliasAliasInterface != nil {
		in.AliasAliasInterface.ResetAliasAliasInterface()
	}
	if in.AliasInterfaceMap != nil {
		in := &in.AliasInterfaceMap
		for _, val := range *in {
			if val != nil {
				val.ResetAliasInterface()
			}
		}
		clear(*in)
		x := (map[string]AliasInterface)(*in)
		map_string_AliasInterface_Pool.Put(x)
	}
	if in.AliasInterfaceSlice != nil {
		in := &in.AliasInterfaceSlice
		for i := range *in {
			if (*in)[i] != nil {
				(*in)[i].ResetAliasInterface()
			}
		}
		s_AliasInterfaceSPool.Put(*in)
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in *Ttest) ResetOnlySelf() {
	if in == nil {
		return
	}
	*in = Ttest{}
	TtestPool.Put(in)
}

var intp_Pool = sync.Pool{New: func() any { return new(int) }}
var AliasStructPool = sync.Pool{New: func() any { return new(AliasStruct) }}
var FooPool = sync.Pool{New: func() any { return new(Foo) }}
var FooAliasPool = sync.Pool{New: func() any { return new(FooAlias) }}
var Foop_Pool = sync.Pool{New: func() any { return new(Foo) }}
var Builtinp_Pool = sync.Pool{New: func() any { return new(Builtin) }}
var StructPool = sync.Pool{New: func() any { return new(Struct) }}
var TtestPool = sync.Pool{New: func() any { return new(Ttest) }}
var map_string_AliasInterface_Pool = sync.Pool{New: func() any { return make(map[string]AliasInterface, 16) }}
var map_string_Foo_Pool = sync.Pool{New: func() any { return make(map[string]Foo, 16) }}
var map_string_Slice_Pool = sync.Pool{New: func() any { return make(map[string]Slice, 16) }}
var map_string_int_Pool = sync.Pool{New: func() any { return make(map[string]int, 16) }}
var s_AliasInterfaceSPool = slicepool.NewSlicePool[AliasInterface]()
var s_FooSPool = slicepool.NewSlicePool[Foo]()
var s_SliceSPool = slicepool.NewSlicePool[Slice]()
var s_intSPool = slicepool.NewSlicePool[int]()
var _ = slicepool.SlicePool[struct{}]{}
