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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Domain) DeepCopyInto(out *Domain) {
	*out = *in
	if in.C != nil {
		in, out := &in.C, &out.C
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain.
func (in *Domain) DeepCopy() *Domain {
	if in == nil {
		return nil
	}
	out := new(Domain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fight) DeepCopyInto(out *Fight) {
	*out = *in
	if in.C != nil {
		in, out := &in.C, &out.C
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fight.
func (in *Fight) DeepCopy() *Fight {
	if in == nil {
		return nil
	}
	out := new(Fight)
	in.DeepCopyInto(out)
	return out
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
	out := new(History)
	in.DeepCopyInto(out)
	return out
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
	out := new(Item)
	in.DeepCopyInto(out)
	return out
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
	out := new(Player)
	in.DeepCopyInto(out)
	return out
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
	out := new(User)
	in.DeepCopyInto(out)
	return out
}
