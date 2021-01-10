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

package str

import "testing"

func TestIsNumerical(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{"1.2"},
			true,
		},
		{
			"2",
			args{"2131.1212121212"},
			true,
		},
		{
			"3",
			args{"1"},
			true,
		},
		{
			"4",
			args{"123232"},
			true,
		},
		{
			"5",
			args{"000.123232"},
			true,
		},
		{
			"6",
			args{"asd.sd"},
			false,
		},
		{
			"7",
			args{"asd"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumerical(tt.args.s); got != tt.want {
				t.Errorf("IsNumerical() = %v, want %v", got, tt.want)
			}
		})
	}
}
