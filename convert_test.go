package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	v := "hi there"

	b := ToBytes(v)
	assert.NotEmpty(t, b)
	assert.Equal(t, v, string(b))

	o := ToString(&b)
	assert.NotEmpty(t, b)
	assert.Equal(t, v, o)
}

func TestBools(t *testing.T) {
	v := []bool{true, false, true, true, false, false}

	b := BoolsToBinary(&v)
	assert.NotEmpty(t, b)
	assert.Equal(t, []byte{0x1, 0x0, 0x1, 0x1, 0x0, 0x0}, b)

	o := BinaryToBools(&b)
	assert.NotEmpty(t, b)
	assert.Equal(t, v, o)
}
func BenchmarkBinaryToBools(b *testing.B) {
	var tmp = []byte{0x1}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinaryToBools(&tmp)
	}
}
func BenchmarkBoolsToBinary(b *testing.B) {
	var tmp = []bool{true}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BoolsToBinary(&tmp)
	}
}

func BenchmarkToString(b *testing.B) {
	var tmp = []byte("xm54Sj0srWlSEctra-yU6ZA6Z2e6pp7c/a/roman/is/da/best/?opt1=true&opt2=false")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToString(&tmp)
	}
}
func BenchmarkToBytes(b *testing.B) {
	var tmp = "xm54Sj0srWlSEctra-yU6ZA6Z2e6pp7c/a/roman/is/da/best/?opt1=true&opt2=false"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToBytes(tmp)
	}
}
