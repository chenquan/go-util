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
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsEmpty 判断字符串是否为空
func IsEmpty(s string) bool {
	return len(s) == 0
}

// IsNotEmpty 判断字符串是否不为空
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// IsAnyEmpty 是否存在空字符串
func IsAnyEmpty(strings ...string) bool {
	if len(strings) == 0 {
		return true
	}
	for _, s := range strings {
		if IsEmpty(s) {
			return true
		}
	}
	return false
}

// IsNoneEmpty 判断是否不存在空字符串
func IsNoneEmpty(strings ...string) bool {
	return !IsAnyEmpty(strings...)
}

// IsBlank 断某字符串是否为空或长度为0或由空白符(whitespace) 构成
func IsBlank(s string) bool {
	runes := []rune(s)
	for _, c := range runes {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

func IsAnyBlank(strings ...string) bool {
	if len(strings) == 0 {
		return true
	}

	for _, s := range strings {
		if IsBlank(s) {
			return true
		}
	}
	return false

}

func IsNoneBlank(strings ...string) bool {
	return !IsAnyBlank(strings...)
}

func Trim(s string) string {
	return strings.TrimSpace(s)
}
func Strip(s string, stripChars string) string {
	if IsEmpty(s) {
		return s
	}
	s = StripStart(s, stripChars)
	s = StripEnd(s, stripChars)
	return s
}
func StripStart(str string, stripChars string) string {
	return strings.TrimLeft(str, stripChars)
}
func StripEnd(str string, stripChars string) string {
	return strings.TrimRight(str, stripChars)
}
func StripAll(strings []string, stripChars string) []string {
	if len(strings) != 0 {
		newStrings := make([]string, len(strings))
		for i, s := range strings {
			newStrings[i] = Strip(s, stripChars)
		}
		return newStrings
	}
	return strings
}

func contains(s string, searchChar string) bool {
	if IsEmpty(s) {
		return false
	}

	return strings.Index(s, searchChar) >= 0
}
func SubstringStart(str string, start int) string {
	if start < 0 {
		start += utf8.RuneCountInString(str)
	}
	if start < 0 {
		start = 0
	}
	if start > utf8.RuneCountInString(str) {
		return str
	} else {
		return string([]rune(str)[start:])
	}
}
func Substring(str string, start int, end int) string {
	strLen := utf8.RuneCountInString(str)
	if strLen == 0 {
		return ""
	}
	if end < 0 {
		end += strLen
	}
	if start < 0 {
		start += strLen
	}
	if end > strLen {
		end = strLen
	}
	if start > end {
		return ""
	} else {
		if start < 0 {
			start = 0
		}
		if end < 0 {
			end = 0
		}
		return string([]rune(str)[start:end])
	}
}
func Left(str string, n int) string {
	if n < 0 {
		return ""
	} else {
		if utf8.RuneCountInString(str) < n {
			return str
		} else {
			return string([]rune(str)[0:n])
		}
	}

}
func Right(str string, n int) string {
	if n <= 0 {
		return ""
	} else {
		strLen := utf8.RuneCountInString(str)
		if strLen < n {
			return str
		} else {
			return string([]rune(str)[strLen-n:])
		}
	}
}

// IsNumber 是否数字
func IsNumber(str string) bool {
	reg, _ := regexp.Compile("^\\d+$")
	return reg.MatchString(str)
}

// IsNumerical 是否数字
func IsNumerical(s string) bool {
	reg, _ := regexp.Compile("^\\d+.?\\d*$")
	return reg.MatchString(s)
}
