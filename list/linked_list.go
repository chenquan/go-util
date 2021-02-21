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

// linkedNode List of Nodes
type linkedNode struct {
	elem collection.Element // elements
	next *linkedNode        // Pointer to next node
	prev *linkedNode        // Pointer to previous node
}

// LinkedList Doubly-linked list implementation of the List and Deque interfaces.
//
// Implements all optional list operations, and permits all types (including nil).
type LinkedList struct {
	size  int         // LinkedList of size
	first *linkedNode // Pointer to first node
	last  *linkedNode // Pointer to last node
}

// NewLinkedList Create a empty linked list.
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// AddFirst Inserts the specified element at the beginning of this list.
func (l *LinkedList) AddFirst(e collection.Element) error {
	f := l.first
	n := &linkedNode{
		elem: e,
		next: f,
		prev: nil,
	}
	// the first element in the list
	l.first = n
	if f == nil {
		l.last = n
	} else {
		f.prev = n
	}
	l.size++
	return nil
}

// AddLast Appends the specified element to the end of this list.
//This method is equivalent to Add.
func (l *LinkedList) AddLast(e collection.Element) error {
	l.linkLast(e)
	return nil
}

// RemoveFirst Removes and returns the first element from this list.
func (l *LinkedList) RemoveFirst() (collection.Element, error) {
	first := l.first
	if first == nil {
		return nil, errs.NoSuchElement
	}
	next := first.next
	l.first = next
	if next == nil {
		// empty linked list
		l.last = nil
	} else {
		// set the previous node of the head node to nil
		l.first.prev = nil
	}
	l.size--
	return first.elem, nil
}

// RemoveLast Removes And returns the last element from this list.
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

// GetFirst Retrieves, but does not remove, the head (first element) of this list.
func (l *LinkedList) GetFirst() (collection.Element, error) {
	if l.first == nil {
		return nil, errs.NoSuchElement
	}
	return l.first.elem, nil
}

// GetLast Returns the last element in this list.
func (l *LinkedList) GetLast() (collection.Element, error) {
	if l.last == nil {
		return nil, errs.NoSuchElement
	}
	return l.last.elem, nil
}

// RemoveFirstOccurrence Removes the first occurrence of the specified element in this list (when traversing the list from head to tail).
//
// If the list does not contain the element, it is unchanged.
func (l *LinkedList) RemoveFirstOccurrence(e collection.Element) (bool, error) {
	return l.Remove(e)
}

// RemoveLastOccurrence Removes the last occurrence of the specified element in this list (when traversing the list from head to tail).
//
// If the list does not contain the element, it is unchanged.
func (l *LinkedList) RemoveLastOccurrence(e collection.Element) (bool, error) {
	for x := l.last; x != nil; x = x.prev {
		if x.elem == e {
			l.unLink(x)
			return true, nil
		}
	}
	return false, nil
}

// unLink unlink then specified node.
func (l *LinkedList) unLink(x *linkedNode) collection.Element {
	elem := x.elem
	prev := x.prev
	next := x.next
	if prev == nil {
		l.first = next
	} else {
		prev.next = next
		x.prev = nil
	}
	if next == nil {
		l.last = prev
	} else {
		next.prev = prev
		x.next = nil
	}
	x.elem = nil
	l.size--
	return elem

}

// Push Pushes an element onto the stack represented by this list.
//
// In other words, inserts the element at the front of this list.
func (l *LinkedList) Push(e collection.Element) error {
	return l.AddFirst(e)
}

// Pop Pops an element from the stack represented by this list.
//
// In other words, removes and returns the first element of this list.
func (l *LinkedList) Pop() (collection.Element, error) {
	return l.RemoveLast()
}

// DescendingIterator Returns a iterator of descending.
func (l *LinkedList) DescendingIterator() collection.Iterator {
	panic("implement me")
}

