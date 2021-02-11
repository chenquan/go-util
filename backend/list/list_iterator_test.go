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

import (
	"github.com/chenquan/go-util/backend/api/collection"
	"github.com/chenquan/go-util/backend/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sliceListIterator(t *testing.T) {
	l := &itrList{
		lastRet: -1,
		cursor:  0,
		data: &SliceList{
			size: 3,
			data: []collection.Element{"1", 2, "3"},
		},
	}
	var hasNext bool
	var next collection.Element
	var err error

	hasNext = l.HasNext()
	next, err = l.Next()
	assert.Equal(t, true, hasNext)
	assert.Equal(t, "1", next)
	assert.Equal(t, nil, err)

	hasNext = l.HasNext()
	next, err = l.Next()
	assert.Equal(t, true, hasNext)
	assert.Equal(t, 2, next)
	assert.Equal(t, nil, err)

	hasNext = l.HasNext()
	next, err = l.Next()
	err2 := l.Remove()
	assert.Equal(t, true, hasNext)
	assert.Equal(t, "3", next)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, err2)
	assert.Equal(t, []collection.Element{"1", 2}, l.data.Slice())

	hasNext = l.HasNext()
	next, err = l.Next()
	assert.Equal(t, false, hasNext)
	assert.Equal(t, nil, next)
	assert.Equal(t, errs.NoSuchElement, err)

	err2 = l.Remove()
	assert.Equal(t, errs.IllegalState, err2)

}
