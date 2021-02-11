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

package collection

// List 有序集合,也称为序列
//
// 该界面的用户可以精确控制列表中每个元素的插入位置.
// 用户可以通过其整数索引(列表中的位置)访问元素，并在列表中搜索元素.
type List interface {
	// Collection List 默认实现集合功能
	Collection
	// AddAllIndex 将指定集合中的所有元素插入此列表中的指定位置
	//
	// 将当前在该位置的元素（如果有）和任何后续元素右移(增加其索引)
	// 新元素将按照指定集合的迭代器返回的顺序显示在此列表中,
	// 如果在操作进行过程中修改了指定的集合,则此操作的行为是不确定的.
	AddAllIndex(index int, c Collection) (bool, error)
	// Get 返回此列表中指定位置的元素
	Get(index int) (Element, error)
	// Set 用指定的元素替换此列表中指定位置的元素
	Set(index int, e Element) (Element, error)
	// AddIndex将指定的元素插入此列表中的指定位置(可选操作)
	//
	// 将当前在该位置的元素(如果有)和任何后续元素右移(即将其索引加一).
	AddIndex(index int, e Element) error
	// RemoveIndex 删除此列表中指定位置的元素(可选操作)
	//
	// 将所有后续元素向左移动(即将其索引中减去1),并返回从列表中删除的元素.
	RemoveIndex(index int) (Element, error)
	// Index 返回指定元素在此列表中首次出现的索引,如果此列表不包含该元素，则返回-1
	Index(e Element) int
	// LastIndex 返回此列表中指定元素的最后一次出现的索引,如果此列表不包含该元素，则返回-1
	LastIndex(e Element) int
	// SubList 返回此列表中指定的fromIndex(包括)和toIndex(不包括)之间的元素
	// 如果fromIndex和toIndex相等, 则返回的列表为空.
	SubList(fromIndex, toIndex int) (List, error)
}

type IteratorList interface {
	Iterator
	HasPrevious() bool
	Previous() bool
	NextIndex() int
	PreviousIndex() int
	Set(e Element)
	Add(e Element)
}
