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
	"github.com/chenquan/go-utils/collection/api/collection"
	"github.com/chenquan/go-utils/collection/api/iterator"
	"github.com/chenquan/go-utils/collection/api/list"
)

var _ list.List = (*SliceList)(nil)

const (
	defaultCapacity = 10
)

func NewSliceListDefault() *SliceList {
	return &SliceList{data: make([]collection.Element, 0)}
}
func NewSliceListWithCollection(c collection.Collection) *SliceList {
	size := c.Size()
	if size != 0 {
		slice := c.Slice()
		elements := make([]collection.Element, size)
		// 深拷贝
		copy(elements, slice)
		return &SliceList{data: elements, size: c.Size()}
	}
	return NewSliceListDefault()
}

func NewSliceList(initialCapacity int) *SliceList {
	if initialCapacity < 0 {
		initialCapacity = defaultCapacity
	}
	return &SliceList{data: make([]collection.Element, 0, initialCapacity)}
}

type SliceList struct {
	size int
	data []collection.Element
}

func (sliceList *SliceList) Iterator() iterator.Iterator {
	panic("implement me")
}

func (sliceList *SliceList) Slice() []collection.Element {
	return sliceList.data
}

func (sliceList *SliceList) Size() int {
	return sliceList.size
}

func (sliceList *SliceList) IsEmpty() bool {
	return sliceList.size == 0
}

func (sliceList *SliceList) Contains(e collection.Element) bool {
	return sliceList.Index(e) >= 0
}

func (sliceList *SliceList) Add(e collection.Element) {
	sliceList.size++
	sliceList.data = append(sliceList.data, e)
}

func (sliceList *SliceList) Remove(e collection.Element) (b bool) {
	for i := 0; i < sliceList.size; i++ {
		if sliceList.data[i] == e {
			sliceList.fastRemove(i)
			return true
		}
	}
	return false
}
func (sliceList *SliceList) fastRemove(index int) {
	if sliceList.size == index+1 {
		sliceList.data = sliceList.data[:index]
	} else {
		sliceList.data = append(sliceList.data[:index], sliceList.data[index+1:]...)
	}

}

func (sliceList *SliceList) ContainsAll(c collection.Collection) {
	slice := c.Slice()
	elements := make([]collection.Element, c.Size())
	// 深拷贝
	copy(elements, slice)
	sliceList.data = append(sliceList.data, elements...)
}

func (sliceList *SliceList) AddAll(c collection.Collection) (b bool) {
	b = c.Size() != 0
	if b {
		slice := c.Slice()
		elements := make([]collection.Element, c.Size())
		// 深拷贝
		copy(elements, slice)
		sliceList.data = append(sliceList.data, elements...)
	}
	return
}

func (sliceList *SliceList) RemoveAll(c collection.Collection) (modified bool) {
	return sliceList.batchRemove(c, false)
}
func (sliceList *SliceList) batchRemove(c collection.Collection, complement bool) (modified bool) {
	data := sliceList.data
	size := sliceList.size
	r, w := 0, 0
	for ; r < size; r++ {
		if c.Contains(data[r]) == complement {
			data[w] = data[r]
			w++
		}
	}
	if w != 0 {
		// w 不为0 说明元素已更改
		// 剔除多余元素
		sliceList.data = sliceList.data[:w]
		modified = true
		sliceList.size = w
	}
	return modified
}
func (sliceList *SliceList) RetainAll(c collection.Collection) (modified bool) {
	return sliceList.batchRemove(c, true)
}
func (sliceList *SliceList) Clear() {
	sliceList.data = make([]collection.Element, 0, defaultCapacity)
	sliceList.size = 0
}

func (sliceList *SliceList) Equals(collection collection.Collection) (b bool) {

	if sliceList == collection {
		return true
	}
	iterator1 := sliceList.Iterator()
	iterator2 := collection.Iterator()
	for iterator1.HasNext() && iterator2.HasNext() {
		e1 := iterator1.Next()
		e2 := iterator2.Next()
		if e1 != e2 {
			return false
		}
	}
	return !(iterator1.HasNext() || iterator2.HasNext())

}

func (sliceList *SliceList) AddAllIndex(index int, c collection.Collection) {
	panic("implement me")
}

func (sliceList *SliceList) Get(index int) (e collection.Element) {
	panic("implement me")
}

func (sliceList *SliceList) Set(index int, e collection.Element) {
	panic("implement me")
}

func (sliceList *SliceList) AddIndex(index int, e collection.Element) {
	panic("implement me")
}

func (sliceList *SliceList) RemoveIndex(index int) {
	panic("implement me")
}

func (sliceList *SliceList) Index(e collection.Element) (index int) {
	panic("implement me")
}

func (sliceList *SliceList) LastIndex(e collection.Element) (index int) {
	panic("implement me")
}

func (sliceList *SliceList) SubList(fromIndex, toIndex int) (list list.List) {
	panic("implement me")
}
