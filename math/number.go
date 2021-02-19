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

package math

import (
	"errors"
)

var (
	ExistNotNumber = errors.New("exist not number")
)

const (
	F64 = iota
	F32
	UI
	I
	UI64
	I64
	UI32
	I32
	UI16
	I16
	UI8
	I8
)

type Number interface {
	Add(n Number) Number
	Delete(n Number) Number
	Multi(n Number) Number
	Div(n Number) Number
	Int() int
	Int64() int64
	Int32() int32
	Int16() int16
	Int8() int8
	Uint() uint
	Uint64() uint64
	Uint32() uint32
	Uint16() uint16
	Uint8() uint8
	Float64() float64
	Float32() float32
	Type() int
}

type Float64 struct {
	n float64
}

func (f Float64) Add(n Number) Number {
	panic("implement me")
}

func (f Float64) Delete(n Number) Number {
	panic("implement me")
}

func (f Float64) Multi(n Number) Number {
	panic("implement me")
}

func (f Float64) Div(n Number) Number {
	panic("implement me")
}

func (f Float64) Int() int {
	panic("implement me")
}

func (f Float64) Int64() int64 {
	panic("implement me")
}

func (f Float64) Int32() int32 {
	panic("implement me")
}

func (f Float64) Int16() int16 {
	panic("implement me")
}

func (f Float64) Int8() int8 {
	panic("implement me")
}

func (f Float64) Uint() uint {
	panic("implement me")
}

func (f Float64) Uint64() uint64 {
	panic("implement me")
}

func (f Float64) Uint32() uint32 {
	panic("implement me")
}

func (f Float64) Uint16() uint16 {
	panic("implement me")
}

func (f Float64) Uint8() uint8 {
	panic("implement me")
}

func (f Float64) Float64() float64 {
	panic("implement me")
}

func (f Float64) Float32() float32 {
	panic("implement me")
}

func (f Float64) Type() int {
	panic("implement me")
}

type Int struct {
	n int
}

func (i *Int) Uint() uint {
	return uint(i.n)
}

func numberType(n1, n2 Number) int {
	if n1.Type() < n2.Type() {
		return n1.Type()
	}
	return n2.Type()
}
func (i *Int) Add(n Number) Number {
	//reflect.TypeOf()
	typeName := numberType(i, n)
	switch typeName {
	case F64:
		num := i.Float64() + n.Float64()
		return &Float64{n: num}
	case F32:
		num := i.Float32() + n.Float32()

	}
}

func (i *Int) Delete(n Number) Number {
	panic("implement me")
}

func (i *Int) Multi(n Number) Number {
	panic("implement me")
}

func (i *Int) Div(n Number) Number {
	panic("implement me")
}

func (i *Int) Int() int {
	panic("implement me")
}

func (i *Int) Int64() int64 {
	panic("implement me")
}

func (i *Int) Int32() int32 {
	panic("implement me")
}

func (i *Int) Int16() int16 {
	panic("implement me")
}

func (i *Int) Int8() int8 {
	panic("implement me")
}

func (i *Int) Uint64() uint64 {
	panic("implement me")
}

func (i *Int) Uint32() uint32 {
	panic("implement me")
}

func (i *Int) Uint16() uint16 {
	panic("implement me")
}

func (i *Int) Uint8() uint8 {
	panic("implement me")
}

func (i *Int) Float64() float64 {
	panic("implement me")
}

func (i *Int) Float32() float32 {
	panic("implement me")
}

func (i *Int) Type() int {
	return I
}

//
//func (n *Number) Float64() (float64, bool) {
//	f, ok := n.number.(float64)
//	return f, ok
//}
//
//type Array struct {
//	numbers []Number
//	size    int
//}
//
//func NewArray(numbers ...Number) (*Array, error) {
//	for _, number := range numbers {
//		if !IsNumber(number) {
//			return nil, ExistNotNumber
//		}
//	}
//	return &Array{numbers: numbers, size: len(numbers)}, nil
//}
//func (a *Array) ArgMax() int {
//	maxIndex := 0
//	for i := 1; i < a.size; i++ {
//		number := a.numbers[maxIndex]
//	}
//	return 0
//}
