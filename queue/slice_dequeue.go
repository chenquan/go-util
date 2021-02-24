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
	if e == nil {
		return false, errs.NilPointer
	}
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

func (s *SliceDeQueue) ContainsAll(collection collection.Collection) (bool, error) {
	panic("implement me")
}

func (s *SliceDeQueue) AddAll(collection collection.Collection) (bool, error) {
	panic("implement me")
}

func (s *SliceDeQueue) RemoveAll(collection collection.Collection) (bool, error) {
	panic("implement me")
}

func (s *SliceDeQueue) RetainAll(collection collection.Collection) (bool, error) {
	panic("implement me")
}

func (s *SliceDeQueue) Clear() error {
	panic("implement me")
}

func (s *SliceDeQueue) Equals(collection collection.Collection) bool {
	panic("implement me")
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
	panic("implement me")
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
	for ; x != nil; x = s.elements[i] {
		if e == x {
			// TODO
			return true, nil
		}
	}
	return false, nil
}

func (s *SliceDeQueue) RemoveLastOccurrence(e collection.Element) (bool, error) {
	panic("implement me")
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
