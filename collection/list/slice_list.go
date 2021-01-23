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
	"fmt"
	"github.com/chenquan/go-utils/collection/api/collection"
)

var _ collection.List = (*SliceList)(nil)

const (
	defaultCapacity = 10
)

var (
	//NotFound        = errors.New("not found")
	IndexOutOfBound = errors.New("index out of bound")
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

func (sliceList *SliceList) Add(e collection.Element) bool {
	sliceList.size++
	sliceList.data = append(sliceList.data, e)
	return true
}

func (sliceList *SliceList) Remove(e collection.Element) (b bool) {
	for i := 0; i < sliceList.size; i++ {
		if sliceList.data[i] == e {
			sliceList.size--
			sliceList.fastRemove(i)
			return true
		}
	}
	return false
}
func (sliceList *SliceList) fastRemove(index int) {
	sliceList.data = append(sliceList.data[:index], sliceList.data[index+1:]...)
}

func (sliceList *SliceList) ContainsAll(c collection.Collection) bool {
	if sliceList == c {
		return true
	}
	for _, e := range c.Slice() {
		if !sliceList.Contains(e) {
			return false
		}
	}
	return true
}

func (sliceList *SliceList) AddAll(c collection.Collection) (b bool) {
	b = c.Size() != 0
	if b {
		slice := c.Slice()
		elements := make([]collection.Element, c.Size())
		// 深拷贝
		copy(elements, slice)
		sliceList.data = append(sliceList.data, elements...)
		sliceList.size += c.Size()
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
	if r != 0 {
		// r 不为0 说明元素已更改
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
		e1, err1 := iterator1.Next()
		if err1 != nil {
			return false
		}
		e2, err2 := iterator2.Next()
		if err2 != nil {
			return false
		}
		if e1 != e2 {
			return false
		}
	}
	return !(iterator1.HasNext() || iterator2.HasNext())

}

func (sliceList *SliceList) AddAllIndex(index int, c collection.Collection) error {
	if index > sliceList.size-1 {
		return IndexOutOfBound
	}
	slice := c.Slice()
	if c.Size() != 0 {
		front := sliceList.data[:index]
		end := sliceList.data[index:]
		sliceList.data = append(front, slice...)
		sliceList.data = append(sliceList.data, end...)
	}
	return nil
}

func (sliceList *SliceList) Get(index int) (e collection.Element, err error) {
	size := sliceList.size
	if index > size {
		e = IndexOutOfBound
	} else {
		e = sliceList.data[index]
	}
	return
}

func (sliceList *SliceList) Set(index int, e collection.Element) {
	size := sliceList.size
	if index > size {
		e = IndexOutOfBound
	} else {
		sliceList.data[index] = e
	}
	return
}

func (sliceList *SliceList) AddIndex(index int, e collection.Element) error {
	if index > sliceList.size-1 {
		return IndexOutOfBound
	}
	front := sliceList.data[:index]
	end := sliceList.data[index:]

	sliceList.data = append(front, e)
	sliceList.data = append(sliceList.data, end...)
	return nil
}

func (sliceList *SliceList) RemoveIndex(index int) (collection.Element, error) {
	if index >= sliceList.size {
		return nil, IndexOutOfBound
	}
	element := sliceList.data[index]
	sliceList.data = append(sliceList.data[:index], sliceList.data[index+1:])
	return element, nil
}

func (sliceList *SliceList) Index(e collection.Element) (index int) {
	size := sliceList.size
	for i := 0; i < size; i++ {
		if sliceList.data[i] == e {
			return i
		}
	}
	return -1
}

func (sliceList *SliceList) LastIndex(e collection.Element) (index int) {
	size := sliceList.size
	for i := size - 1; i >= 0; i-- {
		if sliceList.data[i] == e {
			return i
		}
	}
	return -1
}

func (sliceList *SliceList) SubList(fromIndex, toIndex int) (list collection.List, err error) {
	if fromIndex < 0 {
		return nil, fmt.Errorf("fromIndex = %d", fromIndex)
	}
	if toIndex > sliceList.size {
		return nil, fmt.Errorf("toIndex = %d", toIndex)

	}
	if fromIndex > toIndex {
		return nil, fmt.Errorf("fromIndex(%d) > toIndex(%d)", fromIndex, toIndex)
	}
	return &SliceList{
			size: fromIndex - toIndex,
			data: sliceList.data[fromIndex:toIndex],
		},
		nil
}

func (sliceList *SliceList) Iterator() collection.Iterator {
	return &SliceListIterator{
		lastRet: -1,
		cursor:  0,
		data:    sliceList,
	}
}

type SliceListIterator struct {
	cursor  int //游标,指向下一个元素
	lastRet int
	data    *SliceList
}

func (s *SliceListIterator) HasNext() bool {
	return s.data.size != s.cursor
}

func (s *SliceListIterator) Next() (collection.Element, error) {
	if s.cursor >= s.data.size {
		return nil, errors.New("no such element")
	}
	s.lastRet = s.cursor
	s.cursor++
	return s.data.Get(s.lastRet)
}

func (s *SliceListIterator) Remove() error {
	if s.lastRet < 0 {
		return errors.New("illegal state")
	}
	_, err := s.data.RemoveIndex(s.lastRet)
	return err
}
