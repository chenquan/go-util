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

package logging

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewLog(t *testing.T) {
	NewLogFile("logs/log.txt")
	Debug("log out")
}
func TestNew(t *testing.T) {
	New("logs/log.txt")
	Debug("file and stdout")
}

func TestDebug(t *testing.T) {
	buffer := bytes.Buffer{}
	newLog(&buffer)
	Debug("你好")
	index := strings.Index(buffer.String(), "你好")
	assert.Less(t, 0, index)
}

func TestInfo(t *testing.T) {
	buffer := bytes.Buffer{}
	newLog(&buffer)
	Info("你好")
	index := strings.Index(buffer.String(), "你好")
	assert.Less(t, 0, index)
}

func TestWarn(t *testing.T) {
	buffer := bytes.Buffer{}
	newLog(&buffer)
	Warn("你好")
	index := strings.Index(buffer.String(), "你好")
	assert.Less(t, 0, index)
}

func TestError(t *testing.T) {
	buffer := bytes.Buffer{}
	newLog(&buffer)
	Error("你好")
	index := strings.Index(buffer.String(), "你好")
	assert.Less(t, 0, index)
}
