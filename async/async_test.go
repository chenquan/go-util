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

package async

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestRepeat(t *testing.T) {
	var i int32 = 0
	var startTime time.Time
	var endTime time.Time
	cancelFunc := Repeat(context.TODO(), time.Second*3, func() {
		atomic.AddInt32(&i, 1)
		if i == 1 {
			startTime = time.Now()
		} else if i == 2 {
			endTime = time.Now()
		}
	})
	for {
		if i == 2 {
			cancelFunc()
			ts := endTime.Second() - startTime.Second()
			if ts != 3 {
				t.Error("错误!")
			}
			break
		}
		if i == 3 {
			t.Error("错误!")
		}
	}

}
