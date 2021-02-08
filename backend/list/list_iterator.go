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

// itrList 实现 collection.Iterator 接口
type itrList struct {
	cursor  int             // 游标,指向下一个元素
	lastRet int             // 最近一次返回的下标
	data    collection.List // 数据
}

// HasNext 如果当前迭代还有更多的元素则返回 true,否则返回 false
func (s *itrList) HasNext() bool {
	return s.data.Size() != s.cursor
}

// Next 返回当前迭代中的下一个元素
func (s *itrList) Next() (collection.Element, error) {
	if s.cursor >= s.data.Size() {
		return nil, errs.NoSuchElement
	}
	s.lastRet = s.cursor
	s.cursor++
	return s.data.Get(s.lastRet)
}

// Remove 从基础集合中移除当前迭代器返回的最后一个元素
//
// 每次调用 Next 方法,才可以调用一次此方法.
func (s *itrList) Remove() error {
	if s.lastRet < 0 {
		return errs.IllegalState
	}
	_, err := s.data.RemoveIndex(s.lastRet)
	if err == nil {
		s.cursor = s.lastRet
		s.lastRet = -1
	}
	return err
}
