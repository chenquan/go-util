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

package queue

import (
	"fmt"
	"github.com/chenquan/go-util/backend/collection"
	"github.com/chenquan/go-util/errs"
)

type SliceDeQueue struct {
	elements []collection.Element
	head     int
	tail     int
}

func (s *SliceDeQueue) Size() int {
	return (s.tail-s.head)&len(s.elements) - 1
}

func (s *SliceDeQueue) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SliceDeQueue) Contains(e collection.Element) (bool, error) {
	if e == nil {
		return false, nil
	}
	n := len(s.elements) - 1
	i := s.head
	x := s.elements[i]
	for ; x != nil; x = s.elements[i] {
		if x == e {
			return true, nil
		}
		i = (i + 1) & n
	}
	return false, nil
}

func (s *SliceDeQueue) Add(e collection.Element) (bool, error) {
	return true, s.AddLast(e)
}

func (s *SliceDeQueue) Remove(e collection.Element) (bool, error) {
	iterator := s.Iterator()
	for iterator.HasNext() {
		next, _ := iterator.Next()
		if next == e {
			_ = iterator.Remove()
			return true, nil
		}
	}
	return false, nil
}

func (s *SliceDeQueue) ContainsAll(c collection.Collection) (bool, error) {
	if c == nil {
		return false, errs.NilPointer
	}
	iterator := c.Iterator()
	for iterator.HasNext() {
		next, _ := iterator.Next()
		if contains, _ := s.Contains(next); !contains {
			return false, nil
		}
	}
	return false, nil
}

func (s *SliceDeQueue) AddAll(c collection.Collection) (bool, error) {
	if c == nil {
		return false, errs.NilPointer
	}
	elements := c.Slice()
	modified := false
	for _, element := range elements {
		add, _ := s.Add(element)
		if add {
			modified = true
		}
	}
	return modified, nil
}

func (s *SliceDeQueue) RemoveAll(c collection.Collection) (bool, error) {
	if c == nil {
		return false, errs.NilPointer
	}
	modified := false

	iterator := s.Iterator()
	for iterator.HasNext() {
		e, _ := iterator.Next()
		contains, err := c.Contains(e)
		if err != nil {
			return false, err
		}
		if contains {
			_ = iterator.Remove()
			modified = true
		}
	}
	return modified, nil
}

func (s *SliceDeQueue) RetainAll(c collection.Collection) (bool, error) {
	if c == nil {
		return false, errs.NilPointer
	}
	modified := false

	iterator := s.Iterator()
	for iterator.HasNext() {
		e, _ := iterator.Next()
		contains, err := c.Contains(e)
		if err != nil {
			return false, err
		}
		if !contains {
			_ = iterator.Remove()
			modified = true
		}
	}
	return modified, nil
}

func (s *SliceDeQueue) Clear() error {
	h, t := s.head, s.tail
	if t != h {
		s.tail, s.head = 0, 0
		mask := len(s.elements) - 1
		i := h
		for i != t {
			// help cg
			s.elements[i] = nil
			i = (i + 1) & mask
		}
	}
	return nil
}

