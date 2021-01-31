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

package errs

import (
	"reflect"
	"testing"
)

func TestNewIndexOutOfBoundsError(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want *indexOutOfBoundsError
	}{
		{
			"1",
			args{""},
			&indexOutOfBoundsError{str: ""},
		}, {
			"2",
			args{"errs"},
			&indexOutOfBoundsError{str: "errs"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIndexOutOfBoundsError(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndexOutOfBoundsException() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIndexOutOfBoundsErrorDefault(t *testing.T) {
	tests := []struct {
		name string
		want *indexOutOfBoundsError
	}{
		{
			"1",
			&indexOutOfBoundsError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIndexOutOfBoundsErrorDefault(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndexOutOfBoundsExceptionDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexOutOfBoundsError_Error(t *testing.T) {
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
			i := indexOutOfBoundsError{
				str: tt.fields.str,
			}
			if got := i.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
