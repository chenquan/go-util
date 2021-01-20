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
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want *Stack
	}{
		{
			"1",
			&Stack{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testStack(t *testing.T, stack Stacker) {
	assert.Equal(t, true, stack.IsEmpty())

	stack.Push("1")
	assert.Equal(t, false, stack.IsEmpty())
	assert.Equal(t, 1, stack.Len())

	stack.Push("2")
	assert.Equal(t, false, stack.IsEmpty())
	assert.Equal(t, 2, stack.Len())

	stack.Push("3")
	assert.Equal(t, false, stack.IsEmpty())
	assert.Equal(t, 3, stack.Len())

	peek, err := stack.Peek()
	assert.Equal(t, nil, err)
	assert.Equal(t, "3", peek)
	assert.Equal(t, 3, stack.Len())

	pop, err := stack.Pop()
	assert.Equal(t, nil, err)
	assert.Equal(t, "3", pop)
	assert.Equal(t, 2, stack.Len())

	stack.Clean()
	assert.Equal(t, true, stack.IsEmpty())
	assert.Equal(t, 0, stack.Len())

	pop, err = stack.Pop()
	assert.Equal(t, NotExistErr, err)
	assert.Equal(t, nil, pop)
	assert.Equal(t, 0, stack.Len())

	peek, err = stack.Peek()
	assert.Equal(t, NotExistErr, err)
	assert.Equal(t, nil, peek)
	assert.Equal(t, 0, stack.Len())
}
func TestStack(t *testing.T) {
	stacks := [3]Stacker{
		NewStack(),
		NewDefaultSyncStack(),
		NewSyncStack(NewStack()),
	}
	for _, s := range stacks {

		testStack(t, s)
	}
}
