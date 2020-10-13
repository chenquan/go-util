/*
 *    Copyright 2020 Chen Quan
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

package group

// Grouper 实现该接口进行分组
type Grouper interface {
	Size() int           // 列表大小
	Get(int) interface{} // 通过下标获取某一个元素
	IsEmpty() bool       //列表是否为空
}

// By 列表分组
func By(grouper Grouper, functionBy func(obj interface{}) interface{}) (result map[interface{}][]interface{}) {
	result = make(map[interface{}][]interface{})
	if grouper.IsEmpty() {
		return
	}
	for i := 0; i < grouper.Size(); i++ {
		obj := grouper.Get(i)
		// 分组的key
		key := functionBy(obj)
		if _, ok := result[key]; !ok {
			result[key] = []interface{}{obj}
		} else {
			result[key] = append(result[key], obj)
		}
	}
	return
}

// BySlice 多列表分组
func BySlice(grouper []Grouper, functionBy func(obj interface{}) interface{}) (results []map[interface{}][]interface{}) {

	if len(grouper) == 0 {
		return
	}
	for _, g := range grouper {
		results = append(results, By(g, functionBy))
	}
	return
}

// BySliceCount 多列表分组统计
func BySliceCount(grouper []Grouper, functionBy func(obj interface{}) interface{}) (results []map[interface{}]int) {

	if len(grouper) == 0 {
		return
	}
	for _, g := range grouper {
		results = append(results, ByCount(g, functionBy))
	}
	return
}

// ByCount 列表分组统计
func ByCount(grouper Grouper, functionBy func(obj interface{}) interface{}) (result map[interface{}]int) {
	result = map[interface{}]int{}
	if grouper.IsEmpty() {
		return
	}
	for i := 0; i < grouper.Size(); i++ {
		obj := grouper.Get(i)
		// 分组的key
		key := functionBy(obj)
		result[key] += 1
	}
	return
}