func (s *SliceDeQueue) Equals(c collection.Collection) bool {
	if s == c {
		return true
	}
	if s.Size() != c.Size() {
		return false
	}
	i1 := s.Iterator()
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

func (s *SliceDeQueue) Slice() []collection.Element {
	n := len(s.elements)
	h := s.head
	r := n - h
	elements := make([]collection.Element, n)
	copy(elements[0:r], s.elements[h:n])
	copy(elements[r:r+h], s.elements[0:h])
	return elements
}

func (s *SliceDeQueue) Iterator() collection.Iterator {
	return &sliceDequeueItr{data: s, cursor: s.head, fence: s.tail, lastRet: -1}
}

func (s *SliceDeQueue) Offer(e collection.Element) (bool, error) {
	err := s.AddLast(e)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *SliceDeQueue) Poll() collection.Element {
	element, _ := s.RemoveFirst()
	return element
}

func (s *SliceDeQueue) Delete() (collection.Element, error) {
	return s.RemoveFirst()
}

func (s *SliceDeQueue) Element() (collection.Element, error) {
	return s.GetFirst()
}

func (s *SliceDeQueue) Peek() collection.Element {
	return s.elements[s.head]
}

func (s *SliceDeQueue) AddFirst(e collection.Element) error {
	if e == nil {
		return errs.NilPointer
	}
	s.head = (s.head - 1) & (len(s.elements) - 1)
	s.elements[s.head] = e
	if s.head == s.tail {
		if err := s.doubleCapacity(); err != nil {
			return err
		}
	}
	return nil
}

func (s *SliceDeQueue) AddLast(e collection.Element) error {
	if e == nil {
		return errs.NilPointer
	}
	s.elements[s.tail] = e
	s.tail = (s.tail + 1) & (len(s.elements) - 1)
	if s.tail == s.head {
		if err := s.doubleCapacity(); err != nil {
			return err
		}
	}
	return nil
}

func (s *SliceDeQueue) RemoveFirst() (collection.Element, error) {
	element := s.elements[s.head]
	if element == nil {
		return nil, errs.NoSuchElement
	}
	s.elements[s.head] = nil
	s.head = (s.head + 1) & (len(s.elements) - 1)
	return element, nil
}

func (s *SliceDeQueue) RemoveLast() (collection.Element, error) {
	t := (s.tail - 1) & (len(s.elements) - 1)
	element := s.elements[t]
	if element == nil {
		return nil, errs.NoSuchElement
	}
	s.elements[t] = nil
	s.tail = t
	return element, nil
}

func (s *SliceDeQueue) GetFirst() (collection.Element, error) {
	element := s.elements[s.head]
	if element == nil {
		return nil, errs.NoSuchElement
	}
	return element, nil
}

func (s *SliceDeQueue) GetLast() (collection.Element, error) {
	element := s.elements[(s.tail-1)&(len(s.elements)-1)]
	if element == nil {
		return nil, errs.NoSuchElement
	}
	return element, nil
}

func (s *SliceDeQueue) RemoveFirstOccurrence(e collection.Element) (bool, error) {
	i := s.head
	x := s.elements[i]
	mask := len(s.elements) - 1
	for ; x != nil; x = s.elements[i] {
		if e == x {
			if _, err := s.delete(i); err != nil {
				return false, err
			}
			return true, nil
		}
		i = (i + 1) & mask
	}
	return false, nil
}
func (s *SliceDeQueue) delete(index int) (bool, error) {
	h := s.head
	t := s.tail
	mask := len(s.elements) - 1
	front := (index - h) & mask
	back := (t - index) & mask
	if front >= ((t - h) & mask) {
		return false, errs.ConcurrentModification
	}
	if front < back {
		if h <= index {
			copy(s.elements[h+1:h+1+front+1], s.elements[h:h+front+1])
		} else {
			copy(s.elements[1:index+2], s.elements[:index+1])
			s.elements[0] = mask
			copy(s.elements[h+1:h+mask-h+1], s.elements[h:h+mask-h+1])
		}
		s.elements[h] = nil
		s.head = (h + 1) & mask
		return true, nil
	} else {
		if index < t {
			copy(s.elements[index:index+back+1], s.elements[index+1:index+1+back+1])
			s.tail = t - 1
		} else {
			copy(s.elements[index:index+mask-index+1], s.elements[index+1:index+1+mask-index+1])
			s.elements[mask] = s.elements[0]
			copy(s.elements[:t+1], s.elements[1:1+t+1])
			s.tail = (t - 1) & mask
		}
	}

	return true, nil
}

func (s *SliceDeQueue) RemoveLastOccurrence(e collection.Element) (bool, error) {
	if e == nil {
		return false, nil
	}
	mask := len(s.elements) - 1
	i := (s.tail - 1) & mask

	for x := s.elements[i]; x != nil; x = s.elements[i] {
		if x == e {
			if _, err := s.delete(i); err != nil {
				return false, err
			}
			return true, nil
		}
		i = (i - 1) & mask
	}
	return false, nil
}

func (s *SliceDeQueue) Push(e collection.Element) error {
	return s.AddFirst(e)
}

func (s *SliceDeQueue) Pop() (collection.Element, error) {
	return s.RemoveFirst()
}

func (s *SliceDeQueue) DescendingIterator() collection.Iterator {
	panic("implement me")
}

func (s *SliceDeQueue) doubleCapacity() error {
	n := len(s.elements)
	h := s.head
	r := n - h
	newSize := n << 1
	if newSize < 0 {
		return fmt.Errorf("deque too big")
	}
	elements := make([]collection.Element, newSize)
	copy(elements[0:r], s.elements[h:n])
	copy(elements[r:r+h], s.elements[0:h])
	s.elements = elements
	s.head = 0
	s.tail = n
	return nil
}

type sliceDequeueItr struct {
	data    *SliceDeQueue
	cursor  int
	fence   int
	lastRet int
}

func (s *sliceDequeueItr) HasNext() bool {
	return s.lastRet != s.fence
}

func (s *sliceDequeueItr) Next() (collection.Element, error) {
	if !s.HasNext() {
		return nil, errs.NoSuchElement
	}
	element := s.data.elements[s.cursor]
	if s.HasNext() || element == nil {
		return nil, errs.ConcurrentModification
	}
	s.lastRet = s.cursor
	s.cursor = (s.cursor + 1) & (len(s.data.elements) - 1)
	return element, nil
}

func (s *sliceDequeueItr) Remove() error {
	if s.lastRet < 0 {
		return errs.IllegalState
	}
	if b, err := s.data.delete(s.cursor); err != nil {
		return err
	} else if b {
		s.cursor = (s.cursor - 1) & (len(s.data.elements) - 1)
		s.fence = s.data.tail
	}
	s.lastRet = -1
	return nil
}
