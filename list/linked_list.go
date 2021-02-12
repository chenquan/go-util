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
)

// node 链表节点
type node struct {
	elem collection.Element // 数据元素
	next *node              // 下一个节点
	prev *node              // 上一个节点
}

// LinkedList
type LinkedList struct {
	size  int
	first *node
	last  *node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// AddFirst
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
	l.size++
	return nil
}

// AddLast
func (l *LinkedList) AddLast(e collection.Element) error {
	l.linkLast(e)
	return nil
}

// RemoveFirst
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
	l.size--
	return first.elem, nil
}

// RemoveLast
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

// GetFirst
func (l *LinkedList) GetFirst() (collection.Element, error) {
	if l.first == nil {
		return nil, errs.NoSuchElement
	}
	return l.first.elem, nil
}

// GetLast
func (l *LinkedList) GetLast() (collection.Element, error) {
	if l.last == nil {
		return nil, errs.NoSuchElement
	}
	return l.last.elem, nil
}

func (l *LinkedList) RemoveFirstOccurrence(e collection.Element) (bool, error) {
	return l.Remove(e)
}

// RemoveLastOccurrence
func (l *LinkedList) RemoveLastOccurrence(e collection.Element) (bool, error) {
	for x := l.last; x != nil; x = x.prev {
		if x.elem == e {
			l.unLink(x)
			return true, nil
		}
	}
	return false, nil
}

// unLink 移除节点
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
		l.last = prev
	} else {
		next.prev = prev
		// 断开
		x.next = nil
	}
	x.elem = nil
	l.size--
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

func (l *LinkedList) AddAllIndex(index int, c collection.Collection) (bool, error) {
	if !l.isPositionIndex(index) {
		return false, errs.IndexOutOfBound
	}
	slice := c.Slice()
	numNew := len(slice)
	if numNew == 0 {
		return false, nil
	}
	var (
		prev *node
		p    *node
	)

	if l.size == index {
		p = nil
		prev = l.last
	} else {
		p = l.getNode(index)
		prev = p.prev
	}

	for _, element := range slice {
		n := &node{
			elem: element,
			next: nil,
			prev: prev,
		}
		if prev == nil {
			l.first = n
		} else {
			// 反向关联
			prev.next = n
		}
		prev = n
	}

	if p == nil {
		l.last = prev
	} else {
		prev.next = p
		p.prev = prev
	}
	l.size += numNew
	return true, nil
}
func (l *LinkedList) isPositionIndex(index int) bool {
	return index >= 0 && index <= l.size
}

// checkElementIndex 检查索引
func (l *LinkedList) checkElementIndex(index int) error {
	if index >= 0 && index < l.size {
		return nil
	}
	return errs.IndexOutOfBound
}
func (l *LinkedList) checkPositionIndex(index int) error {
	if index >= 0 && index <= l.size {
		return nil
	}
	return errs.IndexOutOfBound
}
func (l *LinkedList) Get(index int) (collection.Element, error) {
	err := l.checkElementIndex(index)
	if err != nil {
		return nil, err
	}
	n := l.getNode(index)
	return n.elem, nil

}

// getNode 获取索引对应的节点
func (l *LinkedList) getNode(index int) *node {
	if index < (l.size >> 1) {
		n := l.first
		for i := 0; i < index; i++ {
			n = n.next
		}
		return n
	} else {
		n := l.last
		for i := l.size - 1; i > index; i-- {
			n = n.prev
		}
		return n
	}
}

func (l *LinkedList) Set(index int, e collection.Element) (collection.Element, error) {
	err := l.checkElementIndex(index)
	if err != nil {
		return nil, err
	}
	n := l.getNode(index)
	oldElement := n.elem
	n.elem = e
	return oldElement, nil

}

func (l *LinkedList) AddIndex(index int, e collection.Element) error {
	err := l.checkPositionIndex(index)
	if err != nil {
		return err
	}
	if index == l.size {
		return l.AddLast(e)
	}
	l.linkBefore(e, l.getNode(index))
	return nil
}

func (l *LinkedList) linkBefore(e collection.Element, n *node) {
	prev := n.prev
	newNode := &node{
		elem: e,
		next: n,
		prev: prev,
	}
	n.prev = newNode
	if prev == nil {
		l.first = newNode
	} else {
		prev.next = newNode
	}
	l.size++
}

