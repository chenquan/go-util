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

func TestNewLinkedList(t *testing.T) {
	tests := []struct {
		name string
		want *LinkedList
	}{
		{
			"1",
			&LinkedList{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkedList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Add(t *testing.T) {
	list := &LinkedList{}
	_, _ = list.Add("111")
	assert.Equal(t, "111", list.first.elem)
	_, _ = list.Add("222")
	assert.Equal(t, "111", list.first.elem)
	assert.Equal(t, "222", list.first.next.elem)
	_, _ = list.Add("333")
	assert.Equal(t, "111", list.first.elem)
	assert.Equal(t, "333", list.first.next.next.elem)

}

func TestLinkedList_AddAll(t *testing.T) {
	list := &LinkedList{}
	l := &SliceList{size: 3, data: []collection.Element{"1", "2", 3}}
	var (
		b   bool
		err error
	)

	b, err = list.AddAll(l)
	assert.Equal(t, true, b)
	assert.Equal(t, nil, err)
	assert.Equal(t, "1", list.first.elem)
	assert.Equal(t, 3, list.last.elem)
	b, err = list.AddAll(l)
	assert.Equal(t, true, b)
	assert.Equal(t, nil, err)
	assert.Equal(t, "1", list.first.elem)
	assert.Equal(t, 3, list.last.elem)

}

func TestLinkedList_AddAllIndex(t *testing.T) {
	list := &LinkedList{}
	sliceList := &SliceList{
		size: 2,
		data: []collection.Element{1, "2"},
	}
	var (
		isAdd bool
		err   error
	)

	isAdd, err = list.AddAllIndex(0, sliceList)
	assert.Equal(t, true, isAdd)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, list.first.elem)
	assert.Equal(t, "2", list.first.next.elem)
	assert.Equal(t, (*node)(nil), list.first.next.next)

	isAdd, err = list.AddAllIndex(-1, sliceList)
	assert.Equal(t, false, isAdd)
	assert.Equal(t, errs.IndexOutOfBound, err)
	assert.Equal(t, 1, list.first.elem)
	assert.Equal(t, "2", list.first.next.elem)
	assert.Equal(t, (*node)(nil), list.first.next.next)

	sliceList1 := &SliceList{
		size: 0,
		data: []collection.Element{},
	}
	isAdd, err = list.AddAllIndex(1, sliceList1)
	assert.Equal(t, false, isAdd)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, list.first.elem)
	assert.Equal(t, "2", list.first.next.elem)
	assert.Equal(t, (*node)(nil), list.first.next.next)

	sliceList2 := &SliceList{
		size: 0,
		data: []collection.Element{},
	}
	isAdd, err = list.AddAllIndex(1, sliceList2)
	assert.Equal(t, false, isAdd)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, list.first.elem)
	assert.Equal(t, "2", list.first.next.elem)
	assert.Equal(t, (*node)(nil), list.first.next.next)

}

func TestLinkedList_AddFirst(t *testing.T) {
	list := &LinkedList{}
	_ = list.AddFirst("1")
	assert.Equal(t, "1", list.first.elem)
	assert.Equal(t, "1", list.last.elem)
	assert.Equal(t, 1, list.size)

	_ = list.AddFirst("2")
	assert.Equal(t, "2", list.first.elem)
	assert.Equal(t, "1", list.last.elem)
	assert.Equal(t, 2, list.size)

	_ = list.AddFirst("3")
	assert.Equal(t, "3", list.first.elem)
	assert.Equal(t, "1", list.last.elem)
	assert.Equal(t, "2", list.last.prev.elem)
	assert.Equal(t, 3, list.size)

	_ = list.AddFirst("4")
	assert.Equal(t, "4", list.first.elem)
	assert.Equal(t, "3", list.first.next.elem)
	assert.Equal(t, "1", list.last.elem)
	assert.Equal(t, 4, list.size)

}

func TestLinkedList_AddIndex(t *testing.T) {
	list := &LinkedList{}
	var err error
	err = list.AddIndex(0, 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, list.first.elem)
	assert.Equal(t, 1, list.last.elem)
	assert.Equal(t, 1, list.size)

	err = list.AddIndex(0, 2)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, list.first.elem)
	assert.Equal(t, 1, list.last.elem)
	assert.Equal(t, 2, list.size)

	err = list.AddIndex(3, 2)
	assert.Equal(t, errs.IndexOutOfBound, err)
	assert.Equal(t, 2, list.first.elem)
	assert.Equal(t, 1, list.last.elem)
	assert.Equal(t, 2, list.size)

	err = list.AddIndex(2, 3)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, list.first.elem)
	assert.Equal(t, 3, list.last.elem)
	assert.Equal(t, 3, list.size)

}

