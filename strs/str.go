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
	"github.com/chenquan/go-utils/convert"
	"github.com/chenquan/go-utils/math"
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

func Contains(s string, searchChar string) bool {
	if IsEmpty(s) {
		return false
	}

	return strings.Index(s, searchChar) >= 0
}
func SubstringStart(str string, start int) string {
	runes := []rune(str)
	strLen := len(runes)
	if start < 0 {
		start += strLen
	}
	if start < 0 {
		start = 0
	}
	if start > strLen {
		return ""
	} else {
		return string(runes[start:])
	}
}
func Substring(str string, start int, end int) string {
	runes := []rune(str)
	strLen := len(runes)

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
		return string(runes[start:end])
	}
}
func Left(str string, n int) string {
	if n < 0 {
		return ""
	} else {
		runes := []rune(str)
		if len(runes) < n {
			return str
		} else {
			return string(runes[0:n])
		}
	}

}
func Right(str string, n int) string {
	if n <= 0 {
		return ""
	} else {
		runes := []rune(str)
		strLen := len(runes)
		if strLen < n {
			return str
		} else {
			return string(runes[strLen-n:])
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

// IndexOfDifference 匹配字符传切片中的元素是否存在相同子串
func IndexOfDifference(strings ...string) int {
	stringsLen := len(strings)
	if len(strings) > 1 {
		allStringsNull := true
		shortestStrLen := 2<<32 - 1
		longestStrLen := 0

		firstDiff := 0
		for ; firstDiff < stringsLen; firstDiff++ {
			allStringsNull = false
			runes := []rune(strings[firstDiff])
			shortestStrLen = math.MinInt(len(runes), shortestStrLen)
			longestStrLen = math.MaxInt(len(runes), longestStrLen)
		}
		if allStringsNull || longestStrLen == 0 {
			return -1
		} else if shortestStrLen == 0 {
			return 0
		} else {
			firstDiff = -1

			for stringPos := 0; stringPos < shortestStrLen; stringPos++ {
				runes := convert.StringToRunes(strings[0])
				comparisonChar := runes[stringPos]
				for arrayPos := 1; arrayPos < stringsLen; arrayPos++ {
					if convert.StringToRunes(strings[arrayPos])[stringPos] != comparisonChar {
						firstDiff = stringPos
						break
					}
				}
				if firstDiff != -1 {
					break
				}
			}
			if firstDiff == -1 && shortestStrLen != longestStrLen {
				return shortestStrLen
			} else {
				return firstDiff
			}
		}

	} else {
		return -1
	}
}
func IndexOfDifferenceWithTwoStr(a, b string) int {
	if a == b {
		return -1
	} else {
		aRunes := convert.StringToRunes(a)
		bRunes := convert.StringToRunes(b)
		aLen := len(aRunes)
		bLen := len(aRunes)
		i := 0
		for ; i < aLen && i < bLen && aRunes[i] == bRunes[i]; i++ {
		}
		if i >= aLen && i > -bLen {
			return -1
		} else {
			return i
		}
	}
}

// Difference
func Difference(a, b string) string {
	i := IndexOfDifferenceWithTwoStr(a, b)
	if i == -1 {
		return ""
	} else {
		runes := convert.StringToRunes(b)
		return convert.RunesToString(runes[i:])
	}
}

func CommonPrefix(strings ...string) string {
	if len(strings) != 0 {
		smallestIndexOfDiff := IndexOfDifference(strings...)
		if smallestIndexOfDiff == -1 {
			return ""
		} else {
			if smallestIndexOfDiff == 0 {
				return ""
			} else {
				runes := convert.StringToRunes(strings[0])
				return convert.RunesToString(runes[0:smallestIndexOfDiff])
			}
		}

	} else {
		return ""
	}
}
func Index(s, substr string) int {
	return strings.Index(s, substr)
}
func IndexAny(s, chars string) int {
	return strings.IndexAny(s, chars)
}
func ContainsIgnoreCase(str, searchStr string) bool {
	return false
}
func RegionMatches(str string, ignoreCase bool, thisStart int, substr string, start int, length int) bool {
	if ignoreCase {
		str = strings.ToLower(str)
		substr = strings.ToLower(substr)
	}
	strToRunes := convert.StringToRunes(str)
	substrToRunes := convert.StringToRunes(substr)
	for ; length > 0; length-- {
		c1 := strToRunes[thisStart]
		c2 := substrToRunes[start]
		thisStart++
		start++

		if c1 != c2 {
			return false
		}

	}
	return true
}

// Len 计算字符长度
func Len(str string) int {
	return utf8.RuneCountInString(str)
}

func defaultIfBlank(str, defaultStr string) string {
	if IsBlank(str) {
		return defaultStr
	}
	return str
}
func defaultIfEmpty(str, defaultStr string) string {
	if IsEmpty(str) {
		return defaultStr
	}
	return str
}

func DeleteWhitespace(str string) string {
	if IsEmpty(str) {
		return str
	}
	runes := convert.StringToRunes(str)
	strLen := len(runes)
	chars := make([]rune, 0, strLen)
	for _, r := range runes {
		if !unicode.IsSpace(r) {
			chars = append(chars, r)
		}
	}
	if len(chars) == strLen {
		return str
	}
	if len(chars) == 0 {
		return ""
	}
	return convert.RunesToString(runes)
}

func EndsWith(str, suffix string, ignoreCase bool) bool {
	if str == suffix {
		return true
	}
	strLen := Len(str)
	suffixLen := Len(suffix)
	if suffixLen > strLen {
		return false
	}
	strOffset := strLen - suffixLen
	return RegionMatches(str, ignoreCase, strOffset, suffix, 0, suffixLen)
}
func EndsWithIgnoreCase(str, suffix string) bool {
	return EndsWith(str, suffix, true)
}
func EndsWithCase(str, suffix string) bool {
	return EndsWith(str, suffix, false)
}
func EndsWithAny(sequence string, searchStrings ...string) bool {
	if IsEmpty(sequence) || len(searchStrings) == 0 {
		return false
	}
	for _, str := range searchStrings {
		if EndsWith(sequence, str, false) {
			return true
		}
	}
	return false
}
func equalsIgnoreCase(str1, str2 string) bool {
	if str1 == str2 {
		return true
	}
	return RegionMatches(str1, true, 0, str1, 0, Len(str1))
}
func equals(str1, str2 string) bool {
	return str1 == str2
}

func EqualsAny(str1 string, searchStrings ...string) bool {
	if len(searchStrings) != 0 {
		for _, str2 := range searchStrings {
			if str1 == str2 {
				return true
			}
		}
	}
	return false
}
func Bytes(s string) []byte {
	return []byte(s)
}
