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
	"github.com/chenquan/go-utils/collection/api/collection"
	"github.com/chenquan/go-utils/collection/api/iterator"
)

type List interface {
	collection.Collection
	AddAllIndex(index int, c collection.Collection)
	Get(index int) (e collection.Element, err error)
	Set(index int, e collection.Element)
	AddIndex(index int, e collection.Element)
	RemoveIndex(index int)
	Index(e collection.Element) (index int)
	LastIndex(e collection.Element) (index int)
	SubList(fromIndex, toIndex int) (list List)
	RetainAll(c collection.Collection) (modified bool)
}

type Iterator interface {
	iterator.Iterator
	HasPrevious() bool
	Previous() bool
	NextIndex() int
	PreviousIndex() int
	Set(e collection.Element)
	Add(e collection.Element)
}
