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
	"github.com/chenquan/go-util/backend/collection"
	"github.com/chenquan/go-util/errs"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewSliceListDefault(t *testing.T) {
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
		}, {
			"1",
			args{c: &SliceList{
				size: 2,
				data: []collection.Element{1, 2},
			}},
			&SliceList{
				size: 2,
				data: []collection.Element{1, 2},
			},
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
		}, {
			"3",
			args{initialCapacity: -1},
			&SliceList{
				size: 0,
				data: make([]collection.Element, 0, defaultCapacity),
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
			contains, _ := sliceList.Contains(tt.args.e)
			if got := contains; got != tt.want {
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
	_, _ = sliceList.Add("1")
	assert.Equal(t, 1, sliceList.size)
	assert.Equal(t, []collection.Element{"1"}, sliceList.data)

	_, _ = sliceList.Add("2")
	assert.Equal(t, 2, sliceList.size)
	assert.Equal(t, []collection.Element{"1", "2"}, sliceList.data)

	_, _ = sliceList.Add("")
	assert.Equal(t, 3, sliceList.size)
	assert.Equal(t, []collection.Element{"1", "2", ""}, sliceList.data)
	_, _ = sliceList.Add(nil)
	assert.Equal(t, 4, sliceList.size)
	assert.Equal(t, []collection.Element{"1", "2", "", nil}, sliceList.data)

}

func TestSliceList_Remove(t *testing.T) {
	sliceList := &SliceList{
		size: 2,
		data: []collection.Element{"1", 2},
	}
	var remove bool
	remove, _ = sliceList.Remove("1")
	assert.Equal(t, 1, sliceList.size)
	assert.Equal(t, true, remove)
	assert.Equal(t, []collection.Element{2}, sliceList.data)

	remove, _ = sliceList.Remove(12)
	assert.Equal(t, 1, sliceList.size)
	assert.Equal(t, false, remove)
	assert.Equal(t, []collection.Element{2}, sliceList.data)

	remove, _ = sliceList.Remove(2)
	assert.Equal(t, 0, sliceList.size)
	assert.Equal(t, true, remove)
	assert.Equal(t, []collection.Element{}, sliceList.data)

	remove, _ = sliceList.Remove(4)
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
		size: 2,
		data: []collection.Element{"1", "22"},
	}
	containsAll, _ := sliceList.ContainsAll(sliceList)
	assert.Equal(t, true, containsAll)
	containsAll, _ = sliceList.ContainsAll(c1)
	assert.Equal(t, true, containsAll)
	containsAll, _ = sliceList.ContainsAll(c2)
	assert.Equal(t, true, containsAll)
	containsAll, _ = sliceList.ContainsAll(c3)
	assert.Equal(t, false, containsAll)

}

func TestSliceList_AddAll(t *testing.T) {
	sliceList := &SliceList{
		size: 2,
		data: []collection.Element{"1", 2},
	}
	var b bool
	b, _ = sliceList.AddAll(sliceList)
	assert.Equal(t, true, b)
	assert.Equal(t, 4, sliceList.size)
	assert.Equal(t, []collection.Element{"1", 2, "1", 2}, sliceList.data)

	sliceList = &SliceList{
		size: 2,
		data: []collection.Element{"1", 2},
	}
	b, _ = sliceList.AddAll(&SliceList{
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
	modified, _ = sliceList.RemoveAll(sliceList)
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
	modified, _ = sliceList.RemoveAll(c1)
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
	modified, _ = sliceList.RemoveAll(c2)
	assert.Equal(t, true, modified)
	assert.Equal(t, 2, sliceList.size)
	assert.Equal(t, []collection.Element{2, 3}, sliceList.data)

}

func TestSliceList_RetainAll(t *testing.T) {
	sliceList := &SliceList{
		size: 3,
		data: []collection.Element{"1", 2, 3},
	}

	_, _ = sliceList.RetainAll(sliceList)
	assert.Equal(t, 3, sliceList.size)
	assert.Equal(t, []collection.Element{"1", 2, 3}, sliceList.data)

	sliceList1 := &SliceList{
		size: 2,
		data: []collection.Element{"1", 2},
	}
	_, _ = sliceList.RetainAll(sliceList1)
	assert.Equal(t, 2, sliceList.size)
	assert.Equal(t, []collection.Element{"1", 2}, sliceList.data)

	sliceList2 := &SliceList{
		size: 0,
		data: []collection.Element{},
	}
	_, _ = sliceList.RetainAll(sliceList2)
	assert.Equal(t, 0, sliceList.size)
	assert.Equal(t, []collection.Element{}, sliceList.data)

}

func TestSliceList_Clear(t *testing.T) {
	sliceList := &SliceList{
		size: 3,
		data: []collection.Element{"1", 2, 3},
	}
	_ = sliceList.Clear()
	assert.Equal(t, 0, sliceList.size)
	assert.Equal(t, []collection.Element{}, sliceList.data)
	assert.Equal(t, defaultCapacity, cap(sliceList.data))

}

func TestSliceList_Equals(t *testing.T) {
	sliceList := &SliceList{
		size: 3,
		data: []collection.Element{"1", 2, 3},
	}
	sliceList1 := &SliceList{
		size: 3,
		data: []collection.Element{"1", 2, 3},
	}
	sliceList2 := &SliceList{
		size: 2,
		data: []collection.Element{"1", 2},
	}
	sliceList3 := &SliceList{
		size: 3,
		data: []collection.Element{"1", 2, "3"},
	}
	assert.Equal(t, true, sliceList.Equals(sliceList))
	assert.Equal(t, true, sliceList.Equals(sliceList1))
	assert.Equal(t, false, sliceList.Equals(sliceList2))
	assert.Equal(t, false, sliceList.Equals(sliceList3))
}

func TestSliceList_AddAllIndex(t *testing.T) {
	sliceList := &SliceList{
		size: 3,
		data: []collection.Element{"1", 2, 3},
	}
	sliceList2 := &SliceList{
		size: 3,
		data: []collection.Element{"1", 2, 3},
	}
	var err error
	_, err = sliceList.AddAllIndex(0, sliceList2)
	assert.Equal(t, nil, err)
	assert.Equal(t, 6, sliceList.size)
	assert.Equal(t, []collection.Element{"1", 2, 3, "1", 2, 3}, sliceList.data)

	_, err = sliceList.AddAllIndex(sliceList.size, sliceList2)
	assert.Equal(t, nil, err)
	assert.Equal(t, 9, sliceList.size)
	assert.Equal(t, []collection.Element{"1", 2, 3, "1", 2, 3, "1", 2, 3}, sliceList.data)

	_, err = sliceList.AddAllIndex(1, sliceList2)
	assert.Equal(t, nil, err)
	assert.Equal(t, 12, sliceList.size)
	assert.Equal(t, []collection.Element{"1", "1", 2, 3, 2, 3, "1", 2, 3, "1", 2, 3}, sliceList.data)

	_, err = sliceList.AddAllIndex(-1, sliceList2)
	assert.Equal(t, errs.IndexOutOfBound, err)
	assert.Equal(t, 12, sliceList.size)
	assert.Equal(t, []collection.Element{"1", "1", 2, 3, 2, 3, "1", 2, 3, "1", 2, 3}, sliceList.data)

	_, err = sliceList.AddAllIndex(sliceList.size+1, sliceList2)
	assert.Equal(t, errs.IndexOutOfBound, err)
	assert.Equal(t, 12, sliceList.size)
	assert.Equal(t, []collection.Element{"1", "1", 2, 3, 2, 3, "1", 2, 3, "1", 2, 3}, sliceList.data)

	sliceList = &SliceList{
		size: 0,
		data: []collection.Element{},
	}
	_, err = sliceList.AddAllIndex(0, sliceList2)
	assert.Equal(t, nil, err)

}

func TestSliceList_Get(t *testing.T) {
	type fields struct {
		size int
		data []collection.Element
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantE   collection.Element
		wantErr bool
	}{
		{
			"1",
			fields{
				size: 2,
				data: []collection.Element{"1", "2"},
			},
			args{1},
			"2",
			false,
		}, {
			"2",
			fields{
				size: 4,
				data: []collection.Element{"1", "2", "2", "3"},
			},
			args{2},
			"2",
			false,
		}, {
			"3",
			fields{
				size: 0,
				data: []collection.Element{},
			},
			args{0},
			nil,
			true,
		}, {
			"4",
			fields{
				size: 0,
				data: []collection.Element{},
			},
			args{-1},
			nil,
			true,
		}, {
			"5",
			fields{
				size: 2,
				data: []collection.Element{"1", 222},
			},
			args{2},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceList := &SliceList{
				size: tt.fields.size,
				data: tt.fields.data,
			}
			gotE, err := sliceList.Get(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() errs = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotE, tt.wantE) {
				t.Errorf("Get() gotE = %v, want %v", gotE, tt.wantE)
			}
		})
	}
}

func TestSliceList_Set(t *testing.T) {
	sliceList := &SliceList{
		size: 3,
		data: []collection.Element{"1", 2, 3},
	}
	var err error
	_, err = sliceList.Set(0, "111")
	assert.Equal(t, nil, err)
	assert.Equal(t, []collection.Element{"111", 2, 3}, sliceList.data)
	assert.Equal(t, 3, sliceList.size)

	_, err = sliceList.Set(3, "111")
	assert.Equal(t, errs.IndexOutOfBound, err)
	assert.Equal(t, []collection.Element{"111", 2, 3}, sliceList.data)
	assert.Equal(t, 3, sliceList.size)

	sliceList = &SliceList{
		size: 0,
		data: []collection.Element{},
	}
	_, err = sliceList.Set(0, "1")
	assert.Equal(t, errs.IndexOutOfBound, err)
	assert.Equal(t, []collection.Element{}, sliceList.data)
	assert.Equal(t, 0, sliceList.size)
}

func TestSliceList_AddIndex(t *testing.T) {
	sliceList := &SliceList{
		size: 0,
		data: []collection.Element{},
	}
	var err error

	err = sliceList.AddIndex(0, 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, sliceList.size)
	assert.Equal(t, []collection.Element{1}, sliceList.data)

	err = sliceList.AddIndex(0, 1111)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, sliceList.size)
	assert.Equal(t, []collection.Element{1111, 1}, sliceList.data)

	err = sliceList.AddIndex(2, "1111")
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, sliceList.size)
	assert.Equal(t, []collection.Element{1111, 1, "1111"}, sliceList.data)

	err = sliceList.AddIndex(8, "1111")
	assert.Equal(t, errs.IndexOutOfBound, err)
	assert.Equal(t, 3, sliceList.size)
	assert.Equal(t, []collection.Element{1111, 1, "1111"}, sliceList.data)

}

