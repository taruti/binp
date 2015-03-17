package binp

import (
	"encoding/binary"
	"testing"
)

func BenchmarkEncodingBinaryPut(b *testing.B) {
	bs := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint32(bs, 31)
	}
}

func BenchmarkEncodingBinaryGet(b *testing.B) {
	bs := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		binary.BigEndian.Uint32(bs)
	}
}

func BenchmarkGetB32EB(b *testing.B) {
	bs := make([]byte, 16)
	var x, y uint32
	for i := 0; i < b.N; i++ {
		x = binary.BigEndian.Uint32(bs)
		y = binary.BigEndian.Uint32(bs[4:])
		x = binary.BigEndian.Uint32(bs[8:])
		y = binary.BigEndian.Uint32(bs[12:])
	}
	_, _ = x, y
}

func BenchmarkGetB32(b *testing.B) {
	bs := make([]byte, 16)
	var x, y uint32
	for i := 0; i < b.N; i++ {
		p := NewParser(bs)
		p.B32(&x).B32(&y)
		p.B32(&x).B32(&y)
	}
}

func BenchmarkGetN32(b *testing.B) {
	bs := make([]byte, 16)
	var x, y uint32
	for i := 0; i < b.N; i++ {
		p := NewParser(bs)
		p.N32(&x).N32(&y)
		p.N32(&x).N32(&y)
	}
}

func BenchmarkGetByte(b *testing.B) {
	bs := make([]byte, 8)
	var x uint8
	for i := 0; i < b.N; i++ {
		p := NewParser(bs)
		p.Byte(&x)
		p.Byte(&x)
		p.Byte(&x)
		p.Byte(&x)
		p.Byte(&x)
		p.Byte(&x)
		p.Byte(&x)
		p.Byte(&x)
	}
}
