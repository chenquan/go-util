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

import "sync"

var _ Stacker = (*SyncStack)(nil)

// NewDefaultSyncStack 创建默认同步栈
func NewDefaultSyncStack() *SyncStack {
	return &SyncStack{stack: NewStack()}
}

// NewSyncStack 创建同步栈
// stack 实现了 Stacker 接口的非安全栈
func NewSyncStack(stack Stacker) *SyncStack {
	return &SyncStack{stack: stack}
}

// SyncStack 同步栈
type SyncStack struct {
	lock  sync.RWMutex
	stack Stacker
}

// IsEmpty 空栈
func (stack *SyncStack) IsEmpty() bool {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	return stack.stack.IsEmpty()
}

// Len 栈大小
func (stack *SyncStack) Len() int {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	return stack.stack.Len()
}

// Peek 返回栈顶
func (stack *SyncStack) Peek() (interface{}, error) {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	return stack.stack.Peek()
}

// Pop 移除并返回当前栈顶
func (stack *SyncStack) Pop() (interface{}, error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	return stack.stack.Pop()
}

// Push 入栈
func (stack *SyncStack) Push(v interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.stack.Push(v)
}

// Clean 清空栈
func (stack *SyncStack) Clean() {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.stack.Clean()
}
