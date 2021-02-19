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

// Package math 数学运算
package math

// MaxInt 返回最大值
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinInt 返回最小值
func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// MaxInt64 返回最大值
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// MinInt64 返回最小值
func MinInt64(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func MinInt32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func MaxInt16(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

func MinInt16(a, b int16) int16 {
	if a < b {
		return a
	}
	return b
}

func MaxInt8(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

func MinInt8(a, b int8) int8 {
	if a < b {
		return a
	}
	return b
}

func MaxUint64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func MinUint64(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
func MaxUint32(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

func MinUint32(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}
func MaxUint16(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}

func MinUint16(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}
func MaxUint8(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}

func MinUint8(a, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}
func MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func MinFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func MaxFloat32(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func MinFloat32(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}
