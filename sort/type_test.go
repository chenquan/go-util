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

package sort

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestInt64s(t *testing.T) {
	v := Int64s{1, 8, 4, 3}
	sort.Sort(&v)
	assert.Equal(t, []int64(v), []int64{1, 3, 4, 8})
}
func TestUint64s(t *testing.T) {
	v := Uint64s{1, 8, 4, 3}
	sort.Sort(&v)
	assert.Equal(t, []uint64(v), []uint64{1, 3, 4, 8})
}

func TestInt32s(t *testing.T) {
	v := Int32s{1, 8, 4, 3}
	sort.Sort(&v)
	assert.Equal(t, []int32(v), []int32{1, 3, 4, 8})
}
func TestUint32s(t *testing.T) {
	v := Uint32s{1, 8, 4, 3}
	sort.Sort(&v)
	assert.Equal(t, []uint32(v), []uint32{1, 3, 4, 8})
}

func TestUint16s(t *testing.T) {
	v := Uint16s{1, 8, 4, 3}
	sort.Sort(&v)
	assert.Equal(t, []uint16(v), []uint16{1, 3, 4, 8})
}
func TestInt16s(t *testing.T) {
	v := Int16s{1, 8, 4, 3}
	sort.Sort(&v)
	assert.Equal(t, []int16(v), []int16{1, 3, 4, 8})
}
