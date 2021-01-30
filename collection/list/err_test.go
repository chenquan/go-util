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

package list

import (
	"reflect"
	"testing"
)

func TestNewIndexOutOfBoundsException(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want *indexOutOfBoundsException
	}{
		{
			"1",
			args{""},
			&indexOutOfBoundsException{str: ""},
		}, {
			"2",
			args{"error"},
			&indexOutOfBoundsException{str: "error"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIndexOutOfBoundsException(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndexOutOfBoundsException() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIndexOutOfBoundsExceptionDefault(t *testing.T) {
	tests := []struct {
		name string
		want *indexOutOfBoundsException
	}{
		{
			"1",
			&indexOutOfBoundsException{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIndexOutOfBoundsExceptionDefault(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndexOutOfBoundsExceptionDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexOutOfBoundsException_Error(t *testing.T) {
	type fields struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"1",
			fields{str: "1"},
			"index out of bound,1",
		}, {
			"2",
			fields{str: ""},
			"index out of bound",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := indexOutOfBoundsException{
				str: tt.fields.str,
			}
			if got := i.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
