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
	return l.Size() == 0
}

const n = 10000

func TestBy(t *testing.T) {
	list1 := make(List, 0)
	result1 := make(map[interface{}][]interface{})
	for i := 0; i < n; i++ {
		list1 = append(list1, i)
		result1[i%2] = append(result1[i%2], i)
	}
	list2 := make(List, 0)
	result2 := make(map[interface{}][]interface{})
	for i := 0; i < n; i++ {
		list2 = append(list2, i)
		result2[i%3] = append(result2[i%3], i)
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
				list1,
				func(obj interface{}) interface{} {
					return obj.(int) % 2
				},
			},
			result1,
		},
		{
			"2",
			args{
				list2,
				func(obj interface{}) interface{} {
					return obj.(int) % 3
				},
			},
			result2,
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
func TestByCount(t *testing.T) {
	list1 := make(List, 0)
	result1 := make(map[interface{}]int)
	for i := 0; i < n; i++ {
		list1 = append(list1, i)
		result1[i%2] += 1
	}
	list2 := make(List, 0)
	result2 := make(map[interface{}]int)
	for i := 0; i < n; i++ {
		list2 = append(list2, i)
		result2[i%3] += 1
	}

	type args struct {
		grouper    Grouper
		functionBy func(interface{}) interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult map[interface{}]int
	}{
		{
			"1",
			args{
				list1,
				func(obj interface{}) interface{} {
					return obj.(int) % 2
				},
			},
			result1,
		},
		{
			"2",
			args{
				list2,
				func(obj interface{}) interface{} {
					return obj.(int) % 3
				},
			},
			result2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ByCount(tt.args.grouper, tt.args.functionBy); !reflect.DeepEqual(gotResult, tt.wantResult) {
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

func TestBySlice(t *testing.T) {
	type args struct {
		grouper    []Grouper
		functionBy func(obj interface{}) interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantResults []map[interface{}][]interface{}
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := BySlice(tt.args.grouper, tt.args.functionBy); !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("BySlice() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}
