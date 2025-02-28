package keepalive

import (
	"unsafe"
)

type KAState uint16

const (
	KAStateReset KAState = iota
	KAStateDelete
	KAStateRead
	KAStateSave
)

type KAElem struct {
	Data       unsafe.Pointer
	WriteIndex int64
	NoOrigin   int32
	State      KAState
	Wrote      uint16
}

type KeepAlive struct {
	Data      []KAElem
	SliceData [][]KAElem
	WI        []int64
	WSI       []int64
	SDUsed    []int
	DUsed     []int
}

func (ka *KeepAlive) Clear(sdUsed, dUsed []int) {
	clear(ka.Data)
	ka.Data = ka.Data[:0]
	for i := 0; i < len(ka.SliceData); i++ {
		clear(ka.SliceData[i])
		ka.SliceData[i] = ka.SliceData[i][:0]
	}
	ka.SliceData = ka.SliceData[:0]
	ka.WI = ka.WI[:0]
	ka.WSI = ka.WSI[:0]
	copy(ka.SDUsed, sdUsed)
	copy(ka.DUsed, dUsed)
}

func (ka *KeepAlive) AddSliceData() (*[]KAElem, int) {
	if len(ka.SliceData) < cap(ka.SliceData) {
		ka.SliceData = ka.SliceData[:len(ka.SliceData)+1]
		l := len(ka.SliceData) - 1
		return &ka.SliceData[l], l
	}
	ka.SliceData = append(ka.SliceData, make([]KAElem, 0, 16))
	l := len(ka.SliceData) - 1
	return &ka.SliceData[l], l
}

func (ka *KeepAlive) AddData() (*KAElem, int) {
	if len(ka.Data) < cap(ka.Data) {
		ka.Data = ka.Data[:len(ka.Data)+1]
		l := len(ka.Data) - 1
		return &ka.Data[l], l
	}
	ka.Data = append(ka.Data, KAElem{})
	l := len(ka.Data) - 1
	return &ka.Data[l], l
}

func SliceFindKA[T any](d []KAElem, f func(*T) bool) (*T, int) {
	for _, v := range d {
		if f((*T)(v.Data)) {
			if v.State > KAStateDelete {
				return (*T)(v.Data), 1
			}
			if v.State == KAStateDelete {
				return nil, -1
			}
		}
	}
	return nil, 0
}

func SliceFindKAForSave[T any](d []KAElem, f func(*T) bool) (int, bool) {
	for i, v := range d {
		if f((*T)(v.Data)) {
			return i, true
		}
	}
	return 0, false
}

func SliceFind[T any](s []T, f func(*T) bool) (*T, bool) {
	for i := range s {
		v := &s[i]
		if f(v) {
			return v, true
		}
	}
	return nil, false
}

func SliceFindPtr[T any](s []*T, f func(*T) bool) (*T, bool) {
	for _, v := range s {
		if f(v) {
			return v, true
		}
	}
	return nil, false
}

func SliceFindIndex[T any](s []*T, f func(*T) bool) int {
	for i := range s {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

func SwapDelete[T any](s []T, f func(*T) bool) ([]T, bool) {
	for i := range s {
		v := &s[i]
		if f(v) {
			l := len(s)
			s[i], s[l-1] = s[l-1], s[i]
			s = s[:l-1]
			return s, true
		}
	}
	return nil, false
}

func SwapDeletePtr[T any](s []*T, f func(*T) bool) ([]*T, bool) {
	for i := range s {
		v := s[i]
		if f(v) {
			l := len(s)
			s[i], s[l-1] = s[l-1], s[i]
			s = s[:l-1]
			return s, true
		}
	}
	return nil, false
}
