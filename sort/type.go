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

type Uint64s []uint64
type Uint32s []uint32
type Uint16s []uint16

type Int64s []int64
type Int32s []int32
type Int16s []int16

func (s Uint16s) Len() int           { return len(s) }
func (s Uint16s) Less(i, j int) bool { return s[i] < s[j] }
func (s Uint16s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s Int16s) Len() int           { return len(s) }
func (s Int16s) Less(i, j int) bool { return s[i] < s[j] }
func (s Int16s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s Uint32s) Len() int           { return len(s) }
func (s Uint32s) Less(i, j int) bool { return s[i] < s[j] }
func (s Uint32s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s Int32s) Len() int           { return len(s) }
func (s Int32s) Less(i, j int) bool { return s[i] < s[j] }
func (s Int32s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s Uint64s) Len() int           { return len(s) }
func (s Uint64s) Less(i, j int) bool { return s[i] < s[j] }
func (s Uint64s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s Int64s) Len() int           { return len(s) }
func (s Int64s) Less(i, j int) bool { return s[i] < s[j] }
func (s Int64s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
