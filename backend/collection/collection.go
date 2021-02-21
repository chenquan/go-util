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

// Element 集合元素
type Element interface{}

// Collection 集合层次结构中的根接口
// 集合表示一组对象,称为其元素.
// 一些集合允许重复的元素,而另一些则不允许;
// 一些是有序的,而其他则是无序的.
type Collection interface {
	// Size 返回当前集合大小
	//
	// 在32位系统中集合大小最大可为 math.MaxInt32,
	// 在64位系统中集合大小最大可为 math.MaxInt64.
	Size() int
	// IsEmpty 如果当前集合没有存储元素则返回 true,否则返回 false
	IsEmpty() bool
	// Contains 如果当前集合包含元素 e 则返回 true,否则返回 false
	Contains(e Element) (bool, error)
	// Add 添加指定元素
	//
	// 如果当前集合由于调用而更改, 则返回true.
	// 如果此集合不允许重复并且已经包含指定的元素,则返回false.
	Add(e Element) (bool, error)
	// Remove 删除指定元素
	//
	// 如果当前集合中存在指定元素,则删除改元素并返回 true,否则返回 false.
	Remove(e Element) (bool, error)
	// ContainsAll 如果当前集合包含指定集合中的所有元素，则返回 true,否则返回 false.
	ContainsAll(collection Collection) (bool, error)
	// AddAll 将指定集合中的所有元素添加到当前集合中
	//
	// 如果调用 AddAll 改变了集合,则返回 true,否则返回 false.
	AddAll(collection Collection) (bool, error)
	// RemoveAll 删除当前集合中与指定集合相同的所有元素
	//
	// 如果调用 RemoveAll 改变了集合,则返回 true,否则返回 false.
	RemoveAll(collection Collection) (bool, error)
	// RetainAll 仅保留当前集合中包含在指定集合中的元素
	//
	// 换句话说,从当前集合中删除所有未包含在指定集合中的元素.
	RetainAll(collection Collection) (bool, error)
	// Clear 清空集合中所有元素
	//
	// 此方法返回后集合将为空.
	Clear() error
	// Equals 比较指定对象与此集合的相等性
	Equals(collection Collection) bool
	// Slice 返回一个包含此集合中所有元素的切片
	//
	// 如果此集合保证其迭代器返回元素的顺序，则此方法必须按相同的顺序返回元素.
	// 返回的切片是深度拷贝当前集合的数据,因此可以随意修改和删除返回的切片.
	Slice() []Element
	// Iterator 返回当前集合中元素的迭代器
	//
	// 如果当前集合是有序的,则保证返回的迭代器是有序的.
	Iterator() Iterator
}
