package binp

import "testing"

type vs struct {
	v64 uint64
	v32 uint32
	v16 uint16
	v8  byte
	s   string
}

func TestParserNative(t *testing.T) {
	v := vs{0x8877665544332211, 0x01234567, 0xFEDC, 42, "fofof"}
	bs := Out().N64(v.v64).N32(v.v32).N16(v.v16).Byte(v.v8).Align(4).N32String(v.s).Out()
	var r vs

	NewParser(bs).N64(&r.v64).N32(&r.v32).N16(&r.v16).Byte(&r.v8).Align(4).N32String(&r.s).End()

	if v != r {
		t.Fatal("Print->Parse->Equal failed for vs")
	}
}

func TestLen(t *testing.T) {
	var l1, l2 Len
	bs := Out().LenN16(&l1).Byte(0x11).LenStart(&l1).Byte(0x22).Byte(0x33).LenStart(&l2).Byte(0x44).LenN16(&l2).LenN16(&l2).LenDone(&l1).LenDone(&l2).Out()
	//                 0700        11                       22         33                       44         0500        0500
	if string(bs) != "\x07\x00\x11\x22\x33\x44\x05\x00\x05\x00" {
		t.Fatal("TestLen unexpected result")
	}
}
