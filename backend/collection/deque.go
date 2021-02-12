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

// DeQue 双端队列
type DeQue interface {
	// 实现队列接口
	Queue
	AddFirst(e Element) error
	AddLast(e Element) error
	RemoveFirst() (Element, error)
	RemoveLast() (Element, error)
	GetFirst() (Element, error)
	GetLast() (Element, error)
	RemoveFirstOccurrence(e Element) (bool, error)
	RemoveLastOccurrence(e Element) (bool, error)
	// Stack methods
	Push(e Element) error
	Pop() (Element, error)
	DescendingIterator() Iterator
}
