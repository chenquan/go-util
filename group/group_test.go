/*
 *    Copyright 2020 Chen Quan
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 */

package group

import (
	"reflect"
	"testing"
)

type List []int

func (l List) Size() int {
	return len(l)
}

func (l List) Get(i int) interface{} {
	return l[i]
}

func (l List) IsEmpty() bool {
	return len(l) == 0
}

const n = 100000000

func TestBy(t *testing.T) {
	list := make(List, 0)
	result := make(map[interface{}][]interface{})
	for i := 0; i < n; i++ {
		list = append(list, i)
		result[i%2] = append(result[i%2], i)
	}
	type args struct {
		grouper    Grouper
		functionBy func(interface{}) interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult map[interface{}][]interface{}
	}{
		{
			"1",
			args{
				list,
				func(obj interface{}) interface{} {
					return obj.(int) % 2
				},
			},
			result,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := By(tt.args.grouper, tt.args.functionBy); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("By() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func BenchmarkBy(b *testing.B) {
	list := make(List, 0)
	for i := 0; i < n; i++ {
		list = append(list, i)
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		By(list, func(obj interface{}) interface{} {
			return obj.(int) % 2
		})
	}
}
func BenchmarkBySlice(b *testing.B) {
	list := make([]interface{}, 0)
	for i := 0; i < n; i++ {
		list = append(list, i)
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		BySlice(list, func(obj interface{}) interface{} {
			return obj.(int) % 2
		})
	}
}

func BenchmarkByCount(b *testing.B) {
	list := make(List, 0)
	for i := 0; i < n; i++ {
		list = append(list, i)
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ByCount(list, func(obj interface{}) interface{} {
			return obj.(int) % 2
		})
	}
}

func BenchmarkBySliceCount(b *testing.B) {
	list := make([]interface{}, 0)
	for i := 0; i < n; i++ {
		list = append(list, i)
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		BySliceCount(list, func(obj interface{}) interface{} {
			return obj.(int) % 2
		})
	}
}
