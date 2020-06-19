// 数据类型转换工具
package utils

import (
	"reflect"
	"unsafe"
)

// ToString 不重新分配内存将字节切片转换为字符串。
func ToString(b *[]byte) string {
	return *(*string)(unsafe.Pointer(b))
}

// ToBytes 不重新分配内存将字符串转换为字节切片
func ToBytes(v string) (b []byte) {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&v))
	byteHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	byteHeader.Data = strHeader.Data
	l := len(v)
	byteHeader.Len = l
	byteHeader.Cap = l
	return
}

// BinaryToBools 不重新分配内存将字节切片转换为布尔切片
func BinaryToBools(b *[]byte) []bool {
	return *(*[]bool)(unsafe.Pointer(b))
}

// BoolsToBinary 不重新分配内存将布尔切片转换为字节切片
func BoolsToBinary(v *[]bool) []byte {
	return *(*[]byte)(unsafe.Pointer(v))
}
