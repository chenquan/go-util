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

func genLinkedList(elements ...collection.Element) *LinkedList {
	var first, last, p *node
	for _, element := range elements {

		n := &node{
			elem: element,
			next: nil,
			prev: last,
		}
		p = last
		last = n
		if first == nil {
			first = n
		} else {
			p.next = n
		}
	}
	return &LinkedList{
		size:  len(elements),
		first: first,
		last:  last,
	}
}
func linkedToSlice(list *LinkedList) []collection.Element {
	elements := make([]collection.Element, 0, list.size)
	for x := list.first; x != nil; x = x.next {
		elements = append(elements, x.elem)
	}
	return elements
}
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
	var list *LinkedList
	list = genLinkedList([]collection.Element{"1", "2", "3"})
	_ = list.Clear()
	assert.Equal(t, 0, list.size)
	assert.Equal(t, (*node)(nil), list.first)
	assert.Equal(t, (*node)(nil), list.last)

	list = &LinkedList{}
	_ = list.Clear()
	assert.Equal(t, 0, list.size)
	assert.Equal(t, (*node)(nil), list.first)
	assert.Equal(t, (*node)(nil), list.last)
}

func TestLinkedList_Contains(t *testing.T) {
	list := genLinkedList([]collection.Element{"1", 2, 3}...)
	var (
		b   bool
		err error
	)

	b, err = list.Contains("1")
	assert.Equal(t, true, b)
	assert.Equal(t, nil, err)

	b, err = list.Contains(1)
	assert.Equal(t, false, b)
	assert.Equal(t, nil, err)
}

func TestLinkedList_ContainsAll(t *testing.T) {
	list := genLinkedList([]collection.Element{"1", 2, 3}...)

	var (
		b   bool
		err error
	)

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

	list = genLinkedList([]collection.Element{"1", 2, 3}...)

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

	list = genLinkedList([]collection.Element{"1", 2, 3}...)
	equals = list.Equals(list1)
	assert.Equal(t, false, equals)

	list1 = genLinkedList([]collection.Element{"1", 2, 3}...)
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

	list = genLinkedList([]collection.Element{"1", 2, 3}...)
	element, err = list.Get(0)
	assert.Equal(t, nil, err)
	assert.Equal(t, "1", element)

	element, err = list.Get(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, element)

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

	list = genLinkedList([]collection.Element{"1", 2, 3}...)
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

	list = genLinkedList([]collection.Element{"1", 2, 3}...)
	element, err = list.GetLast()
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, element)
}

func TestLinkedList_Index(t *testing.T) {
	list := &LinkedList{}
	list = genLinkedList([]collection.Element{"1", 2, 3}...)

	var index int
	index = list.Index("1")
	assert.Equal(t, 0, index)

	index = list.Index(3)
	assert.Equal(t, 2, index)

	index = list.Index(2)
	assert.Equal(t, 1, index)
	index = list.Index("111")
	assert.Equal(t, -1, index)

}

func TestLinkedList_IsEmpty(t *testing.T) {
	list := &LinkedList{}
	empty := list.IsEmpty()
	assert.Equal(t, true, empty)

	list = genLinkedList([]collection.Element{"1", 2, 3}...)
	empty = list.IsEmpty()
	assert.Equal(t, false, empty)

}

func TestLinkedList_Iterator(t *testing.T) {
	list := &LinkedList{}
	iterator := list.Iterator()
	assert.NotEqual(t, nil, iterator)
}

func TestLinkedList_LastIndex(t *testing.T) {
	list := &LinkedList{}

	var index int
	index = list.LastIndex("1")
	assert.Equal(t, -1, index)

	list = genLinkedList([]collection.Element{"1", 2, 3}...)
	index = list.LastIndex(3)
	assert.Equal(t, 2, index)
}

func TestLinkedList_Offer(t *testing.T) {
	list := &LinkedList{}
	var (
		isAdd bool
		err   error
	)

	isAdd, err = list.Offer("1")
	assert.Equal(t, true, isAdd)
	assert.Equal(t, nil, err)
	assert.Equal(t, []collection.Element{"1"}, linkedToSlice(list))

	isAdd, err = list.Offer("2")
	assert.Equal(t, true, isAdd)
	assert.Equal(t, nil, err)
	assert.Equal(t, []collection.Element{"1", "2"}, linkedToSlice(list))

	isAdd, err = list.Offer(3)
	assert.Equal(t, true, isAdd)
	assert.Equal(t, nil, err)
	assert.Equal(t, []collection.Element{"1", "2", 3}, linkedToSlice(list))

}

func TestLinkedList_Peek(t *testing.T) {
	list := genLinkedList([]collection.Element{"1", 2, 3}...)
	var peek collection.Element

	peek = list.Peek()
	assert.Equal(t, "1", peek)

	list = &LinkedList{}
	peek = list.Peek()
	assert.Equal(t, nil, peek)

}

func TestLinkedList_Poll(t *testing.T) {

	list := genLinkedList([]collection.Element{"1", 2, 3}...)
	var poll collection.Element

	poll = list.Poll()
	assert.Equal(t, "1", poll)

	poll = list.Poll()
	assert.Equal(t, 2, poll)

	poll = list.Poll()
	assert.Equal(t, 3, poll)

	poll = list.Poll()
	assert.Equal(t, nil, poll)

	list = &LinkedList{}
	poll = list.Poll()
	assert.Equal(t, nil, poll)

}

