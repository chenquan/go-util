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
	"fmt"
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	task := NewTask()
	var JobFunc1 JobFunc = func() error {
		fmt.Println("第一个任务，我运行了")
		return nil
	}
	var JobFunc2 JobFunc = func() error {
		fmt.Println("第二个任务，我运行了")
		return nil
	}
	newJob1 := NewJob(JobFunc1)
	newJob2 := NewJob(JobFunc2)
	jobId1 := task.Add(newJob1)
	jobId2 := task.Add(newJob2)
	fmt.Println(jobId1)
	if jobId1 == "" {
		t.Errorf("Add error")

	}
	job1 := task.Get(jobId1)
	job2 := task.Get(jobId2)
	if job1 != newJob1 {
		t.Errorf("Get error")
	}
	err := task.Run(jobId1)
	err = job2.Run()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 5)
	task.Stop(jobId1)
	job2.Stop()
	time.Sleep(time.Second * 5)

	id := job1.JobId()
	task.Delete(id)
	if task.Get(id) != nil {
		t.Errorf("Delete error")
	}
}