func TestLinkedList_AddLast(t *testing.T) {
	list := &LinkedList{}
	_ = list.AddLast("1")
	assert.Equal(t, "1", list.first.elem)
	assert.Equal(t, "1", list.last.elem)
	assert.Equal(t, 1, list.size)

	_ = list.AddLast("2")
	assert.Equal(t, "1", list.first.elem)
	assert.Equal(t, "2", list.first.next.elem)
	assert.Equal(t, (*node)(nil), list.first.next.next)
	assert.Equal(t, (*node)(nil), list.first.prev)
	assert.Equal(t, 2, list.size)

}

func TestLinkedList_Clear(t *testing.T) {

	n1 := &node{
		elem: 1,
		next: nil,
		prev: nil,
	}

	n2 := &node{
		elem: 2,
		next: nil,
		prev: n1,
	}
	n1.next = n2
	list := &LinkedList{2, n1, n2}

	_ = list.Clear()
	assert.Equal(t, 0, list.size)
	assert.Equal(t, (*node)(nil), list.first)
	assert.Equal(t, (*node)(nil), list.last)
}

func TestLinkedList_Contains(t *testing.T) {
	list := &LinkedList{}
	var (
		b   bool
		err error
	)
	b, err = list.AddAll(&SliceList{
		size: 3,
		data: []collection.Element{"1", 2, 3},
	})
	assert.Equal(t, true, b)
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, list.size)

	b, err = list.Contains("1")
	assert.Equal(t, true, b)
	assert.Equal(t, nil, err)

	b, err = list.Contains(1)
	assert.Equal(t, false, b)
	assert.Equal(t, nil, err)
}

func TestLinkedList_ContainsAll(t *testing.T) {
	list := &LinkedList{}
	var (
		b   bool
		err error
	)
	b, err = list.AddAll(&SliceList{
		size: 3,
		data: []collection.Element{"1", 2, 3, 4},
	})
	b, err = list.ContainsAll(list)
	assert.Equal(t, true, b)
	assert.Equal(t, nil, err)

	b, err = list.ContainsAll(&SliceList{
		size: 3,
		data: []collection.Element{"1", 2, 3},
	})
	assert.Equal(t, true, b)
	assert.Equal(t, nil, err)

	b, err = list.ContainsAll(&SliceList{
		size: 3,
		data: []collection.Element{"1", "2", 3},
	})
	assert.Equal(t, false, b)
	assert.Equal(t, nil, err)

}

