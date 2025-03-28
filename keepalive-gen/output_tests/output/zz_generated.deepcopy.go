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

package output

import (
	sync "sync"

	slicepool "github.com/ravinggo/tools/deepcopy-gen/slicepool"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Domain) DeepCopyInto(out *Domain) {
	*out = *in
	out.A = in.A
	if in.C != nil {
		in, out := &in.C, &out.C
		*out = s_stringSPool.Get(len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain.
func (in *Domain) DeepCopy() *Domain {
	if in == nil {
		return nil
	}
	out := DomainPool.Get().(*Domain)
	in.DeepCopyInto(out)
	return out
}

// Reset puts the given value back into the pool.
func (in *Domain) Reset() {
	if in == nil {
		return
	}
	if in.C != nil {
		in := &in.C
		s_stringSPool.Put(*in)
	}
	*in = Domain{}
	DomainPool.Put(in)
}

// ResetNoSelf puts the given field value back into the pool.
func (in *Domain) ResetNoSelf() {
	if in == nil {
		return
	}
	if in.C != nil {
		in := &in.C
		s_stringSPool.Put(*in)
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in *Domain) ResetOnlySelf() {
	if in == nil {
		return
	}
	*in = Domain{}
	DomainPool.Put(in)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fight) DeepCopyInto(out *Fight) {
	*out = *in
	if in.C != nil {
		in, out := &in.C, &out.C
		*out = s_stringSPool.Get(len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fight.
func (in *Fight) DeepCopy() *Fight {
	if in == nil {
		return nil
	}
	out := FightPool.Get().(*Fight)
	in.DeepCopyInto(out)
	return out
}

// Reset puts the given value back into the pool.
func (in *Fight) Reset() {
	if in == nil {
		return
	}
	if in.C != nil {
		in := &in.C
		s_stringSPool.Put(*in)
	}
	*in = Fight{}
	FightPool.Put(in)
}

// ResetNoSelf puts the given field value back into the pool.
func (in *Fight) ResetNoSelf() {
	if in == nil {
		return
	}
	if in.C != nil {
		in := &in.C
		s_stringSPool.Put(*in)
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in *Fight) ResetOnlySelf() {
	if in == nil {
		return
	}
	*in = Fight{}
	FightPool.Put(in)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *History) DeepCopyInto(out *History) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new History.
func (in *History) DeepCopy() *History {
	if in == nil {
		return nil
	}
	out := HistoryPool.Get().(*History)
	in.DeepCopyInto(out)
	return out
}

// Reset puts the given value back into the pool.
func (in *History) Reset() {
	if in == nil {
		return
	}
	*in = History{}
	HistoryPool.Put(in)
}

// ResetNoSelf puts the given field value back into the pool.
func (in *History) ResetNoSelf() {
	if in == nil {
		return
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in *History) ResetOnlySelf() {
	if in == nil {
		return
	}
	*in = History{}
	HistoryPool.Put(in)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Item) DeepCopyInto(out *Item) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Item.
func (in *Item) DeepCopy() *Item {
	if in == nil {
		return nil
	}
	out := ItemPool.Get().(*Item)
	in.DeepCopyInto(out)
	return out
}

// Reset puts the given value back into the pool.
func (in *Item) Reset() {
	if in == nil {
		return
	}
	*in = Item{}
	ItemPool.Put(in)
}

// ResetNoSelf puts the given field value back into the pool.
func (in *Item) ResetNoSelf() {
	if in == nil {
		return
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in *Item) ResetOnlySelf() {
	if in == nil {
		return
	}
	*in = Item{}
	ItemPool.Put(in)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Player) DeepCopyInto(out *Player) {
	*out = *in
	out.User = in.User
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Player.
func (in *Player) DeepCopy() *Player {
	if in == nil {
		return nil
	}
	out := PlayerPool.Get().(*Player)
	in.DeepCopyInto(out)
	return out
}

// Reset puts the given value back into the pool.
func (in *Player) Reset() {
	if in == nil {
		return
	}
	*in = Player{}
	PlayerPool.Put(in)
}

// ResetNoSelf puts the given field value back into the pool.
func (in *Player) ResetNoSelf() {
	if in == nil {
		return
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in *Player) ResetOnlySelf() {
	if in == nil {
		return
	}
	*in = Player{}
	PlayerPool.Put(in)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := UserPool.Get().(*User)
	in.DeepCopyInto(out)
	return out
}

// Reset puts the given value back into the pool.
func (in *User) Reset() {
	if in == nil {
		return
	}
	*in = User{}
	UserPool.Put(in)
}

// ResetNoSelf puts the given field value back into the pool.
func (in *User) ResetNoSelf() {
	if in == nil {
		return
	}
}

// ResetOnlySelf puts the given value back into the pool.
func (in *User) ResetOnlySelf() {
	if in == nil {
		return
	}
	*in = User{}
	UserPool.Put(in)
}

var DomainPool = sync.Pool{New: func() any { return new(Domain) }}
var FightPool = sync.Pool{New: func() any { return new(Fight) }}
var HistoryPool = sync.Pool{New: func() any { return new(History) }}
var ItemPool = sync.Pool{New: func() any { return new(Item) }}
var PlayerPool = sync.Pool{New: func() any { return new(Player) }}
var UserPool = sync.Pool{New: func() any { return new(User) }}
var s_stringSPool = slicepool.NewSlicePool[string]()
var _ = slicepool.SlicePool[struct{}]{}
