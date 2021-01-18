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

package sort

// Uint16s 无符号16位整型切片
type Uint16s []uint16

// Len 实现 sort.Interface 接口
func (s Uint16s) Len() int { return len(s) }

// Len 实现 sort.Interface 接口
func (s Uint16s) Less(i, j int) bool { return s[i] < s[j] }

// Len 实现 sort.Interface 接口
func (s Uint16s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Int16s 16位整型切片
type Int16s []int16

// Len 实现 sort.Interface 接口
func (s Int16s) Len() int { return len(s) }

// Len 实现 sort.Interface 接口
func (s Int16s) Less(i, j int) bool { return s[i] < s[j] }

// Len 实现 sort.Interface 接口
func (s Int16s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Uint32s 无符号32位整型切片
type Uint32s []uint32

// Len 实现 sort.Interface 接口
func (s Uint32s) Len() int { return len(s) }

// Len 实现 sort.Interface 接口
func (s Uint32s) Less(i, j int) bool { return s[i] < s[j] }

// Len 实现 sort.Interface 接口
func (s Uint32s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Int32s 64位整型切片
type Int32s []int32

// Len 实现 sort.Interface 接口
func (s Int32s) Len() int { return len(s) }

// Len 实现 sort.Interface 接口
func (s Int32s) Less(i, j int) bool { return s[i] < s[j] }

// Len 实现 sort.Interface 接口
func (s Int32s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Len 实现 sort.Interface 接口
func (s Uint64s) Len() int { return len(s) }

// Uint64s 无符号64位整型切片
type Uint64s []uint64

// Len 实现 sort.Interface 接口
func (s Uint64s) Less(i, j int) bool { return s[i] < s[j] }

// Len 实现 sort.Interface 接口
func (s Uint64s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Int64s 64位整型切片
type Int64s []int64

// Len 实现 sort.Interface 接口
func (s Int64s) Len() int { return len(s) }

// Len 实现 sort.Interface 接口
func (s Int64s) Less(i, j int) bool { return s[i] < s[j] }

// Len 实现 sort.Interface 接口
func (s Int64s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
