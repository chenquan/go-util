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
	"github.com/chenquan/go-util/backend/api/collection"
	"github.com/chenquan/go-util/backend/errs"
)

var _ collection.List = (*SliceList)(nil)

const (
	// 默认列表容量
	defaultCapacity = 10
)

// NewSliceListDefault 新增切片列表
//
// 默认容量: 10
func NewSliceListDefault() *SliceList {
	return &SliceList{data: make([]collection.Element, 0, defaultCapacity)}
}

// NewSliceListWithCollection 由指定集合创建切片列表
func NewSliceListWithCollection(c collection.Collection) *SliceList {
	size := c.Size()
	if size != 0 {
		slice := c.Slice()
		return &SliceList{data: slice, size: c.Size()}
	}
	return NewSliceListDefault()
}

// NewSliceList 创建指定容量大小的切片
//
// 如果 initialCapacity<0 ,则容量大小使用默认值:10.
func NewSliceList(initialCapacity int) *SliceList {
	if initialCapacity < 0 {
		initialCapacity = defaultCapacity
	}
	return &SliceList{data: make([]collection.Element, 0, initialCapacity)}
}

// SliceList 实现 collection.List 接口
// 注意 SliceList 协程不安全,不能用于高并发.
type SliceList struct {
	size int                  // 列表大小
	data []collection.Element // 数据
}

// Slice 返回当前切片列表所有元素的切片
//
// 返回的切片是安全的,可任意修改不会影响源切片列表
func (sliceList *SliceList) Slice() []collection.Element {
	elements := make([]collection.Element, sliceList.size)
	// 深拷贝
	copy(elements, sliceList.data)
	return elements
}

// Size 返回当前切片列表的大小
func (sliceList *SliceList) Size() int {
	return sliceList.size
}

// IsEmpty 如果不存在元素则返回 true,否则返回 false
func (sliceList *SliceList) IsEmpty() bool {
	return sliceList.size == 0
}

// Contains 如果当前集合包含元素 e 则返回 true,否则返回 false
//
// 当前返回的error接口总为 nil
func (sliceList *SliceList) Contains(e collection.Element) (bool, error) {
	return sliceList.Index(e) >= 0, nil
}

// Add 添加指定元素
//
// 如果当前集合由于调用而更改, 则返回true.
// 如果此集合不允许重复并且已经包含指定的元素,则返回false.
// 当前返回的error接口总为 nil
func (sliceList *SliceList) Add(e collection.Element) (bool, error) {
	sliceList.data = append(sliceList.data, e)
	sliceList.size++
	return true, nil
}

// Remove 删除指定元素
//
// 如果当前集合中存在指定元素,则删除改元素并返回 true,否则返回 false.
// 当前返回的error接口总为 nil
func (sliceList *SliceList) Remove(e collection.Element) (bool, error) {
	for i := 0; i < sliceList.size; i++ {
		if sliceList.data[i] == e {
			sliceList.size--
			sliceList.fastRemove(i)
			return true, nil
		}
	}
	return false, nil
}

// fastRemove 快速删除
func (sliceList *SliceList) fastRemove(index int) {
	sliceList.data = append(sliceList.data[:index], sliceList.data[index+1:]...)
}

// ContainsAll 如果当前集合包含指定集合中的所有元素，则返回 true,否则返回 false.
// 当前返回的error接口总为 nil
func (sliceList *SliceList) ContainsAll(c collection.Collection) (bool, error) {
	if sliceList == c {
		return true, nil
	}
	for _, e := range c.Slice() {

		if contains, _ := sliceList.Contains(e); !contains {
			return false, nil
		}
	}
	return true, nil
}

// AddAll 将指定集合中的所有元素添加到当前集合中
//
// 如果调用 AddAll 改变了集合,则返回 true,否则返回 false.
// 当前返回的error接口总为 nil
func (sliceList *SliceList) AddAll(c collection.Collection) (b bool, err error) {
	b = c.Size() != 0
	if b {
		slice := c.Slice()
		sliceList.data = append(sliceList.data, slice...)
		sliceList.size += c.Size()
	}
	return
}

// RemoveAll 删除当前集合中与指定集合相同的所有元素
//
// 如果调用 RemoveAll 改变了集合,则返回 true,否则返回 false.
func (sliceList *SliceList) RemoveAll(c collection.Collection) (modified bool, err error) {
	return sliceList.batchRemove(c, false), nil
}

