/*
Copyright 2024 The Kubernetes Authors.

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

package outputtests

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"reflect"
	"testing"

	fuzz "github.com/google/gofuzz"
	"k8s.io/apimachinery/pkg/util/dump"

	"github.com/ravinggo/tools/deepcopy-gen/output_tests/aliases"
	"github.com/ravinggo/tools/deepcopy-gen/output_tests/builtins"
	"github.com/ravinggo/tools/deepcopy-gen/output_tests/interfaces"
	"github.com/ravinggo/tools/deepcopy-gen/output_tests/maps"
	"github.com/ravinggo/tools/deepcopy-gen/output_tests/pointer"
	"github.com/ravinggo/tools/deepcopy-gen/output_tests/slices"
	"github.com/ravinggo/tools/deepcopy-gen/output_tests/structs"
)

func TestWithValueFuzzer(t *testing.T) {
	tests := []interface{}{
		aliases.Ttest{},
		builtins.Ttest{},
		interfaces.Ttest{},
		maps.Ttest{},
		pointer.Ttest{},
		slices.Ttest{},
		structs.Ttest{},
	}

	fuzzer := fuzz.New()
	fuzzer.NilChance(0.5)
	fuzzer.NumElements(0, 2)
	fuzzer.Funcs(interfaceFuzzers...)
	for _, test := range tests {
		t.Run(
			fmt.Sprintf("%T", test), func(t *testing.T) {
				N := 1000
				for i := 0; i < N; i++ {
					original := reflect.New(reflect.TypeOf(test)).Interface()

					fuzzer.Fuzz(original)

					reflectCopy := ReflectDeepCopy(original)

					if !reflect.DeepEqual(original, reflectCopy) {
						t.Errorf(
							"original and reflectCopy are different:\n\n  original = %s\n\n  jsonCopy = %s", dump.Pretty(original), dump.Pretty(reflectCopy),
						)
					}

					deepCopy := reflect.ValueOf(original).MethodByName("DeepCopy").Call(nil)[0].Interface()

					if !reflect.DeepEqual(original, deepCopy) {
						t.Fatalf("original and deepCopy are different:\n\n  original = %s\n\n  deepCopy() = %s", dump.Pretty(original), dump.Pretty(deepCopy))
					}

					ValueFuzz(original)

					if !reflect.DeepEqual(reflectCopy, deepCopy) {
						t.Fatalf("reflectCopy and deepCopy are different:\n\n  origin = %s\n\n  jsonCopy() = %s", dump.Pretty(original), dump.Pretty(deepCopy))
					}
					reflect.ValueOf(deepCopy).MethodByName("Reset").Call(nil)
				}
			},
		)
	}
}

func BenchmarkReflectDeepCopy(b *testing.B) {
	fourtytwo := "fourtytwo"
	fourtytwoPtr := &fourtytwo
	var nilMap map[string]string
	var nilSlice []string
	mapPtr := &map[string]string{"0": "fourtytwo", "1": "fourtytwo"}
	slicePtr := &[]string{"fourtytwo", "fourtytwo", "fourtytwo"}
	structPtr := &pointer.Ttest{
		Builtin: &fourtytwo,
		Ptr:     &fourtytwoPtr,
	}

	tests := []interface{}{
		maps.Ttest{
			Byte:         map[string]byte{"0": 42, "1": 42, "3": 42},
			Int16:        map[string]int16{"0": 42, "1": 42, "3": 42},
			Int32:        map[string]int32{"0": 42, "1": 42, "3": 42},
			Int64:        map[string]int64{"0": 42, "1": 42, "3": 42},
			Uint8:        map[string]uint8{"0": 42, "1": 42, "3": 42},
			Uint16:       map[string]uint16{"0": 42, "1": 42, "3": 42},
			Uint32:       map[string]uint32{"0": 42, "1": 42, "3": 42},
			Uint64:       map[string]uint64{"0": 42, "1": 42, "3": 42},
			Float32:      map[string]float32{"0": 42.0, "1": 42.0, "3": 42.0},
			Float64:      map[string]float64{"0": 42, "1": 42, "3": 42},
			String:       map[string]string{"0": "fourtytwo", "1": "fourtytwo", "3": "fourtytwo"},
			StringPtr:    map[string]*string{"0": &fourtytwo, "1": &fourtytwo, "3": &fourtytwo},
			StringPtrPtr: map[string]**string{"0": &fourtytwoPtr, "1": &fourtytwoPtr, "3": &fourtytwoPtr},
			Map:          map[string]map[string]string{"0": nil, "1": {"a": fourtytwo, "b": fourtytwo}, "3": {}},
			MapPtr:       map[string]*map[string]string{"0": nil, "1": {"a": fourtytwo, "b": fourtytwo}, "3": &nilMap},
			Slice:        map[string][]string{"0": nil, "1": {"a", "b"}, "2": {}},
			SlicePtr:     map[string]*[]string{"0": nil, "1": {"a", "b"}, "2": &nilSlice},
			Struct:       map[string]maps.Ttest{"0": {}, "1": {Byte: map[string]byte{"0": 42, "1": 42, "3": 42}}},
			StructPtr:    map[string]*maps.Ttest{"0": nil, "1": {}, "2": {Byte: map[string]byte{"0": 42, "1": 42, "3": 42}}},
			StructTest2:  map[string]maps.Test2{"x": {}, "1": {}},
		},
		slices.Ttest{
			Byte:         []byte{42, 42, 42},
			Int16:        []int16{42, 42, 42},
			Int32:        []int32{42, 42, 42},
			Int64:        []int64{42, 42, 42},
			Uint8:        []uint8{42, 42, 42},
			Uint16:       []uint16{42, 42, 42},
			Uint32:       []uint32{42, 42, 42},
			Uint64:       []uint64{42, 42, 42},
			Float32:      []float32{42.0, 42.0, 42.0},
			Float64:      []float64{42, 42, 42},
			String:       []string{"fourtytwo", "fourtytwo", "fourtytwo"},
			StringPtr:    []*string{&fourtytwo, &fourtytwo, &fourtytwo},
			StringPtrPtr: []**string{&fourtytwoPtr, &fourtytwoPtr, &fourtytwoPtr},
			Map:          []map[string]string{nil, {"a": fourtytwo, "b": fourtytwo}, {}},
			MapPtr:       []*map[string]string{nil, {"a": fourtytwo, "b": fourtytwo}, &nilMap},
			Slice:        [][]string{nil, {"a", "b"}, {}},
			SlicePtr:     []*[]string{nil, {"a", "b"}, &nilSlice},
			Struct:       []slices.Ttest{{}, {Byte: []byte{42, 42, 42}}},
			StructPtr:    []*slices.Ttest{nil, {}, {Byte: []byte{42, 42, 42}}},
		},
		pointer.Ttest{
			Builtin:  &fourtytwo,
			Ptr:      &fourtytwoPtr,
			Map:      &map[string]string{"0": "fourtytwo", "1": "fourtytwo"},
			Slice:    &[]string{"fourtytwo", "fourtytwo", "fourtytwo"},
			MapPtr:   &mapPtr,
			SlicePtr: &slicePtr,
			Struct: &pointer.Ttest{
				Builtin: &fourtytwo,
				Ptr:     &fourtytwoPtr,
			},
			StructPtr: &structPtr,
		},
	}
	tests = tests[:1]
	fuzzer := fuzz.New()
	fuzzer.NilChance(0.5)
	fuzzer.NumElements(0, 2)
	fuzzer.Funcs(interfaceFuzzers...)
	go func() {
		http.ListenAndServe(":10111", nil)
	}()
	for _, test := range tests {
		b.Run(
			fmt.Sprintf("%T", test), func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					switch t := test.(type) {
					case maps.Ttest:
						x := t.DeepCopy()
						x.Reset()
					case slices.Ttest:
						x := t.DeepCopy()
						x.Reset()
					case pointer.Ttest:
						x := t.DeepCopy()
						x.Reset()
					default:
						b.Fatalf("missing type case in switch for %T", t)
					}
				}
			},
		)
	}
}

// func TestMapTTest(t *testing.T) {
// 	go func() {
// 		http.ListenAndServe(":10111", nil)
// 	}()
// 	fourtytwo := "fourtytwo"
// 	fourtytwoPtr := &fourtytwo
// 	var nilMap map[string]string
// 	var nilSlice []string
// 	tm := maps.Ttest{
// 		Byte:         map[string]byte{"0": 42, "1": 42, "3": 42},
// 		Int16:        map[string]int16{"0": 42, "1": 42, "3": 42},
// 		Int32:        map[string]int32{"0": 42, "1": 42, "3": 42},
// 		Int64:        map[string]int64{"0": 42, "1": 42, "3": 42},
// 		Uint8:        map[string]uint8{"0": 42, "1": 42, "3": 42},
// 		Uint16:       map[string]uint16{"0": 42, "1": 42, "3": 42},
// 		Uint32:       map[string]uint32{"0": 42, "1": 42, "3": 42},
// 		Uint64:       map[string]uint64{"0": 42, "1": 42, "3": 42},
// 		Float32:      map[string]float32{"0": 42.0, "1": 42.0, "3": 42.0},
// 		Float64:      map[string]float64{"0": 42, "1": 42, "3": 42},
// 		String:       map[string]string{"0": "fourtytwo", "1": "fourtytwo", "3": "fourtytwo"},
// 		StringPtr:    map[string]*string{"0": &fourtytwo, "1": &fourtytwo, "3": &fourtytwo},
// 		StringPtrPtr: map[string]**string{"0": &fourtytwoPtr, "1": &fourtytwoPtr, "3": &fourtytwoPtr},
// 		Map:          map[string]map[string]string{"0": nil, "1": {"a": fourtytwo, "b": fourtytwo}, "3": {}},
// 		MapPtr:       map[string]*map[string]string{"0": nil, "1": {"a": fourtytwo, "b": fourtytwo}, "3": &nilMap},
// 		Slice:        map[string][]string{"0": nil, "1": {"a", "b"}, "2": {}},
// 		SlicePtr:     map[string]*[]string{"0": nil, "1": {"a", "b"}, "2": &nilSlice},
// 		Struct:       map[string]maps.Ttest{"0": {}, "1": {Byte: map[string]byte{"0": 42, "1": 42, "3": 42}}},
// 		StructPtr:    map[string]*maps.Ttest{"0": nil, "1": {}, "2": {Byte: map[string]byte{"0": 42, "1": 42, "3": 42}}},
// 	}
// 	for {
// 		x := tm.DeepCopy()
// 		x.Reset()
// 	}
// }
