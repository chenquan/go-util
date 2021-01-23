/*
 *    Copyright 2021 Chen Quan
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

package list

import (
	"fmt"
	"github.com/chenquan/go-utils/collection/api/collection"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewSliceListDefault(t *testing.T) {
	l := []int{1, 2, 4}
	ints := l[3:]
	fmt.Println(ints)
}

func TestNewSliceListDefault1(t *testing.T) {
	tests := []struct {
		name string
		want *SliceList
	}{
		{
			"1",
			&SliceList{
				size: 0,
				data: make([]collection.Element, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSliceListDefault(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSliceListDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSliceListWithCollection(t *testing.T) {
	type args struct {
		c collection.Collection
	}
	tests := []struct {
		name string
		args args
		want *SliceList
	}{
		{
			"1",
			args{c: NewSliceListDefault()},
			NewSliceListDefault(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSliceListWithCollection(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSliceListWithCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSliceList(t *testing.T) {
	type args struct {
		initialCapacity int
	}
	tests := []struct {
		name string
		args args
		want *SliceList
	}{
		{
			"1",
			args{initialCapacity: 16},
			&SliceList{
				size: 0,
				data: make([]collection.Element, 0, 16),
			},
		}, {
			"2",
			args{initialCapacity: 0},
			&SliceList{
				size: 0,
				data: make([]collection.Element, 0, 0),
			},
		}, {
			"3",
			args{initialCapacity: 200},
			&SliceList{
				size: 0,
				data: make([]collection.Element, 0, 200),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSliceList(tt.args.initialCapacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSliceList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceList_Slice(t *testing.T) {
	type fields struct {
		size int
		data []collection.Element
	}
	tests := []struct {
		name   string
		fields fields
		want   []collection.Element
	}{
		{
			"1",
			fields{
				size: 0,
				data: make([]collection.Element, 0),
			},
			[]collection.Element{},
		}, {
			"2",
			fields{
				size: 3,
				data: []collection.Element{1, 2, 3},
			},
			[]collection.Element{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceList := &SliceList{
				size: tt.fields.size,
				data: tt.fields.data,
			}
			if got := sliceList.Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceList_Size(t *testing.T) {
	type fields struct {
		size int
		data []collection.Element
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"1",
			fields{
				size: 1,
				data: []collection.Element{1},
			},
			1,
		}, {
			"2",
			fields{
				size: 2,
				data: []collection.Element{1, 2},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceList := &SliceList{
				size: tt.fields.size,
				data: tt.fields.data,
			}
			if got := sliceList.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceList_IsEmpty(t *testing.T) {
	type fields struct {
		size int
		data []collection.Element
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"1",
			fields{
				size: 0,
				data: make([]collection.Element, 0),
			},
			true,
		}, {
			"2",
			fields{
				size: 1,
				data: []collection.Element{1},
			},
			false,
		}, {
			"3",
			fields{
				size: 2,
				data: []collection.Element{1, 2},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceList := &SliceList{
				size: tt.fields.size,
				data: tt.fields.data,
			}
			if got := sliceList.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceList_Contains(t *testing.T) {
	type fields struct {
		size int
		data []collection.Element
	}
	type args struct {
		e collection.Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"1",
			fields{
				size: 2,
				data: []collection.Element{"1", "2"},
			},
			args{e: "1"},
			true,
		}, {
			"2",
			fields{
				size: 2,
				data: []collection.Element{"1", ""},
			},
			args{e: ""},
			true,
		}, {
			"3",
			fields{
				size: 2,
				data: []collection.Element{"1", ""},
			},
			args{e: "2"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceList := &SliceList{
				size: tt.fields.size,
				data: tt.fields.data,
			}
			if got := sliceList.Contains(tt.args.e); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceList_Add(t *testing.T) {
	sliceList := &SliceList{
		size: 0,
		data: []collection.Element{},
	}
	sliceList.Add("1")
	assert.Equal(t, 1, sliceList.size)
	assert.Equal(t, []collection.Element{"1"}, sliceList.data)

	sliceList.Add("2")
	assert.Equal(t, 2, sliceList.size)
	assert.Equal(t, []collection.Element{"1", "2"}, sliceList.data)

	sliceList.Add("")
	assert.Equal(t, 3, sliceList.size)
	assert.Equal(t, []collection.Element{"1", "2", ""}, sliceList.data)
	sliceList.Add(nil)
	assert.Equal(t, 4, sliceList.size)
	assert.Equal(t, []collection.Element{"1", "2", "", nil}, sliceList.data)

}

func TestSliceList_Remove(t *testing.T) {
	sliceList := &SliceList{
		size: 2,
		data: []collection.Element{"1", 2},
	}
	var remove bool
	remove = sliceList.Remove("1")
	assert.Equal(t, 1, sliceList.size)
	assert.Equal(t, true, remove)
	assert.Equal(t, []collection.Element{2}, sliceList.data)

	remove = sliceList.Remove(12)
	assert.Equal(t, 1, sliceList.size)
	assert.Equal(t, false, remove)
	assert.Equal(t, []collection.Element{2}, sliceList.data)

	remove = sliceList.Remove(2)
	assert.Equal(t, 0, sliceList.size)
	assert.Equal(t, true, remove)
	assert.Equal(t, []collection.Element{}, sliceList.data)

	remove = sliceList.Remove(4)
	assert.Equal(t, 0, sliceList.size)
	assert.Equal(t, false, remove)
	assert.Equal(t, []collection.Element{}, sliceList.data)

}

func TestSliceList_ContainsAll(t *testing.T) {
	sliceList := &SliceList{
		size: 2,
		data: []collection.Element{"1", 2},
	}
	c1 := &SliceList{
		size: 2,
		data: []collection.Element{"1", 2},
	}
	c2 := &SliceList{
		size: 1,
		data: []collection.Element{"1"},
	}
	c3 := &SliceList{
		size: 1,
		data: []collection.Element{"1", "22"},
	}

	assert.Equal(t, true, sliceList.ContainsAll(sliceList))
	assert.Equal(t, true, sliceList.ContainsAll(c1))
	assert.Equal(t, true, sliceList.ContainsAll(c2))
	assert.Equal(t, false, sliceList.ContainsAll(c3))

}

func TestSliceList_AddAll(t *testing.T) {
	sliceList := &SliceList{
		size: 2,
		data: []collection.Element{"1", 2},
	}
	var b bool
	b = sliceList.AddAll(sliceList)
	assert.Equal(t, true, b)
	assert.Equal(t, 4, sliceList.size)
	assert.Equal(t, []collection.Element{"1", 2, "1", 2}, sliceList.data)

	sliceList = &SliceList{
		size: 2,
		data: []collection.Element{"1", 2},
	}
	b = sliceList.AddAll(&SliceList{
		size: 0,
		data: []collection.Element{},
	})
	assert.Equal(t, false, b)
	assert.Equal(t, 2, sliceList.size)
	assert.Equal(t, []collection.Element{"1", 2}, sliceList.data)

}

func TestSliceList_RemoveAll(t *testing.T) {
	sliceList := &SliceList{
		size: 3,
		data: []collection.Element{"1", 2, 3},
	}
	var modified bool
	modified = sliceList.RemoveAll(sliceList)
	assert.Equal(t, true, modified)
	assert.Equal(t, 0, sliceList.size)
	assert.Equal(t, []collection.Element{}, sliceList.data)

	sliceList = &SliceList{
		size: 3,
		data: []collection.Element{"1", 2, 3},
	}
	c1 := &SliceList{
		size: 1,
		data: []collection.Element{"1"},
	}
	modified = sliceList.RemoveAll(c1)
	assert.Equal(t, true, modified)
	assert.Equal(t, 2, sliceList.size)
	assert.Equal(t, []collection.Element{2, 3}, sliceList.data)

	sliceList = &SliceList{
		size: 3,
		data: []collection.Element{"1", 2, 3},
	}
	c2 := &SliceList{
		size: 1,
		data: []collection.Element{"1", 4, 34, 34, 34, 34, 1, 43, 4, 34, 3, 4},
	}
	modified = sliceList.RemoveAll(c2)
	assert.Equal(t, true, modified)
	assert.Equal(t, 2, sliceList.size)
	assert.Equal(t, []collection.Element{2, 3}, sliceList.data)

}