// AddAllIndex Inserts all of the elements in the specified collection into this list, starting at the specified position.
//
// Shifts the element currently at that position (if any) and any subsequent elements to the right (increases their indices).
//The new elements will appear in the list in the order that they are returned by the specified collection's iterator.
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
		prev *linkedNode
		p    *linkedNode
	)

	if l.size == index {
		p = nil
		prev = l.last
	} else {
		p = l.getNode(index)
		prev = p.prev
	}

	for _, element := range slice {
		n := &linkedNode{
			elem: element,
			next: nil,
			prev: prev,
		}
		if prev == nil {
			l.first = n
		} else {
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

// isPositionIndex Tells if the argument is the index of a valid position for an iterator or an add operation.
func (l *LinkedList) isPositionIndex(index int) bool {
	return index >= 0 && index <= l.size
}

// checkElementIndex Checks if the argument is the index of a error position.
func (l *LinkedList) checkElementIndex(index int) error {
	if index >= 0 && index < l.size {
		return nil
	}
	return errs.IndexOutOfBound
}

// checkPositionIndex Tells if the argument is the index of a error position for an iterator or an add operation.
func (l *LinkedList) checkPositionIndex(index int) error {
	if index >= 0 && index <= l.size {
		return nil
	}
	return errs.IndexOutOfBound
}

// Get Returns the element at the specified position in this list.
func (l *LinkedList) Get(index int) (collection.Element, error) {
	err := l.checkElementIndex(index)
	if err != nil {
		return nil, err
	}
	n := l.getNode(index)
	return n.elem, nil

}

// getNode Returns the (non-null) Node at the specified element index.
func (l *LinkedList) getNode(index int) *linkedNode {
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

// Set Replaces the element at the specified position in this list with the specified element.
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

// AddIndex Inserts the element at the specified position in this list.
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

// linkBefore Inserts element e before non-nil Node n.
func (l *LinkedList) linkBefore(e collection.Element, n *linkedNode) {
	prev := n.prev
	newNode := &linkedNode{
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

// linkLast Appends the specified element to the end of this list.
func (l *LinkedList) linkLast(e collection.Element) {
	last := l.last
	n := &linkedNode{
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

// RemoveIndex Removes the element at the specified position in this list.
//
// Shifts any subsequent elements to the left (subtracts one from their indices).
// Returns the element that was removed from the list.
func (l *LinkedList) RemoveIndex(index int) (collection.Element, error) {
	if err := l.checkElementIndex(index); err != nil {
		return nil, err
	}
	return l.unLink(l.getNode(index)), nil
}

// Index Returns the index of the first occurrence of the specified element in this list,
// or -1 if this list does not contain the element.
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

// Index Returns the index of the last occurrence of the specified element in this list,
// or -1 if this list does not contain the element.
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

// Size Returns the number of elements in this list.
func (l *LinkedList) Size() int {
	return l.size
}

// IsEmpty Returns true if this collection contains no elements.
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// Contains Returns true if this list contains the specified element.
func (l *LinkedList) Contains(e collection.Element) (bool, error) {
	return l.Index(e) >= 0, nil
}

// Add  Appends the specified element to the end of this list.
func (l *LinkedList) Add(e collection.Element) (bool, error) {
	l.linkLast(e)
	return true, nil
}

// Remove Removes the first occurrence of the specified element from this list, if it is present.
//
// If this list does not contain the element, it is unchanged.
func (l *LinkedList) Remove(e collection.Element) (bool, error) {
	for x := l.first; x != nil; x = x.next {
		if x.elem == e {
			l.unLink(x)
			return true, nil
		}
	}
	return false, nil
}

// ContainsAll  Returns true if this list contains the all elements of specified collection.
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

// AddAll Inserts all of the elements in the specified collection at end of the list.
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

// RemoveAll Removes all of the elements in the specified collection from the list.
func (l *LinkedList) RemoveAll(c collection.Collection) (bool, error) {
	return l.batchRemove(c, true)
}
func (l *LinkedList) batchRemove(c collection.Collection, complement bool) (bool, error) {
	iterator := l.Iterator()
	modified := false
	for iterator.HasNext() {
		next, _ := iterator.Next()
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
	return modified, nil
}

// RetainAll Retains all of the elements in the specified collection from this list.
func (l *LinkedList) RetainAll(c collection.Collection) (bool, error) {
	return l.batchRemove(c, false)
}

// Clear Removes all of the elements from this list.
// The list will be empty after this call returns.
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

// Equals Compares the specified object with this list for equality.
// Returns true if and only if the specified object is also a list, both lists have the same size, and all corresponding pairs of elements in the two lists are equal.
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

// Slice Returns an slice containing all of the elements in this list.
func (l *LinkedList) Slice() []collection.Element {
	elements := make([]collection.Element, 0, l.size)
	for x := l.first; x != nil; x = x.next {
		elements = append(elements, x.elem)
	}
	return elements
}

// Iterator Returns a iterator of list.
func (l *LinkedList) Iterator() collection.Iterator {
	return &itrLinkedList{
		data: l,
		next: l.first,
	}
}

// Offer Adds the specified element as the tail (last element) of this list.
func (l *LinkedList) Offer(e collection.Element) (bool, error) {
	return l.Add(e)
}

// Poll Retrieves and removes the head (first element) of this list.
func (l *LinkedList) Poll() collection.Element {
	first := l.first
	if first == nil {
		return nil
	}
	elem := first.elem
	l.unLink(first)
	return elem
}

// Element Retrieves, but does not remove, the head (first element) of this list.
func (l *LinkedList) Element() (collection.Element, error) {
	return l.GetFirst()
}

// Peek Retrieves, but does not remove, the head (first element) of this list.
func (l *LinkedList) Peek() collection.Element {
	first := l.first
	if first == nil {
		return nil
	}
	return first.elem
}

// Delete Removes and returns the first element from this list.
func (l *LinkedList) Delete() (collection.Element, error) {
	return l.RemoveFirst()
}

// itrLinkedList 实现链表迭代器
type itrLinkedList struct {
	data       *LinkedList
	next       *linkedNode
	lastReturn *linkedNode
	nextIndex  int
}

// HasNext Returns true if this list iterator has more elements when traversing the list in the forward direction.
func (itr *itrLinkedList) HasNext() bool {
	return itr.nextIndex < itr.data.size
}

// Next Returns the next element in the list and advances the cursor position.
func (itr *itrLinkedList) Next() (collection.Element, error) {
	if !itr.HasNext() {
		return nil, errs.IndexOutOfBound
	}
	itr.lastReturn = itr.next
	itr.next = itr.next.next
	itr.nextIndex++
	return itr.lastReturn.elem, nil
}

// Remove Removes from the list the last element that was returned by next or previous (optional operation).
//
// This call can only be made once per call to next or previous.
// It can be made only if add has not been called after the last call to next or previous.
func (itr *itrLinkedList) Remove() error {
	if itr.lastReturn == nil {
		return errs.IllegalState
	}
	lastNext := itr.lastReturn.next
	itr.data.unLink(itr.lastReturn)
	if itr.next == itr.lastReturn {
		itr.next = lastNext
	} else {
		itr.nextIndex--
	}
	itr.lastReturn = nil
	return nil
}
