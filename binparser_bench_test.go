package binp

import "testing"

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
