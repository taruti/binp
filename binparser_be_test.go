package bigendian

import "testing"

func TestParser(t *testing.T) {
	if ntohs(0x1234) != 0x3412 {
		t.Fatal("ntohs")
	}

	var by byte
	NewParser([]byte{11}).Byte(&by).End()
	if by != 11 {
		t.Fatal("Byte")
	}

	var u16 uint16
	NewParser([]byte{0x11, 0x22}).U16(&u16).End()
	if u16 != 0x1122 {
		t.Fatal("U16")
	}

	var u32 uint32
	NewParser([]byte{0x11, 0x22, 0x33, 0x44}).U32(&u32).End()
	if u32 != 0x11223344 {
		t.Fatal("U32")
	}

	var u64 uint64
	NewParser([]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}).U64(&u64).End()
	if u64 != 0x1122334455667788 {
		t.Fatal("U64")
	}

	var s string
	NewParser([]byte("\000\000\000\003foo")).U32String(&s).End()
	if s != "foo" {
		t.Fatal("U32String")
	}
}

func TestEndToEnd(t *testing.T) {
	v := uint64(0x1122334455667788)
	var w uint64
	NewParser(NewPrinter().U64(v).Out()).U64(&w).End()
	if v!=w { t.Fatal("E2E: U64") }
}

