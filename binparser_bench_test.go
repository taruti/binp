package binp

import (
	"encoding/binary"
	"testing"
)

func BenchmarkNativeParser(b *testing.B) {
	v := vs{0x8877665544332211, 0x01234567, 0xFEDC, 42, "fofof"}
	bs := Out().N64(v.v64).N32(v.v32).N16(v.v16).Byte(v.v8).Align(4).N32String(v.s).Out()
	var r vs
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewParser(bs).N64(&r.v64).N32(&r.v32).N16(&r.v16).Byte(&r.v8).Align(4).N32String(&r.s).End()
	}
}

func BenchmarkBEParser(b *testing.B) {
	v := vs{0x8877665544332211, 0x01234567, 0xFEDC, 42, "fofof"}
	bs := Out().B64(v.v64).B32(v.v32).B16(v.v16).Byte(v.v8).Align(4).B32String(v.s).Out()
	var r vs
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewParser(bs).B64(&r.v64).B32(&r.v32).B16(&r.v16).Byte(&r.v8).Align(4).B32String(&r.s).End()
	}
}

func BenchmarkBEParserRaw(b *testing.B) {
	v := vs{0x8877665544332211, 0x01234567, 0xFEDC, 42, "fofof"}
	bs := Out().B64(v.v64).B32(v.v32).B16(v.v16).Byte(v.v8).Align(4).B32String(v.s).Out()
	var r vs
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if len(bs) < 20 {
			panic("fixme")
		}
		r.v64 = binary.BigEndian.Uint64(bs[0:])
		r.v32 = binary.BigEndian.Uint32(bs[8:])
		r.v16 = binary.BigEndian.Uint16(bs[12:])
		r.v8 = bs[14]
		tmp := binary.BigEndian.Uint32(bs[16:])
		if len(bs) < 20+int(tmp) {
			panic("fixme2")
		}
		r.s = string(bs[20 : 20+tmp])
	}
}
