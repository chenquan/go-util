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
	"errors"
	"github.com/chenquan/go-utils/collection/api"
)

var _ api.List = (*SliceList)(nil)

const (
	defaultCapacity = 10
)

var (
	//NotFound        = errors.New("not found")
	IndexOutOfBound = errors.New("index out of bound")
)

func NewSliceListDefault() *SliceList {
	return &SliceList{data: make([]api.Element, 0)}
}
func NewSliceListWithCollection(c api.Collection) *SliceList {
	size := c.Size()
	if size != 0 {
		slice := c.Slice()
		elements := make([]api.Element, size)
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
	return &SliceList{data: make([]api.Element, 0, initialCapacity)}
}

type SliceList struct {
	size int
	data []api.Element
}

func (sliceList *SliceList) Iterator() api.Iterator {
	panic("implement me")
}

func (sliceList *SliceList) Slice() []api.Element {
	return sliceList.data
}

func (sliceList *SliceList) Size() int {
	return sliceList.size
}

func (sliceList *SliceList) IsEmpty() bool {
	return sliceList.size == 0
}

func (sliceList *SliceList) Contains(e api.Element) bool {
	return sliceList.Index(e) >= 0
}

func (sliceList *SliceList) Add(e api.Element) {
	sliceList.size++
	sliceList.data = append(sliceList.data, e)
}

func (sliceList *SliceList) Remove(e api.Element) (b bool) {
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

func (sliceList *SliceList) ContainsAll(c api.Collection) {
	slice := c.Slice()
	elements := make([]api.Element, c.Size())
	// 深拷贝
	copy(elements, slice)
	sliceList.data = append(sliceList.data, elements...)
}

func (sliceList *SliceList) AddAll(c api.Collection) (b bool) {
	b = c.Size() != 0
	if b {
		slice := c.Slice()
		elements := make([]api.Element, c.Size())
		// 深拷贝
		copy(elements, slice)
		sliceList.data = append(sliceList.data, elements...)
	}
	return
}

func (sliceList *SliceList) RemoveAll(c api.Collection) (modified bool) {
	return sliceList.batchRemove(c, false)
}
func (sliceList *SliceList) batchRemove(c api.Collection, complement bool) (modified bool) {
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
func (sliceList *SliceList) RetainAll(c api.Collection) (modified bool) {
	return sliceList.batchRemove(c, true)
}
func (sliceList *SliceList) Clear() {
	sliceList.data = make([]api.Element, 0, defaultCapacity)
	sliceList.size = 0
}

func (sliceList *SliceList) Equals(collection api.Collection) (b bool) {

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

func (sliceList *SliceList) AddAllIndex(index int, c api.Collection) {
	panic("implement me")
}

func (sliceList *SliceList) Get(index int) (e api.Element, err error) {
	size := sliceList.size
	if index > size {
		e = IndexOutOfBound
	} else {
		e = sliceList.data[index]
	}
	return
}

func (sliceList *SliceList) Set(index int, e api.Element) {
	size := sliceList.size
	if index > size {
		e = IndexOutOfBound
	} else {
		sliceList.data[index] = e
	}
	return
}

func (sliceList *SliceList) AddIndex(index int, e api.Element) {
	panic("implement me")
}

func (sliceList *SliceList) RemoveIndex(index int) {
	panic("implement me")
}

func (sliceList *SliceList) Index(e api.Element) (index int) {
	panic("implement me")
}

func (sliceList *SliceList) LastIndex(e api.Element) (index int) {
	panic("implement me")
}

func (sliceList *SliceList) SubList(fromIndex, toIndex int) (list api.List) {
	panic("implement me")
}