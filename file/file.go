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
	"io"
	"io/ioutil"
	"os"
	"path"
)

// GetSize 获取文件大小
func GetSize(f io.Reader) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetExt 获取文件后辍名
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckNotExist 检查文件是否不存在
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckExist 检查文件是否存在
func CheckExist(src string) bool {
	_, err := os.Stat(src)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// CheckPermission 检查文件是否具有权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	//if err == nil {
	//	return true
	//}
	return os.IsPermission(err)
}

// IsNotExistMkDir 创建目录(如果不存在)
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir 创建目录
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open 根据特定模式打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// MustOpen 最大化尝试打开文件
// filePath是文件路径,fileName是文件名
func MustOpen(filepath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}
	dirPath, _ := path.Split(filepath)
	dir = dir + "/" + dirPath
	perm := CheckPermission(dir)
	if perm {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", filepath)
	}

	err = IsNotExistMkDir(dir)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", filepath, err)
	}

	f, err := Open(filepath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to OpenFile :%v", err)
	}

	return f, nil
}
