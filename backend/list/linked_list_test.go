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
	"github.com/chenquan/go-util/backend/api/collection"
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
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		index int
		c     collection.Collection
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
			if err := l.AddAllIndex(tt.args.index, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AddAllIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_AddFirst(t *testing.T) {
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
			if err := l.AddFirst(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("AddFirst() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_AddIndex(t *testing.T) {
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
			if err := l.AddIndex(tt.args.index, tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("AddIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_AddLast(t *testing.T) {
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
			if err := l.AddLast(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("AddLast() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_Clear(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	tests := []struct {
		name    string
		fields  fields
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
			if err := l.Clear(); (err != nil) != tt.wantErr {
				t.Errorf("Clear() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_Contains(t *testing.T) {
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
			got, err := l.Contains(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Contains() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_ContainsAll(t *testing.T) {
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
			got, err := l.ContainsAll(tt.args.collection)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContainsAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ContainsAll() got = %v, want %v", got, tt.want)
			}
		})
	}
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
			got, err := l.Element()
			if (err != nil) != tt.wantErr {
				t.Errorf("Element() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Element() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Equals(t *testing.T) {
	type fields struct {
		size  int
		first *node
		last  *node
	}
	type args struct {
		collection collection.Collection
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
			if got := l.Equals(tt.args.collection); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Get(t *testing.T) {
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
			got, err := l.Get(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_GetFirst(t *testing.T) {
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
			got, err := l.GetFirst()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFirst() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFirst() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_GetLast(t *testing.T) {
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
			got, err := l.GetLast()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLast() got = %v, want %v", got, tt.want)
			}
		})
	}
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
			if err := l.Offer(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("Offer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
			got, err := l.Poll()
			if (err != nil) != tt.wantErr {
				t.Errorf("Poll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Poll() got = %v, want %v", got, tt.want)
			}
		})
	}
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
