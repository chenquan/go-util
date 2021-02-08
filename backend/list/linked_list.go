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
	"github.com/chenquan/go-util/backend/errs"
)

// node 链表节点
type node struct {
	elem collection.Element // 数据元素
	next *node              // 下一个节点
	prev *node              // 上一个节点
}

// LinkedList
type LinkedList struct {
	len   int
	first *node
	last  *node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) AddFirst(e collection.Element) error {
	f := l.first
	n := &node{
		elem: e,
		next: f,
		prev: nil,
	}
	// 链表中第一个元素
	l.first = n
	if f == nil {
		l.last = n
	} else {
		f.prev = n
	}
	l.len++
	return nil
}

func (l *LinkedList) AddLast(e collection.Element) error {
	last := l.last
	n := &node{
		elem: e,
		next: last,
		prev: nil,
	}
	if last == nil {
		l.first = n
	} else {
		last.next = n
	}
	l.len++
	return nil
}

//func (l *LinkedList) OfferFirst(e collection.Element) error {
//	panic("implement me")
//}
//
//func (l *LinkedList) OfferLast(e collection.Element) error {
//	panic("implement me")
//}

func (l *LinkedList) RemoveFirst() (collection.Element, error) {
	first := l.first
	if first == nil {
		return nil, errs.NoSuchElement
	}
	next := first.next
	l.first = next
	if next == nil {
		// 空链表
		l.last = nil
	} else {
		// 将头结点的前驱置为nil
		l.first.prev = nil
	}
	l.len--
	return first.elem, nil
}

func (l *LinkedList) RemoveLast() (collection.Element, error) {
	last := l.last
	if last == nil {
		return nil, errs.NoSuchElement
	}
	prev := last.prev
	l.last = prev
	if prev == nil {
		l.first = nil
	} else {
		l.last.next = nil
	}
	return last.elem, nil
}

func (l *LinkedList) GetFirst() (collection.Element, error) {
	if l.first == nil {
		return nil, errs.NoSuchElement
	}
	return l.first.elem, nil
}

func (l *LinkedList) GetLast() (collection.Element, error) {
	if l.last == nil {
		return nil, errs.NoSuchElement
	}
	return l.last.elem, nil
}

func (l *LinkedList) RemoveFirstOccurrence(e collection.Element) (bool, error) {
	return l.Remove(e)
}

func (l *LinkedList) RemoveLastOccurrence(e collection.Element) (bool, error) {
	for x := l.last; x != nil; x = x.prev {
		if x.elem == e {
			l.unLink(x)
			return true, nil
		}
	}
	return false, nil
}
func (l *LinkedList) unLink(x *node) collection.Element {
	elem := x.elem
	// 前驱结点
	prev := x.prev
	// 后驱结点
	next := x.next
	if prev == nil {
		// x是头结点
		l.first = next
	} else {
		prev.next = next
		// 断开
		x.prev = nil
	}
	if next == nil {
		l.last = nil
	} else {
		next.prev = prev
		// 断开
		x.next = nil
	}
	x.elem = nil
	l.len--
	return elem

}
func (l *LinkedList) Push(e collection.Element) error {
	return l.AddFirst(e)
}

func (l *LinkedList) Pop() (collection.Element, error) {
	return l.RemoveLast()
}

func (l *LinkedList) DescendingIterator() collection.Iterator {
	panic("implement me")
}

func (l *LinkedList) AddAllIndex(index int, c collection.Collection) error {
	if !l.isPositionIndex(index) {
		return errs.IndexOutOfBound
	}
	slice := c.Slice()
	numNew := len(slice)
	if numNew == 0 {
		return nil
	}
	return nil

}
func (l *LinkedList) isPositionIndex(index int) bool {
	return index >= 0 && index <= l.len
}

func (l *LinkedList) Get(index int) (collection.Element, error) {
	panic("implement me")
}

func (l *LinkedList) Set(index int, e collection.Element) error {
	panic("implement me")
}

func (l *LinkedList) AddIndex(index int, e collection.Element) error {
	panic("implement me")
}

func (l *LinkedList) RemoveIndex(index int) (collection.Element, error) {
	panic("implement me")
}

func (l *LinkedList) Index(e collection.Element) int {
	panic("implement me")
}

func (l *LinkedList) LastIndex(e collection.Element) int {
	panic("implement me")
}

func (l *LinkedList) SubList(fromIndex, toIndex int) (collection.List, error) {
	panic("implement me")
}

func (l *LinkedList) Size() int {
	return l.len
}

func (l *LinkedList) IsEmpty() bool {
	return l.len == 0
}

func (l *LinkedList) Contains(e collection.Element) (bool, error) {
	panic("implement me")
}

func (l *LinkedList) Add(e collection.Element) (bool, error) {
	panic("implement me")
}

func (l *LinkedList) Remove(e collection.Element) (bool, error) {
	panic("implement me")
}

func (l *LinkedList) ContainsAll(collection collection.Collection) (bool, error) {
	panic("implement me")
}

func (l *LinkedList) AddAll(collection collection.Collection) (bool, error) {
	panic("implement me")
}

func (l *LinkedList) RemoveAll(collection collection.Collection) (bool, error) {
	panic("implement me")
}

func (l *LinkedList) RetainAll(collection collection.Collection) (bool, error) {
	panic("implement me")
}

func (l *LinkedList) Clear() error {
	panic("implement me")
}

func (l *LinkedList) Equals(collection collection.Collection) bool {
	panic("implement me")
}

func (l *LinkedList) Slice() []collection.Element {
	panic("implement me")
}

func (l *LinkedList) Iterator() collection.Iterator {
	panic("implement me")
}

func (l *LinkedList) Offer(e collection.Element) error {
	panic("implement me")
}

func (l *LinkedList) Poll() (collection.Element, error) {
	panic("implement me")
}

func (l *LinkedList) Element() (collection.Element, error) {
	panic("implement me")
}

func (l *LinkedList) Peek() collection.Element {
	panic("implement me")
}
