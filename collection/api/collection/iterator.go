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

type Iterator interface {
	// HasNext 如果当前迭代还有更多的元素则返回 true,否则返回 false
	HasNext() bool
	// Next 返回当前迭代中的下一个元素
	Next() (Element, error)
	// Remove 从基础集合中移除当前迭代器返回的最后一个元素
	//
	// 每次调用 Next 方法,才可以调用一次此方法.
	Remove() error
}
