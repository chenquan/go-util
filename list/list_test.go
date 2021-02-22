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
	"math/rand"
	"testing"
)

// Benchmark
func BenchmarkNewLinkedList(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewLinkedList()
	}
}
func BenchmarkNewSliceList(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewSliceListDefault()
	}
}
func BenchmarkLinkedList_Add(b *testing.B) {
	list := NewLinkedList()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = list.Add(i)
	}
}

func BenchmarkSliceList_Add(b *testing.B) {
	list := NewSliceListDefault()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = list.Add(i)
	}
}
func BenchmarkLinkedList_AddAll(b *testing.B) {
	list := NewLinkedList()
	listDefault := NewSliceListDefault()
	for i := 0; i < 100; i++ {
		_, _ = listDefault.Add(i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = list.AddAll(listDefault)
	}
}
func BenchmarkSliceList_AddAll(b *testing.B) {
	list := NewSliceListDefault()
	listDefault := NewSliceListDefault()
	for i := 0; i < 100; i++ {
		_, _ = listDefault.Add(i)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = list.AddAll(listDefault)
	}
}

func BenchmarkLinkedList_AddIndex(b *testing.B) {
	list := NewLinkedList()
	rand.Seed(2020)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Int()

		_ = list.AddIndex(i%n, 00)
	}
}

func BenchmarkSliceList_AddIndex111(b *testing.B) {
	list := NewSliceListDefault()
	rand.Seed(2020)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Int()
		_ = list.AddIndex(i%n, i)
	}
}
func BenchmarkLinkedList_Get(b *testing.B) {
	list := NewLinkedList()
	rand.Seed(2020)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Int()
		_, _ = list.Add(i)
		_, _ = list.Get(i & n)
	}
}