func TestLinkedList_Pop(t *testing.T) {
	list := genLinkedList([]collection.Element{"1", 2, 3}...)
	var (
		pop collection.Element
		err error
	)

	pop, err = list.Pop()
	assert.Equal(t, 3, pop)
	assert.Equal(t, nil, err)

	pop, err = list.Pop()
	assert.Equal(t, 2, pop)
	assert.Equal(t, nil, err)

	pop, err = list.Pop()
	assert.Equal(t, "1", pop)
	assert.Equal(t, nil, err)

	pop, err = list.Pop()
	assert.Equal(t, nil, pop)
	assert.Equal(t, errs.NoSuchElement, err)

}

func TestLinkedList_Push(t *testing.T) {
	list := &LinkedList{}

	var err error
	err = list.Push("1")
	assert.Equal(t, nil, err)
	assert.Equal(t, []collection.Element{"1"}, linkedToSlice(list))

	err = list.Push("2")
	assert.Equal(t, nil, err)
	assert.Equal(t, []collection.Element{"2", "1"}, linkedToSlice(list))

	err = list.Push(3)
	assert.Equal(t, nil, err)
	assert.Equal(t, []collection.Element{3, "2", "1"}, linkedToSlice(list))

}

func TestLinkedList_Remove(t *testing.T) {
	list := genLinkedList([]collection.Element{"1", 2, 3}...)
	var (
		isRemove bool
	)

	isRemove, _ = list.Remove("1")
	assert.Equal(t, true, isRemove)
	assert.Equal(t, linkedToSlice(list), []collection.Element{2, 3})

	isRemove, _ = list.Remove(3)
	assert.Equal(t, true, isRemove)
	assert.Equal(t, linkedToSlice(list), []collection.Element{2})

	isRemove, _ = list.Remove(-3)
	assert.Equal(t, false, isRemove)
	assert.Equal(t, linkedToSlice(list), []collection.Element{2})

	isRemove, _ = list.Remove(2)
	assert.Equal(t, true, isRemove)
	assert.Equal(t, linkedToSlice(list), []collection.Element{})

	isRemove, _ = list.Remove(-2)
	assert.Equal(t, false, isRemove)
	assert.Equal(t, linkedToSlice(list), []collection.Element{})

}

func TestLinkedList_RemoveAll(t *testing.T) {
	list := genLinkedList([]collection.Element{"1", 2, 3, 4, 5, 6}...)
	list1 := genLinkedList([]collection.Element{"1", 2}...)
	var (
		isRemoveAll bool
		err         error
	)

	isRemoveAll, err = list.RemoveAll(list1)
	assert.Equal(t, true, isRemoveAll)
	assert.Equal(t, nil, err)
	assert.Equal(t, []collection.Element{3, 4, 5, 6}, linkedToSlice(list))

	isRemoveAll, err = list.RemoveAll(list1)
	assert.Equal(t, false, isRemoveAll)
	assert.Equal(t, nil, err)
	assert.Equal(t, []collection.Element{3, 4, 5, 6}, linkedToSlice(list))

	list2 := genLinkedList([]collection.Element{3, 4, 5, 6}...)
	isRemoveAll, err = list.RemoveAll(list2)
	assert.Equal(t, true, isRemoveAll)
	assert.Equal(t, nil, err)
	assert.Equal(t, []collection.Element{}, linkedToSlice(list))

	list3 := genLinkedList([]collection.Element{}...)
	isRemoveAll, err = list.RemoveAll(list3)
	assert.Equal(t, false, isRemoveAll)
	assert.Equal(t, nil, err)
	assert.Equal(t, []collection.Element{}, linkedToSlice(list))

}

func TestLinkedList_RemoveFirst(t *testing.T) {
	list := genLinkedList([]collection.Element{"1", 2, 3}...)
	var (
		first collection.Element
		err   error
	)
	first, err = list.RemoveFirst()
	assert.Equal(t, "1", first)
	assert.Equal(t, nil, err)
	assert.Equal(t, []collection.Element{2, 3}, linkedToSlice(list))

	first, err = list.RemoveFirst()
	assert.Equal(t, 2, first)
	assert.Equal(t, nil, err)
	assert.Equal(t, []collection.Element{3}, linkedToSlice(list))

	first, err = list.RemoveFirst()
	assert.Equal(t, 3, first)
	assert.Equal(t, nil, err)
	assert.Equal(t, []collection.Element{}, linkedToSlice(list))

	first, err = list.RemoveFirst()
	assert.Equal(t, nil, first)
	assert.Equal(t, errs.NoSuchElement, err)
	assert.Equal(t, []collection.Element{}, linkedToSlice(list))

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

// Benchmark
func BenchmarkNewLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewLinkedList()
	}
}
func BenchmarkLinkedList_Add(b *testing.B) {
	list := NewLinkedList()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = list.Add(i)
	}
}

func BenchmarkLinkedList_AddAll(b *testing.B) {
	list := NewLinkedList()
	listDefault := NewSliceListDefault()
	for i := 0; i < 1; i++ {
		listDefault.Add(i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.AddAll(listDefault)
	}
}