// batchRemove 批量删除指定集合元素
//
// 如果complement等于false,则删除当前集合中与指定集合相同的所有元素.
// 如果complement等于true,仅保留当前集合中包含在指定集合中的元素.
func (sliceList *SliceList) batchRemove(c collection.Collection, complement bool) (modified bool) {
	data := sliceList.data
	size := sliceList.size
	r, w := 0, 0
	for ; r < size; r++ {
		contains, _ := c.Contains(data[r])
		if contains == complement {
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

// RetainAll 仅保留当前集合中包含在指定集合中的元素
//
// 换句话说,从当前集合中删除所有未包含在指定集合中的元素.
func (sliceList *SliceList) RetainAll(c collection.Collection) (modified bool, err error) {
	return sliceList.batchRemove(c, true), nil
}

// Clear 清空集合中所有元素
//
// 此方法返回后集合将为空.
// 当前返回的error接口总为 nil
func (sliceList *SliceList) Clear() error {
	sliceList.data = make([]collection.Element, 0, defaultCapacity)
	sliceList.size = 0
	return nil
}

// Equals 比较指定对象与此集合的相等性
func (sliceList *SliceList) Equals(collection collection.Collection) (b bool) {

	if sliceList == collection {
		return true
	}
	iterator1 := sliceList.Iterator()
	iterator2 := collection.Iterator()
	for iterator1.HasNext() && iterator2.HasNext() {
		e1, _ := iterator1.Next()

		e2, _ := iterator2.Next()

		if e1 != e2 {
			return false
		}
	}
	return !(iterator1.HasNext() || iterator2.HasNext())

}

// rangeCheckForAdd 检查新增操作索引范围
func (sliceList *SliceList) rangeCheckForAdd(index int) error {
	if index > sliceList.size || index < 0 {
		return errs.IndexOutOfBound
	}
	return nil
}

// AddAllIndex 将指定集合中的所有元素插入此列表中的指定位置
//
// 将当前在该位置的元素（如果有）和任何后续元素右移(增加其索引)
// 新元素将按照指定集合的迭代器返回的顺序显示在此列表中,
// 如果在操作进行过程中修改了指定的集合,则此操作的行为是不确定的.
func (sliceList *SliceList) AddAllIndex(index int, c collection.Collection) error {
	if err := sliceList.rangeCheckForAdd(index); err != nil {
		return err
	}
	slice := c.Slice()
	if c.Size() != 0 {
		sliceList.data = append(sliceList.data, slice...)
		copy(sliceList.data[index+c.Size():], sliceList.data[index:])
		copy(sliceList.data[index:], slice)
		sliceList.size += c.Size()
	}
	return nil
}

// Get 返回此列表中指定位置的元素
func (sliceList *SliceList) Get(index int) (e collection.Element, err error) {
	size := sliceList.size
	if index >= size || index < 0 {
		err = errs.IndexOutOfBound
	} else {
		e = sliceList.data[index]
	}
	return
}

// Set 用指定的元素替换此列表中指定位置的元素
func (sliceList *SliceList) Set(index int, e collection.Element) (err error) {
	size := sliceList.size
	if index >= size {
		err = errs.IndexOutOfBound
	} else {
		sliceList.data[index] = e
	}
	return
}

// AddIndex将指定的元素插入此列表中的指定位置(可选操作)
//
// 将当前在该位置的元素(如果有)和任何后续元素右移(即将其索引加一).
func (sliceList *SliceList) AddIndex(index int, e collection.Element) error {
	if err := sliceList.rangeCheckForAdd(index); err != nil {
		return err
	}
	sliceList.data = append(sliceList.data, e)
	copy(sliceList.data[index+1:], sliceList.data[index:])
	sliceList.data[index] = e
	sliceList.size++
	return nil
}

// RemoveIndex 删除此列表中指定位置的元素(可选操作)
//
// 将所有后续元素向左移动(即将其索引中减去1),并返回从列表中删除的元素.
func (sliceList *SliceList) RemoveIndex(index int) (collection.Element, error) {
	if index >= sliceList.size {
		return nil, errs.IndexOutOfBound
	}
	element := sliceList.data[index]
	sliceList.data = append(sliceList.data[:index], sliceList.data[index+1:]...)
	sliceList.size--
	return element, nil
}

// Index 返回指定元素在此列表中首次出现的索引,如果此列表不包含该元素，则返回-1
func (sliceList *SliceList) Index(e collection.Element) (index int) {
	size := sliceList.size
	for i := 0; i < size; i++ {
		if sliceList.data[i] == e {
			return i
		}
	}
	return -1
}

// LastIndex 返回此列表中指定元素的最后一次出现的索引,如果此列表不包含该元素，则返回-1
func (sliceList *SliceList) LastIndex(e collection.Element) (index int) {
	size := sliceList.size
	for i := size - 1; i >= 0; i-- {
		if sliceList.data[i] == e {
			return i
		}
	}
	return -1
}

// subListRangeCheck 检查 SubList 函数的参数范围
func (sliceList *SliceList) subListRangeCheck(fromIndex, toIndex int) error {

	if fromIndex < 0 {
		return errs.NewIndexOutOfBoundsError(fmt.Sprintf("fromIndex = %d", fromIndex))
	}
	if toIndex > sliceList.size {
		return errs.NewIndexOutOfBoundsError(fmt.Sprintf("toIndex = %d", toIndex))

	}
	if fromIndex > toIndex {
		return errs.NewIndexOutOfBoundsError(fmt.Sprintf("fromIndex(%d) > toIndex(%d)", fromIndex, toIndex))
	}
	return nil
}

// SubList 返回此列表中指定的fromIndex(包括)和toIndex(不包括)之间的元素
// 如果fromIndex和toIndex相等, 则返回的列表为空.
func (sliceList *SliceList) SubList(fromIndex, toIndex int) (collection.List, error) {
	if err := sliceList.subListRangeCheck(fromIndex, toIndex); err != nil {
		return nil, err
	}
	return &SliceList{
			size: toIndex - fromIndex,
			data: sliceList.data[fromIndex:toIndex],
		},
		nil
}

// Iterator 返回当前集合中元素的迭代器
//
// 如果当前集合是有序的,则保证返回的迭代器是有序的.
func (sliceList *SliceList) Iterator() collection.Iterator {
	return &itrList{
		lastRet: -1,
		cursor:  0,
		data:    sliceList,
	}
}
