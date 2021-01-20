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

type List interface {
	Collection
	AddAllIndex(index int, c Collection)
	Get(index int) (e Element, err error)
	Set(index int, e Element)
	AddIndex(index int, e Element)
	RemoveIndex(index int)
	Index(e Element) (index int)
	LastIndex(e Element) (index int)
	SubList(fromIndex, toIndex int) (list List)
	RetainAll(c Collection) (modified bool)
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