func TestSliceList_RemoveIndex(t *testing.T) {
	type fields struct {
		size int
		data []collection.Element
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    collection.Element
		data    []collection.Element
		wantErr bool
	}{
		{
			"1",
			fields{
				size: 3,
				data: []collection.Element{"1", 2, "3"},
			},
			args{index: 1},
			2,
			[]collection.Element{"1", "3"},
			false,
		}, {
			"2",
			fields{
				size: 3,
				data: []collection.Element{"1", 2, "3"},
			},
			args{index: 3},
			nil,
			[]collection.Element{"1", 2, "3"},

			true,
		}, {
			"3",
			fields{
				size: 3,
				data: []collection.Element{"1", 2, "3"},
			},
			args{index: 0},
			"1",
			[]collection.Element{2, "3"},

			false,
		}, {
			"4",
			fields{
				size: 0,
				data: []collection.Element{},
			},
			args{index: 0},
			nil,
			[]collection.Element{},

			true,
		}, {
			"5",
			fields{
				size: 3,
				data: []collection.Element{"1", 2, "3"},
			},
			args{index: 2},
			"3",
			[]collection.Element{"1", 2},

			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceList := &SliceList{
				size: tt.fields.size,
				data: tt.fields.data,
			}
			got, err := sliceList.RemoveIndex(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveIndex() errs = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIndex() got = %v, want %v", got, tt.want)
			}
			assert.Equal(t, tt.data, sliceList.data)
		})
	}
}

