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

package _map

import (
	"github.com/chenquan/go-util/backend/collection"
	"github.com/chenquan/go-util/function"
)

type key interface{}
type Value interface{}

// Map
type Map interface {
	Size() int
	IsEmpty() bool
	ContainsKey(k key) (bool, error)
	ContainsValue(v Value) (bool, error)
	Get(k key) (Value, error)
	Put(k key, v Value) (Value, error)
	Remove(k key) (Value, error)
	PutAll(m Map) error
	Clear() error
	KeySet() collection.Set
	Values() collection.Collection
	EntrySet() collection.Set
	Equals(o interface{}) bool
	HashCode() int
	GetOrDefault(k key, defaultValue Value)
}

type Entry interface {
	Key() (key, error)
	Value() (Value, error)
	SetValue(value Value) (Value, error)
	Equals(o interface{}) bool
	HashCode() int
	ComparingByKey() function.Comparator
	ComparingByValue() function.Comparator
}
