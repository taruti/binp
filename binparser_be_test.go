package binp

import "testing"

func TestParserBE(t *testing.T) {
	var by byte
	NewParser([]byte{11}).Byte(&by).End()
	if by != 11 {
		t.Fatal("Byte")
	}

	var u16 uint16
	NewParser([]byte{0x11, 0x22}).B16(&u16).End()
	if u16 != 0x1122 {
		t.Fatal("B16")
	}

	var u32 uint32
	NewParser([]byte{0x11, 0x22, 0x33, 0x44}).B32(&u32).End()
	if u32 != 0x11223344 {
		t.Fatal("B32")
	}

	var u64 uint64
	NewParser([]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}).B64(&u64).End()
	if u64 != 0x1122334455667788 {
		t.Fatal("B64")
	}

	var s string
	NewParser([]byte("\000\000\000\003foo")).B32String(&s).End()
	if s != "foo" {
		t.Fatal("B32String")
	}
}

func TestEndToEnd(t *testing.T) {
	v := uint64(0x1122334455667788)
	var w uint64
	NewParser(Out().B64(v).Out()).B64(&w).End()
	if v != w {
		t.Fatal("E2E: B64")
	}
}
