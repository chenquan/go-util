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

func TestStack_Push(t *testing.T) {
	stack := NewStack()
	strings := []interface{}{"1", 2, 23.23, nil,
		struct {
		}{},
		struct {
			name string
		}{name: "1121"},
	}
	for _, s := range strings {
		stack.Push(s)
		if stack.top.data != s {
			t.Errorf("Push() = %v, want %v", s, stack.top.data)
		}
	}
}

func TestStack_Clean(t *testing.T) {
	stack := NewStack()
	stack.Clean()
	if stack.top != nil || stack.length != 0 {
		t.Errorf("Clean() error")
	}
}
