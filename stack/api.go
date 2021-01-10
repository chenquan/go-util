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

package stack

// Stacker 实现Stacker就拥有栈的功能
type Stacker interface {
	Len() int                   // 大小
	IsEmpty() bool              // 是否空栈
	Peek() (interface{}, error) // 返回栈顶
	Pop() (interface{}, error)  // 移除并返回当前栈顶
	Push(v interface{})         // 入栈
	Clean()                     // 清空栈
}
