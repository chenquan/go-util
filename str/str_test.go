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

import (
	"reflect"
	"testing"
)

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

func TestLeft(t *testing.T) {
	type args struct {
		str string
		n   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				str: "你好！世界！",
				n:   1,
			},
			"你",
		}, {
			"2",
			args{
				str: "你好！世界！",
				n:   3,
			},
			"你好！",
		},
		{
			"3",
			args{
				str: "你好！世界！",
				n:   1000,
			},
			"你好！世界！",
		},

		{
			"4",
			args{
				str: "你好！世界！",
				n:   0,
			},
			"",
		},
		{
			"5",
			args{
				str: "你好！世界！",
				n:   -1,
			},
			"",
		},
		{
			"6",
			args{
				str: "你好！世界！",
				n:   -100,
			},
			"",
		}, {
			"7",
			args{
				str: "你好！世界！",
				n:   6,
			},
			"你好！世界！",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Left(tt.args.str, tt.args.n); got != tt.want {
				t.Errorf("Left() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRight(t *testing.T) {
	type args struct {
		str string
		n   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				str: "你好！世界！",
				n:   0,
			},
			"",
		}, {
			"2",
			args{
				str: "你好！世界！",
				n:   6,
			},
			"你好！世界！",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Right(tt.args.str, tt.args.n); got != tt.want {
				t.Errorf("Right() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumber(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{"112121"},
			true,
		}, {
			"2",
			args{"0112121"},
			true,
		}, {
			"3",
			args{"0112.121"},
			false,
		}, {
			"4",
			args{"0112121."},
			false,
		}, {
			"5",
			args{"a1."},
			false,
		}, {
			"6",
			args{"a1.121s"},
			false,
		}, {
			"7",
			args{"abs"},
			false,
		}, {
			"8",
			args{"1e-6"},
			false,
		}, {
			"9",
			args{"你好"},
			false,
		}, {
			"10",
			args{"@"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumber(tt.args.str); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstring(t *testing.T) {
	type args struct {
		str   string
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				str:   "你好！世界！",
				start: 0,
				end:   0,
			},
			"",
		}, {
			"2",
			args{
				str:   "你好！世界！",
				start: -1,
				end:   -1,
			},
			"",
		}, {
			"3",
			args{
				str:   "你好！世界！",
				start: -6,
				end:   -6,
			},
			"",
		}, {
			"4",
			args{
				str:   "你好！世界！",
				start: -7,
				end:   -6,
			},
			"",
		}, {
			"5",
			args{
				str:   "你好！世界！",
				start: -1000,
				end:   -6,
			},
			"",
		}, {
			"6",
			args{
				str:   "你好！世界！",
				start: -1000,
				end:   -1,
			},
			"你好！世界",
		}, {
			"7",
			args{
				str:   "你好！世界！",
				start: -1000,
				end:   -4,
			},
			"你好",
		}, {
			"7",
			args{
				str:   "你好！世界！",
				start: -1000,
				end:   -4,
			},
			"你好",
		}, {
			"8",
			args{
				str:   "你好！世界！",
				start: 1,
				end:   -4,
			},
			"好",
		}, {
			"9",
			args{
				str:   "你好！世界！",
				start: 1,
				end:   2,
			},
			"好",
		}, {
			"9",
			args{
				str:   "你好！世界！",
				start: 1,
				end:   0,
			},
			"",
		}, {
			"10",
			args{
				str:   "",
				start: 1,
				end:   0,
			},
			"",
		}, {
			"10",
			args{
				str:   "abc",
				start: 0,
				end:   4,
			},
			"abc",
		}, {
			"11",
			args{
				str:   "abc",
				start: -1000,
				end:   -100,
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substring(tt.args.str, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("Substring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAnyBlank(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{[]string{}},
			true,
		}, {
			"2",
			args{[]string{" "}},
			true,
		}, {
			"3",
			args{[]string{" ", "sdasd"}},
			true,
		}, {
			"4",
			args{[]string{"sdasd"}},
			false,
		}, {
			"5",
			args{[]string{"你", " \n"}},
			true,
		}, {
			"6",
			args{[]string{"你\n"}},
			false,
		}, {
			"7",
			args{[]string{"你\n", "\t"}},
			true,
		}, {
			"8",
			args{[]string{"你\n", "\r\t\n"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnyBlank(tt.args.strings...); got != tt.want {
				t.Errorf("IsAnyBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAnyEmpty(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{[]string{}},
			true,
		}, {
			"2",
			args{[]string{""}},
			true,
		}, {
			"3",
			args{[]string{"", " "}},
			true,
		}, {
			"4",
			args{[]string{"", "hello"}},
			true,
		}, {
			"6",
			args{[]string{"", "", "你好世界！"}},
			true,
		}, {
			"7",
			args{[]string{" ", "\n"}},
			false,
		}, {
			"8",
			args{[]string{" ", "\t", "\n"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnyEmpty(tt.args.strings...); got != tt.want {
				t.Errorf("IsAnyEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlank(t *testing.T) {
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
			args{"1"},
			false,
		}, {
			"2",
			args{""},
			true,
		}, {
			"3",
			args{"\n"},
			true,
		}, {
			"4",
			args{"\t"},
			true,
		}, {
			"5",
			args{"\r"},
			true,
		}, {
			"6",
			args{"\v"},
			true,
		}, {
			"7",
			args{"\v"},
			true,
		}, {
			"8",
			args{"\f"},
			true,
		}, {
			"8",
			args{" \f"},
			true,
		}, {
			"9",
			args{"s \f "},
			false,
		}, {
			"10",
			args{" s \f "},
			false,
		}, {
			"11",
			args{"  \f\v\t s "},
			false,
		}, {
			"11",
			args{" s"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run("IsBlank/"+tt.name, func(t *testing.T) {
			if got := IsBlank(tt.args.s); got != tt.want {
				t.Errorf("IsBlank() = %v, want %v", got, tt.want)
			}
		})
		t.Run("IsNotBlank/"+tt.name, func(t *testing.T) {
			if got := IsNotBlank(tt.args.s); got != !tt.want {
				t.Errorf("IsBlank() = %v, want %v", got, !tt.want)
			}
		})
	}
}

func TestEmpty(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"1"}, false},
		{"2", args{""}, true},
	}
	for _, tt := range tests {
		t.Run("IsEmpty/"+tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.s); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
		t.Run("IsNotEmpty"+tt.name, func(t *testing.T) {
			if got := IsNotEmpty(tt.args.s); got != !tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, !tt.want)
			}
		})
	}
}

func TestIsNoneBlank(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]string{"", "1"}}, false},
		{"2", args{[]string{"\n", "1"}}, false},
		{"3", args{[]string{"\v", "1"}}, false},
		{"4", args{[]string{"\t", "1"}}, false},
		{"5", args{[]string{"\r", "1"}}, false},
		{"6", args{[]string{"\r", "\n"}}, false},
		{"7", args{[]string{"1", " \nssdasdad"}}, true},
		{"8", args{[]string{"1"}}, true},
		{"9", args{[]string{"1", "ssdasdad"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNoneBlank(tt.args.strings...); got != tt.want {
				t.Errorf("IsNoneBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNoneEmpty(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{strings: []string{"1", ""}},
			false,
		}, {
			"2",
			args{strings: []string{"1 ", ""}},
			false,
		}, {
			"3",
			args{strings: []string{"1 "}},
			true,
		}, {
			"4",
			args{strings: []string{"\n"}},
			true,
		}, {
			"5",
			args{strings: []string{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNoneEmpty(tt.args.strings...); got != tt.want {
				t.Errorf("IsNoneEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrip(t *testing.T) {
	type args struct {
		s          string
		stripChars string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				s:          "! ss !",
				stripChars: "!",
			},
			" ss ",
		}, {
			"2",
			args{
				s:          "! ss",
				stripChars: "!",
			},
			" ss",
		}, {
			"3",
			args{
				s:          "! ss",
				stripChars: "@",
			},
			"! ss",
		}, {
			"4",
			args{
				s:          "!!!! ss",
				stripChars: "!",
			},
			" ss",
		}, {
			"8",
			args{
				s:          "!!!! ss !",
				stripChars: "!",
			},
			" ss ",
		}, {
			"9",
			args{
				s:          "",
				stripChars: "!",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Strip(tt.args.s, tt.args.stripChars); got != tt.want {
				t.Errorf("Strip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripAll(t *testing.T) {
	type args struct {
		strings    []string
		stripChars string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"1",
			args{
				strings:    []string{"! ss !", "ss !", " ! ! "},
				stripChars: "!",
			},
			[]string{" ss ", "ss ", " ! ! "},
		}, {
			"2",
			args{
				strings:    []string{},
				stripChars: "!",
			},
			[]string{},
		}, {
			"3",
			args{
				strings:    []string{"sdsd", "sda"},
				stripChars: "!",
			},
			[]string{"sdsd", "sda"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripAll(tt.args.strings, tt.args.stripChars); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StripAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripEnd(t *testing.T) {
	type args struct {
		str        string
		stripChars string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				str:        "! ss !",
				stripChars: "!",
			},
			"! ss ",
		}, {
			"2",
			args{
				str:        "! ss",
				stripChars: "!",
			},
			"! ss",
		}, {
			"3",
			args{
				str:        "! ss",
				stripChars: "@",
			},
			"! ss",
		}, {
			"4",
			args{
				str:        "!!!! ss",
				stripChars: "!",
			},
			"!!!! ss",
		}, {
			"8",
			args{
				str:        "!!!! ss !",
				stripChars: "!",
			},
			"!!!! ss ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripEnd(tt.args.str, tt.args.stripChars); got != tt.want {
				t.Errorf("StripEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripStart(t *testing.T) {
	type args struct {
		str        string
		stripChars string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			"1",
			args{
				str:        "! ss !",
				stripChars: "!",
			},
			" ss !",
		}, {
			"2",
			args{
				str:        "! ss",
				stripChars: "!",
			},
			" ss",
		}, {
			"3",
			args{
				str:        "! ss",
				stripChars: "@",
			},
			"! ss",
		}, {
			"4",
			args{
				str:        "!!!! ss",
				stripChars: "!",
			},
			" ss",
		}, {
			"8",
			args{
				str:        "ss !",
				stripChars: "!",
			},
			"ss !",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripStart(tt.args.str, tt.args.stripChars); got != tt.want {
				t.Errorf("StripStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstringStart(t *testing.T) {
	type args struct {
		str   string
		start int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				str:   "12345",
				start: 1,
			},
			"2345",
		}, {
			"2",
			args{
				str:   "你好!",
				start: -1,
			},
			"!",
		}, {
			"3",
			args{
				str:   "你好!",
				start: 9,
			},
			"",
		}, {
			"4",
			args{
				str:   "你好!",
				start: -9,
			},
			"你好!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubstringStart(tt.args.str, tt.args.start); got != tt.want {
				t.Errorf("SubstringStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"1"}, "1"},
		{"2", args{"1 "}, "1"},
		{"3", args{"    1 "}, "1"},
		{"4", args{"  	1 \n  1 \n "}, "1 \n  1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trim(tt.args.s); got != tt.want {
				t.Errorf("Trim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		s          string
		searchChar string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{
				s:          "1",
				searchChar: "2",
			},
			false,
		}, {
			"2",
			args{
				s:          "1",
				searchChar: "1",
			},
			true,
		}, {
			"3",
			args{
				s:          "123432234234",
				searchChar: "1",
			},
			true,
		}, {
			"4",
			args{
				s:          " 您好你",
				searchChar: " ",
			},
			true,
		}, {
			"5",
			args{
				s:          "",
				searchChar: "",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.searchChar); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