func (l *LinkedList) linkLast(e collection.Element) {
	last := l.last
	n := &node{
		elem: e,
		next: nil,
		prev: last,
	}
	l.last = n
	if last == nil {
		l.first = n
	} else {
		last.next = n
	}
	l.size++
}

func (l *LinkedList) RemoveIndex(index int) (collection.Element, error) {
	if err := l.checkElementIndex(index); err != nil {
		return nil, err
	}
	return l.unLink(l.getNode(index)), nil
}

func (l *LinkedList) Index(e collection.Element) int {
	index := 0
	for x := l.first; x != nil; x = x.next {
		if x.elem == e {
			return index
		}
		index++
	}
	return -1
}

func (l *LinkedList) LastIndex(e collection.Element) int {
	index := l.size - 1
	for x := l.last; x != nil; x = x.prev {
		if x.elem == e {
			return index
		}
		index--
	}
	return -1
}

func (l *LinkedList) SubList(fromIndex, toIndex int) (collection.List, error) {
	panic("implement me")
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Contains(e collection.Element) (bool, error) {
	return l.Index(e) >= 0, nil
}

func (l *LinkedList) Add(e collection.Element) (bool, error) {
	l.linkLast(e)
	return true, nil
}

func (l *LinkedList) Remove(e collection.Element) (bool, error) {
	for x := l.first; x != nil; x = x.next {
		if x.elem == e {
			l.unLink(x)
			return true, nil
		}
	}
	return false, nil
}

func (l *LinkedList) ContainsAll(c collection.Collection) (bool, error) {
	iterator := c.Iterator()
	var (
		next     collection.Element
		contains bool
	)
	for iterator.HasNext() {
		next, _ = iterator.Next()
		contains, _ = l.Contains(next)
		if !contains {
			return false, nil
		}
	}
	return true, nil
}

func (l *LinkedList) AddAll(collection collection.Collection) (bool, error) {
	modified := false
	elements := collection.Slice()
	for _, element := range elements {
		if add, _ := l.Add(element); add {
			modified = true
		}
	}
	return modified, nil
}

func (l *LinkedList) RemoveAll(c collection.Collection) (bool, error) {
	return l.batchRemove(c, true)
}
func (l *LinkedList) batchRemove(c collection.Collection, complement bool) (bool, error) {
	iterator := l.Iterator()
	modified := false
	for iterator.HasNext() {
		if next, err := iterator.Next(); err != nil {
			return false, err
		} else {
			if contains, err1 := c.Contains(next); err1 != nil {
				return false, err1
			} else {
				if contains == complement {
					err2 := iterator.Remove()
					modified = true
					if err2 != nil {
						return false, err2
					}
				}
			}
		}

	}
	return modified, nil
}
func (l *LinkedList) RetainAll(c collection.Collection) (bool, error) {
	return l.batchRemove(c, false)
}

func (l *LinkedList) Clear() error {
	for n := l.last; n != nil; {
		next := n.next
		n.elem = nil
		n.prev = nil
		n.next = nil
		n = next
	}
	l.last = nil
	l.first = nil
	l.size = 0
	return nil
}

func (l *LinkedList) Equals(c collection.Collection) bool {
	if l == c {
		return true
	}
	if l.size != c.Size() {
		return false
	}
	i1 := l.Iterator()
	i2 := c.Iterator()
	for i1.HasNext() && i2.HasNext() {
		e1, _ := i1.Next()
		e2, _ := i2.Next()
		if e1 != e2 {
			return false
		}
	}
	return true
}

func (l *LinkedList) Slice() []collection.Element {
	elements := make([]collection.Element, 0, l.size)
	for x := l.first; x != nil; x = x.next {
		elements = append(elements, x.elem)
	}
	return elements
}

func (l *LinkedList) Iterator() collection.Iterator {
	return &itrList{cursor: 0, lastRet: -1, data: l}
}

func (l *LinkedList) Offer(e collection.Element) (bool, error) {
	return l.Add(e)
}

func (l *LinkedList) Poll() collection.Element {
	first := l.first
	if first == nil {
		return nil
	}
	elem := first.elem
	l.unLink(first)
	return elem
}

func (l *LinkedList) Element() (collection.Element, error) {
	return l.GetFirst()
}

func (l *LinkedList) Peek() collection.Element {
	first := l.first
	if first == nil {
		return nil
	}
	return first.elem
}
