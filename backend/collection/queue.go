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

// Queue 队列接口
type Queue interface {
	Collection
	// Add 添加指定元素
	//
	// 如果当前集合由于调用而更改, 则返回true.
	// 如果此集合不允许重复并且已经包含指定的元素,则返回false.
	Add(e Element) (bool, error)
	// Offer 可以在不违反容量限制的情况下立即将指定的元素插入此队列
	//
	// 使用容量受限的队列时,通常最好使用add,因为add可能仅通过引发异常而无法插入元素.
	Offer(e Element) (bool, error)
	// Poll 返回并删除此队列的头
	//
	// 如果此队列为空,则返回nil.
	Poll() Element
	// Delete 检索并删除此队列的头,
	//
	// 此方法与poll不同之处仅在于,如果此队列为空,它将返回err.
	Delete() (Element, error)
	// Element 返回但不删除此队列的头
	//
	// 此方法与peek不同之处仅在于如果此队列为空,它将返回err.
	Element() (Element, error)
	// Peek 返回但不删除此队列的头
	//
	// 如果此队列为空,则返回nil.
	Peek() Element
}
