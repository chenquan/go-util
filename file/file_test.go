/*
 *    Copyright 2020 Chen Quan
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

package file

import (
	"fmt"
	"os"
	"testing"
)

var filePath = "test/test1.txt"

func TestCheckNotExist(t *testing.T) {
	if CheckNotExist("test/test1.txt") {
		t.Error("error")
	}
	if !CheckNotExist("test/test11.txt") {
		t.Error("error")
	}
}

func TestGetSize(t *testing.T) {
	open, _ := os.Open("test/test11.txt")
	size, err := GetSize(open)
	if err != nil {
		fmt.Println(err)
		fmt.Println(size)

	} else {
		t.Error("error")
	}
	open, _ = os.Open("test/test1.txt")
	size, err = GetSize(open)
	if err != nil {
		t.Error("error")
	} else {
		if size != 8 {
			t.Error("error")

		}
	}
}

func TestCheckExist(t *testing.T) {
	fmt.Println(CheckExist("test/test1.txt"))
	if !CheckExist("test/test1.txt") {
		t.Error("error")
	}
	if CheckExist("test/test11.txt") {
		t.Error("error")
	}
}

func TestCheckPermission(t *testing.T) {
	if !CheckPermission("test/test1.txt") {
		t.Error("error")
	}
}

func TestGetExt(t *testing.T) {
	if GetExt(filePath) != ".txt" {
		t.Error("error")
	}
	if GetExt("test/log.log") != ".log" {
		t.Error("error")
	}

}

func TestMustOpen(t *testing.T) {
	if _, err := MustOpen("test/test2.txt"); err != nil {

	}
}