func TestSliceList_Index(t *testing.T) {
	type fields struct {
		size int
		data []collection.Element
	}
	type args struct {
		e collection.Element
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantIndex int
	}{
		{
			"1",
			fields{
				size: 3,
				data: []collection.Element{"2", "1", "1"},
			},
			args{e: "1"},
			1,
		}, {
			"2",
			fields{
				size: 3,
				data: []collection.Element{"2", "1", "1"},
			},
			args{e: "22"},
			-1,
		}, {
			"3",
			fields{
				size: 0,
				data: []collection.Element{},
			},
			args{e: "22"},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceList := &SliceList{
				size: tt.fields.size,
				data: tt.fields.data,
			}
			if gotIndex := sliceList.Index(tt.args.e); gotIndex != tt.wantIndex {
				t.Errorf("Index() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestSliceList_LastIndex(t *testing.T) {
	type fields struct {
		size int
		data []collection.Element
	}
	type args struct {
		e collection.Element
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantIndex int
	}{{
		"1",
		fields{
			size: 3,
			data: []collection.Element{"2", "1", "1"},
		},
		args{e: "1"},
		2,
	}, {
		"2",
		fields{
			size: 3,
			data: []collection.Element{"2", "1", "1"},
		},
		args{e: "22"},
		-1,
	}, {
		"3",
		fields{
			size: 0,
			data: []collection.Element{},
		},
		args{e: "22"},
		-1,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceList := &SliceList{
				size: tt.fields.size,
				data: tt.fields.data,
			}
			if gotIndex := sliceList.LastIndex(tt.args.e); gotIndex != tt.wantIndex {
				t.Errorf("LastIndex() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestSliceList_SubList(t *testing.T) {
	type fields struct {
		size int
		data []collection.Element
	}
	type args struct {
		fromIndex int
		toIndex   int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantList collection.List
		wantErr  bool
	}{
		{
			"1",
			fields{
				size: 0,
				data: []collection.Element{},
			},
			args{
				fromIndex: 0,
				toIndex:   0,
			},
			&SliceList{0, []collection.Element{}},
			false,
		}, {
			"2",
			fields{
				size: 0,
				data: []collection.Element{},
			},
			args{
				fromIndex: 1,
				toIndex:   0,
			},
			nil,
			true,
		}, {
			"3",
			fields{
				size: 0,
				data: []collection.Element{},
			},
			args{
				fromIndex: -1,
				toIndex:   0,
			},
			nil,
			true,
		}, {
			"4",
			fields{
				size: 2,
				data: []collection.Element{1, 2},
			},
			args{
				fromIndex: 0,
				toIndex:   2,
			},
			&SliceList{2, []collection.Element{1, 2}},
			false,
		}, {
			"5",
			fields{
				size: 5,
				data: []collection.Element{1, 2, 3, 4, 5},
			},
			args{
				fromIndex: 2,
				toIndex:   2,
			},
			&SliceList{0, []collection.Element{}},
			false,
		}, {
			"6",
			fields{
				size: 5,
				data: []collection.Element{1, 2, 3, 4, 5},
			},
			args{
				fromIndex: 1,
				toIndex:   2,
			},
			&SliceList{1, []collection.Element{2}},
			false,
		}, {
			"7",
			fields{
				size: 5,
				data: []collection.Element{1, 2, 3, 4, 5},
			},
			args{
				fromIndex: 1,
				toIndex:   5,
			},
			&SliceList{4, []collection.Element{2, 3, 4, 5}},
			false,
		}, {
			"8",
			fields{
				size: 5,
				data: []collection.Element{1, 2, 3, 4, 5},
			},
			args{
				fromIndex: 1,
				toIndex:   6,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceList := &SliceList{
				size: tt.fields.size,
				data: tt.fields.data,
			}
			gotList, err := sliceList.SubList(tt.args.fromIndex, tt.args.toIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubList() errs = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("SubList() gotList = %v, want %v", gotList, tt.wantList)
			}
		})
	}
}
