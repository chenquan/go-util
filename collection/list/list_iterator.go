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

import "github.com/chenquan/go-utils/collection/api/collection"

type listIterator struct {
	cursor  int //游标,指向下一个元素
	lastRet int
	data    collection.List
}

func (s *listIterator) HasNext() bool {
	return s.data.Size() != s.cursor
}

func (s *listIterator) Next() (collection.Element, error) {
	if s.cursor >= s.data.Size() {
		return nil, NoSuchElement
	}
	s.lastRet = s.cursor
	s.cursor++
	return s.data.Get(s.lastRet)
}

func (s *listIterator) Remove() error {
	if s.lastRet < 0 {
		return IllegalState
	}
	_, err := s.data.RemoveIndex(s.lastRet)
	if err == nil {
		s.cursor = s.lastRet
		s.lastRet = -1
	}
	return err
}
