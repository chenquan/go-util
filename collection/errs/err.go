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
// 错误处理包
package errs

// indexOutOfBoundsException 实现 errs 接口
type indexOutOfBoundsError struct {
	str string
}

// NewIndexOutOfBoundsExceptionDefault 新增默认索引超出错误
func NewIndexOutOfBoundsErrorDefault() *indexOutOfBoundsError {
	return &indexOutOfBoundsError{}
}

// NewIndexOutOfBoundsError 新增默认索引超出错误
func NewIndexOutOfBoundsError(str string) *indexOutOfBoundsError {
	return &indexOutOfBoundsError{str: str}
}

// Error 实现 errs 接口
func (i indexOutOfBoundsError) Error() string {
	str := "index out of bound"
	if len(i.str) != 0 {

		return str + "," + i.str
	}
	return str
}
