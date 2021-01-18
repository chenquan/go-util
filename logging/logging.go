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
	"fmt"
	"github.com/chenquan/go-utils/file"
	"io"
	log "log"
	"os"
	"path/filepath"
	"runtime"
)

// Level 存放日志级别
type Level uint8

// 日志级别
const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var (
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

// newLog 新建日志
func newLog(out io.Writer) {
	logger = log.New(out, "", log.LstdFlags)
}

// NewLog 新建标准日志
func NewLog() {
	newLog(os.Stdout)
}

// NewLogFile 新建文件日志
func NewLogFile(filepath string) {
	f, _ := file.MustOpen(filepath)

	newLog(f)
}

// New 新建文件日志和标准日志
func New(filepath string) {
	f, _ := file.MustOpen(filepath)
	writer := io.MultiWriter(f, os.Stdout)
	newLog(writer)
}

// Debug debug级日志输出
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}

// Info info级日志输出
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

// Warn warn级日志输出
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

// Error error级日志输出
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

// Fatal  fatal级日志输出
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

// setPrefix 设置日志输出的前缀
func setPrefix(level Level) {
	var _, fileName, line, ok = runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(fileName), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	if logger == nil {
		NewLog()
	}
	logger.SetPrefix(logPrefix)
}
