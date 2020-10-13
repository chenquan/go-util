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

package task

import (
	"log"
	"testing"
	"time"
)

func TestJob(t *testing.T) {
	var jobFunc JobFunc = func() error {
		time.Sleep(time.Second * 5)
		return nil
	}
	job := NewJob(jobFunc)
	log.Println("1.运行")
	_ = job.Run()
	time.Sleep(time.Second)
	log.Println("2.停止")
	job.Stop()
	time.Sleep(time.Second)
	log.Println("2.运行")
	_ = job.RunWithDeadline(time.Now().Add(time.Second * 2))
	_ = job.RunWithTimeout(time.Second * 2)

}
