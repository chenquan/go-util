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

type indexOutOfBoundsException struct {
	str string
}

func NewIndexOutOfBoundsExceptionDefault() *indexOutOfBoundsException {
	return &indexOutOfBoundsException{}
}

func NewIndexOutOfBoundsException(str string) *indexOutOfBoundsException {
	return &indexOutOfBoundsException{str: str}
}

// Error 实现 error 接口
func (i indexOutOfBoundsException) Error() string {
	str := "index out of bound"
	if len(i.str) != 0 {

		return str + "," + i.str
	}
	return str
}