func TestLinkedList_DescendingIterator(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	tests := []struct {
		name   string
		fields fields
		want   collection.Iterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if got := l.DescendingIterator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DescendingIterator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Element(t *testing.T) {
	list := &LinkedList{}
	var (
		err     error
		element collection.Element
	)
	element, err = list.Element()
	assert.Equal(t, errs.NoSuchElement, err)
	assert.Equal(t, nil, element)
	_, err = list.AddAll(&SliceList{
		size: 3,
		data: []collection.Element{"1", 2, 3, 4},
	})
	element, err = list.Element()
	assert.Equal(t, nil, err)
	assert.Equal(t, "1", element)
}

func TestLinkedList_Equals(t *testing.T) {
	list := &LinkedList{}
	var (
		equals bool
	)

	equals = list.Equals(list)
	assert.Equal(t, true, equals)

	list1 := &LinkedList{}
	equals = list.Equals(list1)
	assert.Equal(t, true, equals)

	list1.Add("1")
	equals = list.Equals(list1)
	assert.Equal(t, false, equals)

	list.Add("1")
	equals = list.Equals(list1)
	assert.Equal(t, true, equals)
}

func TestLinkedList_Get(t *testing.T) {
	list := &LinkedList{}
	var (
		element collection.Element
		err     error
	)

	element, err = list.Get(0)
	assert.Equal(t, errs.IndexOutOfBound, err)
	assert.Equal(t, nil, element)

	_, _ = list.Add("1")
	element, err = list.Get(0)
	assert.Equal(t, nil, err)
	assert.Equal(t, "1", element)

	_, _ = list.Add("2")
	element, err = list.Get(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, "2", element)

}

func TestLinkedList_GetFirst(t *testing.T) {
	list := &LinkedList{}
	var (
		err     error
		element collection.Element
	)
	element, err = list.GetFirst()
	assert.Equal(t, errs.NoSuchElement, err)
	assert.Equal(t, nil, element)
	_, err = list.AddAll(&SliceList{
		size: 4,
		data: []collection.Element{"1", 2, 3, 4},
	})
	element, err = list.GetFirst()
	assert.Equal(t, nil, err)
	assert.Equal(t, "1", element)

}

func TestLinkedList_GetLast(t *testing.T) {

	list := &LinkedList{}
	var (
		err     error
		element collection.Element
	)
	element, err = list.GetLast()
	assert.Equal(t, errs.NoSuchElement, err)
	assert.Equal(t, nil, element)
	_, err = list.AddAll(&SliceList{
		size: 4,
		data: []collection.Element{"1", 2, 3, 4},
	})
	element, err = list.GetLast()
	assert.Equal(t, nil, err)
	assert.Equal(t, 4, element)

}

func TestLinkedList_Index(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		e collection.Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if got := l.Index(tt.args.e); got != tt.want {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if got := l.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Iterator(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	tests := []struct {
		name   string
		fields fields
		want   collection.Iterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if got := l.Iterator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Iterator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_LastIndex(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		e collection.Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if got := l.LastIndex(tt.args.e); got != tt.want {
				t.Errorf("LastIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Offer(t *testing.T) {

}

func TestLinkedList_Peek(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	tests := []struct {
		name   string
		fields fields
		want   collection.Element
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if got := l.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Poll(t *testing.T) {

}

func TestLinkedList_Pop(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	tests := []struct {
		name    string
		fields  fields
		want    collection.Element
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			got, err := l.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Push(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		e collection.Element
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if err := l.Push(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_Remove(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		e collection.Element
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			got, err := l.Remove(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Remove() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_RemoveAll(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		collection collection.Collection
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			got, err := l.RemoveAll(tt.args.collection)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RemoveAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_RemoveFirst(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	tests := []struct {
		name    string
		fields  fields
		want    collection.Element
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			got, err := l.RemoveFirst()
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveFirst() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFirst() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_RemoveFirstOccurrence(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		e collection.Element
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			got, err := l.RemoveFirstOccurrence(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveFirstOccurrence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RemoveFirstOccurrence() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_RemoveIndex(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    collection.Element
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			got, err := l.RemoveIndex(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_RemoveLast(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	tests := []struct {
		name    string
		fields  fields
		want    collection.Element
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			got, err := l.RemoveLast()
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveLast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveLast() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_RemoveLastOccurrence(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		e collection.Element
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			got, err := l.RemoveLastOccurrence(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveLastOccurrence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RemoveLastOccurrence() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_RetainAll(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		collection collection.Collection
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			got, err := l.RetainAll(tt.args.collection)
			if (err != nil) != tt.wantErr {
				t.Errorf("RetainAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RetainAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Set(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		index int
		e     collection.Element
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    collection.Element
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			got, err := l.Set(tt.args.index, tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Size(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if got := l.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Slice(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	tests := []struct {
		name   string
		fields fields
		want   []collection.Element
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if got := l.Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_SubList(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		fromIndex int
		toIndex   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    collection.List
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			got, err := l.SubList(tt.args.fromIndex, tt.args.toIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_checkElementIndex(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if err := l.checkElementIndex(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("checkElementIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_checkPositionIndex(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if err := l.checkPositionIndex(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("checkPositionIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_getNode(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if got := l.getNode(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_isPositionIndex(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if got := l.isPositionIndex(tt.args.index); got != tt.want {
				t.Errorf("isPositionIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_linkBefore(t *testing.T) {

}

func TestLinkedList_linkLast(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		e collection.Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func TestLinkedList_unLink(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		x *node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   collection.Element
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				size:  tt.fields.size,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if got := l.unLink(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
