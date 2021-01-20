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

package api

type Element interface {
}
type Collection interface {
	Size() int
	IsEmpty() bool
	Contains(e Element) (b bool)
	Add(e Element)
	Remove(e Element) (b bool)
	ContainsAll(collection Collection)
	AddAll(collection Collection) (b bool)
	RemoveAll(collection Collection) (b bool)
	Clear()
	Equals(collection Collection) (b bool)
	Slice() []Element
	Iterator() Iterator
}
