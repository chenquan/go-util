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

// 时间处理函数
package time

import (
	"time"
)

const (
	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"
)

// Year 当前年
func Year() int {
	return time.Now().Year()
}

// YearInLocal 当前年
// location 区域信息
func YearInLocal(location *time.Location) int {
	return time.Now().In(location).Year()
}

// Month 当前月
func Month() time.Month {
	return time.Now().Month()
}

// Month 当前月
// location 区域信息
func MonthInLocal(location *time.Location) time.Month {
	return time.Now().In(location).Month()
}

// Day 当前天
func Day() int {
	return time.Now().Day()
}

// DayInLocal 当前天
func DayInLocal(location *time.Location) int {
	return time.Now().In(location).Day()
}

// YearDay 一年之中的当前天数
func YearDay() int {
	return time.Now().YearDay()
}

// YearDayInLocal 一年之中的当前天数
// location 区域信息
func YearDayInLocal(location *time.Location) int {
	return time.Now().In(location).YearDay()
}

// ToDateFormat 时间转日期字符串
// t 时间
func ToDateFormat(t *time.Time) string {
	return t.Format(DateFormat)
}

// ToDateFormatInLocal 时间转日期字符串
// t 时间
// location 区域信息
func ToDateFormatInLocal(t *time.Time, location *time.Location) string {
	return t.In(location).Format(DateFormat)
}

// NowDateFormat 当前日期字符串
func NowDateFormat() string {
	return time.Now().Format(DateFormat)
}

// NowDateFormatInLocal 当前时间日期字符串
// location 区域信息
func NowDateFormatInLocal(location *time.Location) string {
	return time.Now().In(location).Format(DateFormat)
}

// ToDateTimeFormat 时间转长日期字符串
// t 时间
func ToDateTimeFormat(t *time.Time) string {
	return t.Format(DateTimeFormat)
}

// ToDateTimeFormatInLocal 时间转长日期字符串
// t 时间
// location 区域
func ToDateTimeFormatInLocal(t *time.Time, location *time.Location) string {
	return t.In(location).Format(DateTimeFormat)
}

// NowDateTimeFormat 当前时间长日期字符串
func NowDateTimeFormat() string {
	return time.Now().Format(DateTimeFormat)
}

// NowDateTimeFormatInLocal 当前时间长日期字符串
// location 区域
func NowDateTimeFormatInLocal(location *time.Location) string {
	return time.Now().In(location).Format(DateTimeFormat)
}
