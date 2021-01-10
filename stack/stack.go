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

import (
	"errors"
)

// 实现栈
var _ Stacker = (*Stack)(nil)
var (
	NotExistErr = errors.New("not exist")
)

// NewStack 创建栈
// 非协程安全
func NewStack() *Stack {
	return &Stack{}
}

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		data interface{}
		prev *node
	}
)

func (stack *Stack) IsEmpty() bool {
	return stack.Len() == 0
}

// Len 大小
func (stack *Stack) Len() int {
	return stack.length
}

// Peek 返回栈顶
// 当栈顶为空时,返回 nil, NotExistErr
func (stack *Stack) Peek() (interface{}, error) {
	if stack.length == 0 {
		return nil, NotExistErr
	}
	return stack.top.data, nil
}

// Pop 移除并返回当前栈顶
// 当栈顶为空时,返回 nil, NotExistErr
func (stack *Stack) Pop() (interface{}, error) {
	if stack.length == 0 {
		return nil, NotExistErr
	}
	stack.length--
	top := stack.top
	// 将栈顶指向当前栈顶的下一个节点
	stack.top = top.prev
	return stack.top.data, nil
}

// Push 入栈
func (stack *Stack) Push(v interface{}) {
	stack.length++
	n := &node{
		data: v,
		prev: stack.top,
	}
	stack.top = n
}

// Clean 清空栈
func (stack *Stack) Clean() {
	stack.top = nil
	stack.length = 0
}
